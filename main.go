package main

import (
	"net/http"

	"github.com/clintcolding/gogettingstarted/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
