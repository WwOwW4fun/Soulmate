package api

import (
    "encoding/json"
    "io"
    "log"
    "net/http"
    "planner/utils"
)

// ChatGuidedHandler exposes a POST endpoint handling the single-session guided flow.
// Request body: { "message": "user text" }
// Response body: { "reply": "string", "healingPlan": "optional" }
// (healingPlan will be embedded inside reply currently; kept field for potential future split.)
func ChatGuidedHandler() http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            w.WriteHeader(http.StatusMethodNotAllowed)
            return
        }

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

        reply, err := utils.StartOrAdvanceSingle(req.Message)
        if err != nil {
            log.Printf("guided flow error: %v", err)
            http.Error(w, "Service error", http.StatusInternalServerError)
            return
        }

        // Return minimal ChatResponse with only Reply (and HealingPlan if later separated)
        json.NewEncoder(w).Encode(utils.ChatResponse{Reply: reply})
    })
}
