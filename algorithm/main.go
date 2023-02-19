package main

import (
	"log"

	"github.com/chenchen224/easy_golang/algorithm/sort"
)

func main() {
	arr := []int{
		5, 1, 3, 4, 9, 7,
	}
	log.Println(sort.InsertSortOfficial(arr))
}
