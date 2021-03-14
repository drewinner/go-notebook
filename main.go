package main

import (
	"fmt"
	"time"
)

func main() {
	expireAt, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	expireAt = expireAt.Add((24*5 + 2) * time.Hour)
	fmt.Printf("%+v", expireAt)
}
