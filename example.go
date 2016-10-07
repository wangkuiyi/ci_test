package main

import (
	"fmt"
	"os"

	"github.com/topicai/candy"
)

func main() {
	dir, e := os.Getwd()
	candy.Must(e)
	fmt.Println(dir)
}
