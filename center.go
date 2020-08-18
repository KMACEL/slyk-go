package slyk

import (
	"encoding/json"
	"fmt"
	"unsafe"

	"github.com/go-resty/resty/v2"
)

type GenericError struct {
	Message    interface{}
	StatusCode int
}

func (g GenericError) Error() string {
	getError, _ := json.Marshal(g)
	return Byte2String(getError)
}

func (c Client) genericGetQuery(link string, queryParams map[string]string) ([]byte, error) {
	clientReq := resty.New().R()
	if queryParams != nil {
		clientReq.SetQueryParams(queryParams)
	}

	resp, err := clientReq.
		SetHeader(headerApiKey, c.apiKey).
		Get(link)

	fmt.Println(Byte2String(resp.Body()))
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, GenericError{
			Message:    resp.Body(),
			StatusCode: resp.StatusCode(),
		}
	}

	return resp.Body(), nil
}

func (c Client) genericPatchQuery(link string, body interface{}) ([]byte, error) {
	resp, err := resty.New().R().
		SetBody(body).
		SetHeader(headerApiKey, c.apiKey).
		Patch(link)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, GenericError{
			Message:    resp.Body(),
			StatusCode: resp.StatusCode(),
		}
	}

	return resp.Body(), nil
}

func (c Client) genericPostQuery(link string, body interface{}) ([]byte, error) {

	client := resty.New().R()
	resp, err := client.
		SetBody(body).
		SetHeader(headerApiKey, c.apiKey).
		Post(link)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, GenericError{
			Message:    Byte2String(resp.Body()),
			StatusCode: resp.StatusCode(),
		}
	}

	return resp.Body(), nil
}

func (c Client) genericDeleteQuery(link string, body interface{}) error {
	client := resty.New().R()
	resp, err := client.
		SetBody(body).
		SetHeader(headerApiKey, c.apiKey).
		Delete(link)

	if err != nil {
		return err
	}

	if resp.IsError() {
		return GenericError{
			Message:    Byte2String(resp.Body()),
			StatusCode: resp.StatusCode(),
		}
	}
	return nil
}

// merge is
func merge(m interface{}) map[string]string {
	mmap := make(map[string]string)
	switch getMap := m.(type) {
	case []*getUsersFilter:
		for i := range getMap {
			for k, v := range *getMap[i] {
				mmap[k] = v
			}
		}
	case []*getWalletsFilter:
		for i := range getMap {
			for k, v := range *getMap[i] {
				mmap[k] = v
			}
		}
	case []*getWalletActivityWithIDFilter:
		for i := range getMap {
			for k, v := range *getMap[i] {
				mmap[k] = v
			}
		}
	case []*getWalletBalanceWithIDFilter:
		for i := range getMap {
			for k, v := range *getMap[i] {
				mmap[k] = v
			}
		}
	case []*getWalletMovementsFilter:
		for i := range getMap {
			for k, v := range *getMap[i] {
				mmap[k] = v
			}
		}
	case []*getMovementWithIDFilter:
		for i := range getMap {
			for k, v := range *getMap[i] {
				mmap[k] = v
			}
		}
	case []*getWalletTransactionstFilter:
		for i := range getMap {
			for k, v := range *getMap[i] {
				mmap[k] = v
			}
		}
	case []*getWalletActivityFilter:
		for i := range getMap {
			for k, v := range *getMap[i] {
				mmap[k] = v
			}
		}
	case []*getWalletBalanceFilter:
		for i := range getMap {
			for k, v := range *getMap[i] {
				mmap[k] = v
			}
		}
	case []*geTransactionstFilter:
		for i := range getMap {
			for k, v := range *getMap[i] {
				mmap[k] = v
			}
		}
	case []*getaddressesFilter:
		for i := range getMap {
			for k, v := range *getMap[i] {
				mmap[k] = v
			}
		}
	case []*getassetsFilter:
		for i := range getMap {
			for k, v := range *getMap[i] {
				mmap[k] = v
			}
		}
	case []*getRatesFilter:
		for i := range getMap {
			for k, v := range *getMap[i] {
				mmap[k] = v
			}
		}
	case []*getPaymentMethodsFilter:
		for i := range getMap {
			for k, v := range *getMap[i] {
				mmap[k] = v
			}
		}
	case []*getMovementsFilter:
		for i := range getMap {
			for k, v := range *getMap[i] {
				mmap[k] = v
			}
		}
	case []*getInvitesFilter:
		for i := range getMap {
			for k, v := range *getMap[i] {
				mmap[k] = v
			}
		}
	case []*getCategoriesFilter:
		for i := range getMap {
			for k, v := range *getMap[i] {
				mmap[k] = v
			}
		}
	case []*getOrdersFilter:
		for i := range getMap {
			for k, v := range *getMap[i] {
				mmap[k] = v
			}
		}
	case []*getOrderLinesWithIDFilter:
		for i := range getMap {
			for k, v := range *getMap[i] {
				mmap[k] = v
			}
		}
	}
	return mmap
}

func Byte2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
