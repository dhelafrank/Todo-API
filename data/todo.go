package data

import (
	"errors"
	"todo-api/types"
)

var Todos = []types.Todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Code Go", Completed: false},
	{ID: "3", Item: "Read Book", Completed: false},
}

func FetchTodoById(id string) (*types.Todo, error) {
	for i, t := range Todos { //Iterating over all todos
		if t.ID == id {
			return &Todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}
