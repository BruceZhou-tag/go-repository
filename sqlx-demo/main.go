package main

import "fmt"

// 可在终端中执行go run . or go run main.go db.go
func main() {
	if err := initDB(); err != nil {
		fmt.Printf("init DB failed,err:=%v\n", err)
		return
	}
	fmt.Println("connect MySQL success...")

	queryRowDemo()

	// queryMultiRowDemo()
	// insertRowDemo()
	// updateRowDemo()
	// deleteRowDemo()
	// transactionDemo()
}
