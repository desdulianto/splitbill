// Package splitbill defines types and functions for calculating bill
package splitbill

import "errors"

type Money int

type Person string

type People []Person

type Bill struct {
	Amount Money  // total amount including tax (if applicable)
	PaidBy Person // person who paid the bill
	People People // people in the group
}

// SplitEvenly, split the bill evenly by people counts (using integer division)
// returns bill splitted evenly for each person
func (bill Bill) SplitEvenly() (Money, error) {
	if bill.Amount <= Money(0) {
		return Money(0), errors.New("Amount must not zero or negative")
	}
	return bill.Amount / Money(len(bill.People)), nil
}

// GetPeople returns all the people names in the group excluding paidBy
func (bill Bill) GetPeople() People {
	var people People

	for _, person := range bill.People {
		if bill.PaidBy != person {
			people = append(people, person)
		}
	}

	return people
}
