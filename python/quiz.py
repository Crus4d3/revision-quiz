#!/usr/bin/env python3

import random

def runQuiz(questions):
    response = ""
    score = 0
    total = 0
    correct = False
    print("Please enter all answers as concisely as possible.")
    print("Type 'quit' to quit at any time.")
    random.shuffle(questions)
    for question in questions:
        correct = False
        print(question.prompt)
        response = input()
        if response.lower().strip() in ["q", "quit"]:
            break
        for answer in question.answers:
            if response.lower().strip() == str(answer).lower().strip():
                print("Correct!")
                correct = True
                score += 1
        if correct == False:
            print("Incorrect")
            print("The correct answers were", question.answers)
        total += 1
    print("You got:", score, "out of", total, "questions correct.")
