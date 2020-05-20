package model

import (
	"fmt"
	"testing"
)
// go test
// go test -v


//func TestUser_AddUser(t *testing.T) {
func testAddUser(t *testing.T) {
	fmt.Println("test add user:")
	user := &User{}
	user.AddUser()
	user.AddUser2()
}

func testGetUsers(t *testing.T) {
	fmt.Println("test get user:")
	user := &User{}
	users, _ := user.GetUsers()
	fmt.Println(users)
}

// before test
func TestMain(m *testing.M) {
	fmt.Println("before test exc")
	m.Run()
}

// child test
func TestUser(t *testing.T)  {
	fmt.Println("test user:")
	//t.Run("test add user:", testAddUser)
	t.Run("test get users:", testGetUsers)
}