package store

import (
	"context"

	"github.com/artnoi43/todong/datamodel"
)

// DataGateway abstracts high-level storage of datamode.User and datamodel.Todo
type DataGateway interface {
	// Shutdown is needed for graceful shutdown
	Shutdown()

	CreateUser(
		ctx context.Context,
		user *datamodel.User,
	) error
	CreateTodo(
		ctx context.Context,
		todo *datamodel.Todo,
	) error
	GetUserByUuid(
		ctx context.Context,
		uuid string,
		dst *datamodel.User,
	) error
	GetUserByUsername(
		ctx context.Context,
		username string,
		dst *datamodel.User,
	) error
	GetOneTodo(
		ctx context.Context,
		where *datamodel.Todo,
		dst *datamodel.Todo,
	) error
	GetUserTodos(
		ctx context.Context,
		where *datamodel.Todo,
		dst interface{},
	) error
	ChangePassword(
		ctx context.Context,
		where *datamodel.User,
		hashedPassword []byte,
	) error
	UpdateTodo(
		ctx context.Context,
		where *datamodel.Todo,
		todo *datamodel.Todo,
	) error
	DeleteUser(
		ctx context.Context,
		where *datamodel.User,
	) error
	DeleteTodo(
		ctx context.Context,
		where *datamodel.Todo,
	) error
}
