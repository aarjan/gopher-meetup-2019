package main

import (
	"fmt"
	"log"
)

type Notifier interface {
	Notify() error
}

type User struct {
	name  string
	email string
}

func (u *User) Notify() error {
	if u.name == "" || u.email == "" {
		return fmt.Errorf("err : name or email is nil")
	}
	fmt.Printf("The name is: %s and the email is: %s \n", u.name, u.email)
	return nil
}

func SendNotification(n Notifier) {
	err := n.Notify()
	if err != nil {
		log.Println(err)
	}
}

type Admin struct {
	*User
	level string
}

func (a *Admin) Notify() error {
	if a.name == "" || a.email == "" || a.level == "" {
		return fmt.Errorf("err : name or email or level is nil")
	}
	fmt.Printf("The name is: %s, with email: %s and level: %s", a.name, a.email, a.level)
	return nil

}
func main() {
	u := &User{"aarjan", "aarjan.baskota@dragonlaw.hk"}
	SendNotification(u)

	a := &Admin{u, ""}
	SendNotification(a)
	a.User.Notify()
}
