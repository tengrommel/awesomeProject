package main

import (
	"container/list"
	"fmt"
)

// link list
func main()  {
	var elementList list.List
	elementList.PushBack(11)
	elementList.PushBack("22")
	elementList.PushBack(true)
	for element := elementList.Front(); element!=nil; element = element.Next() {
		fmt.Printf("element: %+v\n", element)
	}
}
