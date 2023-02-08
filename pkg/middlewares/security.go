package middlewares

import "github.com/gofiber/fiber/v2"

func Security(c *fiber.Ctx) error {
	c.Set(fiber.HeaderXXSSProtection, "1; mode=block")
	c.Set(fiber.HeaderXContentTypeOptions, "nosniff")
	c.Set(fiber.HeaderXDownloadOptions, "noopen")
	c.Set(fiber.HeaderStrictTransportSecurity, "max-age=5184000")
	c.Set(fiber.HeaderXFrameOptions, "DENY")
	c.Set(fiber.HeaderXDNSPrefetchControl, "off")
	c.Set(fiber.HeaderAccessControlAllowMethods, "GET, POST, PUT, DELETE, PATCH")
	c.Set(fiber.HeaderContentSecurityPolicy, "default-src https:")
	return c.Next()
}
