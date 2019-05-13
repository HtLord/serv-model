package servmodel

import (
	"net/url"
	"time"
)

type Member struct {
	Name string
	Email string
	PhoneNumber string
	RegisterTime time.Time
}


func MemberWrapper(form url.Values) Member {
	name := form.Get("name")
	email := form.Get("email")
	phoneNumber := form.Get("phone-number")
	now := time.Now()
	return Member{name, email, phoneNumber, now}
}

