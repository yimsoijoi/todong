package store

// NOTE: might not be needed - we can use gomock to mock store.Store instead

import (
	"context"

	"github.com/yimsoijoi/todong/datamodel"
	"github.com/yimsoijoi/todong/test"
)

type mockStore struct{}

func NewMockStore() GormStore {
	return &mockStore{}
}

func (m *mockStore) First(ctx context.Context, where interface{}, item interface{}) error {
	switch item := item.(type) {
	case []*datamodel.Todo:
		switch where := where.(type) {
		case *datamodel.Todo:
			item[0] = &datamodel.Todo{
				UUID:     where.UUID,
				UserUUID: where.UserUUID,
			}
		}
	case *datamodel.Todo:
		switch where := where.(type) {
		case *datamodel.Todo:
			*item = datamodel.Todo{
				UUID:     where.UUID,
				UserUUID: where.UserUUID,
			}
		}
	case *datamodel.User:
		switch where := where.(type) {
		case *datamodel.User:
			*item = datamodel.User{
				UUID:     test.JwtIss,
				Username: where.Username,
				Password: test.HashedPW,
			}
		}
	}
	return nil
}
func (m *mockStore) Find(ctx context.Context, where interface{}, item interface{}) error {
	switch where := where.(type) {
	case *datamodel.Todo:
		switch item := item.(type) {
		case *datamodel.Todo:
			*item = datamodel.Todo{
				UUID:     where.UUID,
				UserUUID: where.UserUUID,
			}
		case []*datamodel.Todo:
			item[0] = &datamodel.Todo{
				UUID:     where.UUID,
				UserUUID: where.UUID,
			}
		}
	case *datamodel.User:
		switch item := item.(type) {
		case *datamodel.User:
			*item = datamodel.User{
				UUID:     test.JwtIss,
				Username: where.Username,
				Password: test.HashedPW,
			}
		}
	}
	return nil
}
func (m *mockStore) Create(ctx context.Context, item interface{}) error {
	return nil
}
func (m *mockStore) Updates(ctx context.Context, where interface{}, item interface{}) error {
	return nil
}
func (m *mockStore) Delete(ctx context.Context, where interface{}, item interface{}) error {
	return nil
}
