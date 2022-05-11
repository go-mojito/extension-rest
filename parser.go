package rest

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"io/ioutil"

	"github.com/go-mojito/mojito"
)

// parseRestBody will parse any JSON or XML bodies into the given object
// or error if the body is in an unknown format or invalid
func parseRestBody[T any](ctx mojito.Context, obj *T) (err error) {
	bytes, err := ioutil.ReadAll(ctx.Request().GetRequest().Body)
	if err != nil {
		return
	}

	if ctx.Request().HasContentType("application/json") {
		err = json.Unmarshal(bytes, obj)
		return
	}

	if ctx.Request().HasContentType("application/xml") || ctx.Request().HasContentType("text/xml") {
		err = xml.Unmarshal(bytes, obj)
		return
	}
	err = errors.New("Unknown body content format, please use JSON or XML")
	return
}
