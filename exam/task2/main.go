package main

import (
	"fmt"
	"sort"
)

type BrainrotMeme struct {
	Name       string
	TrendLevel int
	Category   string
	Views      float64
}

func FindTopTrending(memes []BrainrotMeme, minViews float64) []BrainrotMeme {
	var result []BrainrotMeme

	for _, meme := range memes {
		if meme.Views > minViews {
			result = append(result, meme)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].TrendLevel > result[j].TrendLevel
	})
	return result
}

func CalculateCategoryImpact(memes []BrainrotMeme) map[string]float64 {
	impact := make(map[string]float64)

	for _, meme := range memes {
		impact[meme.Category] += meme.Views
	}

	return impact
}

func FilterByComplexCondition(memes []BrainrotMeme) []string {
	var names []string

	for _, meme := range memes {
		condition1 := meme.TrendLevel >= 7
		condition2 := meme.Views > 50 && meme.Category == "Sigma"

		if condition1 || condition2 {
			names = append(names, meme.Name)
		}
	}

	return names
}

func (m BrainrotMeme) String() string {
	return fmt.Sprintf("«%s» [%s] | Тренд: %d | Просмотры: %.1fM",
		m.Name, m.Category, m.TrendLevel, m.Views)
}

func main() {
	memes := []BrainrotMeme{
		{"Skibidi Toilet", 9, "Skibidi", 200.5},
		{"Sigma Male Grindset", 8, "Sigma", 120.3},
		{"Subo Bratik Dance", 6, "Subo Bratik", 45.7},
		{"TUNTUNTUNSAHUR Remix", 10, "TUNTUNTUNSAHUR", 300.0},
		{"Mewing Tutorial", 5, "Mewing", 30.2},
		{"Alpha Lifting", 7, "Sigma", 60.8},
		{"Gyatt in Ohio", 9, "Other", 80.0},
		{"Rizz God Compilation", 4, "Other", 25.5},
		{"Sigma Grind Never Stops", 6, "Sigma", 55.1},
	}

	fmt.Println("Исходный набор Brainrot-мемов:")
	for _, m := range memes {
		fmt.Println(m)
	}

	topTrending := FindTopTrending(memes, 50.0)
	fmt.Println("\nТоп-тренды (просмотры > 50M), сортировка по трендовости (убывание):")
	for _, m := range topTrending {
		fmt.Println("-", m)
	}

	impact := CalculateCategoryImpact(memes)
	fmt.Println("\nСуммарные просмотры по категориям (в миллионах):")
	var categories []string
	for cat := range impact {
		categories = append(categories, cat)
	}

	for _, cat := range categories {
		fmt.Printf("- %s: %.1fM\n", cat, impact[cat])
	}

	filteredNames := FilterByComplexCondition(memes)
	fmt.Println("\nМемы по сложному условию (TrendLevel ≥ 7 ИЛИ (Views > 50 И Sigma)): ")
	for i, name := range filteredNames {
		fmt.Printf("%d. %s\n", i+1, name)
	}
}
