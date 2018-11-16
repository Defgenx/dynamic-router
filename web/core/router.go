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

func actionRouterByName(myStruct interface{}, funcName string) func( http.ResponseWriter, *http.Request) {
	val := reflect.ValueOf(myStruct)
	t := val.Type()
	if t.Kind() != reflect.Ptr {
		panic("Error 'myStruct' parameter should be a struct pointer.")
	}
	method := val.MethodByName(funcName)
	if !method.IsValid() {
		panic("Error method name is not valid.")
	}
	return method.Interface().(func(http.ResponseWriter, *http.Request))
}
