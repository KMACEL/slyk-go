package slyk

import (
	"encoding/json"
	"testing"
)

func getClient() *Client {
	cl := New("ehaIILDIbu0kzNi5CqWYOlYZru3yHXX0")
	return cl
}

func OnlyError(t *testing.T, tst string, err error) {
	if err != nil {
		t.Errorf("%s error is not null : %s", tst, err.Error())
	} else {
		t.Skipf("OK : %s \n", tst)
	}
}

func ReturnAndError(t *testing.T, tst string, returnValue interface{}, err error) {

	if err != nil {
		t.Errorf("%s error is not null : %s", tst, err.Error())
	}

	if returnValue == "" || returnValue == nil {
		t.Errorf("%s return value is null ! ", tst)
	} else {
		jsonMarshal, _ := json.Marshal(returnValue)
		t.Skipf("OK : %s : %s\n", tst, string(jsonMarshal))
	}
}
