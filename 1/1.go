package main

// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

import (
	"fmt"
)

type Human struct {
	Name   string
	Age    int
	Gender string
}

// конструктор Human
func NewHuman(name string, age int, gender string) *Human {
	return &Human{Name: name, Age: age, Gender: gender}
}

func (h Human) String() string {
	return fmt.Sprintf("Name: %s, Age: %d, Gender: %s", h.Name, h.Age, h.Gender)
}

func (h *Human) SetName(name string) {
	h.Name = name
}

func (h *Human) SetAge(age int) error {
	if age < 0 {
		return fmt.Errorf("age must be a positive integer")
	}
	h.Age = age
	return nil
}

type Action struct {
	Name     string
	Duration int
	Human
}

// конструктор Action
func NewAction(name string, duration int, human Human) *Action {
	return &Action{Name: name, Duration: duration, Human: human}
}

func (a Action) String() string {
	return fmt.Sprintf("Name: %s, Duration: %d, Human: %s", a.Name, a.Duration, a.Human)
}

func (a *Action) SetName(name string) {
	a.Name = name
}

func (a *Action) SetAge(age int) error {
	if age < 0 {
		return fmt.Errorf("age must be a positive integer")
	}
	a.Human.Age = age
	return nil
}

func (a *Action) SetDuration(d int) error {
	if d < 0 {
		return fmt.Errorf("duration must be a positive integer")
	}
	a.Duration = d
	return nil
}

func main() {
	// создаем экземпляр структуры Human
	h := NewHuman("John Doe", 30, "Male")
	fmt.Println(h)

	// меняем поле age экземпляра h структуры Human
	err := h.SetAge(31)
	if err != nil {
		fmt.Println(err)
	}

	// меняем поле name экземпляра h структуры Human
	h.SetName("Mark Down")
	fmt.Println(h)

	// создаем экземпляр структуры Action
	a := NewAction("Running", 50, *h)
	fmt.Println(a)

	// меняем поле Duration экземпляра a структуры Action
	err = a.SetDuration(60)
	if err != nil {
		fmt.Println(err)
	}

	// напрямую меняем поле name экземпляра a структуры Action
	a.Name = "Swimming"
	// видно, что изменения не коснулись вложенной структуры Human
	fmt.Println(a)

	// меняем поле name экземпляра a структуры Action
	a.SetName("Walking")
	// аналогично видим, что изменения не коснулись полей вложенной структуры
	fmt.Println(a)

	// изменяем поле name структуры Human внутри Action
	a.Human.SetName("John Smith")
	a.SetAge(32)

	fmt.Println(a)

	// коллизии не происходит
	fmt.Println(h)

	// изменяем поле name структуры Human внутри Action
	a.Human.SetAge(33)

	fmt.Println(a)

	// видим, что хотя мы и присвоили h экземпляру Action,
	// изменнеия внутри вложенной структуры Human не повлияли
	// на внешний экземпляр Human,
	// так как поле Human внутри структуры Action
	// не является указателем на внешний экземпляр Human
	fmt.Println(h)
}
