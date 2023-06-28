package cyoa

// Story represents a Choose Your Own Adventure story.
type Story map[string]Arc

// Arc represents a story arc within the Choose Your Own Adventure story.
type Arc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

// Option represents a choice or option for the reader to proceed in the story.
type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
