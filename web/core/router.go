package core

import (
	"github.com/defgenx/dynamic-router/web/config"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"reflect"
)

type Router struct {
	*mux.Router
}

type Controller struct {}

func NewRouter() *Router {
	return &Router{
		mux.NewRouter()}
}

func (r *Router) SetRoutes(routeList []config.Route) {
	for i := 0; i < len(routeList); i++ {
		log.Println("Route Loaded: " + routeList[i].Method)
		r.Router.HandleFunc(routeList[i].Url, actionRouterByName(new(Controller), routeList[i].Method)).Name(routeList[i].Name).Methods(routeList[i].HttpMethod)
	}
}

func actionRouterByName(myStruct interface{}, funcName string) func(http.ResponseWriter, *http.Request) {
	myStructValue := reflect.ValueOf(myStruct)
	m := myStructValue.MethodByName(funcName)
	return m.Interface().(func(http.ResponseWriter, *http.Request))
}
