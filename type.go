package main

type SaveMd struct {
	Title string    `json:"Title"`
	Time  int       `json:"Time"`
	Desc  string    `json:"Desc"`
	Data  []Message `json:"Data"`
}
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
type Data struct {
	Model     string  `json:"model"`
	CreatedAt string  `json:"created_at"`
	Message   Message `json:"message"`
	Done      bool    `json:"done"`
}

type ModelData struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}
