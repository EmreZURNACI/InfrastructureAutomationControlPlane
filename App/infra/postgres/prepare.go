package postgres

func (h *DB) Prepare() error {
	if err := h.createExtension(); err != nil {
		return err
	}

	//	h.migrate()
	return nil
}
