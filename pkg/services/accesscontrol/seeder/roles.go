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
	{
		Name:    "roles:adminUsers:viewer",
		Version: 1,
		Permissions: []accesscontrol.Permission{
			{
				Permission: accesscontrol.ActionUsersAuthTokenList,
				Scope:      "*",
			},
			{
				Permission: accesscontrol.ActionUsersQuotasList,
				Scope:      "*",
			},
		},
	},
	{
		Name:    "roles:adminUsers:editor",
		Version: 1,
		Permissions: []accesscontrol.Permission{
			{
				Permission: accesscontrol.ActionUsersAuthTokenList,
				Scope:      "*",
			},
			{
				Permission: accesscontrol.ActionUsersPasswordUpdate,
				Scope:      "*",
			},
			{
				Permission: accesscontrol.ActionUsersCreate,
				Scope:      "*",
			},
			{
				Permission: accesscontrol.ActionUsersDelete,
				Scope:      "*",
			},
			{
				Permission: accesscontrol.ActionUsersEnable,
				Scope:      "*",
			},
			{
				Permission: accesscontrol.ActionUsersDisable,
				Scope:      "*",
			},
			{
				Permission: accesscontrol.ActionUsersPermissionsUpdate,
				Scope:      "*",
			},
			{
				Permission: accesscontrol.ActionUsersLogout,
				Scope:      "*",
			},
			{
				Permission: accesscontrol.ActionUsersAuthTokenUpdate,
				Scope:      "*",
			},
			{
				Permission: accesscontrol.ActionUsersQuotasList,
				Scope:      "*",
			},
			{
				Permission: accesscontrol.ActionUsersQuotasUpdate,
				Scope:      "*",
			},
		},
	},
}
