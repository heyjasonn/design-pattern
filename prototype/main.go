package prototype

type Prototype interface {
	Clone() Prototype
}

type User struct {
	ID   string
	Name string
}

func (p *User) Clone() Prototype {
	return &User{
		ID:   p.ID,
		Name: p.Name,
	}
}
