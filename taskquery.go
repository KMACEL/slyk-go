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

// GetTaskWithID is
// https://developers.slyk.io/slyk/reference/endpoints#get-tasks-id
func (c Client) GetTaskWithID(tasksID string) (*Task, error) {
	getBody, err := c.GenericGetQuery(linkTasks+"/"+tasksID, nil)
	if err != nil {
		return nil, err
	}

	var task Task
	errUnmarshal := json.Unmarshal(getBody, &task)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &task, nil
}

// UpdateTask is
// https://developers.slyk.io/slyk/reference/endpoints#patch-tasks-id
func (c Client) UpdateTask(tasksID string, updateTaskDataBody *UpdateTaskDataBody) (*Task, error) {
	getBody, err := c.GenericPatchQuery(linkTasks+"/"+tasksID, updateTaskDataBody)
	if err != nil {
		return nil, err
	}

	var task Task
	errUnmarshal := json.Unmarshal(getBody, &task)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &task, nil
}

// CreateTask is
// https://developers.slyk.io/slyk/reference/endpoints#post-tasks
func (c Client) CreateTask(createTaskDataBody *CreateTaskDataBody) (*Task, error) {
	getBody, err := c.GenericPostQuery(linkTasks, createTaskDataBody)
	if err != nil {
		return nil, err
	}

	var task Task
	errUnmarshal := json.Unmarshal(getBody, &task)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &task, nil
}
