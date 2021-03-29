package seeder

import "github.com/grafana/grafana/pkg/services/accesscontrol"

var builtInRoles = []accesscontrol.RoleDTO{
	{
		Name:    "grafana:builtin:users:read:self",
		Version: 1,
		Permissions: []accesscontrol.Permission{
			{
				Permission: "users:read",
				Scope:      "users:self",
			},
			{
				Permission: "users.tokens:list",
				Scope:      "users:self",
			},
			{
				Permission: "users.teams:read",
				Scope:      "users:self",
			},
		},
	},
}

var predefinedRoles = []accesscontrol.RoleDTO{
	{
		Name:    "roles:adminUsers:viewer",
		Version: 1,
		Permissions: []accesscontrol.Permission{
			{
				Permission: "users.authtoken:list",
				Scope:      "*",
			},
		},
	},
	{
		Name:    "roles:adminUsers:editor",
		Version: 1,
		Permissions: []accesscontrol.Permission{
			{
				Permission: "users.password.update",
				Scope:      "*",
			},
			{
				Permission: "users:delete",
				Scope:      "*",
			},
			{
				Permission: "users:enable",
				Scope:      "*",
			},
			{
				Permission: "users:disable",
				Scope:      "*",
			},
			{
				Permission: "users.permissions.update",
				Scope:      "*",
			},
			{
				Permission: "users:logout",
				Scope:      "*",
			},
			{
				Permission: "users.authtoken:update",
				Scope:      "*",
			},
		},
	},
}
