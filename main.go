package main

import "fmt"

// func main() {
// 	fmt.Println("Hi!")
// }

// type contact struct {
// 	email string
// 	zip   int
// }
// type person struct {
// 	name        string
// 	age         string
// 	contactinfo contact
// }

func main() {

	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf745",
	// }

	colors := make(map[string]string)
	colors[10] = "#fff"

	delete(colors, 10)

	fmt.Println(colors)
	// test := person{name: "chand", age: "tush"}

	// fmt.Println(test)

	// var test person

	// test.name = "testm"
	// fmt.Println(test)

	// fmt.Printf("%+v", test)

	// test := person{
	// 	name: "u",
	// 	age:  "23",
	// 	contactinfo: contact{
	// 		email: "xyz@gmail.com",
	// 		zip:   123,
	// 	},
	// }

	// fmt.Println("%+v", test)
	// testPointer := &test
	// testPointer.updatName("testnew")
	// test.print()
	// var card string = "Ace of Cards"
	// cards := deck{"ace", newCard()}
	// cards = append(cards, "six")
	// cards.print()
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	// cards := newDeck()
	// // fmt.Println(cards.toString())
	// cards.saveToFile("mycards")

	// cards := newDeckFromFile("mycards")
	// cards.print()
	// cards := newDeck()
	// cards.shuffle()
	// cards.print()
	// hand, remainingCards := deal(cards, 5)
	// hand.print()

	// remainingCards.print()
	// test := "Hello the"

	// fmt.Println([]byte(test))
}

// func newCard() string {
// 	return "new cards"
// }
//pass by value
// func (pointerToPerson *person) updatName(namenew string) {

// 	(*pointerToPerson).name = namenew
// }

// func (p person) print() {
// 	fmt.Printf("%+v", p)
// }

func printf(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
