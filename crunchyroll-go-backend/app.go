package main

import (
	"fmt"
)

type TestStr struct {
	firstName string
}

func main() {
	firsNm := TestStr{"George"}
	secNm := TestStr{"John"}
	fmt.Printf("Here is the workable app made by %v and %v.", firsNm, secNm)
}
