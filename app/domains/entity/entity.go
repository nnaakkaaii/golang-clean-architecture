package entity

func NewUser(id, firstName, lastName string, birthDate Date) *User {
	return &User{
		id:        id,
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
	}
}

type User struct {
	id        string
	firstName string
	lastName  string
	birthDate Date
}

func NewDate(year, month, day int) *Date {
	return &Date{
		year:  year,
		month: month,
		day:   day,
	}
}

type Date struct {
	year  int
	month int
	day   int
}

func (u *User) GetId() string {
	return u.id
}

func (u *User) GetFirstName() string {
	return u.firstName
}

func (u *User) GetLastName() string {
	return u.lastName
}

func (u *User) GetBirthDate() *Date {
	return &u.birthDate
}

func (d *Date) GetYear() int {
	return d.year
}

func (d *Date) GetMonth() int {
	return d.month
}

func (d *Date) GetDay() int {
	return d.day
}
