package interactor

import "fmt"

type UserIn struct {
	FirstName  string
	LastName   string
	BirthYear  int
	BirthMonth int
	BirthDay   int
}

func (i *UserIn) FormatDate() string {
	return fmt.Sprintf("%04d/%02d/%02d", i.BirthYear, i.BirthMonth, i.BirthDay)
}

type UserOut struct {
	Id         string
	FirstName  string
	LastName   string
	BirthYear  int
	BirthMonth int
	BirthDay   int
}
