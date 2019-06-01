package main

import "fmt"

type Speaker interface {
	Speak()
}

type SpeakerPrefixer struct {
	Speaker
}

func (sp SpeakerPrefixer) Speak() {
	fmt.Print("Prefix: ")
	sp.Speaker.Speak()
}

type Cat struct {
}

func (c Cat) Speak() {
	fmt.Println("Meow")
}

type Dog struct {
}

type Husky struct {
	Speaker
}

func (d Dog) Speak() {
	fmt.Println("woof")
}

func main() {
	// h := Dog{}
	// h := Husky{Dog{}}
	// h := Husky{Cat{}}
	h := Husky{SpeakerPrefixer{Cat{}}}
	h.Speak()
}
