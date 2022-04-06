package store

import (
	"context"
	"log"

	"github.com/yimsoijoi/todong/datamodel"
	"gorm.io/gorm"
)

type gormDataGateway struct {
	store GormStore
}

func NewGormDataGateway(g GormStore) DataGateway {
	return &gormDataGateway{
		store: g,
	}
}

func (g *gormDataGateway) GetUserByUsername(
	ctx context.Context,
	username string,
	dst *datamodel.User,
) error {
	where := &datamodel.User{
		Username: username,
	}
	return g.store.First(ctx, where, dst)
}

func (g *gormDataGateway) GetUserByUuid(
	ctx context.Context,
	uuid string,
	dst *datamodel.User,
) error {
	where := &datamodel.User{
		UUID: uuid,
	}
	return g.store.First(ctx, where, dst)
}

func (g *gormDataGateway) DeleteUser(
	ctx context.Context,
	where *datamodel.User,
) error {
	return g.store.Delete(ctx, where, &datamodel.User{})
}

func (g *gormDataGateway) DeleteTodo(
	ctx context.Context,
	where *datamodel.Todo,
) error {
	return g.store.Delete(ctx, where, &datamodel.Todo{})
}

func (g *gormDataGateway) CreateUser(
	ctx context.Context,
	user *datamodel.User,
) error {
	return g.store.Create(ctx, user)
}

func (g *gormDataGateway) CreateTodo(
	ctx context.Context,
	todo *datamodel.Todo,
) error {
	return g.store.Create(ctx, todo)
}

func (g *gormDataGateway) GetOneTodo(
	ctx context.Context,
	where *datamodel.Todo,
	dst *datamodel.Todo,
) error {
	return g.store.First(ctx, where, dst)
}

func (g *gormDataGateway) GetUserTodos(
	ctx context.Context,
	where *datamodel.Todo,
	dst interface{},
) error {
	return g.store.Find(ctx, where, dst)
}

func (g *gormDataGateway) UpdateTodo(
	ctx context.Context,
	where *datamodel.Todo,
	todo *datamodel.Todo,
) error {
	return g.store.Updates(ctx, where, todo)
}

func (g *gormDataGateway) ChangePassword(
	ctx context.Context,
	where *datamodel.User,
	hashedPassword []byte,
) error {
	return g.store.Updates(ctx, where, &datamodel.User{
		Password: hashedPassword,
	})
}

func (g *gormDataGateway) Shutdown() {
	switch store := g.store.(type) {
	case *gormStore:
		switch db := store.db.(type) {
		case *gorm.DB:
			sqlDB, err := db.DB()
			if err != nil {
				log.Println("failed to get sqlDB")
			}
			sqlDB.Close()
		}
	}
}
