package domain

type Step string

const (
	Start           Step = "start"
	MainMenu        Step = "mainMenu"
	Faq             Step = "faq"
	Quiz            Step = "quiz"
	Tips            Step = "tips"
	Scams           Step = "scam"
	Error           Step = "error"
	QuizQuestion    Step = "quizQuestion"
	QuizAnswer      Step = "quizAnswer"
	QuizExplanation Step = "quizExplanation"
	QuizFeedback    Step = "quizFeedback"
)
