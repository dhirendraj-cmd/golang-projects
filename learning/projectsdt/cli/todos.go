package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Todo struct{
	id int
	task string
	isCompleted bool
}

var currentID int = 0


// todos constructor
func NewTodo(task string) *Todo{

	currentID++

	newTask := Todo{
		id: currentID,
		task: task,
		isCompleted: false,
	}
	return &newTask
}

// add builtin stringer interface
func (t *Todo) String() string{
	return fmt.Sprintf("{ID: %d, Task: %s, Completed: %t}", t.id, t.task, t.isCompleted)
}

func CreateTask(){
	var addTask string
	var todos []*Todo

	i:=0
	
	reader := bufio.NewReader(os.Stdin)
	for {
		
		fmt.Print("Enter a task (or type 'exit' to finish): ")
		addTask, _ = reader.ReadString('\n')
		addTask = strings.TrimSpace(addTask)

		if strings.ToLower(addTask) == "exit"{
			break
		}

		task := NewTodo(addTask)
		todos = append(todos, task)


		fmt.Println("Task is : ", todos[i])
		i++
	}

	fmt.Println(todos)

}

