package poker

type Suit int

const (
	Hearts Suit = iota
	Diamonds
	Clubs
	Spades
)

type Rank int

const (
	Two Rank = iota + 2
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

type Card struct {
	Rank Rank
	Suit Suit
}

var rankNames = []string{
	"", "", "Двойка", "Тройка", "Четвёрка", "Пятёрка", "Шестёрка", "Семёрка",
	"Восьмёрка", "Девятка", "Десятка", "Валет", "Дама", "Король", "Туз",
}

var suitNames = []string{
	"Черви", "Бубны", "Трефы", "Пики",
}

func (c Card) String() string {
	return rankNames[c.Rank] + " " + suitNames[c.Suit]
}
