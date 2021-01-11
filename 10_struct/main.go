package main

import (
	"10_struct/management"
	"fmt"
)

func main() {

	// instansiasi objek struct
	user := management.User{}
	user2 := management.User{}
	// cara 2 memberikan sebuah nilai properti, properti boleh tidak berurut
	user3 := management.User{ID: 111111, FirstName: "Nufika", LastName: "Fitriani", Email: "fika@gmail.com", IsActive: false}
	// memberikan nilai properti tanpa key, syarat harus berurut dengan properti nya nya
	user4 := management.User{1234, "juleha", "safitri", "juleha@gmail.com", true}

	// cara 1 memberikan sebuah nilai properti struct
	user.ID = 3216607
	user.FirstName = "Danil"
	user.LastName = "Syah"
	user.Email = "danilsyaharihardjo@gmail.com"
	user.IsActive = true

	user2.ID = 3216608
	user2.FirstName = "Haykal"
	user2.LastName = "Dafiansyah"
	user2.Email = "haykal@gmail.com"
	user2.IsActive = true

	// menampilkan seluruh properties
	fmt.Println(user)
	// // menampilkan berdasarkan properti
	fmt.Println("nama saya ", user.FirstName)
	fmt.Println("email :", user.Email)
	fmt.Println("anak saya bernama :", user2.FirstName, user2.LastName)
	fmt.Println("Istri saya bernama :", user3.FirstName)
	// struct sebagai parameter function
	displayUser1 := displayUser(user)
	displayUser2 := displayUser(user3)

	fmt.Println(displayUser1)
	fmt.Println(displayUser2)

	// menambahkan user ke dalam group struct
	userGamers := []management.User{user, user2}
	userBrowsers := []management.User{user3, user4}

	groupGamer := management.Group{"Gamer", user, userGamers, true}
	groupBrowser := management.Group{"Browser", user3, userBrowsers, false}

	displayGroup(groupGamer)
	displayGroup(groupBrowser)

	fmt.Println("=============================")
	// memanggil method User struct
	resultUser1 := user.Display()
	resultUser2 := user2.Display()
	fmt.Println(resultUser1)
	fmt.Println(resultUser2)

	fmt.Println("================================")
	// memanggil method Group struct
	groupGamer.Display()
	groupBrowser.Display()
}

// struct sebagai parameter
func displayUser(user management.User) string {
	var active string
	if user.IsActive {
		active = "aktif"
	} else {
		active = "tidak aktif"
	}
	return fmt.Sprintf("ID : %d , Name : %s %s , Email : %s, IsActive : %s ", user.ID, user.FirstName, user.LastName, user.Email, active)
}

func displayGroup(group management.Group) {
	fmt.Println("Name :", group.Name)
	fmt.Println("Admin :", group.Admin.FirstName)
	fmt.Println("Member Count :", len(group.Users), " = ")
	for _, user := range group.Users {
		fmt.Println("-> ", user.FirstName)
	}
	fmt.Println("IsAvailable : ", group.IsAvailable)
	fmt.Println("==================================")
}
