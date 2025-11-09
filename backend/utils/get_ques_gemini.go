package utils

import (
	"fmt"
	"strings"
)

func GetQuesGemini(userGoal string) ([]string, error) {

	// Vai trò chuyên nghiệp và mục tiêu của 2 câu hỏi đầu tiên
	prompt := "You are a highly empathetic and professional clinical psychologist. A user has expressed their situation/concern: \"" + userGoal + "\". You need to gently guide them to open up further. Give 2 single, open-ended, follow-up questions that probe the user's core feelings and the impact of the situation on them. (Just send a list of 2 questions, each on a new line, and nothing else). Use simple, direct, and encouraging English. Avoid overly formal or academic terms. Use short, actionable sentences."

	// Temperature cao (1.0) khuyến khích các câu hỏi sâu sắc, sáng tạo hơn
	responseText, err := callGemini(prompt, 1.0) 
	if err != nil {
		return []string{}, fmt.Errorf("failed to get initial probing questions: %w", err)
	}

	// Trích xuất các câu hỏi
	return strings.Split(responseText, "\n"), nil
}