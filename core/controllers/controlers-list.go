// Register all your controller instances here
package controllers

import "github.com/gin-gonic/gin"

// ControllerEntity represents a Controlelr
type ControllerEntity struct {
	routeGroup     string
	authMiddleWere string
	routeName      string
	method         string
	handler        func(c *gin.Context)
}

var routes = []ControllerEntity{
	ControllerEntity{"", "", "ping", "GET", Ping},
}
