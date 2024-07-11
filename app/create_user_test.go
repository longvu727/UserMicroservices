package app

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	mockdb "github.com/longvu727/FootballSquaresLibs/DB/db/mock"
	"github.com/longvu727/FootballSquaresLibs/util"
	"github.com/longvu727/FootballSquaresLibs/util/resources"
	"github.com/stretchr/testify/suite"
)

type CreateUserTestSuite struct {
	suite.Suite
}

func (suite *CreateUserTestSuite) SetupTest() {
}

func (suite *CreateUserTestSuite) TestCreateUser() {
	randomUser := randomUser()

	ctrl := gomock.NewController(suite.T())
	defer ctrl.Finish()

	mockMySQL := mockdb.NewMockMySQL(ctrl)

	mockMySQL.EXPECT().
		CreateUser(gomock.Any(), gomock.Any()).
		Times(1).
		Return(int64(randomUser.UserID), nil)

	config, err := util.LoadConfig("../env", "app", "env")
	suite.NoError(err)

	resources := resources.NewResources(config, mockMySQL, context.Background())

	createSquareParams := CreateUserParams{
		IP:         randomUser.Ip.String,
		DeviceName: randomUser.DeviceName.String,
		UserName:   randomUser.UserName.String,
		Alias:      randomUser.Alias.String,
	}
	user, err := NewUserApp().CreateDBUser(createSquareParams, resources)
	suite.NoError(err)

	suite.Equal(randomUser.UserID, int32(user.UserID))
}

func TestCreateUserTestSuite(t *testing.T) {
	suite.Run(t, new(CreateUserTestSuite))
}
