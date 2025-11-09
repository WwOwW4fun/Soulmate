package utils

import (
    "fmt"
    "strings"
    "sync"
)

// Dynamic single-session guided flow with Gemini-generated follow-up questions.
// Behavior:
// - First call: user message becomes the problem description; ask first AI question.
// - Next calls: record answer, ask next AI question.
// - After 3 total answers: return healing plan (healingPlan field only).
var (
    gsMu      sync.Mutex
    gsProblem string
    gsQs      []string // asked questions
    gsAns     []string // answers corresponding to gsQs
    gsStep    int      // number of answers recorded
)

// StartOrAdvanceSingle returns (reply, finished, healingPlan, error)
// reply: next question or initial acknowledgement (finished=false)
// finished=true: healingPlan populated, reply empty
func StartOrAdvanceSingle(message string) (string, bool) {
    gsMu.Lock()
    defer gsMu.Unlock()

    // First user message initializes the problem and asks first question
    if gsProblem == "" {
        gsProblem = strings.TrimSpace(message)
        firstQ, err := generateNextQuestionGemini(gsProblem, nil, nil)
        if err != nil || strings.TrimSpace(firstQ) == "" {
            // Fallback question if model fails
            firstQ = "What sensations or thoughts come with that feeling?"
        }
        gsQs = []string{firstQ}
        gsAns = []string{}
        gsStep = 0
        return fmt.Sprintf("Thanks for sharing. Let's explore this together. %s", firstQ), false
    }

    // Record answer to current question
    if gsStep < len(gsQs) {
        gsAns = append(gsAns, strings.TrimSpace(message))
        gsStep++
    }

    // After collecting 3 answers, build plan
    const requiredAnswers = 3
    if gsStep >= requiredAnswers {
        plan, err := buildHealingPlan()
        if err != nil {
            // Reset state even on error
            gsProblem, gsQs, gsAns, gsStep = "", nil, nil, 0
            return "", false
        }
        gsProblem, gsQs, gsAns, gsStep = "", nil, nil, 0
        return plan, true
    }

    // Ask next AI question
    nextQ, err := generateNextQuestionGemini(gsProblem, gsQs, gsAns)
    if err != nil || strings.TrimSpace(nextQ) == "" {
        // Fallback follow-up to ensure continuity
        nextQ = "When do you notice it feeling a little better or worse?"
    }
    gsQs = append(gsQs, nextQ)
    return nextQ, false
}

// buildHealingPlan composes the plan from collected Q&A.
func buildHealingPlan() (string, error) {
    queries := make([]Query, 0, len(gsQs))
    for i := range gsQs {
        ans := ""
        if i < len(gsAns) {
            ans = gsAns[i]
        }
        queries = append(queries, Query{Question: gsQs[i], Answer: ans})
    }
    planStr, err := GetPlanGemini(gsProblem, queries)
    if err != nil {
        return "", err
    }
    return planStr, nil
}

// generateNextQuestionGemini asks Gemini for ONE best next follow-up question.
// Returns just the question text.
func generateNextQuestionGemini(problem string, asked []string, answers []string) (string, error) {
    ctx := ""
    for i := range asked {
        a := ""
        if i < len(answers) {
            a = answers[i]
        }
        ctx += fmt.Sprintf("Q%d: %s\nA%d: %s\n", i+1, asked[i], i+1, a)
    }
    prompt := "You are a highly empathetic, supportive mental health assistant. The user's main problem: \"" + problem + "\". " +
        "Conversation so far (if any):\n" + ctx +
        "Ask ONE short, compassionate follow-up question to deepen understanding. Constraints:\n" +
        "- Under 18 words\n- No diagnosis or judgment\n- No prefacing/explanation\n- Return ONLY the question text"
    q, err := callGemini(prompt, 0.85)
    if err != nil {
        return "", err
    }
    return strings.TrimSpace(q), nil
}

