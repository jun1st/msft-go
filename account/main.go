package main

type Changeable interface {
	ChangeName()
}

type Account struct {
	FirstName, LastName string
}
