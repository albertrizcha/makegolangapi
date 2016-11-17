package Model

import (
	"time"
)

type Article struct {
	Title   string    `json:"title"`
	Desc    string    `json:"desc"`
	Content string    `json:"content"`
	Due     time.Time `json:"due"`
}

type Articles []Article

type Test_struct struct {
	Test string
}
