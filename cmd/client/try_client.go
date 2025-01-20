package main

import (
	"fmt"

	"github.com/maximakhatov/key-value-store/client"
)

func main() {
	client, err := client.NewClient("localhost:6379")
	if err != nil {
		panic(err)
	}
	defer client.Close()

	err = client.Set("test", "1")
	if err != nil {
		panic(err)
	}
	fmt.Println("`test` is set to 1")

	result, _, err := client.Get("test")
	if err != nil {
		panic(err)
	}
	fmt.Println("`test` equals to:", result)

	_, isNull, err := client.Get("empty_key")
	if err != nil {
		panic(err)
	}
	fmt.Println("`empty_key` is null:", isNull)
}
