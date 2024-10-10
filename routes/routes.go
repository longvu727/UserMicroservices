package routes

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"usermicroservices/app"

	"github.com/longvu727/FootballSquaresLibs/util/resources"
)

type RoutesInterface interface {
	Register(resources *resources.Resources) *http.ServeMux
}

type Routes struct {
	Apps app.User
}

type Handler = func(writer http.ResponseWriter, request *http.Request, resources *resources.Resources)

func NewRoutes() RoutesInterface {
	return &Routes{
		Apps: app.NewUserApp(),
	}
}

func (routes *Routes) Register(resources *resources.Resources) *http.ServeMux {
	log.Println("Registering routes")
	mux := http.NewServeMux()

	routesHandlersMap := map[string]Handler{
		"/":                                 routes.home,
		http.MethodPost + " /CreateUser":    routes.createUser,
		http.MethodPost + " /GetUser":       routes.getUser,
		http.MethodPost + " /GetUserByGUID": routes.getUserByGUID,
	}

	for route, handler := range routesHandlersMap {
		mux.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
			handler(w, r, resources)
		})
	}

	return mux
}

func (routes *Routes) home(writer http.ResponseWriter, _ *http.Request, resources *resources.Resources) {
	fmt.Fprintf(writer, "{\"Acknowledged\": true}")
}

func (routes *Routes) createUser(writer http.ResponseWriter, request *http.Request, resources *resources.Resources) {
	log.Printf("Received request for %s\n", request.URL.Path)

	writer.Header().Set("Content-Type", "application/json")

	var createUserParams app.CreateUserParams
	json.NewDecoder(request.Body).Decode(&createUserParams)

	createUserResponse, err := routes.Apps.CreateDBUser(createUserParams, resources)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		createUserResponse.ErrorMessage = `Unable to create user` + err.Error()
		writer.Write(createUserResponse.ToJson())
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(createUserResponse.ToJson())
}

func (routes *Routes) getUser(writer http.ResponseWriter, request *http.Request, resources *resources.Resources) {
	log.Printf("Received request for %s\n", request.URL.Path)

	writer.Header().Set("Content-Type", "application/json")

	var getUserParams app.GetUserParams
	json.NewDecoder(request.Body).Decode(&getUserParams)

	getUserResponse, err := routes.Apps.GetDBUser(getUserParams, resources)

	if err != nil && err == sql.ErrNoRows {
		getUserResponse.ErrorMessage = `User not found`
	} else if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		getUserResponse.ErrorMessage = `Unable to get user` + err.Error()
		writer.Write(getUserResponse.ToJson())
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(getUserResponse.ToJson())
}

func (routes *Routes) getUserByGUID(writer http.ResponseWriter, request *http.Request, resources *resources.Resources) {
	log.Printf("Received request for %s\n", request.URL.Path)

	writer.Header().Set("Content-Type", "application/json")

	var getUserByGUIDParams app.GetUserByGUIDParams
	json.NewDecoder(request.Body).Decode(&getUserByGUIDParams)

	getUserResponse, err := routes.Apps.GetUserByGUID(getUserByGUIDParams, resources)

	if err != nil && err == sql.ErrNoRows {
		getUserResponse.ErrorMessage = `User not found`
	} else if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		getUserResponse.ErrorMessage = `Unable to get user` + err.Error()
		writer.Write(getUserResponse.ToJson())
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(getUserResponse.ToJson())
}
