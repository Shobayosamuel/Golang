// package main

// import (
// 	"fmt"
// 	// "egg.com/events/articles"
// )

// type User struct{
// 	ID int
// 	FirstName string
// 	LastName string
// 	Email string
// }

// type Group struct {
// 	role string
// 	users []User
// 	newestUser User
// 	spaceAvailable bool

// }

// func describeUser(u User) string {
// 	desc := fmt.Sprintf("Name: %s %s, Email: %s", u.FirstName, u.LastName, u.Email)
// 	return desc
// }

// func describeGroup(new Group) string {
// 	if len(new.users) >= 2 {
// 		new.spaceAvailable = false
// 	}
// 	ans := fmt.Sprintf("This user group has %d users. The newest user is %s %s. Accepting new users: %t", len(new.users), new.newestUser.FirstName, new.newestUser.LastName, new.spaceAvailable)
// 	return ans
// }

// func main() {
// 	u := User{2, "Shobs", "Samuel", "Shobayo@gmail.com"}
// 	u2 := User{1, "Peter", "Daniel", "danny@gmail.com"}
// 	g := Group{
// 		role: "Admin",
// 		users: []User{u, u2},
// 		newestUser: u2,
// 		spaceAvailable: true,
// 	}
// 	fmt.Println(describeUser(u))
// 	fmt.Println(describeGroup(g))
// 	fmt.Println(g)
// }