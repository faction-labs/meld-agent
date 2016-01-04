package rust

func (r *RustServer) Exec(command string) (string, error) {
	return r.RunCommand(command)
}
