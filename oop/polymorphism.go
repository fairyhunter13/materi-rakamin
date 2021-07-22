package main

import "fmt"

type Animal struct {
	Name     string
	Gender   string
	IsMammal bool
}

func (a *Animal) Speak() {
	var text string

	text = fmt.Sprintf("I am %s. My gender is %s. I am not a mammal.", a.Name, a.Gender)
	if a.IsMammal {
		text = fmt.Sprintf("I am %s. My gender is %s. I am a mammal.", a.Name, a.Gender)
	}
	fmt.Println(text)
}

type Cat struct {
	Animal
	Breed string
}

func (a *Cat) Speak() {
	a.Animal.Speak()
	fmt.Printf("My breed is %s.\n", a.Breed)
}

func polymorphism() {
	animalDiona := Animal{
		Name:     "Diona",
		Gender:   "female",
		IsMammal: true,
	}
	animalDiona.Speak()

	catKecing := Cat{
		Animal: Animal{
			Name:     "Kecing",
			Gender:   "female",
			IsMammal: true,
		},
		Breed: "Anggora",
	}
	catKecing.Speak()
}
