package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"kv/stack"
)

func main() {

	s := stack.New()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("kv> ")
		text, _ := reader.ReadString('\n')

		operation := strings.Fields(text)

		switch operation[0] {
		case "EXIT", "STOP":
			fmt.Println("stopped.")
			return

		// Action
		case "SET":
			s.Set(operation[1], operation[2])
		case "GET":
			s.Get(operation[1])
		case "DELETE":
			s.Delete(operation[1])

		// Transaction Control
		case "BEGIN":
			s.Push()
		case "END":
			s.End()
		case "ROLLBACK":
			s.Roll()
		case "COMMIT":
			s.Commit()
		default:
			fmt.Println("unknown action")
		}
	}
}
