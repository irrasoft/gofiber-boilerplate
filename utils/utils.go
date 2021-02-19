package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// GetIntParams ...
func GetIntParams(c *fiber.Ctx, name string) (num int) {
	str := c.Params(name)
	if v, err := strconv.Atoi(str); err == nil {
		num = v
	}

	return num
}

// GetParams ...
func GetParams(c *fiber.Ctx) (query map[string]string, order []string, offset int, limit int) {
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")

	// limit: 10 (default is 10)
	if v, err := strconv.Atoi(limitStr); err == nil {
		limit = v
	} else {
		limit = 10
	}
	// offset: 0 (default is 0)
	if v, err := strconv.Atoi(offsetStr); err == nil {
		offset = v
	}

	// order: desc,asc
	if v := c.Query("order"); v != "" {
		v = strings.Replace(v, ":", " ", 1)
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	var q = make(map[string]string)
	if v := c.Query("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				fmt.Println("Error: invalid query key/value pair")
			} else {
				k, v := kv[0], kv[1]
				q[k] = v
			}
		}
		query = q
	}

	fmt.Println("order", order)
	return query, order, offset, limit
}
