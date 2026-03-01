package handler

import (
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
)

func Whois(c fiber.Ctx) error {
	host := c.Query("host")
	if host == "" {
		host = c.IP()
	}
	if strings.Contains(host, "localhost") || strings.Contains(host, "127.0.0.1") {
		host = "1.1.1.1"
	}
	raw, err := whois.Whois(host)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	parsed, err := whoisparser.Parse(raw)
	if err != nil {
		return c.JSON(fiber.Map{
			"whois": parseRaw(raw),
		})
	}

	return c.JSON(fiber.Map{
		"whois": fiber.Map{
			"domain": fiber.Map{
				"name":        parsed.Domain.Name,
				"status":      parsed.Domain.Status,
				"created":     parsed.Domain.CreatedDate,
				"updated":     parsed.Domain.UpdatedDate,
				"expires":     parsed.Domain.ExpirationDate,
				"nameservers": parsed.Domain.NameServers,
			},
			"registrar": fiber.Map{
				"name":  parsed.Registrar.Name,
				"email": parsed.Registrar.Email,
				"phone": parsed.Registrar.Phone,
			},
			"registrant": fiber.Map{
				"name":    parsed.Registrant.Name,
				"org":     parsed.Registrant.Organization,
				"country": parsed.Registrant.Country,
				"email":   parsed.Registrant.Email,
			},
		},
	})
}

func parseRaw(raw string) map[string]any {
	result := make(map[string]any)
	for _, line := range strings.Split(raw, "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "%") {
			continue
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])
		switch existing := result[key].(type) {
		case nil:
			result[key] = val
		case string:
			result[key] = []string{existing, val}
		case []string:
			result[key] = append(existing, val)
		}
	}
	return result
}
