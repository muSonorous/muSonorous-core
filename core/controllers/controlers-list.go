// Package controllers Register all your controller instances here
package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/muSonorous-core/core/middlewere"
)

// ControllerEntity represents a Controlelr
type ControllerEntity struct {
	routeGroup  string
	middleweres []middlewere.Middlewere
	routeName   string
	method      string
	handler     gin.HandlerFunc
}

// Array of all the available routes which needs to be registered
var routes = []ControllerEntity{
	ControllerEntity{"", []middlewere.Middlewere{}, "ping", "GET", Ping},
}
