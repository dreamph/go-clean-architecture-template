package middleware

import (
	fibercasbin "github.com/gofiber/contrib/casbin"
	"github.com/gofiber/fiber/v2"
)

type AuthPermissions interface {
	RequiresPermissions(permissions []string, opts ...fibercasbin.Option) fiber.Handler
	RequiresRoles(roles []string, opts ...fibercasbin.Option) fiber.Handler
	RoutePermission() fiber.Handler
	PermissionsList() []string
}

type authPermissions struct {
	authz           *fibercasbin.Middleware
	permissionsList []string
	enable          bool
}

func NewAuthPermissions(authz *fibercasbin.Middleware, enable bool) AuthPermissions {
	var permissionsList []string
	return &authPermissions{
		authz:           authz,
		permissionsList: permissionsList,
		enable:          enable,
	}
}

func (a *authPermissions) PermissionsList() []string {
	return a.permissionsList
}

func (a *authPermissions) RequiresPermissions(permissions []string, opts ...fibercasbin.Option) fiber.Handler {
	a.permissionsList = append(a.permissionsList, permissions...)

	if !a.enable {
		return func(c *fiber.Ctx) error {
			return c.Next()
		}
	}

	return a.authz.RequiresPermissions(permissions, opts...)
}

func (a *authPermissions) RequiresRoles(roles []string, opts ...fibercasbin.Option) fiber.Handler {
	return a.authz.RequiresRoles(roles, opts...)
}

func (a *authPermissions) RoutePermission() fiber.Handler {
	return a.authz.RoutePermission()
}
