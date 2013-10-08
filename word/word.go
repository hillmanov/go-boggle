package word

type Word struct {
	letters    string
	usedCoords [][]int
}

func New(letter string, row, col int) *Word {
	w := new(Word)
	w.letters = letter
	w.usedCoords = append(w.usedCoords, []int{row, col})
	return w
}

func NewFromWord(otherWord *Word) *Word {
	w := new(Word)
	w.letters = otherWord.letters
	w.usedCoords = otherWord.usedCoords
	return w
}

func (w *Word) ContainsCoords(row, col int) bool {
	for _, currentCords := range w.usedCoords {
		if (currentCords[0] == row) && (currentCords[1] == col) {
			return true
		}
	}
	return false
}

func (w *Word) AddLetter(letter string, row, col int) {
	w.letters = w.letters + letter
	w.usedCoords = append(w.usedCoords, []int{row, col})
}

func (w *Word) Word() string {
	return w.letters
}
