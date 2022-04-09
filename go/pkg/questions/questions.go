package questions

type Quiz struct {
    SubjectList []Subject
}

type Subject struct {
    Name string
    Questions []Question
}

type Question struct {
    Prompt string
    Answers []string
}

var historyQuestions = []map[string][]string {
    {"When was the Paris Peace Conference?" : {"1919"}},
    {"Which central power was the Treaty of Versailles with?" : {"germany"}},
    {"Which central power was the Treaty of Saint Germain with?" : {"austria"}},
    {"Which central power was the Treaty of Neuilly with?" : {"bulgaria"}},
    {"Which central power was the Treaty of Sevres with?" : {"turkey"}},
    {"Which central power was the Treaty of Lausanne with?" : {"turkey"}},
    {"Who was the leader of Britain in the Paris peace conference?" : {"david lloyd george", "lloyd george"}},
    {"Who was the leader of France in the Paris peace conference?" : {"georges clemenceau", "clemenceau"}},
    {"Who was the leader of the USA in the Paris peace conference?" : {"woodrow wilson", "wilson"}},
}

var scienceQuestions = []map[string][]string {
    {"What biological process removes carbon dioxide from the atmosphere?" : {"photosynthesis"}},
    {"How is energy lost along a food chain?" : {"respiration", "urine", "faeces", "not all biomass eaten", "not all biomass consumed"}},
}

var maths = map[string]interface{} {
    "name": "maths",
    "questions": []map[string][]string {
        {"square root 9": []string {"3", "-3"}},
    },
}

var english = map[string]interface{} {
    "name": "english",
    "questions": []map[string][]string {
        {"how long on an AO2": []string {"12", "12 mins", "12 minutes"}},
    },
}

var SubjectList = []map[string]interface{} {maths, english}

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
    quiz := Quiz{subjectList}
    return quiz
}
