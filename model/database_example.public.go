package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	tableName struct{} `pg:"public.users"`
	Id        uuid.UUID
	Nama      string
	Email     string
	Username  string
	Password  string
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
