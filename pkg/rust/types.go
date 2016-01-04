package rust

type (
	ServerStatus struct {
		Hostname    string `json:"hostname"`
		Description string `json:"description"`
		Version     string `json:"version"`
		Map         string `json:"map"`
		Players     int    `json:"players"`
	}

	StartArgs struct {
		Hostname     string
		Description  string
		Identity     string
		ServerIP     string
		RconIP       string
		ServerPort   int
		RconPort     int
		RconPassword string
		MaxPlayers   int
		Level        string
		Seed         int
		WorldSize    int
		SaveInterval int
		URL          string
		WorkDir      string
	}
)
