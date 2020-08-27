<div>
	<br>
	<br>
	<p align="center"><a href="https://slyk.io/"><img width="150" src="media/slyk.png" alt="Slyk"></a></p>
	<br>
	<p align="center"><a href="https://slyk.io/"><img width="50" src="media/logo.png" alt="Slyk"></a> - GO Library</p>
	<br>
</div>

# Slyk GO

## Index
1. [Create Client](#create-client)
2. [Working Principle](#working-principle)
	1. [Get Functions](#get-functions)
	2. [Update Functions](#update-functions)
	3. [Create Functions](#create-functions)
3. [User](#user)
	1. [Get User](#get-user)
	2. [Get User With ID](#get-user-with-id)
	3. [Update User](#update-user)
	4. [Create User](#create-user)
	5. [Set User Approve](#set-user-approve)
	6. [Set User Block](#set-user-block)
	7. [Set User Unblock](#set-user-unblock)
	8. [Change Password](#change-password)
3. [Wallet](#wallet)
	1. [Get Wallets](#get-wallets)
	2. [Get Wallet With ID](#get-wallet-with-id)
	3. [Get Wallet Activity](#get-wallet-activity)
	4. [Get Wallet Activity With ID](#get-wallet-activity-with-id)
	5. [Get Wallet Balance](#get-wallet-balance)
	6. [Get Wallet Balance With ID](#get-wallet-balance-with-id)
	7. [Get Wallet Movements](#get-wallet-movements)
	8. [Get Wallet Transactions](#get-wallet-transactions)
	9. [Update Wallet](#update-wallet)
	10. [Create Wallet](#create-wallet)
4. [Transaction](#transaction)
	1. [Get Transactions](#get-transactions)
	2. [Get Transactions With ID](#get-transactions-with-id)
	3. [Set Transaction Approve With ID](#set-transaction-approve-with-id)
	4. [Set Transaction Confirm With ID](#set-transaction-confirm-with-id)
	5. [Set Transaction Fail With ID](#set-transaction-fail-with-id)
	6. [Set Transaction Reject With ID](#set-transaction-reject-with-id)
	7. [Create Transaction Deposit](#create-transaction-deposit)
	8. [Create Transaction Pay](#create-transaction-pay)
	9. [Create Transaction Transfer](#create-transaction-transfer)
	10. [Create Transaction Withdrawa](#create-transaction-withdrawa)
5. [Asset](#asset)
	1. [Get Assets](#get-assets)
	2. [Get Asset With Code](#get-asset-with-code)
	3. [Update Assets With Code](#update-assets-with-code)
	4. [Create Asset](#create-asset)
6. [Movement](#movement)
	1. [Get Movements](#get-movements)
	2. [Get Movement With ID](#get-movement-with-id)
7. [Payment Method](#payment-method)
	1. [Get Payment Methods](#get-payment-methods)
	2. [Get Payment Method With Slug](#get-payment-method-with-slug)
8. [Invites](#invites)
	1. [Get Invites](#get-invites)
	2. [Get Invite With Code](#get-invite-with-code)
	3. [Get Invite With Code For Validate](#get-invite-with-code-for-validate)
	4. [Create Invite](#create-invite)
	5. [Cancel Invite](#cancel-invite)
8. [Rates](#rates)
	1. [Get Rates](#get-rates)
	2. [Get Rates With Base Asset Code And Quote Asset Code](#get-rates-with-base-asset-code-and-quote-asset-code)
	3. [Update Rate](#update-rate)
	4. [Create Rate](#create-rate)
	5. [Delete Rate](#delete-rate)
9. [Address](#address)
	1. [Get Addresses](#get-addresses)
	2. [Get Address With ID](#get-address-with-id)
	3. [Create Address](#create-address)
10. [Category](#category)
	1. [Get Categories](#get-categories)
	1. [Get Category With ID](#get-category-with-id)
	1. [Update Category](#update-category)
	1. [Create Category](#create-category)
	1. [Category Reorder](#category-reorder)
	1. [Delete Category](#delete-category)


------------------------------------------------------------------------------------------------------------
------------------------------------------------------------------------------------------------------------

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

------------------------------------------------------------------------------------------------------------
------------------------------------------------------------------------------------------------------------

## Working Principle 

This library has been prepared with reference to "https://developers.slyk.io/slyk/reference/endpoints". The functions in this library are used for communication with the endpoints at this address.

### Get Functions 
Get functions are used to fetch information.

#### Get Functions Filter  

Get functions can take filters. These filters are created by functions. For example, let's examine the Get User system.

This is the function that brings up the User list. **"filter ...\*getUsersFilter"**

```go
func (c Client) GetUsers(filter ...*getUsersFilter) (*Users, error) {
	.
	.
	.
}
```

"..." means bringing the data filtered or unfiltered, thanks to its Section. So if we don't filter this function, it will return all data.

```go
client.GetUsers()
```

But if you want to create a filter, a "getUsersFilter" type filter must be created. But since this is a private value, a filter should be created through constructor methods.

```go
client.GetUsers(slyk.GetUsersFilter().SetEmail("mertacel@gmail.com")
```
Thus, this function will bring up the user whose "email" address is "mertacel@gmail.com".

However, it is possible to add more than one filter as below.

```go
client.GetUsers(slyk.GetUsersFilter().
		SetEmail("mertacel@gmail.com").
		SetBlocked(true).
		SetApproved(true).
		SetName("Mert"))
```

If there is a filter in "slyk endpoint" that is not in this library, you can also query it with "SetGenericQueryParameter".

```go
client.GetUsers(slyk.GetUsersFilter().
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

Note for page filters :
* PageSize = Defines the number of results per page. Default = 30.
* Page Number = Defines the number of the page to retrieve. Default = 1

> **NOTE : In this document, the filters are located under "Filter List". "For Create;" Then comes the filter suitable for that function. "Filters That Can Be Added;" After, come filters that can be added to that filter.** 

> **Example = client.GetUsers(slyk.GetUsersFilter().SetEmail("mertacel@gmail.com").SetBlocked(true))**

	client.GetUsers = API

	slyk.GetUsersFilter = Filter Creator

	SetEmail : Filter ()

	SetBlocked : Filter




### Update Functions

Functions used to update data. There are two types of usage.

Take the "UpdateUser" function for example. The first parameter is the id of the data to be updated. The second data is the struct to update.

```go
func (c Client) UpdateUser(userID string, updateUserData *UpdateUserDataBody) (*User, error) {
	.
	.
	.
}
```

#### Method 1 - With Struct

A struct is required for each "update" function. You can enter data directly with struct suitable for Update function and perform update operations.

Update Struct : 

```go
type UpdateUserDataBody struct {
	Name       string      `json:"name,omitempty"`
	Locale     string      `json:"locale,omitempty"`
	CustomData interface{} `json:"customData,omitempty"`
}
```

Use;

```go
client.UpdateUser("cf99e4d8-bc64-4a5c-80a4-dd1e25e2018d", &slyk.UpdateUserDataBody{
	Name:       "Mert",
	Locale:     "en",
	CustomData: map[string]interface{}{"Test": "123"},
})
```

#### Method 2 - With Methods
You can assign parameters to be updated with a constructor method.

```go
client.UpdateUser("cf99e4d8-bc64-4a5c-80a4-dd1e25e2018d",
		slyk.UpdateUserDataForBody().
			SetName("Mert").
			SetLocale("en"))
```

### Create Functions

Used to make a new recording.  There are two types of usage.

Take the "CreateUser" function for example. As data, the struct to be created is taken.

```go
func (c Client) CreateUser(createUserdata *CreateUserDataBody) (*User, error) {
	.
	.
	.
}
```

#### Method 1 - With Struct


A struct is required for each "create" function. You can enter data directly with struct suitable for Create function and perform create operations.

Struct : 

```go
type CreateUserDataBody struct {
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
client.CreateUser(&sly.CreateUserDataBody{
		Email:    "mertacel@gmail.com",
		Password: "123456789.aA",
	})
```

#### Method 2 - With Methods

You can assign parameters to be created with a constructor method.


```go
client.CreateUser(slyk.CreateUserDataForBody().
		SetName("Mert").
		SetEmail("mertacel@gmail.com").
		SetPassword("123456789.aA").
		SetCustomData(map[string]interface{}{"name": "mert"}).
		SetVerified(true).
		SetApproved(true).
		SetBlocked(false).
		SetLocale("en"))
```

------------------------------------------------------------------------------------------------------------
------------------------------------------------------------------------------------------------------------
------------------------------------------------------------------------------------------------------------

## User

### Get User

GetUser is returns the Slyk user list. This function returns user information and error. 

if you do not use a filter, '30' users' information is received.

```go
response,err := client.GetUsers({{OPTIONAL_FILTER}})
```

#### Filter List

For Create;

```go
slyk.GetUsersFilter()
```

Filters That Can Be Added;

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
response,err := client.UpdateUser({{USER_ID}}, *UpdateUserDataBody) 
```

#### Struct

```go
type UpdateUserDataBody struct {
	Name       string      `json:"name,omitempty"`
	Locale     string      `json:"locale,omitempty"`
	CustomData interface{} `json:"customData,omitempty"`
}
```

#### Body Function

For Create;

```go
slyk.UpdateUserDataForBody()
```

Append List;

```go
SetName(name string)
SetLocale(locale string)
SetCustomData(customData interface{})
```

### Create User

Used to register new users.

Note ;
* Login password. It must contain: 1 capital letter; 1 lower letter; 1 digit; at least 8 characters.

Function;

```go
response,err := client.CreateUser(*CreateUserDataBody)
```


#### Struct

```go
type CreateUserDataBody struct {
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

#### Body Function

For Create;

```go
slyk.CreateUserDataForBody()
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
An user that is not approved is unable to login.

```go
err := client.SetUserApprove({{USER_ID}})
```

### Set User Block

It performs the block process. It takes "user id" as parameter.
An user that is blocked is unable to login.

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
New password. It must contain: 1 capital letter; 1 lower letter; 1 digit; at least 8 characters.

```go
err := client.ChangePassword({{USER_ID}},{{NEW_PASSWORD}})
```

------------------------------------------------------------------------------------------------------------
------------------------------------------------------------------------------------------------------------

## Wallet

### Get Wallets

Brings up the wallet list.

```go
response,err := client.GetWallets({{OPTIONAL_FILTER}})
```

#### Filter List

For Create;

```go
slyk.GetWalletsFilter()
```

Filters That Can Be Added;

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
Convenience endepoint endpoint to get a paginated list of transactions where the wallet id may be either the destinationWalletId or the originWalletId of the Transaction object.


```go
response,err := client.GetWalletWithID({{WALLET_ID}}) 
```

### Get Wallet Activity

Brings wallet activities.

```go
response,err := client.GetWalletActivity({{OPTIONAL_FILTER}})
```


#### Filter List

For Create;

```go
slyk.GetWalletActivityFilter()
```

Filters That Can Be Added;

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
response,err := client.GetWalletActivityWithID({{WALLET_ID}}},{{OPTIONAL_FILTER}})
```

#### Filter List

For Create;

```go
slyk.GetWalletActivtyWithIDFilter()
```

Filters That Can Be Added;

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
Retrieves a payspace's balance for the various assets it contains. Only assets that were ever transacted on the wallet will be shown (even if their balance is currently zero).

```go
response,err := client.GetWalletBalance({{OPTIONAL_FILTER}})
```


#### Filter List

For Create;

```go
slyk.GetWalletBalanceWithIDFilter()
```

Filters That Can Be Added;

```go
SetGenericQueryParameter(key string, value interface{}) 
SetAssetCode(assetCode string) 
SetAssetCodeWithIN(assetCode ...string) 
SetAssetCodeWithNIN(assetCode ...string) 
```

### Get Wallet Balance With ID

Its wallet id brings the balance information given.


```go
response,err := client.GetWalletBalanceWithID({{WALLET_ID}}},{{OPTIONAL_FILTER}})
```


#### Filter List

For Create;

```go
slyk.GetWalletBalanceWithIDFilter()
```

Filters That Can Be Added;

```go
SetGenericQueryParameter(key string, value interface{}) 
SetAssetCode(assetCode string) 
SetAssetCodeWithIN(assetCode ...string) 
SetAssetCodeWithNIN(assetCode ...string) 
```

### Get Wallet Movements

Brings up a list of wallet Movements.


```go
response,err := client.GetWalletMovements({{WALLET_ID}}},{{OPTIONAL_FILTER}})
```


#### Filter List

For Create;

```go
slyk.GetWalletMovementsFilter()
```

Filters That Can Be Added;

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
response,err := client.GetWalletTransactions({{WALLET_ID}}},{{OPTIONAL_FILTER}})
```


#### Filter List

For Create;

```go
slyk.GetWalletTransactionsFilter()
```

Filters That Can Be Added;

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

#### Struct

```go
type UpdateWalletDataBody struct {
	Locked     bool        `json:"locked,omitempty"`
	OwnerID    string      `json:"ownerId,omitempty"`
	CustomData interface{} `json:"customData,omitempty"`
}
```

#### Body Function

For Create;

```go
slyk.UpdateWalletDataForBody()
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

#### Struct

```go
type CreateWalletDataBody struct {
	Name       string      `json:"name,omitempty"`
	Locked     bool        `json:"locked,omitempty"`
	OwnerID    string      `json:"ownerId,omitempty"`
	CustomData interface{} `json:"customData,omitempty"`
}
```

#### Body Function

For Create;

```go
slyk.CreateWalletDataForBody()
```

Append List;

```go
SetName(name string)
SetOwnerID(ownerID string) 
SetLocked(locked bool) 
SetCustomData(customData interface{})
```

------------------------------------------------------------------------------------------------------------
------------------------------------------------------------------------------------------------------------

## Transaction

### Get Transactions

Returns transactions data.

```go
response,err := client.GetTransactions({{OPTIONAL_FILTER}})
```

#### Filter List

For Create;

```go
slyk.GetTransactionsFilter()
```

Filters That Can Be Added;

```go
SetGenericQueryParameter(key string, value interface{})
SetAssetCode(assetCode string)
SetAssetCodeWithIN(assetCode ...string)
SetAssetCodeWithNIN(assetCode ...string)
SetCode(code string)
SetCodeWithIN(code ...string)
SetCodeWithNIN(code ...string)
SetCreated(date string)
SetCreatedAtWithGTE(date string)
SetCreatedAtWithLTE(date string)
SetDescription(description string)
SetDestinationWalletID(destinationWalletID string)
SetDestinationWalletIdWithIN(destinationWalletId ...string)
SetDestinationWalletIdWithNIN(destinationWalletId ...string)
SetExternalId(externalId string)
SetExternalIdWithIN(externalId ...string)
SetExternalIdWithNIN(externalId ...string)
SetExternalReferenceWithNIN(externalReference string)
SetID(id string)
SetIDWithIN(id ...string)
SetIDWithNIN(id ...string)
SetOriginWalletID(originWalletId string)
SetOriginWalletIDWithIN(originWalletId ...string)
SetOriginWalletIdWithNIN(originWalletId ...string)
SetRSeference(reference string)
SetStatus(status string)
SetStatusWithIN(status ...string)
SetStatusWithNIN(status ...string)
SetType(typeParam string)
SetTypeWithIN(typeParam ...string)
SetTypeWithNIN(typeParam ...string)
SetWalletID(walletID string)
SetWalletIdWithIN(walletId ...string)
SetWalletIdWithNIN(walletId ...string)
SetSortWithAmount()
SetSortWithAmountReverse()
SetSortWithCreatedAt()
SetSortWithCreatedAtReverse()
SetPageSize(size int)
SetPageNumber(number int)
```

### Get Transactions With ID

Returns 1 Transaction data given its id.

```go
response,err := client.GetTransactionsWithID({{TRANSACTIN_ID}} )
```

### Set Transaction Approve With ID

It is used to approve the transaction in pending.

```go
response,err := client.SetTransactionApprove({{TRANSACTIN_ID}} )
```

### Set Transaction Confirm With ID

It is used to confirm the transaction in pending.

```go
response,err := client.SetTransactionConfirm({{TRANSACTIN_ID}} )
```

### Set Transaction Fail With ID

It is used to fail the transaction in pending.

```go
response,err := client.SetTransactionFail({{TRANSACTIN_ID}},*TransactionFailDataBody)
```

#### Struct

```go
type TransactionFailDataBody struct {
	Reason string `json:"reason"`
}
```

#### Body Function

For Create;

```go
slyk.TransactionFailDataForBody()
```

Append List;

```go
SetReason(reason string)
```

### Set Transaction Reject With ID

It is used to Reject the transaction in pending.

```go
response,err := client.SetTransactionReject({{TRANSACTIN_ID}}, *TransactionRejectDataBody)
```

#### Struct

```go
type TransactionRejectDataBody struct {
	Reason string `json:"reason"`
}
```

#### Body Function

For Create;

```go
slyk.TransactionRejectDataForBody()
```

Append List;

```go
SetReason(reason string)
```


### Create Transaction Deposit

Used to create deposit to the transection.

Function;

```go
response,err := client.CreateTransactionDeposit({{TRANSACTIN_ID}},*CreateTransactionDepositDataBody)
```

#### Struct

```go
type CreateTransactionDepositDataBody struct {
	Amount              string      `json:"amount"`
	AssetCode           string      `json:"assetCode"`
	Code                string      `json:"code"`
	CustomData          interface{} `json:"customData,omitempty"`
	Data                interface{} `json:"data"`
	Description         string      `json:"description,omitempty"`
	DestinationAddress  string      `json:"destinationAddress,omitempty"`
	DestinationWalletID string      `json:"destinationWalletId"`
	Commit              bool        `json:"commit,omitempty"`
	ExternalReference   string      `json:"externalReference,omitempty"`
}
```

#### Body Function

For Create;

```go
slyk.CreateTransactionDepositDataForBody()
```

Append List;

```go
SetAmount(amount string)
SetAssetCode(assetCode string)
SetCode(code string)
SetCommit(commit bool)
SetCustomData(customData interface{})
SetDescription(description string)
SetDestinationAddress(destinationAddress string)
SetDestinationWalletID(destinationWalletID string)
SetExternalReference(externalReference string)
```

### Create Transaction Pay

Function;

```go
response,err := client.CreateTransactionPay(*CreateTransactionPayDataBody)
```

#### Struct

```go
type CreateTransactionPayDataBody struct {
	Amount         string      `json:"amount"`
	AssetCode      string      `json:"assetCode"`
	CustomData     interface{} `json:"customData,omitempty"`
	Description    string      `json:"description,omitempty"`
	OriginWalletID string      `json:"originWalletId"`
}
```

#### Body Function

For Create;

```go
slyk.CreateTransactionPayDataForBody()
```

Append List;

```go
SetAmount(amount string) 
SetAssetCode(assetCode string) 
SetCustomData(customData interface{}) 
SetDescription(description string) 
SetOriginWalletID(originWalletID string) 
```

### Create Transaction Transfer

Function;
```go
response,err := client.CreateTransactionTransfer(*CreateTransactionTransferDataBody)
```

#### Struct

```go
type CreateTransactionTransferDataBody struct {
	Amount              string      `json:"amount"`
	AssetCode           string      `json:"assetCode"`
	Code                string      `json:"code"`
	Commit              bool        `json:"commit,omitempty"`
	CustomData          interface{} `json:"customData,omitempty"`
	Description         string      `json:"description,omitempty"`
	DestinationAddress  string      `json:"destinationAddress,omitempty"`
	DestinationWalletID string      `json:"destinationWalletId"`
	ExternalReference   string      `json:"externalReference,omitempty"`
	OriginAddress       string      `json:"originAddress,omitempty"`
	OriginWalletID      string      `json:"originWalletId"`
}
```

#### Body Function

For Create;

```go
slyk.CreateTransactionTransferDataForBody()
```

Append List;

```go
SetAmount(amount string) 
SetAssetCode(assetCode string) 
SetCode(code string) 
SetCommit(commit bool) 
SetCustomData(customData interface{}) 
SetDescription(description string) 
SetDestinationAddress(destinationAddress string) 
SetDestinationWalletID(destinationWalletID string) 
SetExternalReference(externalReference string) 
SetOriginWalletID(originWalletID string) 
SetOriginAddress(originAddress string) 
```

### Create Transaction Withdrawal

Function;
```go
response,err := client.CreateTransactionWithdrawal(*CreateTransactionWithdrawalDataBody)
```

#### Struct

```go
type CreateTransactionWithdrawalDataBody struct {
	Amount            string      `json:"amount"`
	AssetCode         string      `json:"assetCode"`
	Code              string      `json:"code"`
	Commit            bool        `json:"commit,omitempty"`
	CustomData        interface{} `json:"customData,omitempty"`
	Data              interface{} `json:"data,omitempty"`
	Description       string      `json:"description,omitempty"`
	ExternalReference string      `json:"externalReference,omitempty"`
	OriginAddress     string      `json:"originAddress,omitempty"`
	OriginWalletID    string      `json:"originWalletId"`
}
```

#### Body Function

For Create;

```go
slyk.CreateTransactionWithdrawalDataForBody()
```

Append List;

```go

SetAmount(amount string) 
SetAssetCode(assetCode string) 
SetCode(code string) 
SetCommit(commit bool) 
SetCustomData(customData interface{}) 
SetDescription(description string) 
SetExternalReference(externalReference string) 
SetOriginWalletID(originWalletID string) 
SetOriginAddress(originAddress string) 
```

------------------------------------------------------------------------------------------------------------
------------------------------------------------------------------------------------------------------------

## Asset

### Get Assets

Brings up the assets list.

```go
response,err := client.GetAssets({{OPTIONAL_FILTER}})
```

#### Filter List

For Create;

```go
slyk.GetAssetsFilter()
```

Filters That Can Be Added;

```go
SetGenericQueryParameter(key string, value interface{}) 
SetCode(code string) 
SetCodeWithIN(code ...string) 
SetCodeWithNIN(code ...string) 
SetEnabled(enabled bool) 
SetName(name string) 
SetSystem(system bool) 
SetType(getType string) 
SetTypeWithIN(getType ...string) 
SetTypeWithNIN(getType ...string) 
SetSortWithCreatedAt() 
SetSortWithCreatedAtReverse() 
SetSortWithEnabled() 
SetSortWithEnabledReverse() 
SetSortWithSystem() 
SetSortWithSystemReverse() 
SetSortWithType() 
SetSortWithTypeReverse() 
SetPageSize(size int) 
SetPageNumber(number int)
```

### Get Asset With Code

It brings 1 asset information whose code is given.

```go
response,err := client.GetAssetsWithCode({{ASSET_CODE}})
```

### Update Assets With Code

Updates the asset information.

Note : System assets cannot be updated.

Function;

```go
response,err := client.UpdateAssets({{ASSET_CODE}}, *UpdateAssetDataBody) 
```

#### Struct

```go
type UpdateAssetDataBody struct {
	DecimalPlaces int         `json:"decimalPlaces,omitempty"`
	Name          string      `json:"name,omitempty"`
	Contract      struct{}    `json:"contract,omitempty"`
	CustomData    interface{} `json:"customData,omitempty"`
	Enabled       bool        `json:"enabled,omitempty"`
	Logo          string      `json:"logo,omitempty"`
	Symbol        string      `json:"symbol,omitempty"`
}
```

#### Body Function

For Create;

```go
slyk.UpdateAssetDataForBody()
```

Append List;

```go
SetName(name string) 
SetDecimalPlaces(decimalPlaces int) 
SetContract(contract struct{}) 
SetCustomData(customData interface{}) 
SetEnabled(enabled bool) 
SetLogo(logo string) 
SetSymbol(symbol string) 
```

### Create Asset

Updates the asset information.

Function;

```go
response,err := client.CreateAsset(*CreateAssetDataBody) 
```

#### Struct

```go
type CreateAssetDataBody struct {
	Code          string      `json:"code"`
	Contract      struct{}    `json:"contract,omitempty"`
	CustomData    interface{} `json:"customData,omitempty"`
	DecimalPlaces int         `json:"decimalPlaces"`
	Enabled       bool        `json:"enabled,omitempty"`
	Name          string      `json:"name"`
	Symbol        string      `json:"symbol,omitempty"`
	Type          string      `json:"type"`
}
```

#### Body Function

For Create;

```go
slyk.CreateAssetDataForBody()
```

Append List;

```go
SetCode(code string) 
SetName(name string) 
SetContract(contract struct{}) 
SetCustomData(customData interface{}) 
SetDecimalPlaces(decimalPlaces int) 
SetEnabled(enabled bool) 
SetSymbol(sysmbol string) 
SetType(typeParam string) 
```

------------------------------------------------------------------------------------------------------------
------------------------------------------------------------------------------------------------------------

## Movement

### Get Movements

Brings up the movements list.

```go
response,err := client.GetMovements({{OPTIONAL_FILTER}})
```

#### Filter List

For Create;

```go
slyk.GetMovementsFilter()
```

Filters That Can Be Added;

```go
SetGenericQueryParameter(key string, value interface{})
SetAssetCode(assetCode string)
SetAssetCodeWithIN(assetCode ...string)
SetAssetCodeWithNIN(assetCode ...string)
SetCreatedAtWithGTE(date string)
SetCreatedAtWithLTE(date string)
SetID(id string)
SetIDWithIN(id ...string)
SetIDWithNIN(id ...string)
SetTransaction(transactionID string)
SetTransactionIDWithIN(transactionID ...string)
SetTransactionIDWithNIN(transactionID ...string)
SetWalletID(walletID string)
SetWalletIDWithIN(walletID ...string)
SetWalletIDWithNIN(walletID ...string)
SetSortWithCreatedAt()
SetSortWithCreatedAtReverse()
SetSortWithAmount(amount string)
SetSortWithAmountReverse(amount string)
SetAvailableTransactionWithUser()
SetPageSize(size int)
SetPageNumber(number int)
```

### Get Movement With ID

It brings 1 movement information whose id is given.

```go
response,err := client.GetMovementWithID({{MOVEMENT_ID}},{{OPTIONAL_FILTER}})
```

#### Filter List

For Create;

```go
slyk.GetMovementWithIDFilter()
```

Filters That Can Be Added;

```go
SetGenericQueryParameter(key string, value interface{}) 
SetAvailableTransactionWithUser() 
```

------------------------------------------------------------------------------------------------------------
------------------------------------------------------------------------------------------------------------

## Payment Method

### Get Payment Methods

Brings up the payment medhods list.

```go
response,err := client.GetPaymentMethods({{OPTIONAL_FILTER}})
```

#### Filter List

For Create;

```go
slyk.GetPaymentMedhodsFilter()
```

Filters That Can Be Added;

```go
SetGenericQueryParameter(key string, value interface{})
SetCreatedAtWithGTE(date string)
SetCreatedAtWithLTE(date string)
SetEnabledWithIN(enabled bool)
SetEnabledWithNIN(enabled bool)
SetSlug(slug string)
SetSlugWithIN(slug ...string) 
SetSlugWithNIN(slug ...string)
```

### Get Payment Method With Slug

It brings 1 payment medhod information whose slug is given.

```go
response,err := client.GetPaymentMethodsWithSlug({{SLUG}})
```

------------------------------------------------------------------------------------------------------------
------------------------------------------------------------------------------------------------------------

## Invites

### Get Invites

Brings up the invites list.

```go
response,err := client.GetInvites({{OPTIONAL_FILTER}})
```

#### Filter List

For Create;

```go
slyk.GetInvitesFilter()
```

Filters That Can Be Added;

```go
SetGenericQueryParameter(key string, value interface{}) 
SetCode(code string) 
SetCodeWithIN(code ...string) 
SetCodeWithNIN(code ...string) 
SetExpiredAtWithGTE(date string) 
SetExpiredAtWithLTE(date string) 
SetInvitedEmail(invitedEmail string) 
SetInvitedEmailWithIN(invitedEmail ...string) 
SetInvitedEmailWithNIN(invitedEmail ...string) 
SetInvitedUserID(invitedEmail string) 
SetInvitedUserIDWithIN(invitedUserID ...string) 
SetInvitedUserIDWithNIN(invitedUserID ...string) 
SetInviterUserID(inviterUserID string) 
SetInviterUserIDWithIN(inviterUserID ...string) 
SetInviterUserIDWithNIN(inviterUserID ...string) 
SetStatus(status string) 
SetStatusWithIN(status ...string) 
SetStatusWithNIN(status ...string) 
SetType(typeInvite string) 
SetSortWithCreatedAt() 
SetSortWithCreatedAtReverse() 
SetSortWithExpiredAt() 
SetSortWithExpiredAtReverse() 
SetSortWithUpdatedAt() 
SetSortWithUpdatedAtReverse() 
SetPageSize(size int) 
SetPageNumber(number int) 
```

### Get Invite With Code

It brings 1 invite information whose slug is given.

```go
response,err := client.GetInviteWithCode({{INVITE_CODE}})
```

### Get Invite With Code For Validate


```go
response,err := client.GetInviteWithCodeForValidate({{INVITE_CODE}})
```


### Create Invite

Create the asset information.

Function;

```go
response,err := client.CreateInvite(*CreateInviteDataBody) 
```

#### Struct

```go
type CreateInviteDataBody struct {
	Email         string `json:"email,omitempty"`
	InviterUserID string `json:"inviterUserId,omitempty"`
}
```

#### Body Function

For Create;

```go
slyk.CreateInviteBody()
```

Append List;

```go
SetEmail(email string)
SetInviterUserID(inviterUserID string)
```

### Cancel Invite

```go
response,err := client.CancelInvite({{INVITE_CODE}})
```

------------------------------------------------------------------------------------------------------------
------------------------------------------------------------------------------------------------------------

## Rates

### Get Rates

Brings up the rates list.

```go
response,err := client.GetRates({{OPTIONAL_FILTER}})
```

#### Filter List

For Create;

```go
slyk.GetRatesFilter()
```

Filters That Can Be Added;

```go
SetGenericQueryParameter(key string, value interface{})
SetBaseAssetCode(baseAssetCode string)
SetBaseAssetCodeWithIN(baseAssetCode ...string)
SetBaseAssetCodeWithNIN(baseAssetCode ...string)
SetQuoteAssetCode(quoteAssetCode string)
SetQuoteAssetCodeWithIN(quoteAssetCode ...string)
SetQuoteAssetCodeWithNIN(quoteAssetCode ...string)
SetRateWithGT(rate string)
SetRateWithGTE(rate string)
SetRateWithLT(rate string)
SetRateWithLTE(rate string)
SetSortWithCreatedAt()
SetSortWithCreatedAtReverse()
SetSortWithUpdatedAt()
SetSortWithUpdatedAtReverse()
SetSortWitRate()
SetSortWitRateReverse()
SetPageSize(size int)
SetPageNumber(number int)
```

### Get Rates With Base Asset Code And Quote Asset Code

```go
response,err := client.GetRatesWithBaseAssetCodeAndQuoteAssetCode({{BASE_ASSET_CODE}},{{QUOTE_ASSET_CODE}})
```

### Update Rate

Function;

```go
response,err := client.UpdateRate({{BASE_ASSET_CODE}},{{QUOTE_ASSET_CODE}},*UpdateRateBodyData)
```
#### Struct

```go
type UpdateRateDataBody struct {
	Rate       string      `json:"rate"`
	CustomData interface{} `json:"customData,omitempty"`
}
```


#### Body Function

For Create;

```go
slyk.UpdateRateDataForBody()
```

Append List;

```go
SetRate(rate string) 
SetCustomData(customData interface{}) 
```

### Create Rate

Function;

```go
response,err := client.CreateRate(*CreateRateDataBody)
```

#### Struct

```go
type CreateRateDataBody struct {
	BaseAssetCode  string      `json:"baseAssetCode"`
	QuoteAssetCode string      `json:"quoteAssetCode"`
	Rate           string      `json:"rate"`
	CustomData     interface{} `json:"customData,omitempty"`
}
```


#### Body Function

For Create;

```go
slyk.CreateRateDataForBody()
```

Append List;

```go
SetBaseAssetCode(baseAssetCode string)
SetQuoteAssetCode(quoteAssetCode string)
SetRate(rate string)
SetCustomData(customData interface{})
```

### Delete Rate

```go
response,err := client.DeleteRate({{BASE_ASSET_CODE}},{{QUOTE_ASSET_CODE}})
```

------------------------------------------------------------------------------------------------------------
------------------------------------------------------------------------------------------------------------

## Address

Note : Address APIs become active only after "CoinBase" integration. It is "Coinbase" as the Only Provider.

### Get Addresses

Brings up the address list.

```go
response,err := client.GetAddresses({{OPTIONAL_FILTER}})
```

#### Filter List

For Create;

```go
slyk.GetAddressesFilter()
```

Filters That Can Be Added;

```go
SetGenericQueryParameter(key string, value interface{})
SetAddress(address string)
SetAssetCode(assetCode string)
SetAssetCodeWithIN(assetCode ...string)
SetAssetCodeWithNIN(assetCode ...string)
SetID(walletID string)
SetIDWithIN(walletID ...string)
SetIDWithNIN(walletID ...string)
```

### Get Address With ID

```go
response,err := client.GetAddressWithID({{ADDRESS_ID}})
```

### Create Address

Function;

```go
response,err := client.CreateAddress(*CreateAddressBody)
```

#### Struct

```go
type CreateAddressBody struct {
	Address    string      `json:"address,omitempty"`
	AssetCode  string      `json:"assetCode,omitempty"`
	CustomData interface{} `json:"customData,omitempty"`
	Provider   string      `json:"provider,omitempty"`
	WalletID   string      `json:"walletId,omitempty"`
}
```

#### Body Function

For Create;

```go
slyk.CreateAddressForBody()
```

Append List;

```go
SetAddress(address string) 
SetAssetCode(assetCode string) 
SetCustomData(customData interface{}) 
SetProvider(provider string) 
SetWalletID(walletID string) 
```

------------------------------------------------------------------------------------------------------------
------------------------------------------------------------------------------------------------------------

## Category

### Get Categories

Brings up the categories list.

```go
response,err := client.GetCategories({{OPTIONAL_FILTER}})
```

#### Filter List

For Create;

```go
slyk.GetCategoriesFilter()
```

Filters That Can Be Added;

```go
SetGenericQueryParameter(key string, value interface{}) 
SetAvailableProducts(availableProducts interface{}) 
SetDescription(description string) 
SetID(id string) 
SetIDWithIN(id ...string) 
SetIDWithNIN(id ...string) 
SetOrder(order string) 
SetOrderWithGTE(order string) 
SetOrderWithLTE(order string) 
SetTitle(title string) 
SetVisibleProducts(visibleProducts interface{}) 
SetSortWithCreatedAt() 
SetSortWithCreatedAtReverse() 
SetSortWithOrder() 
SetSortWithOrderReverse() 
SetSortWithTitle() 
SetSortWithTitleReverse() 
SetPageSize(size int) 
SetPageNumber(number int) 
```

### Get Category With ID

```go
response,err := client.GetCategoryWithID({{CATEGORY_ID}})
```

### Update Category

Function;

```go
response,err := client.UpdateCategory({{CATEGORY_ID}},*UpdateCategoryDataBody)
```
#### Struct

```go
type UpdateCategoryDataBody struct {
	Description string      `json:"description,omitempty"`
	Image       string      `json:"image,omitempty"`
	Title       string      `json:"title,omitempty"`
	CustomData  interface{} `json:"customData"`
	Order       string      `json:"order,omitempty"`
}
```


#### Body Function

For Create;

```go
slyk.UpdateCategoryDataForBody()
```

Append List;

```go
SetDescription(description string)
SetImage(image string) 
SetTitle(title string) 
SetCustomData(customData string) 
SetOrder(order string) }
```

### Create Category

Function;

```go
response,err := client.CreateCategory(*CreateCategoryDataBody)
```

#### Struct

```go
type CreateCategoryDataBody struct {
	Description string      `json:"description,omitempty"`
	Image       string      `json:"image,omitempty"`
	Title       string      `json:"title"`
	CustomData  interface{} `json:"customData"`
	Order       string      `json:"order,omitempty"`
}
```


#### Body Function

For Create;

```go
slyk.CreateCategoryDataForBody()
```

Append List;

```go
SetDescription(description string)
SetImage(image string) 
SetTitle(title string) 
SetCustomData(customData string) 
SetOrder(order string) }
```

### Category Reorder

```go
response,err := client.CategoryReorder({{CATEGORY_ID}},*CategoryReorderDataBody)
```

#### Struct

```go
type CategoryReorderDataBody struct {
	SubsequentID string `json:"subsequentId,omitempty"`
}
```

#### Body Function

For Create;

```go
slyk.CategoryReorderDataForBody()
```

Append List;

```go
SetSubsequentID(subsequentID string)
```

### Delete Category

```go
err := client.DeleteCategory({{CATEGORY_ID}})
```

