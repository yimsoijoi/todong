package cachehelper

import (
	"fmt"

	"github.com/yimsoijoi/todong/datamodel"
)

type TodoKey struct {
	UserUUID string
}

func (k TodoKey) String() string {
	return fmt.Sprintf("todo:%s", k.UserUUID)
}

func KeyFromTodo(todo *datamodel.Todo) TodoKey {
	return TodoKey{
		UserUUID: todo.UserUUID,
	}
}
