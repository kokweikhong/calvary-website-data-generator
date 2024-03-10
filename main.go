package main

import "fmt"

func main() {
	keys, err := ListObjects("")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(keys)
}
