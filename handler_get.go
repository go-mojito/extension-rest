package rest

import "github.com/go-mojito/mojito"

func handlerGet[T any](ctrl Controller[T]) func(ctx mojito.Context) error {
	return func(ctx mojito.Context) error {
		obj, err := ctrl.ResolveIdentifier(ctx, ctx.Request().Param(routeParamID))
		if err != nil {
			ctx.Response().WriteHeader(500)
			return err
		}

		// The resource identifier did not fail, but also did not return a resource
		// The requested resource could not be found.
		if obj == nil {
			ctx.Response().WriteHeader(404)
			return err
		}

		newObj, err := ctrl.Get(ctx, obj)
		if err != nil {
			return err
		}
		return ctx.JSON(newObj)
	}
}
