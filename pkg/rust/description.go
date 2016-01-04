package rust

func (r *RustServer) Description() (string, error) {
	d, err := r.RunCommand("server.description")
	if err != nil {
		return "", err
	}

	desc, err := r.parseStandardResponse(d)
	if err != nil {
		return "", err
	}
	return desc, nil
}
