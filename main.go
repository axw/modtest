package main

import (
	"fmt"

	v2_version "github.com/axw/modtest/v2/version"
	v3_version "github.com/axw/modtest/v3/version"
)

func main() {
	fmt.Println(v2_version.Version)
	fmt.Println(v3_version.Version)
}
