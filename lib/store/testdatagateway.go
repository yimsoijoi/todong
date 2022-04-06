package store

import (
	"context"

	"github.com/yimsoijoi/todong/datamodel"
)

type testDataGateway struct{}

func NewTestDataGateway() DataGateway { return &testDataGateway{} }

func (g *testDataGateway) DeleteUser(
	ctx context.Context,
	where *datamodel.User,
) error {
	return nil
}

func (g *testDataGateway) DeleteTodo(
	ctx context.Context,
	where *datamodel.Todo,
) error {
	return nil
}

func (g *testDataGateway) CreateUser(
	ctx context.Context,
	user *datamodel.User,
) error {
	return nil
}

func (g *testDataGateway) CreateTodo(
	ctx context.Context,
	todo *datamodel.Todo,
) error {
	return nil
}

// This function is called in handler.Login and handler.Register
// So calling this method will fail either of the tests.
func (g *testDataGateway) GetUserByUuid(
	ctx context.Context,
	uuid string,
	dst *datamodel.User,
) error {
	return nil
}

func (g *testDataGateway) GetUserByUsername(
	ctx context.Context,
	username string,
	dst *datamodel.User,
) error {
	return nil
}

func (g *testDataGateway) GetOneTodo(
	ctx context.Context,
	where *datamodel.Todo,
	dst *datamodel.Todo,
) error {
	return nil
}

func (g *testDataGateway) GetUserTodos(
	ctx context.Context,
	where *datamodel.Todo,
	dst interface{},
) error {
	return nil
}

func (g *testDataGateway) UpdateTodo(
	ctx context.Context,
	where *datamodel.Todo,
	todo *datamodel.Todo,
) error {
	return nil
}

func (g *testDataGateway) ChangePassword(
	ctx context.Context,
	where *datamodel.User,
	hashedPassword []byte,
) error {
	return nil
}

func (g *testDataGateway) Shutdown() {}
