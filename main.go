package main

import (
	"fmt"

	"github.com/kashifkhan0771/utils/slugger"
)

func main() {
	s := slugger.New(map[string]string{}, false, true)
	fmt.Println(s.Slug("❤️", "/"))
}
