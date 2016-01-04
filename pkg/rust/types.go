package rust

type (
	ServerStatus struct {
		Hostname    string `json:"hostname"`
		Description string `json:"description"`
		Version     string `json:"version"`
		Map         string `json:"map"`
		Players     int    `json:"players"`
	}
)
