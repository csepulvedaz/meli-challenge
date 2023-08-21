package middlewares

import (
	"crypto/sha256"
	"crypto/subtle"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func validateAPIKey(_ *fiber.Ctx, key string) (bool, error) {
	hashedAPIKey := sha256.Sum256([]byte(os.Getenv("API_KEY")))
	hashedKey := sha256.Sum256([]byte(key))

	if subtle.ConstantTimeCompare(hashedAPIKey[:], hashedKey[:]) == 1 {
		return true, nil
	}
	return false, keyauth.ErrMissingOrMalformedAPIKey
}

// FiberMiddleware provide Fiber's built-in middlewares.
// See: https://docs.gofiber.io/category/-middleware/
func FiberMiddleware(a *fiber.App) {
	a.Use(
		// Add CORS to each route.
		cors.New(),
		// Add simple logger.
		logger.New(),
		// Add recover middleware.
		recover.New(),
		// Add keyauth middleware.
		keyauth.New(keyauth.Config{
			Validator: validateAPIKey,
		}),
	)
}
