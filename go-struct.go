package main

type user struct {
	firstName    string
	lastName     string
	emailAddress string
	password     string
}

func main() {
	// --------------------

	// Different ways of defining structs

	// Key value pairs - doesn't need to be in order
	aUser := user{
		firstName:    "A",
		lastName:     "B",
		password:     "password",
		emailAddress: "admin",
	}

	// Order necessary
	bUser := user{"A", "B", "admin", "password"}

	// --------------------
}
