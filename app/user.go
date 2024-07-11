package app

import "github.com/longvu727/FootballSquaresLibs/util/resources"

type User interface {
	GetDBUser(getUserParams GetUserParams, resources *resources.Resources) (*GetUserResponse, error)
	CreateDBUser(createUserParams CreateUserParams, resources *resources.Resources) (*CreateUserResponse, error)
}

type UserApp struct{}

func NewUserApp() User {
	return &UserApp{}
}
