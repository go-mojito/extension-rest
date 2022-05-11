package rest

import "github.com/go-mojito/mojito"

func handlerUpdate[T any](ctrl Controller[T]) func(ctx mojito.Context) error {
	return func(ctx mojito.Context) error {
		existingObj, err := ctrl.ResolveIdentifier(ctx, ctx.Request().Param(routeParamID))
		if err != nil {
			ctx.Response().WriteHeader(500)
			return err
		}

		if existingObj == nil {
			ctx.Response().WriteHeader(404)
			return err
		}

		obj := *existingObj
		if err := parseRestBody(ctx, &obj); err != nil {
			return err
		}

		newObj, err := ctrl.Update(ctx, existingObj, &obj)
		if err != nil {
			return err
		}

		return ctx.JSON(newObj)
	}
}
