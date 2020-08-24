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
