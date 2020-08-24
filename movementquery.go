package slyk

import (
	"encoding/json"
)

// GetMovements
// https://developers.slyk.io/slyk/reference/endpoints#get-movements
func (c Client) GetMovements(filter ...*getMovementsFilter) (*Movements, error) {
	getBody, err := c.GenericGetQuery(linkMovements, merge(filter))
	if err != nil {
		return nil, err
	}

	var movements Movements
	errUnmarshal := json.Unmarshal(getBody, &movements)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &movements, nil
}

// GetMovementWithID is
// https://developers.slyk.io/slyk/reference/endpoints#get-movements-id
func (c Client) GetMovementWithID(movementID string, getMovementWithIDFilter ...*getMovementWithIDFilter) (*Movement, error) {
	getBody, err := c.GenericGetQuery(linkMovements+"/"+movementID, merge(getMovementWithIDFilter))
	if err != nil {
		return nil, err
	}

	var movement Movement
	errUnmarshal := json.Unmarshal(getBody, &movement)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &movement, nil
}
