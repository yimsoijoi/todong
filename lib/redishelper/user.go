package redishelper

import (
	"encoding/json"
	"fmt"

	"github.com/artnoi43/todong/datamodel"
)

type UserKey struct {
	UUID string
}

type UserVal struct {
	UUID     string `json:"uuid"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (k UserKey) String() string {
	return fmt.Sprintf("username:%s", k.UUID)
}

func (v UserVal) Marshal() string {
	b, _ := json.Marshal(v)
	return string(b)
}

func ValUser(user *datamodel.User) *UserVal {
	return &UserVal{
		UUID:     user.UUID,
		Username: user.Username,
		Password: string(user.Password),
	}
}

func FromUser(user *datamodel.User) (*UserKey, *UserVal) {
	return &UserKey{
		UUID: user.UUID,
	}, ValUser(user)
}

func CopyUser(src UserVal, dst *datamodel.User) {
	*dst = datamodel.User{
		UUID:     src.UUID,
		Username: src.Username,
		Password: []byte(src.Password),
	}
}
