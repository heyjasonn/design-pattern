package builder

import "errors"

type User struct {
	Name     string
	Age      string
	Email    string
	Password string
	Address  string
	Phone    string
	Gender   string
	Role     string
	IsActive bool
}

type UserBuilder struct {
	user *User
}

func NewUserBuilder() *UserBuilder {
	return &UserBuilder{
		user: &User{},
	}
}

func (b *UserBuilder) Name(name string) *UserBuilder {
	b.user.Name = name
	return b
}

func (b *UserBuilder) Age(age string) *UserBuilder {
	b.user.Age = age
	return b
}

func (b *UserBuilder) Email(email string) *UserBuilder {
	b.user.Email = email
	return b
}

func (b *UserBuilder) Password(password string) *UserBuilder {
	b.user.Password = password
	return b
}

func (b *UserBuilder) Address(address string) *UserBuilder {
	b.user.Address = address
	return b
}

func (b *UserBuilder) Phone(phone string) *UserBuilder {
	b.user.Phone = phone
	return b
}

func (b *UserBuilder) Gender(gender string) *UserBuilder {
	b.user.Gender = gender
	return b
}

func (b *UserBuilder) Role(role string) *UserBuilder {
	b.user.Role = role
	return b
}

func (b *UserBuilder) IsActive(isActive bool) *UserBuilder {
	b.user.IsActive = isActive
	return b
}

func (b *UserBuilder) Build() (*User, error) {
	if b.user.Name == "" {
		return nil, errors.New("name is required")
	}

	return b.user, nil
}
