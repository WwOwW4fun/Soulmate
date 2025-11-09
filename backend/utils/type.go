package utils

type ChatRequest struct {
    Message string `json:"message"`
}

type ChatResponse struct {
	Reply       string `json:"reply,omitempty"`
	HealingPlan string `json:"healingPlan,omitempty"`
}
