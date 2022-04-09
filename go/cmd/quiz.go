package quiz

import (
    "bufio"
    "fmt"
    "time"
    "os"
    "math/rand"
    "strings"
    "github.com/crus4d3/revision-quiz/go/pkg/utils"
)

func RunQuiz(questions []utils.Question) {
    response := ""
    score := 0
    total := 0
    correct := false
    reader := bufio.NewReader(os.Stdin)
    quitMSG := []string{"q", "quit"}

    fmt.Println("Please enter all answers as concisely as possible.")
    fmt.Println("Type 'quit' to quit at any time.")

    rand.Seed(time.Now().UnixNano())
    rand.Shuffle(len(questions), func(i, j int) { questions[i], questions[j] = questions[j], questions[i] })

    for _, question := range(questions) {
        correct = false
        fmt.Println(question.Prompt)
        response, _ = reader.ReadString('\n')
        response = strings.TrimSuffix(response, "\n")
        response = strings.ReplaceAll(response, " ", "")
        response = strings.ReplaceAll(response, "\t", "")
        response = strings.ToLower(response)
        if utils.Contains(strings.ToLower(response), quitMSG) {
            break
        }
        for _, answer := range question.Answers {
            answer = strings.ReplaceAll(answer, " ", "")
            answer = strings.ReplaceAll(answer, "\t", "")
            answer = strings.ToLower(answer)

            if response == answer {
                fmt.Println("Correct!")
                correct = true
                score += 1
            }
        }
        if correct == false {
            fmt.Println("Incorrect")
            fmt.Printf("The correct answers were %+q\n", question.Answers)
        }
        total += 1
    }
    fmt.Printf("You got: %d out of %d questions correct.\n", score, total)
}
