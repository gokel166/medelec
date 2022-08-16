package main

import (
	myCustomHandlers "crunchydata-app/handlers"
	mymodel "crunchydata-app/models"
	"fmt"
)

func main() {
	genString := "My made up string."
	mymodel.PrintString(genString)
	getMyStaticVideoCollections()
	fmt.Println("Hello and Welcome to the Golang App.")
}

func getMyStaticVideoCollections() {
	fmt.Println(myCustomHandlers.GetVideoTitleCollection())
	for _, anm := range myCustomHandlers.GetVideoTitleCollection() {
		fmt.Println(anm.VideoTitle)
	}
}
