package main

import (
	"net/http"

	"github.com/mikhailtse/go-simple-webservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
