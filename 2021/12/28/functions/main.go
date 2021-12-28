package main

type User struct {
	Name      string
	BirthYear uint
}

func main() {
	user := User{
		Name:      "mario",
		BirthYear: 1900,
	}
}

// BirthYear ...
func (u User) GetBirthYear() uint {
	return u.BirthYear
}

// SetBirthYear ...
func (u *User) SetBirthYear(birthYear uint) {
	u.BirthYear = birthYear
}
