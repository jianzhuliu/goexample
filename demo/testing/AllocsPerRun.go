package main

import (
	"fmt"
	"testing"
	"strings"
)

var x []byte = []byte{1023:'x'}
var y []byte = []byte{1023:'y'}
var s string 

func fc() {
	s = (" " + string(x) + string(y))[1:]
}

func fd(){
	s = string(x) + string(y)
}

func fe(){
	t := strings.Builder{}
	t.Write(x)
	t.Write(y)
	s = t.String()
}

func main(){
	fmt.Println("fc:",testing.AllocsPerRun(1, fc))
	fmt.Println("fe:",testing.AllocsPerRun(1, fe))
	fmt.Println("fd:",testing.AllocsPerRun(1, fd))
}