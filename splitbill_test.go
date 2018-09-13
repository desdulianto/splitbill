package splitbill

import (
	"errors"
	"testing"
)

type Want struct {
	money  Money
	people People
}

// contains returns true if person is included in these people, false otherwise
func (people People) contains(person Person) bool {
	found := false
	for _, p := range people {
		if person == p {
			found = true
			break
		}
	}

	return found
}

// ContainsAll returns true if all otherPeople is same with these people, false otherwise
func (people People) containsAll(otherPeople People) bool {
	found := true
	for _, p := range otherPeople {
		found = found && people.contains(p)
	}
	return found
}

func TestSplitBill(t *testing.T) {
	cases := []struct {
		in   Bill
		want Want
	}{
		{
			Bill{100000, "A", People{"A", "B", "C", "D", "E"}},
			Want{Money(20000), People{"B", "C", "D", "E"}},
		},
		{
			Bill{100000, "A", People{"A", "B", "C"}},
			Want{Money(33333), People{"B", "C"}},
		},
		{
			Bill{200000, "", People{"A", "B", "C"}},
			Want{Money(66666), People{"A", "B", "C"}},
		},
	}

	for _, c := range cases {
		money, ok := c.in.SplitEvenly()
		if ok == nil {
			if money != c.want.money {
				t.Errorf("{%v}.SplitEvenly() == %v, want %v", c.in, money, c.want.money)
			}
		}

		people := c.in.GetPeople()

		if !c.want.people.containsAll(people) {
			t.Errorf("{%v}.GetPeople() == %v, want %v", c.in, people, c.want.people)
		}
	}
}

func TestWithEmptyPeople(t *testing.T) {
	bill := Bill{1000000, "", People{}}

	_, ok := bill.SplitEvenly()
	if ok == nil {
		t.Errorf("{%v}.SplitEvenly() should error with %v", bill, "No people in group")
	}

	people := bill.GetPeople()
	if people == nil || len(people) != 0 {
		t.Errorf("{%v}.GetPeople() should return with empty list, got %v", bill, people)
	}
}

func TestWithOnlyOnePerson(t *testing.T) {
	cases := []struct {
		in   Bill
		want Want
	}{
		{
			in:   Bill{100000, "A", People{"A"}},
			want: Want{Money(100000), People{}},
		},
		{
			in:   Bill{100000, "A", People{"B"}},
			want: Want{Money(100000), People{"B"}},
		},
	}

	for _, c := range cases {
		money, err := c.in.SplitEvenly()
		switch {
		case err != nil:
			t.Errorf("{%v}.SplitEvenly() should not error, got %v", c.in, err.Error())
		case c.want.money != money:
			t.Errorf("{%v}.SplitEvenly() should return %v, got %v", c.in, c.want.money, money)
		}

		people := c.in.GetPeople()
		if !c.want.people.containsAll(people) {
			t.Errorf("{%v}.GetPeople() should return %v, got %v", c.in, c.want.people, people)
		}
	}
}

func TestSplitBillError(t *testing.T) {
	cases := []struct {
		in   Bill
		want error
	}{
		{Bill{100000, "A", People{"A", "B", "C", "D", "E"}}, nil},
		{Bill{0, "A", People{"A", "B", "C", "D", "E"}}, errors.New("Amount must not zero or negative")},
		{Bill{-1, "A", People{"A", "B", "C", "D", "E"}}, errors.New("Amount must not zero or negative")},
	}

	for _, c := range cases {
		_, ok := c.in.SplitEvenly()
		if (ok == nil && ok != c.want) && ok.Error() != c.want.Error() {
			t.Errorf("{%v}.evenSplit() failed with %v error expected %v error", c.in, ok, c.want)
		}
	}
}
