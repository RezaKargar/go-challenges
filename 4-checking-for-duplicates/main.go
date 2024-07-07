package main

import "fmt"

type Developer struct {
	Name string
	Age  int
}

func FilterUnique(developers []Developer) []string {
	var seen = make(map[string]Developer)
	var result []string
	for _, developer := range developers {
		if _, ok := seen[developer.Name]; !ok {
			seen[developer.Name] = developer
			result = append(result, developer.Name)
		}
	}
	return result
}

func main() {
	fmt.Println("Filter Unique Challenge")

	developers := []Developer{
		Developer{Name: "Elliot"},
		Developer{Name: "Alan"},
		Developer{Name: "Jennifer"},
		Developer{Name: "Graham"},
		Developer{Name: "Paul"},
		Developer{Name: "Alan"},
	}

	filtered := FilterUnique(developers)
	fmt.Println(filtered)
}
