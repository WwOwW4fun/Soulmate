package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"planner/utils"
)

func GetPlanHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			bodyBytes, _ := io.ReadAll(r.Body)
			var body utils.UserInfo

			// Đây là lỗi 1 (lỗi giải nén JSON)
			err := json.Unmarshal(bodyBytes, &body)
			log.Println("Received user's information successfully")
			
			if err != nil {
				// In lỗi ra terminal VÀ dừng lại
				log.Printf("Error parsing JSON: %v", err)
				http.Error(w, "Bad request", http.StatusBadRequest)
				return // Dừng lại!
			}

			// Gọi Gemini
			plan, err := utils.GetPlanGemini(body.Usergoal, body.Queries)
			
			// Đây là lỗi 2 (lỗi từ Gemini)
			if err != nil {
				// In lỗi ra terminal VÀ dừng lại
				log.Printf("Error from GetPlanGemini: %v", err) // <--- LỖI SẼ HIỆN Ở ĐÂY
				http.Error(w, "Service error", http.StatusInternalServerError)
				return // Dừng lại!
			}

			// Nếu không có lỗi, log và gửi kết quả
			log.Println("Generate the answer successfully!")
			w.Write([]byte(plan))
		}
	})
}