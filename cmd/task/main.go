package main

import "fmt"

type Task struct {
	ID              int
	raw             string
	done            bool
	priority        string
	completion_date string
	creation_date   string
	todo            string
	project         string
	context         string
	tags            map[string]string
}

func main() {
	fmt.Println("Hello Tempus!")
}
