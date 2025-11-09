package utils

import (
	"encoding/json"
	"fmt"
	"strings"
)

func GetPlanGemini(usergoal string, information []Query) (string, error) {

	information_json, err := json.Marshal(information)
	if err != nil {
		return "", fmt.Errorf("failed to marshal query information: %w", err)
	}
    
    // Quy tắc cốt lõi cho lời khuyên cuối cùng
    // Chúng ta buộc Gemini đưa ra Lời động viên và Lời khuyên riêng biệt.
    system_role := "You are a warm, supportive, and compassionate mentor. Use simple, direct, and encouraging English. Avoid overly formal or academic terms. Use short, actionable sentences. Your response must be deeply empathetic. BEFORE giving advice, provide validation and encouragement (or congratulations if appropriate). DO NOT diagnose. ALWAYS include a final disclaimer to seek professional help for severe issues."

	// Prompt MỚI: Yêu cầu 3 thành phần (Động viên, Lời khuyên, Hành động) trong JSON
	tempprompt := system_role + " The user's main concern is: \"" + usergoal + "\". The detailed Q&A for analysis is: " + string(information_json) + "Create an advice for this users, suggesting organizations or hotline they could use to help. Predict any problem the users may have in plain text with NO Markdown, no asterisks (*), and no bullet points"
	prompt := strings.ReplaceAll(tempprompt, "*", "")
	// Nhiệt độ thấp (0.1) để có cấu trúc JSON chính xác và phân tích ổn định.
	responseText, err := callGemini(prompt, 0.1)
	if err != nil {
		return "", fmt.Errorf("failed to generate final advice: %w", err)
	}

	// Xử lý kết quả
	return strings.ReplaceAll(responseText, "\n", ""), nil
}