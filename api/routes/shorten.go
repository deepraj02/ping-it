package routes

import (
	"fmt"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/deepraj02/pingit/database"
	"github.com/deepraj02/pingit/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
}

type response struct {
	URL             string        `json:"url"`
	CustomShort     string        `json:"short"`
	Expiry          time.Duration `json:"expiry"`
	XRateRemaining  int           `json:"rate_limit"`
	XRateLimitReset time.Duration `json:"rate-limit-reset"`
}

func ShortenURL(c *fiber.Ctx) error {

	body := new(request)
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// Rate Limiting

	r2 := database.CreateClient(1)
	defer r2.Close()
	val, err := r2.Get(database.Ctx, c.IP()).Result()
	if err == redis.Nil {
		_ = r2.Set(database.Ctx, c.IP(), 10, 30*60*time.Second).Err()
		fmt.Print(val)
	} else {
		val, _ := r2.Get(database.Ctx, c.IP()).Result()
		valInt, _ := strconv.Atoi(val)
		if valInt <= 0 {
			limit, _ := r2.TTL(database.Ctx, c.IP()).Result()
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error": "Rate Limit Exceeded", "retry-after": limit})
		}

	}
	// Checking if the URL is Valid
	if !govalidator.IsURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid URL"})
	}

	// Check for Domain Errors
	if !helpers.RemoveDomainError(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Domain"})
	}

	// Enforce SSL, HTTPS
	body.URL = helpers.EnforceHTTP(body.URL)

	var id string
	fmt.Println(id)
	if body.CustomShort == "" {
		id = uuid.New().String()[:6]

	} else {
		id = body.CustomShort
	}
	r := database.CreateClient(0)
	defer r.Close()

	val, _ = r.Get(database.Ctx, id).Result()
	if val != "" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Short URL already exists"})
	}
	if body.Expiry == 0 {
		body.Expiry = 24
	}
	e := r.Set(database.Ctx, id, body.URL, body.Expiry*3600*time.Second).Err()
	if e != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	r2.Decr(database.Ctx, c.IP())

	return nil
}
