package models

type Block struct {
	Index     int
	Timestamp string
	Content   string
	Hash      string
	PrevHash  string
}
