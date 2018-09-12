# SplitBill

## Description

Simple Go package for calculating bill for group.

## Owner

SRE

## Contact and On-Call Information

- Des Dulianto (des.dulianto@bukalapak.com) SRE

## Prerequisite

- Go

## Build

Steps for building the library:

1. clone from git repository https://github.com/desdulianto/splitbill into $GOPATH/src directory
2. run `go build github.com/desdulianto/splitbill` to build the library
3. run `go install github.com/desdulianto/splitbill` to install the library

## Running Test

Run test using `go test github.com/desdulianto/splitbill`

## Usage

```go
import "github.com/desdulianto/splitbill"

// create bill
bill = splitbill.Bill{
    Amount: 1200000, // bill amount
    PaidBy: "Abang", // the one who paid the bill
    People: splitbill.People{"Abang", "Adek", "Kakak"}, // poeple in the group
}

// split bill evenly
eachAmount = bill.SplitEvenly() // returns 400000

// get list of people who owes money to whom paid the bill
people = bill.GetPeople() // returns ["Adek", "Kakak"]
```

## Links

- [Golang](https://tour.golang.org/basics/1)
- [Write Go Code](https://golang.org/doc/code.html)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go Code Review](https://github.com/golang/go/wiki/CodeReviewComments)