package app

import (
	"context"
	"database/sql"
	"math/rand"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	db "github.com/longvu727/FootballSquaresLibs/DB/db"
	mockdb "github.com/longvu727/FootballSquaresLibs/DB/db/mock"
	"github.com/longvu727/FootballSquaresLibs/util"
	"github.com/longvu727/FootballSquaresLibs/util/resources"
	"github.com/stretchr/testify/suite"
)

type GetUserTestSuite struct {
	suite.Suite
}

func (suite *GetUserTestSuite) SetupTest() {
}

func (suite *GetUserTestSuite) TestGetUser() {
	randomUser := randomUser()

	ctrl := gomock.NewController(suite.T())
	defer ctrl.Finish()

	mockMySQL := mockdb.NewMockMySQL(ctrl)

	mockMySQL.EXPECT().
		GetUser(gomock.Any(), gomock.Eq(randomUser.UserID)).
		Times(1).
		Return(randomUser, nil)

	config, err := util.LoadConfig("../env", "app", "env")
	suite.NoError(err)

	resources := resources.NewResources(config, mockMySQL, context.Background())

	getUserParams := GetUserParams{UserID: int(randomUser.UserID)}
	user, err := NewUserApp().GetDBUser(getUserParams, resources)
	suite.NoError(err)

	suite.Equal(randomUser.UserID, int32(user.UserID))
	suite.Equal(randomUser.UserGuid, user.UserGUID)
	suite.Equal(randomUser.Ip.String, user.IP)
	suite.Equal(randomUser.DeviceName.String, user.DeviceName)
	suite.Equal(randomUser.UserName.String, user.UserName)
	suite.Equal(randomUser.DeviceName.String, user.DeviceName)

}

func (suite *GetUserTestSuite) TestGetUserByGUID() {
	randomUser := randomUserByGUID()

	ctrl := gomock.NewController(suite.T())
	defer ctrl.Finish()

	mockMySQL := mockdb.NewMockMySQL(ctrl)

	mockMySQL.EXPECT().
		GetUserByGUID(gomock.Any(), gomock.Eq(randomUser.UserGuid)).
		Times(1).
		Return(randomUser, nil)

	config, err := util.LoadConfig("../env", "app", "env")
	suite.NoError(err)

	resources := resources.NewResources(config, mockMySQL, context.Background())

	getUserParams := GetUserByGUIDParams{UserGUID: randomUser.UserGuid}
	user, err := NewUserApp().GetUserByGUID(getUserParams, resources)
	suite.NoError(err)

	suite.Equal(randomUser.UserID, int32(user.UserID))
	suite.Equal(randomUser.UserGuid, user.UserGUID)
	suite.Equal(randomUser.Ip.String, user.IP)
	suite.Equal(randomUser.DeviceName.String, user.DeviceName)
	suite.Equal(randomUser.UserName.String, user.UserName)
	suite.Equal(randomUser.DeviceName.String, user.DeviceName)
}

func randomUser() db.GetUserRow {
	return db.GetUserRow{
		UserID:     rand.Int31n(1000),
		UserGuid:   uuid.NewString(),
		Ip:         sql.NullString{String: "", Valid: true},
		DeviceName: sql.NullString{String: "", Valid: true},
		UserName:   sql.NullString{String: "", Valid: true},
		Alias:      sql.NullString{String: "", Valid: true},
	}
}

func randomUserByGUID() db.GetUserByGUIDRow {
	return db.GetUserByGUIDRow{
		UserID:     rand.Int31n(1000),
		UserGuid:   uuid.NewString(),
		Ip:         sql.NullString{String: "", Valid: true},
		DeviceName: sql.NullString{String: "", Valid: true},
		UserName:   sql.NullString{String: "", Valid: true},
		Alias:      sql.NullString{String: "", Valid: true},
	}
}

func TestGetUserTestSuite(t *testing.T) {
	suite.Run(t, new(GetUserTestSuite))
}
