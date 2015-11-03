package main

import (
	"fmt"
)
import (
	"./services"
)


func main() {
	// fmt.Printf("study test")

	listService := services.ListService{}
	// data := listService.Service()

	// listService.Cluster()
	data := listService.Cluster()

	fmt.Println("====================================")
	fmt.Println(data)
	
	
}
