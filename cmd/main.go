package main

import (
	"fmt"
	"spider/internal"
)

func main() {
	fmt.Println("spider going now!")
	spider := internal.NewSpiderEngine()
	spider.Start()
}
