package book1

import (
	"github.com/RileyGabrielson/text-adventure/model"
)

var Book1 = model.Book{
	Title:       "Book 1",
	Description: "Book Description",
	Chapters:    []model.Chapter{IntroductionChapter},
}
