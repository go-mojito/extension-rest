package rest

import "github.com/go-mojito/mojito"

func handlerCreate[T any](ctrl Controller[T]) func(ctx mojito.Context) error {
	return func(ctx mojito.Context) error {
		var obj T
		if err := parseRestBody(ctx, &obj); err != nil {
			return err
		}

		newObj, err := ctrl.Create(ctx, &obj)
		if err != nil {
			return err
		}

		return ctx.JSON(newObj)
	}
}
