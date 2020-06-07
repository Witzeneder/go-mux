// main.go

package main

import "os"

func main() {
	a := App{}
	a.Initialize(
		"postgres",
		"",
		"postgres")

	a.Run(":8010")
}
