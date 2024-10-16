package routes

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"usermicroservices/app"
	mockuserapp "usermicroservices/app/mock"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type RoutesTestSuite struct {
	suite.Suite
}

func TestRoutesTestSuite(t *testing.T) {
	suite.Run(t, new(RoutesTestSuite))
}

func (suite *RoutesTestSuite) TestCreateUser() {

	url := "/CreateUser"
	ctrl := gomock.NewController(suite.T())

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(`
	{
		"ip": "127.0.0.1",
		"device_name": "AppleWebKit/5",
		"user_name": "longvu",
		"alias": "lvu"
	}`)))
	suite.NoError(err)

	httpRecorder := httptest.NewRecorder()

	mockUser := mockuserapp.NewMockUser(ctrl)
	mockUser.EXPECT().
		CreateDBUser(gomock.Any(), gomock.Any()).
		Times(1).
		Return(&app.CreateUserResponse{UserID: 10, UserGUID: uuid.NewString()}, nil)

	routes := Routes{Apps: mockUser}
	serveMux := routes.Register(nil)

	handler, pattern := serveMux.Handler(req)
	suite.Equal(http.MethodPost+" "+url, pattern)

	handler.ServeHTTP(httpRecorder, req)

	suite.Equal(httpRecorder.Code, http.StatusOK)
}

func (suite *RoutesTestSuite) getTestError() error {
	return errors.New("test error")
}

func (suite *RoutesTestSuite) TestCreateUserError() {

	url := "/CreateUser"
	ctrl := gomock.NewController(suite.T())

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(`
	{
		"ip": "127.0.0.1",
		"device_name": "AppleWebKit/5",
		"user_name": "longvu",
		"alias": "lvu"
	}`)))
	suite.NoError(err)

	httpRecorder := httptest.NewRecorder()

	mockUser := mockuserapp.NewMockUser(ctrl)
	mockUser.EXPECT().
		CreateDBUser(gomock.Any(), gomock.Any()).
		Times(1).
		Return(&app.CreateUserResponse{}, suite.getTestError())

	routes := Routes{Apps: mockUser}
	serveMux := routes.Register(nil)

	handler, pattern := serveMux.Handler(req)
	suite.Equal(http.MethodPost+" "+url, pattern)

	handler.ServeHTTP(httpRecorder, req)

	suite.Equal(httpRecorder.Code, http.StatusInternalServerError)
}

func (suite *RoutesTestSuite) TestGetUser() {

	url := "/GetUser"
	ctrl := gomock.NewController(suite.T())

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(`{"user_id":10}`)))
	suite.NoError(err)

	httpRecorder := httptest.NewRecorder()

	returnUser := &app.GetUserResponse{}
	returnUser.UserID = 10
	returnUser.UserGUID = uuid.NewString()
	returnUser.DeviceName = "AppleWebKit/5"
	returnUser.UserName = "longvu"
	returnUser.Alias = "lvu"

	mockUser := mockuserapp.NewMockUser(ctrl)
	mockUser.EXPECT().
		GetDBUser(gomock.Any(), gomock.Any()).
		Times(1).
		Return(returnUser, nil)

	routes := Routes{Apps: mockUser}
	serveMux := routes.Register(nil)

	handler, pattern := serveMux.Handler(req)
	suite.Equal(http.MethodPost+" "+url, pattern)

	handler.ServeHTTP(httpRecorder, req)

	suite.Equal(httpRecorder.Code, http.StatusOK)
}

func (suite *RoutesTestSuite) TestGetUserError() {

	url := "/GetUser"
	ctrl := gomock.NewController(suite.T())

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(`{}`)))
	suite.NoError(err)

	httpRecorder := httptest.NewRecorder()

	mockUser := mockuserapp.NewMockUser(ctrl)
	mockUser.EXPECT().
		GetDBUser(gomock.Any(), gomock.Any()).
		Times(1).
		Return(&app.GetUserResponse{}, suite.getTestError())

	routes := Routes{Apps: mockUser}
	serveMux := routes.Register(nil)

	handler, pattern := serveMux.Handler(req)
	suite.Equal(http.MethodPost+" "+url, pattern)

	handler.ServeHTTP(httpRecorder, req)

	suite.Equal(httpRecorder.Code, http.StatusInternalServerError)
}

func (suite *RoutesTestSuite) TestGetUserByGUID() {
	testGuid := "f838b751-2553-46bc-a19a-cfb3bbac49a5"

	url := "/GetUserByGUID"
	ctrl := gomock.NewController(suite.T())

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(`{"user_guid":"`+testGuid+`"}`)))
	suite.NoError(err)

	httpRecorder := httptest.NewRecorder()

	returnUser := &app.GetUserResponse{}
	returnUser.UserID = 10
	returnUser.UserGUID = testGuid
	returnUser.DeviceName = "AppleWebKit/5"
	returnUser.UserName = "longvu"
	returnUser.Alias = "lvu"

	mockUser := mockuserapp.NewMockUser(ctrl)
	mockUser.EXPECT().
		GetUserByGUID(gomock.Any(), gomock.Any()).
		Times(1).
		Return(returnUser, nil)

	routes := Routes{Apps: mockUser}
	serveMux := routes.Register(nil)

	handler, pattern := serveMux.Handler(req)
	suite.Equal(http.MethodPost+" "+url, pattern)

	handler.ServeHTTP(httpRecorder, req)

	suite.Equal(httpRecorder.Code, http.StatusOK)
}

func (suite *RoutesTestSuite) TestGetUserByGUIDError() {
	url := "/GetUserByGUID"
	ctrl := gomock.NewController(suite.T())

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(`{}`)))
	suite.NoError(err)

	httpRecorder := httptest.NewRecorder()

	mockUser := mockuserapp.NewMockUser(ctrl)
	mockUser.EXPECT().
		GetUserByGUID(gomock.Any(), gomock.Any()).
		Times(1).
		Return(&app.GetUserResponse{}, suite.getTestError())

	routes := Routes{Apps: mockUser}
	serveMux := routes.Register(nil)

	handler, pattern := serveMux.Handler(req)
	suite.Equal(http.MethodPost+" "+url, pattern)

	handler.ServeHTTP(httpRecorder, req)

	suite.Equal(httpRecorder.Code, http.StatusInternalServerError)
}

func (suite *RoutesTestSuite) TestHome() {

	url := "/"
	req, err := http.NewRequest(http.MethodPost, url, nil)
	suite.NoError(err)

	httpRecorder := httptest.NewRecorder()

	routes := NewRoutes()
	serveMux := routes.Register(nil)

	handler, pattern := serveMux.Handler(req)
	suite.Equal(url, pattern)

	handler.ServeHTTP(httpRecorder, req)

	suite.Equal(httpRecorder.Code, http.StatusOK)
}
