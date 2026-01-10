package postgres

func (h *DB) PrepareDB() error {
	if err := h.createExtension(); err != nil {
		return err
	}

	h.migrate()
	h.FillRoles()
	h.FillPermissions()
	h.FillRolesPermissions()
	return nil
}
