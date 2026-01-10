package postgres

import "github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/domain"

var permissions []string = []string{"vm.list", "vm.create", "vm.edit", "vm.terminate",
	"vm.detail", "vm.start", "vm.stop", "vm.restart", "vm.connect",
	"ebs.list", "ebs.detail", "ebs.edit", "ebs.create", "ebs.delete",
	"ebs.attach", "ebs.detach", "ebs.snap.create", "ebs.snap.delete"}

func (h *DB) FillRoles() {
	roles := []string{"admin", "moderator", "user"}
	for _, value := range roles {
		role := domain.Role{
			Name: value,
		}
		h.db.Model(&domain.Role{}).Create(&role)
	}

}
func (h *DB) FillPermissions() {

	for _, value := range permissions {
		permission := domain.Permission{
			Name: value,
		}
		h.db.Model(&domain.Permission{}).Create(&permission)
	}
}
func (h *DB) FillRolesPermissions() {
	roles_permissions := map[string][]string{
		"admin": permissions,

		"moderator": {
			"vm.list", "vm.detail", "vm.create", "vm.edit",
			"vm.start", "vm.stop", "vm.restart", "vm.connect",

			"ebs.list", "ebs.detail", "ebs.edit", "ebs.create",
			"ebs.attach", "ebs.detach",
		},

		"user": {
			"vm.list",
			"vm.detail",
			"ebs.list",
			"ebs.detail",
			"vm.connect",
		},
	}

	for role, permissions := range roles_permissions {
		var role_id string
		h.db.Model(&domain.Role{}).Where("name = ?", role).Select("id").Scan(&role_id)
		for _, permission := range permissions {
			var permission_id string
			h.db.Model(&domain.Permission{}).Where("name = ?", permission).Select("id").Scan(&permission_id)
			h.db.Model(&domain.RolesPermission{}).Create(map[string]interface{}{
				"roles_id":       role_id,
				"permissions_id": permission_id,
			})
		}
	}

}
