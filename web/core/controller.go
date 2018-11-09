package core

import (
	"fmt"
	"net/http"
)

func (ctrl *Controller) HelloWorld(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("Hello World !")
	w.Write([]byte("Hello World !"))
}

func (ctrl *Controller) HelloFrance(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("Hello France !")
	w.Write([]byte("Hello France !"))
}
