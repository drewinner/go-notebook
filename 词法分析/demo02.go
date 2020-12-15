package main

import (
	"fmt"
	"os"
)

func envOr(key, value string) string {
	if x := os.Getenv(key); x != "" {
		return x
	}
	return value
}
func main() {
	fmt.Println(envOr("GOARCH", ""))
}
