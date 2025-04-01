package main

import (
	"fmt"
)

func main() {
	var frst int
	var zee int
	var jeff string
	var joy int
	tasks := []string{}
	for {

		fmt.Println("\n\nwelcome to the todo list app")
		fmt.Println("1. Create your task")
		fmt.Println("2. Check your task")
		fmt.Println("3. Remove a task")
		fmt.Println("4. Exit the app")
		fmt.Print("Enter your option:")

		_, err := fmt.Scan(&frst)
		if err != nil {
			fmt.Print("error reading input", err)
			continue
		}

		if frst == 1 {
			fmt.Print("How many task do you want to enter:")
			_, err := fmt.Scan(&zee)
			if err != nil {
				fmt.Print("error reading input", err)
				continue
			}

			for i := 0; i < zee; {
				fmt.Printf("enter task %d: ", i+1)
				_, err := fmt.Scan(&jeff)
				if err != nil {
					fmt.Print("error reading input", err)
					break
				}
				if jeff == "" {
					fmt.Println("this cannot be empty please")
					continue
				}
				tasks = append(tasks, jeff)
				i++
			}

		} else if frst == 2 {
			fmt.Println("Your todolist")
			if len(tasks) == 0 {
				fmt.Println("no task available ")
			} else {
				for i, task := range tasks {
					fmt.Printf("%d. %s\n", i+1, task)
				}
			}

		} else if frst == 3 {
			fmt.Print("What task number do you want to remove : ")
			_, err := fmt.Scan(&joy)
			if err != nil || joy < 1 || joy > len(tasks) {
				fmt.Print("error reading input", err)
				continue
			}
			indexremove := joy - 1
			tasks = append(tasks[:indexremove], tasks[indexremove+1:]...)
			fmt.Println("Task removed successfully")
			fmt.Println("Your todolist")
			for i, task := range tasks {
				fmt.Printf("%d. %s\n", i+1, task)
			}

		} else if frst == 4 {
			fmt.Println("Exiting the todo list app")
			return
		} else {
			fmt.Println("enter a valid number")
		}

	}

}
