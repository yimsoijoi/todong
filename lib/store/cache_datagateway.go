package store

import (
	"context"
	"fmt"
	"reflect"

	"github.com/patrickmn/go-cache"
	"github.com/pkg/errors"

	"github.com/artnoi43/todong/datamodel"
	"github.com/artnoi43/todong/lib/cachehelper"
)

var (
	noExp = cache.NoExpiration
)

type cacheDataGateway struct {
	cache cacheDB
}

func NewCacheDataGateway() DataGateway {
	c := cache.New(noExp, noExp)
	return &cacheDataGateway{
		cache: c,
	}
}

func (c *cacheDataGateway) CreateUser(
	ctx context.Context,
	user *datamodel.User,
) error {
	c.cache.Set(user.Username, user, noExp)
	c.cache.Set(user.UUID, user.Username, noExp)
	return nil
}

func (c *cacheDataGateway) CreateTodo(
	ctx context.Context,
	todo *datamodel.Todo,
) error {
	todoKey := cachehelper.KeyFromTodo(todo)
	keyString := todoKey.String()
	v, found := c.cache.Get(keyString)
	if !found {
		new := []*datamodel.Todo{
			todo,
		}
		c.cache.Set(keyString, new, noExp)
	}
	v, found = c.cache.Get(keyString)
	if !found {
		return ErrRecordNotFound
	}
	existing, ok := v.([]*datamodel.Todo)
	if !ok {
		fmt.Printf("%s\n", reflect.TypeOf(v))
		return ErrtypeAssertionFailed
	}
	existing = append(existing, todo)
	c.cache.Set(keyString, existing, noExp)
	return nil
}

func (c *cacheDataGateway) GetUserByUuid(
	ctx context.Context,
	uuid string,
	dst *datamodel.User,
) error {
	username, found := c.cache.Get(uuid)
	if !found {
		return ErrRecordNotFound
	}
	usernameStr := username.(string)
	return c.GetUserByUsername(ctx, usernameStr, dst)
}

func (c *cacheDataGateway) GetUserByUsername(
	ctx context.Context,
	username string,
	dst *datamodel.User,
) error {
	u, found := c.cache.Get(username)
	if !found {
		return ErrRecordNotFound
	}
	user, ok := u.(*datamodel.User)
	if !ok {
		return ErrtypeAssertionFailed
	}
	*dst = *user
	return nil
}

func (c *cacheDataGateway) GetOneTodo(
	ctx context.Context,
	where *datamodel.Todo,
	dst *datamodel.Todo,
) error {
	key := cachehelper.KeyFromTodo(where)
	keyStr := key.String()
	todos, found := c.cache.Get(keyStr)
	if !found {
		return ErrRecordNotFound
	}
	sliceTodo, ok := todos.([]*datamodel.Todo)
	if !ok {
		return ErrtypeAssertionFailed
	}
	for _, todo := range sliceTodo {
		if todo.UUID == where.UUID {
			*dst = *todo
		}
	}
	return nil
}

func (c *cacheDataGateway) GetUserTodos(
	ctx context.Context,
	where *datamodel.Todo,
	dst interface{}, //  *[]datamodel.Todo{}
) error {
	key := cachehelper.KeyFromTodo(where)
	keyStr := key.String()
	t, found := c.cache.Get(keyStr)
	if !found {
		return ErrRecordNotFound
	}
	todos, ok := t.([]*datamodel.Todo)
	if !ok {
		return ErrtypeAssertionFailed
	}
	dst, ok = dst.(*[]datamodel.Todo)
	if !ok {
		return ErrtypeAssertionFailed
	}
	answers := []datamodel.Todo{}
	switch dst := dst.(type) {
	case *[]datamodel.Todo:
		for _, todo := range todos {
			answers = append(answers, *todo)
		}
		*dst = answers
	default:
		return errors.New("interface{} conversion to *[]datamodel.Todo failed: dst is of wrong type")
	}
	return nil
}

func (c *cacheDataGateway) ChangePassword(
	ctx context.Context,
	where *datamodel.User,
	hashedPassword []byte,
) error {
	username, found := c.cache.Get(where.UUID)
	if !found {
		return ErrRecordNotFound
	}
	usernameStr, ok := username.(string)
	if !ok {
		return ErrtypeAssertionFailed
	}
	v, found := c.cache.Get(usernameStr)
	if !found {
		return ErrRecordNotFound
	}
	user, ok := v.(*datamodel.User)
	if !ok {
		return ErrtypeAssertionFailed
	}
	user.Password = hashedPassword
	c.cache.Set(usernameStr, user, noExp)
	return nil
}

func (c *cacheDataGateway) UpdateTodo(
	ctx context.Context,
	where *datamodel.Todo,
	todo *datamodel.Todo,
) error {
	key := cachehelper.KeyFromTodo(where)
	keyStr := key.String()
	v, found := c.cache.Get(keyStr)
	if !found {
		return ErrRecordNotFound
	}
	todos, ok := v.([]*datamodel.Todo)
	if !ok {
		return ErrtypeAssertionFailed
	}
	targetIdx := -1
	for idx, todo := range todos {
		if todo.UUID == where.UUID {
			targetIdx = idx
		}
	}
	if targetIdx == -1 {
		return ErrRecordNotFound
	}
	// [100, 200, 300, 400]
	// [100, 300]
	todos = append(todos[:targetIdx], todos[targetIdx+1:]...)
	todos = append(todos, todo)
	c.cache.Set(keyStr, todos, noExp)
	return nil
}

func (c *cacheDataGateway) DeleteUser(
	ctx context.Context,
	where *datamodel.User,
) error {
	v, found := c.cache.Get(where.UUID)
	if !found {
		return ErrRecordNotFound
	}
	username, ok := v.(string)
	if !ok {
		return ErrtypeAssertionFailed
	}
	key := cachehelper.KeyFromTodo(&datamodel.Todo{
		UserUUID: where.UUID,
	})
	keyStr := key.String()
	c.cache.Delete(where.UUID)
	c.cache.Delete(username)
	c.cache.Delete(keyStr)
	return nil
}

func (c *cacheDataGateway) DeleteTodo(
	ctx context.Context,
	where *datamodel.Todo,
) error {
	key := cachehelper.KeyFromTodo(where)
	keyStr := key.String()
	v, found := c.cache.Get(keyStr)
	if !found {
		return ErrRecordNotFound
	}
	todos := v.([]*datamodel.Todo)
	targetIdx := 0
	for idx, todo := range todos {
		if todo.UUID == where.UUID {
			targetIdx = idx
		}
	}
	todos = append(todos[:targetIdx], todos[targetIdx+1:]...)
	c.cache.Set(keyStr, todos, noExp)
	return nil
}

func (c *cacheDataGateway) Shutdown() {
	c.cache.Flush()
}
