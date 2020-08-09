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

## User

### Get User

GetUser is returns the Slyk user list. This function returns user information and error. 

if you do not use a filter, '30' users' information is received.
```go
response,err:=client.GetUsers()
```

#### Filter

Example Using:
```go
response,err:=client.GetUsers(slyk.GetUserFilter().SetEmail("mert@monoji.com").SetName("Mert"))
```
