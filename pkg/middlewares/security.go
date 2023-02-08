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
	c.Set(fiber.HeaderReferrerPolicy, "no-referrer")
	c.Set(fiber.HeaderPermissionsPolicy, "accelerometer=(), ambient-light-sensor=(), autoplay=(), battery=(), camera=(), cross-origin-isolated=(), display-capture=(), document-domain=(), encrypted-media=(), execution-while-not-rendered=(), execution-while-out-of-viewport=(), fullscreen=(), geolocation=(), gyroscope=(), magnetometer=(), microphone=(), midi=(), navigation-override=(), payment=(), picture-in-picture=(), publickey-credentials-get=(), screen-wake-lock=(), sync-xhr=(self), usb=(), web-share=(), xr-spatial-tracking=()")
	return c.Next()
}
