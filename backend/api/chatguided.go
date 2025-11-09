package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"planner/utils"
)

func ChatGuidedHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// --- Handle CORS preflight ---
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		// --- Read request body ---
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("read body error: %v", err)
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		var req utils.ChatRequest
		if err := json.Unmarshal(bodyBytes, &req); err != nil {
			log.Printf("unmarshal error: %v", err)
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		// --- Process the chat message ---
		reply, finished := utils.StartOrAdvanceSingle(req.Message)

		res := utils.ChatResponse{
			Reply: reply,
		}

		// --- Return JSON response ---
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)

		log.Printf("Reply: %s | Finished: %v", reply, finished)
	})
}
