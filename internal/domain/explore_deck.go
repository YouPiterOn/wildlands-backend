package domain

import "math/rand"

type ExploreDeck struct {
	Cards       []ExploreCard
	ShuffleSeed int64
}

func NewShuffledDeck(cards []ExploreCard, seed int64) *ExploreDeck {
	deck := &ExploreDeck{
		Cards:       cards,
		ShuffleSeed: seed,
	}
	deck.Shuffle()
	return deck
}

func (d *ExploreDeck) DrawCard() *ExploreCard {
	card := d.Cards[0]
	d.Cards = d.Cards[1:]
	return &card
}

func (d *ExploreDeck) AddCardsAndShuffle(cards []ExploreCard, seed int64) {
	d.Cards = append(d.Cards, cards...)
	d.Shuffle()
}

func (d *ExploreDeck) Shuffle() {
	rng := rand.New(rand.NewSource(d.ShuffleSeed))
	for i := range d.Cards {
		j := rng.Intn(i + 1)
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}
}
