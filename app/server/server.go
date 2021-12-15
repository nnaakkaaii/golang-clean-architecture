package server

import (
	"clean_architecture/app/controller"
	"fmt"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

type UserServer struct {
	e          *echo.Echo
	controller *controller.UserControllerInterface
}

func NewUserServer(controller *controller.UserControllerInterface) *UserServer {
	server := &UserServer{
		e:          echo.New(),
		controller: controller,
	}
	for _, route := range handleRoutes(controller) {
		switch route.method {
		case http.MethodGet:
			server.e.GET(route.path, route.handler)
		case http.MethodPost:
			server.e.POST(route.path, route.handler)
		case http.MethodPut:
			server.e.PUT(route.path, route.handler)
		case http.MethodDelete:
			server.e.DELETE(route.path, route.handler)
		default:
			panic("unknown route")
		}
	}
	return server
}

func (s *UserServer) Start(port int) {
	log.Fatal(s.e.Start(fmt.Sprintf(":%d", port)))
}

// route は method (リクエストメソッド), path, handler (実行するHandler)の情報を含む
type route struct {
	method  string
	path    string
	handler echo.HandlerFunc
}

func handleRoutes(controller *controller.UserControllerInterface) []route {
	return []route{
		{
			method:  http.MethodGet,
			path:    "/users",
			handler: (*controller).ListUsers,
		},
		{
			method:  http.MethodPost,
			path:    "/users",
			handler: (*controller).CreateUser,
		},
		{
			method:  http.MethodGet,
			path:    "/users/:id",
			handler: (*controller).RetrieveUser,
		},
		{
			method:  http.MethodPut,
			path:    "/users/:id",
			handler: (*controller).UpdateUser,
		},
		{
			method:  http.MethodDelete,
			path:    "/users/:id",
			handler: (*controller).DeleteUser,
		},
	}
}
