package domain

import "math/rand"

type ExploreDeckCard struct {
	Card    ExploreCard
	IsRuins bool
}

func toDeckCards(exploreCards []ExploreCard) []ExploreDeckCard {
	deckCards := make([]ExploreDeckCard, 0, len(exploreCards))
	for _, card := range exploreCards {
		deckCards = append(deckCards, ExploreDeckCard{
			Card:    card,
			IsRuins: false,
		})
	}
	return deckCards
}

type ExploreDeck struct {
	cards       []ExploreDeckCard
	rng         *rand.Rand
	ruinsNumber int
}

func NewShuffledDeck(cards []ExploreCard, ruinsNumber int, seed int64) *ExploreDeck {
	deck := &ExploreDeck{
		cards:       toDeckCards(cards),
		rng:         rand.New(rand.NewSource(seed)),
		ruinsNumber: ruinsNumber,
	}
	deck.shuffle()
	deck.resetRuinsAndShuffle()
	return deck
}

func (d *ExploreDeck) DrawCard() *ExploreDeckCard {
	card := d.cards[0]
	d.cards = d.cards[1:]
	if card.IsRuins && d.ruinsNumber > 0 {
		d.ruinsNumber--
	}
	return &card
}

func (d *ExploreDeck) AddCardsAndShuffle(cards []ExploreCard, seed int64) {
	deckCards := toDeckCards(cards)
	d.cards = append(d.cards, deckCards...)
	d.resetRuinsAndShuffle()
}

func (d *ExploreDeck) shuffle() {
	d.rng.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}

func (d *ExploreDeck) resetRuinsAndShuffle() {
	for i := 0; i < len(d.cards); i++ {
		if i < d.ruinsNumber {
			d.cards[i].IsRuins = true
		} else {
			d.cards[i].IsRuins = false
		}
	}
	d.shuffle()
}
