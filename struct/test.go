package main

import "fmt"

func main() {
	table := "`user`"
	query := fmt.Sprintf("delete from %s where `id` = ?", table)

	fmt.Println("Hello, World!", query)
}
