package main

import (
	"github.com/dannyroes/golangkwblog/blog"
)

func main() {
	err := blog.ConnectDB()
	if err != nil {
		panic(err)
	}

	blog.Run()
}
