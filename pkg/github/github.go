package github

const (
	TemplateURL = "https://raw.githubusercontent.com/ryanadiputraa/ggen/main"
	TagURL      = "https://api.github.com/repos/ryanadiputraa/ggen/tags"
)

type Tag struct {
	Name string `json:"name"`
}
