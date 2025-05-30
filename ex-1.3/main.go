package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	res := ""
	res = strings.Join(os.Args[1:], " ")
	fmt.Println(res)
}
