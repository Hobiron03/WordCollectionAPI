package models

type Word struct {
	ID        int
	UserID    int
	Word      string
	Mean      string
	Pronounce string
	Genre     string
	color     string
}

func (u *User) CreateWord(word string, mean string, pronounce string, genre string, color string) (err error) {
	// cmd := `insert into words `
}
