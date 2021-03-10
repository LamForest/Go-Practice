// Issues prints a table of GitHub issues matching the search terms.
//TODO 最多只有30个issue返回，问题在哪？
package main

import (
	"fmt"
	"gopl/chapter4/issues/github"
	"log"
	"os"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range *result {
		fmt.Printf("#%-5d %9.9s %5s, %.55s\n",
			item.Number, item.User.Login, item.State, item.Title)
	}
	fmt.Printf("%d issues in Total.\n", len(*result))

}
