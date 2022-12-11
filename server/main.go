package main

import (
	"server/config"
	"server/controller"
)

func main() {

	// load environment, config logging
	config.Init()

	// initialize database connection
	db := config.InitDB()

	// initialize Router
	r := config.InitRouter()

	// initialize Controller
	c := controller.InitController(db)

	// Add routes
	v1 := r.Group("/v1")
	{
		v1.GET("request", c.ReadRequests)
		v1.GET("request/:id", c.ReadRequest)
		v1.POST("request", c.CreateRequest)
		v1.PUT("request/:id", c.UpdateRequest)
		v1.DELETE("request/:id", c.DeleteRequest)
	}

	r.Run()

}
