/*
自己写的 基于互斥锁的函数缓存机制，源于GOPL第9章 9.7节，非常有意思的代码

有几个要点:
1) 为了防止数据竞态，使用了sync.Mutex，这点很普通。
2）为了实现重复抑制（即保证不会有2个goroutine同时看到某个cache[key]为空，然后都调用了慢函数f(key))
先在cache[key]放了一个占位的，然后用通道关闭，广播f(key)调用完成
3） 注意到，Get中仅仅上了一次锁，其余2次对共享变量this的访问没有加锁。第1次是因为只可能有1个goroutine在写item，第2次是因为接触item的都是读goroutine
4) 在Get的第2局，使用了item := this.cache[key]，从此我们尽量避免使用this.cache[key](因为要上锁)，而是使用item
5) 问题：同时对map的不同key进行写入，是否需要上锁？必须，虽然看上去map[key1] 与 map[key2]不在同一个位置，没有冲突，但是实际上可能都会
修改map的某个属性，比如若导致rehash，则直接乱套
*/
package main

import "sync"

type result struct {
	ret        interface{}
	err        error
	isFinished chan struct{}
}

type Func func(string) (interface{}, error)

type Memo struct {
	mutex sync.Mutex
	cache map[string]*result
	f     Func
}

func NewMemo(f Func) *Memo {
	return &Memo{cache: make(map[string]*result), f: f}
}

func (this *Memo) Get(key string) (interface{}, error) {

	this.mutex.Lock()
	item := this.cache[key]
	// 这里不能解锁，因为某个goroutine发现该cache[key]后
	// 应该立刻占住位置，否则会有多个goroutine试图调用慢函数 f(重复抑制失败)
	// this.mutex.Unlock()

	//只有第1个goroutine进入这个if，所以if中的操作不用上锁
	if item == nil {
		//占个位置，但是还没有真正计算
		item = &result{isFinished: make(chan struct{})}
		this.cache[key] = item
		this.mutex.Unlock()

		newRet, newErr := this.f(key)

		//这里不需要上锁，因为只可能有一个goroutine在操作，其他都在等待isFinished
		// this.mutex.Lock()
		// 这里不应该用this.cache[key].ret，而是item.ret，避免对map的同时写入
		// this.cache[key].ret, this.cache[key].err = newRet, newErr
		item.ret, item.err = newRet, newErr
		// 广播其他等待的goroutine，慢函数调用完成了
		close(item.isFinished)
		// this.mutex.Unlock()

	} else {
		this.mutex.Unlock()
		//1)对于通道的读写不需要上锁
		//2)此刻，虽然cache[key]不为空，但不能保证f调用完成，需要等待第一个goroutine的广播
		// 这里用关闭信道模拟广播
		<-item.isFinished
	}

	//这里也不需要上锁，因为都是读goroutine，没有写goroutine
	// this.mutex.Lock()

	// this.mutex.Unlock()
	return item.ret, item.err
}
