package utils

// Query định nghĩa cấu trúc của một câu hỏi và câu trả lời
type Query struct {
	Question string `json:"planner"`
	Answer   string `json:"user"`
}

// UserInfo định nghĩa cấu trúc dữ liệu gửi lên từ frontend
type UserInfo struct {
	Usergoal string  `json:"usergoal"`
	Time     string  `json:"time"`
	Queries  []Query `json:"queries"` 
}