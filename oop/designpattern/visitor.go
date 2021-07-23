package designpattern

import "fmt"

type Park struct{}

func (p *Park) Accept(v IVisitor) {
	v.VisitForPark(p)
}

type Zoo struct{}

func (z *Zoo) Accept(v IVisitor) {
	v.VisitForZoo(z)
}

type AmusementPark struct{}

func (ap *AmusementPark) Accept(v IVisitor) {
	v.VisitForAmusementPark(ap)
}

type TouristSpot interface {
	Accept(v IVisitor)
}

type IVisitor interface {
	VisitForPark(*Park)
	VisitForZoo(*Zoo)
	VisitForAmusementPark(*AmusementPark)
}

type Tourist struct {
	Name   string
	Age    int
	Gender string
}

func NewTourist(name string, age int, gender string) *Tourist {
	return &Tourist{
		Name:   name,
		Age:    age,
		Gender: gender,
	}
}

func (t *Tourist) VisitForPark(p *Park) {
	fmt.Printf("Tourist named %s, with age %d and gender %s, is having a picnic in the park.\n", t.Name, t.Age, t.Gender)
}

func (t *Tourist) VisitForZoo(z *Zoo) {
	fmt.Printf("Tourist named %s, with age %d and gender %s, is observing the animals in the zoo.\n", t.Name, t.Age, t.Gender)
}

func (t *Tourist) VisitForAmusementPark(ap *AmusementPark) {
	fmt.Printf("Tourist named %s, with age %d and gender %s, is enjoying the entertainments in the amusement park.\n", t.Name, t.Age, t.Gender)
}

func Visitor() {
	tourist := NewTourist("John", 35, "male")

	var spot TouristSpot
	spot = new(Zoo)
	spot.Accept(tourist)
	spot = new(AmusementPark)
	spot.Accept(tourist)
	spot = new(Park)
	spot.Accept(tourist)
}
