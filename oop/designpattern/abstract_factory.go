package designpattern

import "fmt"

type IBlacksmith interface {
	makeSword() ISword
}

type ISword interface {
	SetSheath(sheath string)
	SetCurve(curveness int)
	SetHolder(holder string)
	Status()
}

type modernBlacksmith struct{}

func (mb *modernBlacksmith) makeSword() ISword {
	return new(claymore)
}

type sword struct {
	sheath    string
	curveness int
	holder    string
}

type claymore struct {
	sword
}

func (c *claymore) SetSheath(sheath string) {
	c.sheath = sheath
}

func (c *claymore) SetCurve(curveness int) {
	c.curveness = curveness
}

func (c *claymore) SetHolder(holder string) {
	c.holder = holder
}

func (c *claymore) Status() {
	fmt.Printf("Claymore status: \nsheath: %s\ncurveness: %d\nholder: %s\n-----------------------------\n", c.sheath, c.curveness, c.holder)
}

type classicBlacksmith struct{}

func (cb *classicBlacksmith) makeSword() ISword {
	return new(katana)
}

type katana struct {
	sword
}

func (k *katana) SetSheath(sheath string) {
	k.sheath = sheath
}

func (k *katana) SetCurve(curveness int) {
	k.curveness = curveness
}

func (k *katana) SetHolder(holder string) {
	k.holder = holder
}

func (k *katana) Status() {
	fmt.Printf("Katana status: \nsheath: %s\ncurveness: %d\nholder: %s\n-----------------------------\n", k.sheath, k.curveness, k.holder)
}

func getBlacksmith(kind string) IBlacksmith {
	if kind == "modern" {
		return &modernBlacksmith{}
	}
	return &classicBlacksmith{}
}

func AbstractFactory() {
	modernBs := getBlacksmith("modern")
	classicBs := getBlacksmith("classic")

	claymore := modernBs.makeSword()
	katana := classicBs.makeSword()

	claymore.SetCurve(15)
	claymore.SetHolder("iron")
	claymore.SetSheath("leather")

	katana.SetCurve(35)
	katana.SetHolder("wood")
	katana.SetSheath("bamboo")

	printSwordDetails(claymore)
	printSwordDetails(katana)
}

func printSwordDetails(sword ISword) {
	sword.Status()
}
