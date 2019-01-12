# SendPulse REST client library
A SendPulse REST client library and example for Go (Golang).

API Documentation [https://sendpulse.com/api](https://sendpulse.com/api)

[![Build Status](https://travis-ci.org/dimuska139/sendpulse-sdk-go.svg?branch=master)](https://travis-ci.org/dimuska139/sendpulse-sdk-go)
[![codecov](https://codecov.io/gh/dimuska139/sendpulse-sdk-go/branch/master/graph/badge.svg)](https://codecov.io/gh/dimuska139/sendpulse-sdk-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/dimuska139/sendpulse-sdk-go)](https://goreportcard.com/report/github.com/dimuska139/sendpulse-sdk-go)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/dimuska139/sendpulse-sdk-go/blob/master/LICENSE)

### Download

```shell
go get -u github.com/dimuska139/sendpulse-sdk-go
```

### Example
```go
package main

import (
	"fmt"
	"github.com/dimuska139/sendpulse-sdk-go"
)
const ApiUserId = "12345"
const ApiSecret = "12345"
const ApiTimeout = 5

func main() {
	addressBookId := 12345

	client, e := sendpulse.ApiClient(ApiUserId, ApiSecret, ApiTimeout)
	if e != nil {
		switch err := e.(type) {
		case *sendpulse.ResponseError: // Http error
			fmt.Println(err.HttpCode)
			fmt.Println(err.Url)
			fmt.Println(err.Body)
		default: // Another errors
			fmt.Println(e)
		}
	}

	// Get address book info by id
	bookInfo, e := client.Books.Get(uint(addressBookId))
	if e != nil {
		switch err := e.(type) {
		case *sendpulse.ResponseError: // Http error
			fmt.Println(err.HttpCode)
			fmt.Println(err.Url)
			fmt.Println(err.Body)
		default: // Another errors
			fmt.Println(e)
		}
	} else {
		fmt.Println(*bookInfo)
	}

	// Get address books list
	limit := 10
	offset := 20
	books, err := client.Books.List(uint(limit), uint(offset))
	if err != nil {
		switch err := e.(type) {
		case *sendpulse.ResponseError: // Http error
			fmt.Println(err.HttpCode)
			fmt.Println(err.Url)
			fmt.Println(err.Body)
		default: // Another errors
			fmt.Println(e)
		}
	} else {
		fmt.Println(*books)
	}

	// Add emails to address book
	emails := []sendpulse.Email{
		sendpulse.Email{
			Email:     "alex@test.net",
			Variables: map[string]string{
				"name": "Alex",
				"age": "25",
			},
		},
		sendpulse.Email{
			Email:     "dima@test.net",
			Variables: make(map[string]string),
		},
	}
	
	extraParams := make(map[string]string)
	
	err = client.Books.AddEmails(uint(addressBookId), emails, extraParams)
	if err != nil {
		switch err := e.(type) {
		case *sendpulse.ResponseError: // Http error
			fmt.Println(err.HttpCode)
			fmt.Println(err.Url)
			fmt.Println(err.Body)
		default: // Another errors
			fmt.Println(e)
		}
	}
}
```
