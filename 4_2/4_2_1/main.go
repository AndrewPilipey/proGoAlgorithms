//It's simple linearSearch

package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4, 5} //slice

	key := 6

	if linearSearch(list, key) == false {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}

func linearSearch(list []int, key int) bool {
	for _, value := range list {
		if value == key {
			return true
		}
	}
	return false
}
