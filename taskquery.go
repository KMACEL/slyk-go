package slyk

import "encoding/json"

// GetTasks is
// https://developers.slyk.io/slyk/reference/endpoints#get-invites-9
func (c Client) GetTasks(filter ...*getTasksFilter) (*Tasks, error) {
	getBody, err := c.GenericGetQuery(linkTasks, merge(filter))
	if err != nil {
		return nil, err
	}

	var tasks Tasks
	errUnmarshal := json.Unmarshal(getBody, &tasks)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &tasks, nil
}
