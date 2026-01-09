package main

import "fmt"

type PepeSchnele struct {
	Speed, Charisma, Wisdom int
}

func NewPepeSchnele(speed, charisma, wisdom int) *PepeSchnele {
	return &PepeSchnele{
		Speed:    speed,
		Charisma: charisma,
		Wisdom:   wisdom,
	}
}

func (p *PepeSchnele) GetRating() int {
	return (p.Speed * 2) + (p.Charisma * 3) + p.Wisdom
}

func (p *PepeSchnele) String() string {
	return fmt.Sprintf("Пепе Шнеле [Скорость: %d, Харизма: %d, Мудрость: %d] | Рейтинг: %d",
		p.Speed, p.Charisma, p.Wisdom, p.GetRating())
}

func main() {
	pepe1 := NewPepeSchnele(80, 90, 70)
	pepe2 := NewPepeSchnele(95, 60, 85)

	fmt.Println(pepe1)
	fmt.Println(pepe2)

	if pepe1.GetRating() > pepe2.GetRating() {
		fmt.Println("Пепе №1 круче!")
	} else if pepe2.GetRating() > pepe1.GetRating() {
		fmt.Println("Пепе №2 круче!")
	} else {
		fmt.Println("Оба Пепе одинаково круты!")
	}
}
