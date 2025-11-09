package utils

// Query holds a single question and its corresponding answer collected during the guided flow.
type Query struct {
    Question string `json:"question"`
    Answer   string `json:"answer"`
}
