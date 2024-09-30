package entities

type Home struct {
	HomeData
}

type HomeData struct {
	Sections *Sections `json:"sections"`
}

type HomeParams struct {
	UserID   string `json:"user_id"`
	Sections []string
}

type Sections struct {
	*SearchSection `json:"search"`
}

type SearchSection struct {
	Payload string `json:"payload"`
	Name    string `json:"name"`
	Result  string `json:"result"`
}

type SearchSectionParams struct {
}
