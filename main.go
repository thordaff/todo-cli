package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/thordaff/todo-cli/todo"

)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Gunakan perintah: add/list/done")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Gunakan: add <Judul tugas>")
			return
		}
		title := os.Args[2]
		if err := todo.AddTask(title); err != nil {
			fmt.Println("Error:", err)
		}

	case "list":
		if err := todo.ListTasks(); err != nil {
			fmt.Println("Error:", err)
		}

	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Gunakan: done <id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("ID harus angka.")
			return
		}
		if err := todo.CompleteTask(id); err != nil {
			fmt.Println("Error:", err)
		}

	default:
		fmt.Println("Perintah tidak dikenal. Gunakan: add, list, done")
	}
}