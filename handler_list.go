package rest

import "github.com/go-mojito/mojito"

func handlerList[T any](ctrl Controller[T]) func(ctx mojito.Context) error {
	return func(ctx mojito.Context) error {
		newObj, err := ctrl.List(ctx, parseRestListFilter(ctx))
		if err != nil {
			return err
		}
		return ctx.JSON(newObj)
	}
}

func parseRestListFilter(ctx mojito.Context) ListFilter {
	return ListFilter{}
}
