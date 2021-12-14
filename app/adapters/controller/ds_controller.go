package controller

import (
	"clean_architecture/app/domains/entity"
	"time"
)

type requestBody struct {
	FirstName string    `json:"first_name" form:"first_name"`
	LastName  string    `json:"last_name" form:"last_name"`
	BirthDate time.Time `json:"birth_date" form:"birth_name"`
}

func (r *requestBody) convert() (user *entity.User) {
	date := *entity.NewDate(r.BirthDate.Year(), int(r.BirthDate.Month()), r.BirthDate.Day())
	user = entity.NewUser("", r.FirstName, r.LastName, date)
	return
}
