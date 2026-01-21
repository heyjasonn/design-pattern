package composite

import (
	"context"
	"fmt"
)

type Rule interface {
	Eval(ctx context.Context) bool
}

type AdminRule struct {
	role string
}

func (r *AdminRule) Eval(ctx context.Context) bool {
	return r.role == "admin"
}

type AgeRule struct {
	age int
}

func (r *AgeRule) Eval(ctx context.Context) bool {
	return r.age >= 18
}

type AndRule struct {
	rules []Rule
}

func (r *AndRule) Eval(ctx context.Context) bool {
	for _, rule := range r.rules {
		if !rule.Eval(ctx) {
			return false
		}
	}
	return true
}

type OrRule struct {
	rules []Rule
}

func (r *OrRule) Eval(ctx context.Context) bool {
	for _, rule := range r.rules {
		if rule.Eval(ctx) {
			return true
		}
	}
	return false
}

func Main() {
	rule := &AndRule{
		rules: []Rule{
			&AdminRule{role: "admin"},
			&OrRule{rules: []Rule{
				&AgeRule{age: 19},
				&AgeRule{age: 60},
			}},
		},
	}

	result := rule.Eval(context.Background())
	fmt.Println(result)
}
