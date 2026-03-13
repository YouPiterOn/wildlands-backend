package domain

import "math/rand"

type Deck struct {
	Cards       []Card
	ShuffleSeed int64
}

func NewShuffledDeck(cards []Card, seed int64) *Deck {
	deck := &Deck{
		Cards:       cards,
		ShuffleSeed: seed,
	}
	deck.Shuffle()
	return deck
}

func (d *Deck) DrawCard() *Card {
	card := d.Cards[0]
	d.Cards = d.Cards[1:]
	return &card
}

func (d *Deck) AddCardsAndShuffle(cards []Card, seed int64) {
	d.Cards = append(d.Cards, cards...)
	d.Shuffle()
}

func (d *Deck) Shuffle() {
	rng := rand.New(rand.NewSource(d.ShuffleSeed))
	for i := range d.Cards {
		j := rng.Intn(i + 1)
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}
}
