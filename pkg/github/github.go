package github

const (
	// TODO: revert
	TemplateURL = "https://raw.githubusercontent.com/ryanadiputraa/ggen/refactor/structure"
	TagURL      = "https://api.github.com/repos/ryanadiputraa/ggen/tags"
)

type Tag struct {
	Name string `json:"name"`
}
