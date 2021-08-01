package oop

type state interface {
	Progress()
	Submit()
	Approve()
	Reject()
	Publish()
}

type dummy struct{}

func (d *dummy) Progress() {}
func (d *dummy) Submit()   {}
func (d *dummy) Approve()  {}
func (d *dummy) Reject()   {}
func (d *dummy) Publish()  {}

type (
	// Add the logic in the statuses below
	Draft struct {
		dummy
		doc *Document
	}
	Completed struct {
		dummy
		doc *Document
	}
	Submitted struct {
		dummy
		doc *Document
	}
	Approved struct {
		dummy
		doc *Document
	}
	Rejected struct {
		dummy
		doc *Document
	}
	Published struct {
		dummy
		doc *Document
	}
)

type Document struct {
	state
	stateMap map[string]state
}

func NewDocument() *Document {
	doc := &Document{}
	doc.stateMap = map[string]state{
		"draft":     &Draft{doc: doc},
		"completed": &Completed{doc: doc},
		"submitted": &Submitted{doc: doc},
		"approved":  &Approved{doc: doc},
		"rejected":  &Rejected{doc: doc},
		"published": &Published{doc: doc},
	}
	return doc
}
func (d *Document) GetStateFromMap(input string) state { return d.stateMap[input] }
func (d *Document) GetState() state                    { return d.state }
func (d *Document) SetState(s state)                   { d.state = s }

// Complete the logic in the functions below
func (d *Document) Progress() {}
func (d *Document) Submit()   {}
func (d *Document) Approve()  {}
func (d *Document) Reject()   {}
func (d *Document) Publish()  {}
