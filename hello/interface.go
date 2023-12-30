package main

import (
	"fmt"
)

type Shooter interface {
	Shoot()
}

type NPC struct{}

func (n NPC) Shoot() {
	fmt.Println("npc is shooting")
}

func (n NPC) Stab() {
	fmt.Println("npc is stabbed")
}

type Player struct{}

func (p Player) Shoot() {
	fmt.Println("Player is shooting")
}

func (p Player) Stab() {
	fmt.Println("Player is stabbed")
}

func main() {
	var s Shooter

	s = NPC{} // So everything can be implemented by shooter can be assigned to s

	s.Shoot()
	s = Player{}
	s.Shoot()
	// s.Stab()         // Doesn't work because Shooter doesn't have the method Stab which can be implemented
}
