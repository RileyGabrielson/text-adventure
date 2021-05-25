package book1

import (
	"github.com/RileyGabrielson/textAdventure/model"
)

var IntroductionChapter = model.Chapter{
	Title:          "Journey of the Fallen",
	Description:    "Alone and forgotten.",
	StartingChoice: theForest,
}

var theForest = model.Choice{
	Location: "Elenthal Forest",
	Description: "Test text 1 " +
		"Test text 2",
	Options: []model.Choice{theCave},
}

var theCave = model.Choice{
	Location: "Elenthal Cave",
	Description: "Test text 1 " +
		"Test text 2",
}
