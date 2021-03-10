// Movie prints Movies as JSON.
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Movie struct {
	Title string
	Year  int  `json:"released_year"`
	Color bool `json:"color,omitempty"`
	Actor []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actor: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actor: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actor: []string{"Steve McQueen", "Jacqueline Bisset"}},
}

func jsonDumpstoFile(filename string) {

	data, err := json.MarshalIndent(movies, "", "	")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s\n", err)
	}
	fmt.Printf("%s\n", data)

	ioutil.WriteFile(filename, data, 0777)
}

func jsonLoadsfromFile(filename string) {
	var loadResult []Movie
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Err occured when reading [%s] : %v", filename, err)
	}

	err = json.Unmarshal(data, &loadResult)
	if err != nil {
		fmt.Printf("Err occured when json.Unmarshal : %v", err)
	}

	for _, item := range loadResult {
		fmt.Printf("T: %10.10s Y: %d A: %10.10s C:%v\n", item.Title, item.Year, item.Actor, item.Color)
	}
	return

}

func jsonLoadsPartObject(filename string) {
	type PartOfMovie struct {
		Title string
		Year  int `json:"released_year"`
	}
	var loadResult []PartOfMovie
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Err occured when reading [%s] : %v", filename, err)
	}

	err = json.Unmarshal(data, &loadResult)
	if err != nil {
		fmt.Printf("Err occured when json.Unmarshal : %v", err)
	}

	for _, item := range loadResult {
		fmt.Printf("T: %10.10s Y: %d \n", item.Title, item.Year)
	}
	return
}

func main() {
	// {
	// 	data, err := json.Marshal(movies)
	// 	if err != nil {
	// 		log.Fatalf("JSON marshaling failed: %s", err)
	// 	}
	// 	fmt.Printf("%s\n", data)
	// }

	// {
	filename := "output.json"
	jsonDumpstoFile(filename)
	jsonLoadsfromFile(filename)
	jsonLoadsPartObject(filename)

	// 	var titles []struct{ Title string }
	// 	if err := json.Unmarshal(data, &titles); err != nil {
	// 		log.Fatalf("JSON unmarshaling failed: %s", err)
	// 	}
	// 	fmt.Println(titles)
	// }
}
