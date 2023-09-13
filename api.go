package anki

type CreateCategory struct {
	Title string `json:"title"`
}

type CreateQuestion struct {
	Name         string `json:"name"`
	QuestionText string `json:"question-text"`
	Answer       string `json:"answer"`
}

type UpdateQuestion struct {
	Name   string `json:"update_name,omitempty"`
	Question_Text string `json:"update_question,omitempty"`
	Answer string `json:"update_answer,omitempty"`
}
