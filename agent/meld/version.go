package meld

func (m *Meld) Version(args int, resp *int) error {
	*resp = Version
	return nil
}
