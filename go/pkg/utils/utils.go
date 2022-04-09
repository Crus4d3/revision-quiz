package utils

import (
    "bufio"
    "fmt"
    "gopkg.in/yaml.v3"
    "io/ioutil"
    "os"
    "strings"
    "github.com/crus4d3/revision-quiz/go/pkg/questions"
)

func Version() {
    fmt.Println("Version 0.4.0")
    fmt.Println("This project is still in active development")
    fmt.Println("Visit https://github.com/crus4d3/revision-quiz for the latest version")
}

func Contains(needle string, haystack []string) bool {
    for _, value := range haystack {
        if value == needle {
            return true
        }
    }
    return false
}

func GetSubjects(quiz Quiz) []Subject {
    finished := false
    reader := bufio.NewReader(os.Stdin)

    all := []string{"a", "all"}
    done := []string{"d", "done"}
    quitMSG := []string{"q", "quit"}

    var inputSubject string
    var chosenSubjects []Subject
    var subjectNames []string
    for _, s := range(quiz) {
        subjectNames = append(subjectNames, s.Name)
    }

    for finished == false {
        fmt.Print("Which subjects do you want to revise: ")
        fmt.Print(strings.Join(subjectNames, ", "))
        fmt.Println(" or all?")
        fmt.Println("Type 'done' to continue")
        inputSubject, _ = reader.ReadString('\n')
        inputSubject = strings.TrimSuffix(inputSubject, "\n")

        if Contains(strings.ToLower(inputSubject), quitMSG) {
            os.Exit(0)
        } else if Contains(strings.ToLower(inputSubject), all) {
            fmt.Println("All chosen!\n")
            chosenSubjects = nil
            finished = true
        } else if Contains(strings.ToLower(inputSubject), done) {
            if chosenSubjects != nil {
                fmt.Println("Subjects chosen!\n")
                finished = true
            } else {
                fmt.Println("You must choose at least one subject")
            }
        } else {
            found := false
            for _, subject := range(quiz) {
                if subject.Name == inputSubject {
                    chosenSubjects = append(chosenSubjects, subject)
                    found = true
                    fmt.Printf("\n%s chosen\n", inputSubject)
                }
            }
            if found == false {
                fmt.Printf("Unkown option: %s\n", inputSubject)
            }
        }
    }
    return chosenSubjects
}

type Quiz []Subject

type Subject struct {
    Name string
    Questions []Question
}

type Question struct {
    Prompt string
    Answers []string
}

func MakeQuiz(subjects []map[string]interface{}) Quiz {
    var subjectList []Subject

    for _, subject := range subjects {
        rawName := subject["name"]
        name := rawName.(string)
        rawQuestions := subject["questions"]
        questions := rawQuestions.([]map[string][]string)
        var questionList []Question

        for _, question := range questions {
            for prompt, answers := range question {
                fmtQuestion := Question{prompt, answers}
                questionList = append(questionList, fmtQuestion)
            }
        }
        fmtSubject := Subject{name, questionList}
        subjectList = append(subjectList, fmtSubject)
    }
    quiz := subjectList
    return quiz
}

//func ReadYaml(quiz *Quiz) *Quiz {}

func GetQuestions() Quiz {
    useDefaultQuestions := func() Quiz {
        fmt.Println("Using default question set")
        quiz := MakeQuiz(questions.SubjectList)
        return quiz
    }

    //validateFile := func(filename string) (bool, error) {
    //    return true, nil
    //}

    file, err := ioutil.ReadFile("questions.yaml")
    if err != nil {
        fmt.Printf("Error: %#v\n", err)
        return useDefaultQuestions()
    }
    quiz := &Quiz{}
    err = yaml.Unmarshal(file, quiz)
    if err != nil {
        fmt.Printf("Error: %#v\n", err)
        return useDefaultQuestions()
    }
    fmt.Println("Using custom question set in questions.yaml")
    return *quiz
}
