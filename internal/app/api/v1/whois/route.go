package whois

import (
	"strings"

	"checkhost-backend/internal/httpx"

	"github.com/gofiber/fiber/v3"
	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
)

func GET(c fiber.Ctx) error {
	raw, err := whois.Whois(httpx.TargetHost(c))
	if err != nil {
		return httpx.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	parsed, err := whoisparser.Parse(raw)
	if err != nil {
		return c.JSON(fiber.Map{"whois": parseRaw(raw)})
	}

	return c.JSON(fiber.Map{"whois": format(parsed)})
}

func format(parsed whoisparser.WhoisInfo) fiber.Map {
	domain := parsed.Domain
	if domain == nil {
		domain = &whoisparser.Domain{}
	}
	registrar := contactOrEmpty(parsed.Registrar)
	registrant := contactOrEmpty(parsed.Registrant)

	return fiber.Map{
		"domain": fiber.Map{
			"name":        domain.Name,
			"status":      domain.Status,
			"created":     domain.CreatedDate,
			"updated":     domain.UpdatedDate,
			"expires":     domain.ExpirationDate,
			"nameservers": domain.NameServers,
		},
		"registrar": fiber.Map{
			"name":  registrar.Name,
			"email": registrar.Email,
			"phone": registrar.Phone,
		},
		"registrant": fiber.Map{
			"name":    registrant.Name,
			"org":     registrant.Organization,
			"country": registrant.Country,
			"email":   registrant.Email,
		},
	}
}

func contactOrEmpty(contact *whoisparser.Contact) *whoisparser.Contact {
	if contact == nil {
		return &whoisparser.Contact{}
	}
	return contact
}

func parseRaw(raw string) map[string]any {
	result := make(map[string]any)
	for line := range strings.SplitSeq(raw, "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "%") {
			continue
		}
		key, val, ok := strings.Cut(line, ":")
		if !ok {
			continue
		}
		key, val = strings.TrimSpace(key), strings.TrimSpace(val)
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
