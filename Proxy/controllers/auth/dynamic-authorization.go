package auth

import (
	"strings"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/controllers"
	"github.com/gofiber/fiber/v2"
)

func DynamicAuthorization() fiber.Handler {
	return func(c *fiber.Ctx) error {

		method := c.Method()
		path := c.Path() // /api/v1/instance/123/start
		//path := c.Path() // /api/v1/ebs

		permission := resolvePermission(method, path)
		if permission == "" {
			return c.JSON(controllers.FailureResponse(fiber.ErrForbidden.Code, "Forbidden", fiber.ErrForbidden.Message))
		}

		// mevcut authorization fonksiyonunu kullan
		return Authorization(permission)(c)
	}
}
func resolvePermission(method, path string) string {

	path = strings.TrimPrefix(path, "/api/v1")
	parts := strings.Split(strings.Trim(path, "/"), "/")

	if len(parts) == 0 {
		return ""
	}

	resource := parts[0] // instance, ebs, network

	switch resource {

	case "instance":
		return instancePermission(method, parts)
	case "network":
		return networkPermission(method, parts)
	case "image":
		return imagePermission(method, parts)
	case "ebs":
		return ebsPermission(method, parts)
	case "key":
		return keyPermission(method, parts)
	}

	return ""
}
func instancePermission(method string, parts []string) string {

	// /instance
	if len(parts) == 1 {
		switch method {
		case fiber.MethodPost:
			return "vm.create"
		case fiber.MethodGet:
			return "vm.list"
		}
	}

	///instance-types dahil etmekgerekiyor

	// /instance/:id
	if len(parts) == 2 {
		if method == fiber.MethodGet {
			return "vm.detail"
		}
		if method == fiber.MethodPut {
			return "vm.edit"
		}
	}

	// /instance/start | stop | restart | terminate
	if len(parts) == 2 && method == fiber.MethodPost {
		return "vm." + parts[1]
	}

	return ""
}
func networkPermission(method string, parts []string) string {

	// /network/vpcs
	if len(parts) == 2 {
		switch method {
		case fiber.MethodGet:
			return "vpc.list"
		}
	}
	if len(parts) == 3 {
		switch method {
		case fiber.MethodGet:
			return "subnet.list"
		}
	}

	return ""
}
func imagePermission(method string, parts []string) string {

	// /image/
	if len(parts) == 1 {
		switch method {
		case fiber.MethodGet:
			return "image.list"
		}
	}

	return ""
}
func keyPermission(method string, parts []string) string {

	// /key/
	if len(parts) == 1 {
		switch method {
		case fiber.MethodGet:
			return "key.list"
		}
	}

	return ""
}
func ebsPermission(method string, parts []string) string {

	if len(parts) == 1 {
		switch method {
		case fiber.MethodPost:
			return "ebs.create"
		case fiber.MethodGet:
			return "ebs.list"
		}
	}

	if len(parts) == 2 {
		switch method {
		case fiber.MethodGet:
			return "ebs.detail"
		case fiber.MethodPut:
			return "ebs.edit"
		case fiber.MethodPatch:
			return "ebs.attach"
		case fiber.MethodDelete:
			return "ebs.delete"
		case fiber.MethodPost:
			return "ebs.snap.create"
		}

	}

	if len(parts) == 3 {
		switch method {
		case fiber.MethodDelete:
			return "ebs.snap.delete"
		case fiber.MethodPatch:
			return "ebs.detach"
		}

	}

	return ""
}
