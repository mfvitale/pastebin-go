[![Go](https://github.com/mfvitale/pastebin-go/actions/workflows/go.yml/badge.svg)](https://github.com/mfvitale/pastebin-go/actions/workflows/go.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/mfvitale/pastebin-go.svg)](https://pkg.go.dev/github.com/mfvitale/pastebin-go)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/mfvitale/pastebin-go?style=flat-square)
![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/mfvitale/pastebin-go?style=flat-square)
![GitHub](https://img.shields.io/github/license/mfvitale/pastebin-go?style=flat-square)

# What is pastebin-go?
pastebin-go is a Go library to interact with [Pastebin](https://pastebin.com/)

## Functions
* Create a paste (guest, logged user)
* List Pastes Created By A User
* Delete A Paste Created By A User
* Get raw paste output of users' pastes including 'private' pastes
* Get raw paste output of any 'public' & 'unlisted' pastes

# Getting Started 

## Installing

To start using pastebin-go run:

```shell
go get -u https://github.com/mfvitale/pastebin-go
```

## Instantiate anonymous client

With this client you need only the API_DEV_KEY but you can only create 'Guest' paste.

```
package main

import "https://github.com/mfvitale/pastebin-go"

func main() {
	
    client := pastebin.AnonymousClient(API_DEV_KEY)
}
```

## User client

With this client you need:
* API_DEV_KEY 
* username
* password

but in this case you have full functions.

```
package main

import "github.com/mfvitale/pastebin-go"

func main() {
	
    client := pastebin.Client(API_DEV_KEY, USERNAME, PASSWORD)
}
```

## Create a Paste

If you instantiate an anonymous client the 'CreatePaste' function will create a 'Guest' paste otherwise it will create a past for the logged user. 

```
package main

import (
    "fmt"
    "github.com/mfvitale/pastebin-go"
)

func main() {
	
    client := pastebin.Client(API_DEV_KEY, USERNAME, PASSWORD)

    paste := model.FullPaste("Test paste bin", model.Public, "npm run", "10M", "bash")

    pasteLink := client.CreatePaste(paste)

    fmt.Println(pasteLink)
}
```
output:
```shell
https://pastebin.com/<paste_code>
```

## List pastes of logged user

This will return a list of paste of the logged user.

```
package main

import (
    "fmt"
    "github.com/mfvitale/pastebin-go"
)

func main() {
	
    client := pastebin.Client(API_DEV_KEY, USERNAME, PASSWORD)

    fmt.Println(client.GetPastes())
}
```

output:
```shell
[{5xM0VBj1 1651314995 Untitled 7 0 0 Bash bash https://pastebin.com/5xM0VBj1 17}]
```
