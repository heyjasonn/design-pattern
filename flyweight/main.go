package flyweight

import (
	"context"
	"fmt"
)

type Rule interface {
	Eval(ctx context.Context, code string) bool
}

type Permission struct {
	Code string
}

func (p *Permission) Eval(ctx context.Context, code string) bool {
	return p.Code == code
}

type Role struct {
	rules []Rule
}

func (r *Role) Eval(ctx context.Context, code string) bool {
	for _, rule := range r.rules {
		if !rule.Eval(ctx, code) {
			return false
		}
	}
	return true
}

type RoleRegistry struct {
	cache map[string]*Role
}

func NewRoleRegistry() *RoleRegistry {
	return &RoleRegistry{
		cache: make(map[string]*Role),
	}
}

func (r *RoleRegistry) GetRole(code string) *Role {
	if role, ok := r.cache[code]; ok {
		return role
	}

	role := &Role{
		rules: []Rule{&Permission{Code: code}},
	}
	r.cache[code] = role
	return role
}

func Main() {
	registry := NewRoleRegistry()
	role := registry.GetRole("admin")
	fmt.Println(role.Eval(context.Background(), "user"))
}
