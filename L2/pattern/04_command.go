package main

import "fmt"

type Command interface {
	Execute() string
}
type CheckCPU struct {
}

func (c *CheckCPU) Execute() string {
	return "CPU alive and well"
}

type CheckGPU struct {
}

func (c *CheckGPU) Execute() string {
	return "GPU alive and well"
}

//Запрос к бизнес логике
func request(part string) string {
	commands := map[string]Command{
		"CPU": &CheckCPU{},
		"GPU": &CheckGPU{},
	}
	if command := commands[part]; command == nil {
		return "No such command found"
	} else {
		return command.Execute()
	}
}
func main() {
	fmt.Println(request("CPU"))
}
