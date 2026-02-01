package visitor

import "fmt"

type PermissionNode interface {
	Accept(v PermissionVisitor)
}

type Role struct {
	Name string
}

func (r *Role) Accept(v PermissionVisitor) {
	v.VisitRole(r)
}

type PermissionVisitor interface {
	VisitRole(r *Role)
}

type Admin struct {
	Allowed bool
}

func (a *Admin) VisitRole(r *Role) {
	a.Allowed = r.Name == "admin"
}

func Main() {
	role := &Role{Name: "admin"}
	admin := &Admin{}
	role.Accept(admin)
	fmt.Println(admin.Allowed)
}
