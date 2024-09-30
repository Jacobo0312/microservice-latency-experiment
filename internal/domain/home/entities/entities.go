package entities

type Home struct {
	Sections Sections `json:"sections"`
}

type HomeParams struct {
	UserID string `json:"user_id"`
}

type Sections struct {
	*SearchSection `json:"search"`
}

type SearchSection struct {
	Payload string `json:"payload"`
}
