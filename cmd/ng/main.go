package main

import (
	"fmt"
	"os"
	"strings"
)

func main () {
	env, _ := os.Getwd()
	fmt.Println(env)
	fmt.Println(os.Getenv("GOPATH"))
	fmt.Println(strings.TrimPrefix(env, os.Getenv("GOPATH")+"/src/"))
}
