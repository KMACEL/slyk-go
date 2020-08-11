# Slyk GO

## Create Client

To use the slyk-go library, it is necessary to create a client. You need to give this client an api key.

Api Key : https://developers.slyk.io/slyk/developing-with-slyk/quick-start-guide

```go
package main

import (
	"fmt"

	"github.com/KMACEL/slyk-go"
)

func main() {
	client := slyk.New("{{API_KEY}}")
}
```

## Working Principle

This library has been prepared with reference to "https://developers.slyk.io/slyk/reference/endpoints". The functions in this library are used for communication with the endpoints at this address.

### Get Functions

Get functions are used to fetch information.

#### Get Functions Filter 

Get functions can take filters. These filters are created by functions. For example, let's examine the Get User system.

This is the function that brings up the User list. **"filter ...*getUserFilter**"
```go
func (c Client) GetUsers(filter ...*getUserFilter) (*Users, error) {
	getBody, err := c.genericGetQuery(linkUsers, merge(filter))
	if err != nil {
		return nil, err
	}

	var users Users
	errUnmarshal := json.Unmarshal(getBody, &users)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &users, nil
}
```

"..." means bringing the data filtered or unfiltered, thanks to its Section. So if we don't filter this function, it will return all data.

```go
client.GetUsers()
```

But if you want to create a filter, a "getUserFilter" type filter must be created. But since this is a private value, a filter should be created through constructor methods.

```go
client.GetUsers(slyk.GetUserFilter().SetEmail("mertacel@gmail.com")
```
Thus, this function will bring up the user whose "email" address is "mertacel@gmail.com".

However, it is possible to add more than one filter as below.

```go
client.GetUsers(slyk.GetUserFilter().
		SetEmail("mertacel@gmail.com").
		SetBlocked(true).
		SetApproved(true).
		SetName("Mert"))
```

If there is a filter in "slyk endpoint" that is not in this library, you can also query it with this function.

```go
client.GetUsers(slyk.GetUserFilter().
		SetEmail("mertacel@gmail.com").
		SetBlocked(true).
		SetApproved(true).
		SetName("Mert").
		SetGenericQueryParameter("filter[OTHER]", "TEST"))
```

## User

### Get User

GetUser is returns the Slyk user list. This function returns user information and error. 

if you do not use a filter, '30' users' information is received.

```go
response,err:=client.GetUsers()
```

#### Filter List

For Create;

```go
slyk.GetUserFilter()
```

Append;

```go
SetGenericQueryParameter(key string, value interface{})
SetApproved(approved bool)
SetBlocked(blocked bool)
SetEmail(email string) 
SetID(id string) 
SetIDWithIN(id ...string) 
SetName(name string) 
SetPhone(phone string) 
SetPrimaryWalletId(primaryWalletId string) 
SetReferralCode(referralCode string) 
SetReferralUserID(referralUserID string) 
SetReferralUserIDWithIN(referralUserID ...string) 
SetRole(role string) 
SetVerified(verified bool) 
SetSortWithCreatedAt() 
SetSortWithCreatedAtReverse() 
SetSortWithUpdatedAt() 
SetSortWithUpdatedAtReverse() 
SetPageSize(size int) 
SetPageNumber(number int) 
```