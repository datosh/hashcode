package main

import (
	"log"
	"os"

	onepizza "github.com/datosh/hashcode/onepizza"
)

func main() {
	content, err := os.ReadFile("./one_pizza_input/a_an_example.in.txt")
	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}
	preferences := onepizza.Parse(string(content))
	log.Printf("%v", preferences)
	allIngredients := preferences.AllIngredients()
	log.Printf("Ingredients (%d): %s", len(allIngredients), allIngredients)

	inputs := onepizza.Ingredients{
		onepizza.Ingredient("cheese"),
		onepizza.Ingredient("peppers"),
		// onepizza.Ingredient("basil"),
		// onepizza.Ingredient("pineapple"),
		onepizza.Ingredient("mushrooms"),
	}
	log.Printf("How many would eat %v? %d", inputs, preferences.HowManyWouldEat(inputs))
}
