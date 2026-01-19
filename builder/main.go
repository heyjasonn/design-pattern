package builder

import (
	"encoding/json"
	"fmt"
)

func Main() {
	user, err := NewUserBuilder().
		Name("John Doe").
		Age("30").
		Email("john.doe@example.com").
		Password("password").
		Address("123 Main St").
		Phone("123-456-7890").
		Gender("Male").
		Role("Admin").
		IsActive(true).
		Build()
	if err != nil {
		panic(err)
	}

	b, _ := json.Marshal(user)
	fmt.Println(string(b))
}
