package management

import "fmt"

// struct di ibaratkan seperti model , sebuah object mengandung properties
type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

// Method struct User
func (user User) Display() string {
	return fmt.Sprintf("Name : %s %s, Email : %s", user.FirstName, user.LastName, user.Email)
}

// embedded struct / struct di dalam struct
type Group struct {
	Name        string
	Admin       User
	Users       []User
	IsAvailable bool
}

// method Group struct
func (group Group) Display() {
	fmt.Printf("Name Group : %s ", group.Name)
	fmt.Println("")
	fmt.Printf("Admin : %s %s", group.Admin.FirstName, group.Admin.LastName)
	fmt.Println("")
	fmt.Printf("Member count : %d", len(group.Users))
	fmt.Println("")
	fmt.Printf("Users Member : ")
	fmt.Println("")
	for _, user := range group.Users {
		fmt.Println("-> ", user.FirstName)
	}
}
