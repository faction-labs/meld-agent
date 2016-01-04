package rust

func (r *RustServer) Hostname() (string, error) {
	host, err := r.RunCommand("server.hostname")
	if err != nil {
		return "", err
	}

	h, err := r.parseStandardResponse(host)
	if err != nil {
		return "", err
	}

	return h, nil
}
