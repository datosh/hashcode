package onepizza

type Ingredient string

type Ingredients []Ingredient

func (ingredients Ingredients) FromStrings(strings []string) Ingredients {
	for _, s := range strings {
		ingredients = append(ingredients, Ingredient(s))
	}
	return ingredients
}

func (ingredients Ingredients) Contains(seek Ingredient) bool {
	for _, ingredient := range ingredients {
		if ingredient == seek {
			return true
		}
	}
	return false
}
