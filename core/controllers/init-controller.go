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

	restController.initRoutes()
	// TODO Register all the routes here
	return restController
}

// RegisterMiddlewere register  all middleweres
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

// Internal method that Registers all the controller Handlers dynamicaly from routes array
func (c *RestController) initRoutes() {

	// Iterate over all the route and regieter to route group
	for _, route := range routes {
		c.engine.Handle(route.method, route.routeName, route.handler)
	}
}

// PrintAllRoutes method to print all registered Routes
func (c *RestController) PrintAllRoutes() {
	for groupName, group := range c.routeGroups {
		fmt.Println(groupName, "==>", group)
	}
}

// Start start the controller server
func (c *RestController) Start(port string) {
	c.engine.Run(fmt.Sprintf(":%s", port))
}

// Singleton Rest Controller instance
var restController *RestController = nil
var once sync.Once
