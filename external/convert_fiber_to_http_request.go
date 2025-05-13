package external

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// ConvertFiberToHTTPRequest converts fiber.Ctx to *http.Request
func ConvertFiberToHTTPRequest(c *fiber.Ctx) (*http.Request, error) {
	uri := c.OriginalURL()
	method := string(c.Method())

	r, err := http.NewRequest(method, uri, nil)
	if err != nil {
		return nil, err
	}

	// Salin header
	c.Request().Header.VisitAll(func(key, value []byte) {
		r.Header.Add(string(key), string(value))
	})

	return r, nil
}
