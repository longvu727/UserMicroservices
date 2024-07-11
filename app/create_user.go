package app

import (
	"database/sql"
	"encoding/json"

	"github.com/longvu727/FootballSquaresLibs/DB/db"
	"github.com/longvu727/FootballSquaresLibs/util/resources"

	"github.com/google/uuid"
)

type CreateUserParams struct {
	IP         string `json:"ip"`
	DeviceName string `json:"device_name"`
	UserName   string `json:"user_name"`
	Alias      string `json:"alias"`
}

type CreateUserResponse struct {
	UserID       int64  `json:"user_id"`
	UserGUID     string `json:"user_guid"`
	ErrorMessage string `json:"error_message"`
}

func (response CreateUserResponse) ToJson() []byte {
	jsonStr, _ := json.Marshal(response)
	return jsonStr
}

func (userApp *UserApp) CreateDBUser(createUserParams CreateUserParams, resources *resources.Resources) (*CreateUserResponse, error) {
	var createUserResponse CreateUserResponse

	userGuid := (uuid.New()).String()

	userID, err := resources.DB.CreateUser(resources.Context, db.CreateUserParams{
		UserGuid:   userGuid,
		Ip:         sql.NullString{String: createUserParams.IP, Valid: true},
		DeviceName: sql.NullString{String: createUserParams.DeviceName, Valid: true},
		UserName:   sql.NullString{String: createUserParams.UserName, Valid: true},
		Alias:      sql.NullString{String: createUserParams.Alias, Valid: true},
	})
	if err != nil {
		return &createUserResponse, err
	}

	createUserResponse.UserGUID = userGuid
	createUserResponse.UserID = userID

	return &createUserResponse, nil
}
