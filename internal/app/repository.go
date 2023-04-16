package app

import (
	"context"
	"errors"
)

var NotFoundErr = errors.New("Bad luck, nothing found")
var AlreadyExistErr = errors.New("Bad luck, it's already exist")
var HavntUpd = errors.New("Bad luck, have not updates")
var WrongLink = errors.New("Bad luck, wrong link")
var UserDoesNotExist = errors.New("Bad luck, user does not exist")

type Repository interface {
	CreateUser(ctx context.Context) error
	CreateSource(ctx context.Context) error
	DeleteSource(ctx context.Context) error
	GetSrcsByChat(ctx context.Context) error
	GetRSSBySource(ctx context.Context) error
}
