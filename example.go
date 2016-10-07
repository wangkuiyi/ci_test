package main

import (
	"fmt"
	"os"

	"github.com/topicai/candy"
)

// CurrentDir returns the current directory.
func CurrentDir() string {
	dir, e := os.Getwd()
	candy.Must(e)
	return dir
}

func main() {
	fmt.Println(CurrentDir)
}
