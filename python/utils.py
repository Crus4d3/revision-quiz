#!/usr/bin/env python3

import yaml
from jsonschema import validate, exceptions

from questions import Questions

def version():
    print("Version 1.0.0")
    print("This project is still in active development")
    print("Visit https://github.com/crus4d3/revision-quiz for the latest version")

def getSubjects(quiz):
    done = False
    chosenSubjects = []
    subjectNames = [s.name for s in quiz.subjectList]

    while done == False:
        print("Which subjects do you want to revise", *subjectNames, "or all?", sep=", ")
        print("Type 'done' to continue")
        inputSubject = input()
        inputSubject = inputSubject.strip().lower()

        if inputSubject in ["q", "quit"]:
            quit()
        elif inputSubject in ["a", "all"]:
            print("All chosen!")
            print()
            chosenSubjects = []
            done = True
        elif inputSubject in ["d", "done"]:
            if chosenSubjects != []:
                done = True
                print("Subjects chosen!")
                print()
            else:
                print()
                print("You must choose at least one subject")
        else:
            found = False
            for subject in quiz.subjectList:
                if subject.name == inputSubject:
                    chosenSubjects.append(subject)
                    found = True
                    print()
                    print(inputSubject, "chosen!")
            if found == False:
                print("Unkown option:", inputSubject)
                print()

    return chosenSubjects

class Quiz():
    def __init__(self, subjectList):
        self.subjectList = self.createSubjects(subjectList)

    def createSubjects(self, subjectList):
        fmtSubjects = []
        for subject in subjectList:
            fmtSubject = Subject(subject)
            fmtSubjects.append(fmtSubject)
        return fmtSubjects

class Subject():
    def __init__(self, questions):
        self.name = questions["name"]
        self.questions = self.createQuestions(questions)

    def createQuestions(self, questionList):
        questionList.pop("name")
        questionList = questionList["questions"]
        fmtQuestions = []
        for question in questionList:
            prompt = question["prompt"]
            answers = question["answers"]
            fmtQuestion = Question(prompt, answers)
            fmtQuestions.append(fmtQuestion)
        return fmtQuestions

class Question():
    def __init__(self, prompt, answers):
        self.prompt = prompt
        self.answers = answers

def getQuestions():
    def useDefaultQuestions():
        print("Using default question set")
        defaultQuestions = Questions()
        quiz = Quiz(defaultQuestions.subjectList)
        return quiz

    def validateFile(questions):
        schema = {
            "type": "array",
            "items": {
                "type": "object",
                "properties": {
                    "name": {"type": "string"},
                    "questions": {
                        "type": "array",
                        "items": {
                            "type": "object",
                            "properties": {
                                "prompt": {"type": "string"},
                                "answers": {
                                    "type": "array",
                                    "items": {
                                        "type": ["string", "number"]
                                    }
                                }
                            },
                            "required": ["prompt", "answers"],
                            "additionalProperties": False
                        }
                    }
                },
                "required": ["name", "questions"],
                "additionalProperties": False
            }
        }
        try:
            error = validate(questions, schema)
            return True, error
        except exceptions.ValidationError as error:
            return False, error

    try:
        with open("questions.yaml", "r") as questions:
            try:
                subjects = yaml.safe_load(questions)
                ok, error = validateFile(subjects)
                if ok:
                    quiz = Quiz(subjects)
                    print("Using custom question set in questions.yaml")
                    return quiz
                else:
                    print("Error:", error)
                    print()
                    quiz = useDefaultQuestions()
                    return quiz
            except yaml.scanner.ScannerError as error:
                print("Error:", error)
                print()
                quiz = useDefaultQuestions()
                return quiz
            except yaml.YAMLError as error:
                print("Error:", error)
                print()
                quiz = useDefaultQuestions()
                return quiz
    except FileNotFoundError as error:
        print("Error:", error)
        print()
        quiz = useDefaultQuestions()
        return quiz
