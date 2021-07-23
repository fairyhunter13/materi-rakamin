package designpattern

import "fmt"

type state interface {
	turnOn()
	login(user, password string)
	playGames()
	turnOff()
}

const (
	StrNotImplemented = "Error not implemented!"
)

type baseState struct{}

func (bs *baseState) turnOn() {
	fmt.Println(StrNotImplemented)
}

func (bs *baseState) login(username, password string) {
	fmt.Println(StrNotImplemented)
}

func (bs *baseState) playGames() {
	fmt.Println(StrNotImplemented)
}

func (bs *baseState) turnOff() {
	fmt.Println(StrNotImplemented)
}

type turnedOff struct {
	baseState
	computer *computer
}

func (to *turnedOff) turnOn() {
	fmt.Println("Turn on the computer!")
	to.computer.setState(to.computer.getState("booted"))
}

type booted struct {
	baseState
	computer *computer
}

func (b *booted) login(username, password string) {
	if !b.computer.isLoginValid(username, password) {
		fmt.Printf("Error login for the username: %s and password: %s.\n", username, password)
		return
	}

	fmt.Printf("Succeed in login to the computer with username: %s.\n", username)
	b.computer.setState(b.computer.getState("ready"))
}

type ready struct {
	baseState
	computer *computer
}

func (r *ready) playGames() {
	fmt.Println("Playing all games in the computer!")
}

func (r *ready) turnOff() {
	fmt.Println("Turning off the computer.")
	r.computer.setState(r.computer.getState("turnedOff"))
}

type computer struct {
	stateColl map[string]state
	state

	username string
	password string
}

func newComputer() (c *computer) {
	c = new(computer)
	baseState := new(baseState)
	stateColl := map[string]state{
		"turnedOff": &turnedOff{
			baseState: *baseState,
			computer:  c,
		},
		"booted": &booted{
			baseState: *baseState,
			computer:  c,
		},
		"ready": &ready{
			baseState: *baseState,
			computer:  c,
		},
	}
	c.stateColl = stateColl
	c.username = "hello"
	c.password = "test"
	c.setState(c.getState("turnedOff"))
	return
}

func (c *computer) getState(stateStr string) (s state) {
	s = c.stateColl[stateStr]
	return
}

func (c *computer) isLoginValid(username, password string) bool {
	return c.username == username && c.password == password
}

func (c *computer) turnOn() {
	c.state.turnOn()
}

func (c *computer) playGames() {
	c.state.playGames()
}

func (c *computer) login(username, password string) {
	c.state.login(username, password)
}

func (c *computer) turnOff() {
	c.state.turnOff()
}

func (c *computer) setState(s state) {
	c.state = s
}

func State() {
	computer := newComputer()

	computer.turnOn()
	computer.login("hello", "test")
	computer.playGames()
	computer.turnOff()
}
