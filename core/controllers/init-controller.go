package controllers

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
)

// RestController rest controller
type RestController struct {
	engine      *gin.Engine
	routeGroups map[string]*gin.RouterGroup
}

// New returns a singleton RestController instance
func New() *RestController {
	// This piece of code will be executed only once
	once.Do(func() {
		restController = &RestController{
			gin.Default(),
			map[string]*gin.RouterGroup{},
		}
	})

	// TODO Register all the routes here
	return restController
}

//
func RegisterMiddlewere() (registeredMiddleweres []string) {
	// TODO: Implement this method

	return registeredMiddleweres
}

// !--
// All member methods of RestController

// AddGroup adds a RouterGroup to the current
func (c *RestController) AddGroup(groupName string) *gin.RouterGroup {
	c.routeGroups[groupName] = c.engine.Group(groupName)
	return c.routeGroups[groupName]
}

// PrintAllRoutes method to print all registered Routes
func (c *RestController) PrintAllRoutes() {
	for groupName, group := range c.routeGroups {
		fmt.Println(groupName, "==>", group)
	}
}

// Singleton Rest Controller instance
var restController *RestController = nil
var once sync.Once
