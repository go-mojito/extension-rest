package rest

import (
	"github.com/go-mojito/mojito"
	"github.com/go-mojito/mojito/pkg/router"
)

const (
	// routeParamID defines the name of the route parameter used for a resource's
	// unique identifier
	routeParamID = "id"
)

// ServeController will register a REST controller on a given router with the given path
// and apply a special middleware to provide injection features (not dependency injection!)
func ServeController[T any](path string, r mojito.Router, ctrl Controller[T]) error {
	return r.WithGroup(path, func(router router.Group) {
		router.DELETE("/:"+routeParamID, handlerDelete(ctrl))

		router.GET("", handlerList(ctrl))
		router.GET("/", handlerList(ctrl))

		router.GET("/:"+routeParamID, handlerGet(ctrl))

		router.POST("", handlerCreate(ctrl))
		router.POST("/", handlerCreate(ctrl))

		router.PUT("/:"+routeParamID, handlerUpdate(ctrl))
	})
}
