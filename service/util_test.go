package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findCards(t *testing.T) {
	type args struct {
		codes []string
	}
	tests := []struct {
		name string
		args args
		want []*Card
	}{
		{
			name: "find cards",
			args: args{codes: []string{"AS", "2C", "5H", "6H", "2D"}},
			want: []*Card{
				{Value: "ACE", Suit: "SPADES", Code: "AS"},
				{Value: "2", Suit: "CLUBS", Code: "2C"},
				{Value: "5", Suit: "HEARTS", Code: "5H"},
				{Value: "6", Suit: "HEARTS", Code: "6H"},
				{Value: "2", Suit: "DIAMONDS", Code: "2D"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cards := findCards(tt.args.codes...)
			for i, card := range cards {
				assert.Equal(t, tt.want[i].Value, card.Value)
				assert.Equal(t, tt.want[i].Code, card.Code)
				assert.Equal(t, tt.want[i].Suit, card.Suit)
			}
		})
	}
}

func Test_allCards(t *testing.T) {
	cards := allCards()
	assert.Equal(t, 52, len(cards))
}

func Test_copyCardList(t *testing.T) {
	cards := allCards()
	dup := copyCardList(cards)
	dup[0] = &Card{Code: "test", Value: "test", Suit: "test"}
	assert.NotEqual(t, cards[0].Code, dup[0].Code)
}
