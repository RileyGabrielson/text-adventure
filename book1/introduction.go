package book1

import (
	"github.com/RileyGabrielson/text-adventure/model"
)

var IntroductionChapter = model.Chapter{
	Title:          "Journey of the Fallen",
	Description:    "Alone and forgotten.",
	StartingChoice: theForest,
}

var theForest = model.Choice{
	Location: "Elenthal Forest",
	Decision: "",
	Description: "The forest is littered with leaves and shadow. " +
		"Before you lies a path to a cave.",
	Options: []model.Choice{theCave},
}

var theCave = model.Choice{
	Location:    "Elenthal Cave",
	Decision:    "Journey to the Cave",
	Description: "This is a really stinky cave. ",
}
