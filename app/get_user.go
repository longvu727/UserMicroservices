package app

import (
	"encoding/json"

	usermicroservices "github.com/longvu727/FootballSquaresLibs/services/user_microservices"
	"github.com/longvu727/FootballSquaresLibs/util/resources"
)

type GetUserParams struct {
	UserID int `json:"user_id"`
}

type GetUserResponse struct {
	usermicroservices.User
	ErrorMessage string `json:"error_message"`
}

func (response GetUserResponse) ToJson() []byte {
	jsonStr, _ := json.Marshal(response)
	return jsonStr
}

func (userApp *UserApp) GetDBUser(getUserParams GetUserParams, resources *resources.Resources) (*GetUserResponse, error) {
	var getUserResponse GetUserResponse

	userRow, err := resources.DB.GetUser(resources.Context, int32(getUserParams.UserID))
	if err != nil {
		return &getUserResponse, err
	}

	getUserResponse.UserID = int(userRow.UserID)
	getUserResponse.UserGUID = userRow.UserGuid
	getUserResponse.IP = userRow.Ip.String
	getUserResponse.DeviceName = userRow.DeviceName.String
	getUserResponse.UserName = userRow.UserName.String
	getUserResponse.Alias = userRow.Alias.String

	return &getUserResponse, nil
}
