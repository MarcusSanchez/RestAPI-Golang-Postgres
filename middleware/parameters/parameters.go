package parameters

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func replaceHyphens(c *fiber.Ctx) error {
	fmt.Println(c.Request().URI().Password())
	//c.LoadParams()
	for _, param := range c.Route().Params {
		value := c.Params(param)
		value = strings.ReplaceAll(value, "-", " ")
		//c.SetParams(param, value)
	}
	return c.Next()
}

func New() func(*fiber.Ctx) error {
	return replaceHyphens
}
