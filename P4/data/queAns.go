package data

import (
	"encoding/json"
	"github.com/pkg/errors"
)

//Question
type Question struct {
	Value int
}

func (q *Question) ToJson() ([]byte, error) {
	value, err := json.Marshal(q)
	if err != nil {
		return []byte{}, errors.New("Cannot encode Question to Json")
	}
	return value, nil
}

func QuestionFromJson(inputJson []byte) (Question, error) {
	ques := Question{}
	err := json.Unmarshal(inputJson, &ques)
	if err != nil {
		return ques, errors.New("Cannot decode Json to Question")
	}
	return ques, nil
}

//Answer
type Answer struct {
	Value  int
	Result string
}

func (a *Answer) ToJson() ([]byte, error) {
	value, err := json.Marshal(a)
	if err != nil {
		return []byte{}, errors.New("Cannot encode Answer to Json")
	}
	return value, nil
}

func AnswerFromJson(inputJson []byte) (Answer, error) {
	ans := Answer{}
	err := json.Unmarshal(inputJson, &ans)
	if err != nil {
		return ans, errors.New("Cannot decode Json to Answer")
	}
	return ans, nil
}
