package main

import (
	"testing"
)

type People interface{ Rename(name string) }
type Teacher struct{ Name string }
type Stu struct{ Name string }

func (t Teacher) Rename(name string) {
	t.Name = name
}

func (s *Stu) Rename(name string) {
	s.Name = name
}

func changeName(p People) {
	p.Rename("Bob")
}

func Test_ValueBind(t *testing.T) {
	teacher := Teacher{Name: "CP"}

	changeName(teacher)
	if teacher.Name != "CP" {
		t.Errorf("expect CP, actual %s", teacher.Name)
	}

	changeName(&teacher)
	if teacher.Name != "CP" {
		t.Errorf("expect CP, actual %s", teacher.Name)
	}
}

func Test_PointerBind(t *testing.T) {
	stu := Stu{Name: "CP"}

	// changeName(stu) error

	changeName(&stu)
	if stu.Name != "Bob" {
		t.Errorf("expect Bob, actual %s", stu.Name)
	}
}
