package middleware

import (
	"io"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

const countryCodeCyprus = "CY"

func CheckIsCyprus(c *fiber.Ctx) error {
	if c.Method() == fiber.MethodGet {
		return c.Next()
	}
	ipApiClient := http.Client{}
	req, err := http.NewRequest("GET", "https://ipapi.co/country/", nil)
	req.Header.Set("User-Agent", "ipapi.co/#go-v1.5")
	resp, err := ipApiClient.Do(req)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if string(body) != countryCodeCyprus {
		c.Status(fiber.StatusForbidden)
		return nil
	}
	return c.Next()
}
