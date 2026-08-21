package domain

type Key struct {
	ID, Algorithm, Subject string
	Public                 []byte
	Active                 bool
}
