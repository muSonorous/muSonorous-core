package middlewere

import "github.com/gin-gonic/gin"

// Middlewere represents Middlewere entity..
type Middlewere struct {
	name    string
	handler gin.HandlerFunc
}
