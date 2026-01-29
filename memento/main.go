package memento

import "fmt"

type Saga struct {
	Step int
}

func (s *Saga) Save() int {
	return s.Step
}

func (s *Saga) Restore(step int) {
	s.Step = step
}

func step1(s *Saga) error {
	fmt.Println("Create Order")
	s.Step = 1
	return nil
}

func step2(s *Saga) error {
	fmt.Println("Pay Order")
	s.Step = 2
	return fmt.Errorf("payment failed")
}

func step3(s *Saga) error {
	fmt.Println("Ship Order")
	s.Step = 3
	return nil
}

func Main() {
	saga := &Saga{Step: 0}
	history := []int{}

	run := func(fn func(*Saga) error) error {
		history = append(history, saga.Save())
		return fn(saga)
	}

	if err := run(step1); err != nil {
		return
	}
	if err := run(step2); err != nil {
		fmt.Println("ERROR → rollback")
		last := history[len(history)-1]
		saga.Restore(last)
		fmt.Println("Rollback to step:", saga.Step)
		return
	}
	if err := run(step3); err != nil {
		return
	}
}
