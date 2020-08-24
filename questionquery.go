package slyk

import "encoding/json"

// GetQuestions is
//https://developers.slyk.io/slyk/reference/endpoints#get-invites-7
func (c Client) GetQuestions(filter ...*getQuestionsFilter) (*Questions, error) {
	getBody, err := c.genericGetQuery(linkQuestions, merge(filter))
	if err != nil {
		return nil, err
	}

	var questions Questions
	errUnmarshal := json.Unmarshal(getBody, &questions)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &questions, nil
}

// GetQuestionWithID is
// https://developers.slyk.io/slyk/reference/endpoints#get-questions-id
func (c Client) GetQuestionWithID(questionID string) (*Question, error) {
	getBody, err := c.genericGetQuery(linkQuestions+"/"+questionID, nil)
	if err != nil {
		return nil, err
	}

	var question Question
	errUnmarshal := json.Unmarshal(getBody, &question)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &question, nil
}

// GetQuestionsTypes is
// https://developers.slyk.io/slyk/reference/endpoints#get-invites-8
func (c Client) GetQuestionsTypes(filter ...*getQuestionsTypesFilter) (*QuestionTypes, error) {
	getBody, err := c.genericGetQuery(linkQuestionTypes, merge(filter))
	if err != nil {
		return nil, err
	}

	var questionTypes QuestionTypes
	errUnmarshal := json.Unmarshal(getBody, &questionTypes)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &questionTypes, nil
}
