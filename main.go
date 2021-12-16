package main

import (
	"fmt"

	"gitlab.com/url-builder/go-admin/src"
)

func main() {
	r := src.BootStrap(".env")

	if err := r.Run(); err != nil {
		fmt.Println("Exception occurred: " + err.Error())
	}
}
