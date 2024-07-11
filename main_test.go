package main

import (
	"context"
	"fmt"
	"net/http"
	mockroutes "usermicroservices/routes/mock"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/longvu727/FootballSquaresLibs/util"
	"github.com/longvu727/FootballSquaresLibs/util/resources"
	"github.com/stretchr/testify/suite"
)

type MainTestSuite struct {
	suite.Suite
}

func (suite *MainTestSuite) TestStart() {

	config, err := util.LoadConfig("./env", "app", "env")
	suite.NoError(err)
	resources := resources.NewResources(config, nil, context.Background())

	ctrl := gomock.NewController(suite.T())

	mockRoutes := mockroutes.NewMockRoutesInterface(ctrl)
	mockRoutes.EXPECT().
		Register(gomock.Any()).
		Times(1).
		Return(nil)

	httpServer := &http.Server{}

	api := &api{routes: mockRoutes, address: "::::", server: httpServer}
	err = api.start(resources)
	suite.Error(err)
}

func (suite *MainTestSuite) TestGetResources() {
	type testCase struct {
		testName   string
		path       string
		configName string
		configType string
	}
	testCases := struct {
		successes []testCase
		failures  []testCase
	}{
		successes: []testCase{
			{testName: "Test Success", path: "./env", configName: "app", configType: "env"},
		},
		failures: []testCase{
			{testName: "Test Fail DB connect", path: "./env", configName: "app.test", configType: "env"},
			{testName: "Test Fail Config", path: "./env", configName: "app.env", configType: "test"},
		},
	}

	for _, test := range testCases.successes {
		fmt.Println(test.testName)

		_, err := getResourcesFromConfigFile(test.path, test.configName, test.configType)
		suite.NoError(err)
	}

	for _, test := range testCases.failures {
		fmt.Println(test.testName)

		_, err := getResourcesFromConfigFile(test.path, test.configName, test.configType)
		suite.Error(err)
	}
}

func TestMainTestSuite(t *testing.T) {
	suite.Run(t, new(MainTestSuite))
}
