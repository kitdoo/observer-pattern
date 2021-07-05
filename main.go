package main

func main() {
	bt21Toy := newItem("BT21 RJ pillow")

	armyFirst := &customer{email: "namjoons.wife@gamil.com"}
	armySecond := &customer{email: "army0613@gamil.com"}

	bt21Toy.subscribe(armyFirst)
	bt21Toy.subscribe(armySecond)

	bt21Toy.updateAvailability()
}
