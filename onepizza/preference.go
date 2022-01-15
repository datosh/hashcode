package onepizza

import "fmt"

type Preference struct {
	Likes    Ingredients
	Dislikes Ingredients
}

func (p Preference) String() string {
	return fmt.Sprintf("Likes: %v\nDislikes: %v", p.Likes, p.Dislikes)
}

// TODO: Check all likes are satisfied
func (p Preference) WouldEat(ingredients Ingredients) bool {
	for _, ingredient := range ingredients {
		if p.Dislikes.Contains(ingredient) {
			return false
		}
	}
	return true
}

type Preferences []Preference

func (ps Preferences) String() string {
	s := "-----\n"
	for _, p := range ps {
		s += p.String() + "\n------\n"
	}
	return s
}

func (ps Preferences) AllIngredients() Ingredients {
	ingredients := make(map[Ingredient]bool)
	for _, p := range ps {
		for _, like := range p.Likes {
			ingredients[like] = true
		}
		for _, dislike := range p.Dislikes {
			ingredients[dislike] = true
		}
	}
	keys := make(Ingredients, 0, len(ingredients))
	for k := range ingredients {
		keys = append(keys, k)
	}
	return keys
}

func (ps Preferences) HowManyWouldEat(ingredients Ingredients) (amount int) {
	for _, p := range ps {
		if p.WouldEat(ingredients) {
			amount++
		}
	}
	return amount
}
