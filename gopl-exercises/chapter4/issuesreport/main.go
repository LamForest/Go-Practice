// Issuesreport prints a report of issues matching the search terms.
package main

import (
	"fmt"
	"gopl/chapter4/issues/github"
	"log"
	"os"
	"text/template"
	"time"
)

// 比较值得注意的是循环的写法，以及管道的使用
const templ = `
{{range .}}---------------------------------------
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age: {{.CreatedAt | daysAgo}} days
{{end}}`

//传入模板作为函数
func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func main() {
	var report = template.Must(template.New("issuelist").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(templ))
	result, err := github.SearchIssues(os.Args[1:])

	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d items in total.", len(*result))
}

func noMust() {
	report, err := template.New("report").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(templ)
	if err != nil {
		log.Fatal(err)
	}
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
