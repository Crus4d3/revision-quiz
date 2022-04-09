package main

import (
    "flag"
    "math/rand"
    "time"
    "github.com/crus4d3/revision-quiz/go/cmd"
    "github.com/crus4d3/revision-quiz/go/pkg/tests"
    "github.com/crus4d3/revision-quiz/go/pkg/utils"
)

func main() {
    version := flag.Bool("version", false, "output version information and exit")
    test := flag.Bool("test", false, "run tests")
    history := flag.Bool("history", false, "include history questions")
    historyS := flag.Bool("h", false, "include history questions")
    science := flag.Bool("science", false, "include science questions")
    scienceS := flag.Bool("s", false, "include science questions")
    flag.Parse()

    quizObj := utils.GetQuestions()
    var questionList []utils.Question

    if *version {
        utils.Version()
        return
    }
    if *test {
        tests.Test()
        return
    }
    if *science || *scienceS {
        for _, subject := range(quizObj) {
            if subject.Name == "science" {
                for _, question := range(subject.Questions) {
                    questionList = append(questionList, question)
                }
            }
        }
    }
    if *history || *historyS {
        for _, subject := range(quizObj) {
            if subject.Name == "history" {
                for _, question := range(subject.Questions) {
                    questionList = append(questionList, question)
                }
            }
        }
    } else if questionList == nil {
        subjectList := utils.GetSubjects(quizObj)
        if subjectList == nil {
            for _, subject := range(quizObj) {
                for _, question := range(subject.Questions) {
                    questionList = append(questionList, question)
                }
            }
        } else {
            for _, subject := range subjectList {
                for _, question := range(subject.Questions) {
                    questionList = append(questionList, question)
                }
            }
        }
    }
    rand.Seed(time.Now().UnixNano())
    quiz.RunQuiz(questionList)
}
