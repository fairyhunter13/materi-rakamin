package designpattern

import "fmt"

type Weapon interface {
	Attack()
}

type Damage struct {
	Hit  int
	Crit int
}

func (d *Damage) Status() string {
	return fmt.Sprintf("hit: %d, critical: %d", d.Hit, d.Crit)
}

type AK47 struct {
	Damage
}

func (ak *AK47) Attack() {
	fmt.Printf("AK47 attacks, %s.\n", ak.Status())
}

type Sword struct {
	Damage
}

func (s *Sword) Attack() {
	fmt.Printf("Sword attacks, %s.\n", s.Status())
}

type Bow struct {
	Damage
}

func (b *Bow) Attack() {
	fmt.Printf("Bow attacks, %s.\n", b.Status())
}

func Factory() {
	weapon := getWeapon("ak47")
	attack(weapon)

	weapon = getWeapon("bow")
	attack(weapon)
}

func getWeapon(weaponType string) (weapon Weapon) {
	switch weaponType {
	case "ak47":
		weapon = &AK47{
			Damage: Damage{
				Hit:  45,
				Crit: 75,
			},
		}
	case "sword":
		weapon = &Sword{
			Damage: Damage{
				Hit:  65,
				Crit: 125,
			},
		}
	default:
		weapon = &Bow{
			Damage: Damage{
				Hit:  30,
				Crit: 65,
			},
		}
	}
	return
}

func attack(weapon Weapon) {
	weapon.Attack()
}
