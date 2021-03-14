//本文件复现了第五章 5.6节末尾的警告：捕获迭代变量
package main

import (
	"fmt"
	"os"
)

const prefix = "./temp"

func tempDirs() []string {
	var dirs []string
	for i := 0; i < 10; i++ {
		dirs = append(dirs, fmt.Sprintf("%s/dir_%d", prefix, i))
	}
	return dirs

}

func lsDir() {
	dirsEntry, err := os.ReadDir(prefix)
	if err != nil {
		fmt.Printf("When lsDir, error: %v, return...", err)
		return
	}
	if len(dirsEntry) == 0 {
		fmt.Printf("Dir %s is empty\n", prefix)
	}
	for i, entry := range dirsEntry {
		fmt.Printf("[%d]-th dir: %s\n", i, entry.Name())
	}

}

func main() {
	type rmdirFunc func()
	var rmdirs []rmdirFunc
	fmt.Printf("rmdirs : %v %T\n", rmdirs, rmdirs)
	for _, tempDir := range tempDirs() {
		// tempDir := d
		os.MkdirAll(tempDir, 0755)
		rmdirs = append(rmdirs, func() {
			os.RemoveAll(tempDir)
		})
	}

	fmt.Println("---- Before Remove ----")
	lsDir()

	for _, rmdir := range rmdirs {
		rmdir()
	}
	fmt.Println("---- Remove Finished ----")
	lsDir()
}
