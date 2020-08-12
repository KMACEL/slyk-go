# Slyk GO

## Index
1. [Create Client](#CreateClient)
2. [Working Principle](#WorkingPrinciple)
	1. [Get Functions](#GetFunctions)
	2. [Update Functions](#UpdateFunctions)
	3. [Create Functions](#CreateFunctions)
	


## Create Client  <a name="CreateClient"></a>

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

## Working Principle <a name="WorkingPrinciple"></a>

This library has been prepared with reference to "https://developers.slyk.io/slyk/reference/endpoints". The functions in this library are used for communication with the endpoints at this address.

### Get Functions <a name="GetFunctions"></a>

Get functions are used to fetch information.

#### Get Functions Filter  

Get functions can take filters. These filters are created by functions. For example, let's examine the Get User system.

This is the function that brings up the User list. **"filter ...\*getUserFilter"**

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

If there is a filter in "slyk endpoint" that is not in this library, you can also query it with "SetGenericQueryParameter".

```go
client.GetUsers(slyk.GetUserFilter().
		SetEmail("mertacel@gmail.com").
		SetBlocked(true).
		SetApproved(true).
		SetName("Mert").
		SetGenericQueryParameter("filter[OTHER]", "TEST"))
```

#### Filter Information

Some filters come in 3 types. These mean "=, in, of". Its equivalents are, for example,

```go
SetAssetCode(assetCode string) 
SetAssetCodeWithIN(assetCode ...string) 
SetAssetCodeWithNIN(assetCode ...string) 
```

1. "=": it is checked if the data is equal to the value. 1 piece of data is requested.
2. "in": There can be many data entries. A list of those matching the entered parameters is brought.
3. "nin": There can be many data entries. The list of non-parameters entered is brought.


### Update Functions <a name="UpdateFunctions"></a>

Functions used to update data. There are two types of usage.

Take the "UpdateUser" function for example. The first parameter is the id of the data to be updated. The second data is the struct to update.

```go
func (c Client) UpdateUser(userID string, updateUserData *UpdateUserData) (*User, error) {
	getBody, err := c.genericPatchQuery(linkUsers+"/"+userID, updateUserData)
	if err != nil {
		return nil, err
	}

	var user User
	errUnmarshal := json.Unmarshal(getBody, &user)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &user, nil
}
```

#### Method 1 - With Struct

Update Struct : 

```go
type UpdateUserData struct {
	Name       string      `json:"name,omitempty"`
	Locale     string      `json:"locale,omitempty"`
	CustomData interface{} `json:"customData,omitempty"`
}
```

Use;

```go
client.UpdateUser("cf99e4d8-bc64-4a5c-80a4-dd1e25e2018d", &slyk.UpdateUserData{
	Name:       "Mert",
	Locale:     "en",
	CustomData: map[string]interface{}{"Test": "123"},
})
```

#### Method 2 - With Methods

```go
client.UpdateUser("cf99e4d8-bc64-4a5c-80a4-dd1e25e2018d",
		slyk.UpdateUserParam().
			SetName("Mert").
			SetLocale("en"))
```


### Create Functions <a name="CreateFunctions"></a>

Used to make a new recording.  There are two types of usage.

Take the "CreateUser" function for example. As data, the struct to be created is taken.

```go
func (c Client) CreateUser(createUserdata *CreateUserData) (*User, error) {
	getBody, err := c.genericPostQuery(linkUsers, createUserdata)
	if err != nil {
		return nil, err
	}

	var user User
	errUnmarshal := json.Unmarshal(getBody, &user)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &user, nil
}
```



#### Method 1 - With Struct

Struct : 

```go
type CreateUserData struct {
	Email           string      `json:"email"`
	Code            string      `json:"code,omitempty"`
	Locale          string      `json:"locale,omitempty"`
	Name            string      `json:"name,omitempty"`
	Password        string      `json:"password,"`
	Approved        bool        `json:"approved,omitempty"`
	Blocked         bool        `json:"blocked,omitempty"`
	CustomData      interface{} `json:"customData,omitempty"`
	PrimaryWalletID string      `json:"primaryWalletId,omitempty"`
	Verified        bool        `json:"verified,omitempty"`
}
```

Use;

```go
client.CreateUser(&sly.CreateUserData{
		Email:    "mertacel@gmail.com",
		Password: "123456789.aA",
	})
```

#### Method 2 - With Methods

```go
client.CreateUser(slyk.CreateUserParameter().
		SetName("Mert").
		SetEmail("mertacel@gmail.com").
		SetPassword("123456789.aA").
		SetCustomData(map[string]interface{}{"name": "mert"}).
		SetVerified(true).
		SetApproved(true).
		SetBlocked(false).
		SetLocale("en"))
```

## User

### Get User

GetUser is returns the Slyk user list. This function returns user information and error. 

if you do not use a filter, '30' users' information is received.

```go
response,err := client.GetUsers({{OPTINAL_FILTER}})
```

#### Filter List

For Create;

```go
slyk.GetUserFilter()
```

Append Filters;

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


### Get User With ID

Fetches information about a user. User Id takes as parameter.

```go
response,err := client.GetUserWithID({{USER_ID}})
```


### Update User

Used to update the user's data.

Function;

```go
response,err := client.UpdateUser({{USER_ID}}, *UpdateUserData) 
```

### Struct

```go
type UpdateUserData struct {
	Name       string      `json:"name,omitempty"`
	Locale     string      `json:"locale,omitempty"`
	CustomData interface{} `json:"customData,omitempty"`
}
```

### Body Function

For Create;

```go
slyk.UpdateUserParam()
```

Append List;

```go
SetName(name string)
SetLocale(locale string)
SetCustomData(customData interface{})
```

### Create User

Used to register new users.


Function;

```go
response,err := client.CreateUser(*CreateUserData)
```


### Struct

```go
type CreateUserData struct {
	Email           string      `json:"email"`
	Code            string      `json:"code,omitempty"`
	Locale          string      `json:"locale,omitempty"`
	Name            string      `json:"name,omitempty"`
	Password        string      `json:"password,"`
	Approved        bool        `json:"approved,omitempty"`
	Blocked         bool        `json:"blocked,omitempty"`
	CustomData      interface{} `json:"customData,omitempty"`
	PrimaryWalletID string      `json:"primaryWalletId,omitempty"`
	Verified        bool        `json:"verified,omitempty"`
}
```

### Body Function

For Create;

```go
slyk.CreateUserParameter()
```

Append List;

```go
SetName(name string) 
SetEmail(email string) 
SetPassword(password string) 
SetLocale(locale string) 
SetCustomData(customData interface{}) 
SetApproved(approved bool) 
SetBlocked(blocked bool) 
SetCode(code string) 
SetPrimaryWalletID(primaryWalletID string) 
SetVerified(verified bool) 
```

### Set User Approve

It performs the approval process. It takes "user id" as parameter.

```go
err := client.SetUserApprove({{USER_ID}})
```

### Set User Block

It performs the block process. It takes "user id" as parameter.

```go
err := client.SetUserBlock({{USER_ID}})
```

### Set User Unblock

It performs the block process. It takes "user id" as parameter.

```go
err := client.SetUserUnblock({{USER_ID}})
```

### Change Password

it is used to change the user password.

```go
err := client.ChangePassword({{USER_ID}},{{NEW_PASSWORD}})
```

## Wallet

### Get Wallets

Brings up the wallet list.

```go
response,err := client.GetWallets({{OPTINAL_FILTER}})
```

#### Filter List

For Create;

```go
slyk.GetWalletFilter()
```

Append Filters;

```go
SetGenericQueryParameter(key string, value interface{})
SetID(id string)
SetIDWithIN(id ...string)
SetIDWithNIN(id ...string)
SetLocked(locked bool)
SetName(name string)
SetOwnerID(ownerId string)
SetReference(reference string)
SetSortWithCreatedAt()
SetSortWithCreatedAtReverse()

```

### Get Wallet With ID

The wallet id returns the information of a given wallet.

```go
response,err := client.GetWalletWithID({{WALLET_ID}}) 
```

### Get Wallet Activity

Brings wallet activities.

```go
response,err := client.GetWalletActivity({{OPTINAL_FILTER}})
```


#### Filter List

For Create;

```go
slyk.GetWalletActivityFilter()
```

Append Filters;

```go
SetGenericQueryParameter(key string, value interface{})
SetAssetCode(assetCode string)
SetAssetCodeWithIN(assetCode ...string)
SetAssetCodeWithNIN(assetCode ...string)
SetCode(code string)
SetCodeWithIN(code ...string)
SetCodeWithNIN(code ...string)
SetStatus(status string)
SetStatusWithIN(status ...string)
SetStatusWithNIN(status ...string)
SetType(getType string)
SetTypeWithIN(getType ...string)
SetTypeWithNIN(getType ...string)
SetSortWithAmount()
SetSortWithAmountReverse()
SetSortWithCreatedAt()
SetSortWithCreatedAtReverse()
SetAvailableIncludeWithUser()
SetAvailableIncludeWithWallets()
```

### Get Wallet Activity With ID

Returns the activities of 1 wallet given the id.

```go
response,err := client.GetWalletActivityWithID({{WALLET_ID}}},{{OPTINAL_FILTER}})
```

#### Filter List

For Create;

```go
slyk.GetWalletActivtyWithIDFilter()
```

Append Filters;

```go
SetGenericQueryParameter(key string, value interface{}) 
SetAssetCode(assetCode string) 
SetAssetCodeWithIN(assetCode ...string) 
SetAssetCodeWithNIN(assetCode ...string) 
SetCode(code string) 
SetCodeWithIN(code ...string) 
SetCodeWithNIN(code ...string) 
SetStatus(status string) 
SetStatusWithIN(status ...string) 
SetStatusWithNIN(status ...string) 
SetType(getType string) 
SetTypeWithIN(getType ...string) 
SetTypeWithNIN(getType ...string) 
SetSortWithAmount() 
SetSortWithAmountReverse() 
SetSortWithCreatedAt() 
SetSortWithCreatedAtReverse() 
SetAvailableIncludeWithUser() 
SetAvailableIncludeWithWallets() 
```

### Get Wallet Balance

Brings balance information.

```go
response,err := client.GetWalletBalance({{OPTINAL_FILTER}})
```


#### Filter List

For Create;

```go
slyk.GetWalletBalanceWithIDFilter()
```

Append Filters;

```go
SetGenericQueryParameter(key string, value interface{}) 
SetAssetCode(assetCode string) 
SetAssetCodeWithIN(assetCode ...string) 
SetAssetCodeWithNIN(assetCode ...string) 
```

### Get Wallet Balance With ID

Its wallet id brings the balance information given.


```go
response,err := client.GetWalletBalanceWithID({{WALLET_ID}}},{{OPTINAL_FILTER}})
```


#### Filter List

For Create;

```go
slyk.GetWalletBalanceWithIDFilter()
```

Append Filters;

```go
SetGenericQueryParameter(key string, value interface{}) 
SetAssetCode(assetCode string) 
SetAssetCodeWithIN(assetCode ...string) 
SetAssetCodeWithNIN(assetCode ...string) 
```

### Get Wallet Movements

Brings up a list of wallet Movements.


```go
response,err := client.GetWalletMovements({{WALLET_ID}}},{{OPTINAL_FILTER}})
```


#### Filter List

For Create;

```go
slyk.GetWalletMovementFilter()
```

Append Filters;

```go
SetGenericQueryParameter(key string, value interface{}) 
SetAssetCode(assetCode string) 
SetAssetCodeWithIN(assetCode ...string) 
SetAssetCodeWithNIN(assetCode ...string) 
SetSortWithAmount() 
SetSortWithAmountReverse() 
SetSortWithCreatedAt() 
SetSortWithCreatedAtReverse() 
SetAvailableTransactionWithUser() 
```

### Get Wallet Transactions

Brings up a list of Wallet Transactions


```go
response,err := client.GetWalletTransactions({{WALLET_ID}}},{{OPTINAL_FILTER}})
```


#### Filter List

For Create;

```go
slyk.GetWalletTransactionsFilter()
```

Append Filters;

```go
SetGenericQueryParameter(key string, value interface{}) 
SetAssetCode(assetCode string) 
SetAssetCodeWithIN(assetCode ...string) 
SetAssetCodeWithNIN(assetCode ...string) 
SetCode(code string) 
SetCodeWithIN(code ...string) 
SetCodeWithNIN(code ...string) 
SetStatus(status string) 
SetStatusWithIN(status ...string) 
SetStatusWithNIN(status ...string) 
SetType(getType string) 
SetTypeWithIN(getType ...string) 
SetTypeWithNIN(getType ...string) 
SetSortWithAmount() 
SetSortWithAmountReverse() 
SetSortWithCreatedAt() 
SetSortWithCreatedAtReverse() 
```

### Update Wallet

Used to update wallet information.

Function;

```go
response,err := client.UpdateWallet({{USER_ID}}, *UpdateWalletData) 
```

### Struct

```go
type UpdateWalletData struct {
	Locked     bool        `json:"locked,omitempty"`
	OwnerID    string      `json:"ownerId,omitempty"`
	CustomData interface{} `json:"customData,omitempty"`
}
```

### Body Function

For Create;

```go
slyk.UpdateWalletBody()
```

Append List;

```go
SetOwnerID(ownerID string)
SetLocked(locked bool)
SetCustomData(customData interface{})
```

### Create Wallet

Used to create a new wallet.

Function;

```go
response,err := client.CreateWallet(*CreateWalletData)
```

### Struct

```go
type CreateWalletData struct {
	Name       string      `json:"name,omitempty"`
	Locked     bool        `json:"locked,omitempty"`
	OwnerID    string      `json:"ownerId,omitempty"`
	CustomData interface{} `json:"customData,omitempty"`
}
```

### Body Function

For Create;

```go
slyk.CreateWalletBody()
```

Append List;

```go
SetName(name string)
SetOwnerID(ownerID string) 
SetLocked(locked bool) 
SetCustomData(customData interface{})
```