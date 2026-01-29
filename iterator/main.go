package iterator

import "fmt"

type Iterator interface {
	HasNext() bool
	Next() *User
}

type User struct {
	ID   int
	Name string
}

type UserIterator struct {
	users []*User
	index int
}

func (i *UserIterator) HasNext() bool {
	return i.index < len(i.users)
}

func (i *UserIterator) Next() *User {
	user := i.users[i.index]
	i.index++
	return user
}

func Main() {
	users := []*User{
		{ID: 1, Name: "John"},
		{ID: 2, Name: "Jane"},
		{ID: 3, Name: "Jim"},
	}
	iterator := &UserIterator{users: users, index: 0}
	for iterator.HasNext() {
		user := iterator.Next()
		fmt.Println("User: ", user.Name, " ID: ", user.ID)
	}
}
