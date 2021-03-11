// Issueshtml prints an HTML table of issues matching the search terms.
package main

import (
	"fmt"
	"gopl/chapter4/issues/github"
	"html/template"
	"log"
	"net/http"
	"sort"
	"strings"
)

var issueList = template.Must(template.ParseFiles("template.html"))

const ip = "0.0.0.0:8080"

// 不能为其他包内的类型定义其方法
// func (l *github.IssuesSearchResult) Len() int { return len(*l) }
// 所以这里为github.IssuesSearchResult在这个包里重新命名了

// 以下代码用于自定义排序
type lessFunc func(x, y *github.Issue) bool
type customSort struct {
	issues github.IssuesSearchResult
	less   lessFunc
}

//下面为customSort实现了sort.Interface
func (result customSort) Len() int { return len(result.issues) }
func (result customSort) Swap(i, j int) {
	(result.issues)[i], (result.issues)[j] = (result.issues)[j], (result.issues)[i]
}
func (result customSort) Less(i, j int) bool {
	return result.less((result.issues)[i], (result.issues)[j])
}

//issues是一个slice
//切片拷贝时，仅拷贝sliceHeader，底层是共享的
func sortIfRequired(request *http.Request, issues github.IssuesSearchResult) {
	querys := request.URL.Query()
	sortKey, ok := querys["sortkey"]
	if !ok || len(sortKey) == 0 {
		fmt.Printf("sortKey not in URL query parameter, no sort needed\n")
		return
	}

	fmt.Printf("sortkey = %v, will be sorted\n", sortKey)

	var tempFunc lessFunc
	switch sortKey[0] {
	case "title":
		tempFunc = func(x, y *github.Issue) bool {
			return strings.ToLower(x.Title) < strings.ToLower(y.Title)
		}
	case "user":
		tempFunc = func(x, y *github.Issue) bool {
			return strings.ToLower(x.User.Login) < strings.ToLower(y.User.Login)
		}
	default:
		fmt.Printf("invalid sortkey = %v, will not be sorted, return\n", sortKey)

	}
	fmt.Println("Starting Sorting")
	sort.Sort(customSort{
		issues,
		tempFunc,
	})
	fmt.Println("Finished Sorting")
	return
}

func sortResultHandler(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Println("Receive request... access github API")
	githubReturn, err := github.SearchIssues(nil)
	fmt.Println("Received github API Result")
	issues := *githubReturn
	//构造customSort对象
	sortIfRequired(request, issues)

	// sort.Sort(result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Analyzing HTML template")
	if err := issueList.Execute(responseWriter, issues); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Analyze finished")
}

func main() {

	// http.Handle("/", http.FileServer(http.Dir("/home/ubuntu/web/go/code/gopl/Go-Practice/gopl-exercises/chapter7/sort/sort.html")))
	http.HandleFunc("/sort", sortResultHandler)
	fmt.Println("ListenAndServe")
	log.Fatal(http.ListenAndServe(ip, nil))
}
