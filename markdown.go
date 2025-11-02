package route

type Project struct {
	Id   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

var InMemoryProjectDB []Project

 
type Markdown struct {
	Id      string `json:"id" validate:"required"`
	Content string `json:"content" validate:"required"`
	Project string `json:"project" validate:"required"`
}

var InMemoryMarkdownDB []Markdown

func init() {
	InMemoryProjectDB = ResetableDatabase(&InMemoryProjectDB, []Project{
		{Id: "proj1", Name: "Project 1"},
	})

	InMemoryMarkdownDB = ResetableDatabase(&InMemoryMarkdownDB, []Markdown{
		{Id: "1", Content:  "# Markdown 1", Project: "proj1"},
		{Id: "2", Content:  "# Markdown 2", Project: "proj1"},
	})

	RegisterRouter(&InMemoryMarkdownDB, "/api/markdowns")
	RegisterRouter(&InMemoryProjectDB, "/api/projects")
}