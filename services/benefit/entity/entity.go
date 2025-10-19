package entity

type (
	Benefit struct {
		ID        int     `json:"id"`
		Title     string  `json:"title"`
		Content   string  `json:"content"`
		Bonus     string  `json:"bonus"`
		VideoURL  *string `json:"video_url"`
		SourceURL *string `json:"source_url"`
	}
)
