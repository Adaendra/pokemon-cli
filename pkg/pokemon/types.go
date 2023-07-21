package pokemon

import "strings"

const (
	NormalType   string = "Normal"
	FireType     string = "Fire"
	WaterType    string = "Water"
	GrassType    string = "Grass"
	ElectricType string = "Electric"
	IceType      string = "Ice"
	FightingType string = "Fighting"
	PoisonType   string = "Poison"
	GroundType   string = "Ground"
	FlyingType   string = "Flying"
	PsychicType  string = "Psychic"
	BugType      string = "Bug"
	RockType     string = "Rock"
	GhostType    string = "Ghost"
	DragonType   string = "Dragon"
	DarkType     string = "Dark"
	SteelType    string = "Steel"
	FairyType    string = "Fairy"
)

var (
	TypeList = []string{
		NormalType,
		FireType,
		WaterType,
		GrassType,
		ElectricType,
		IceType,
		FightingType,
		PoisonType,
		GroundType,
		FlyingType,
		PsychicType,
		BugType,
		RockType,
		GhostType,
		DragonType,
		DarkType,
		SteelType,
		FairyType,
	}
)

func CheckType(t string) bool {
	isCorrect := false

	for _, tt := range TypeList {
		if strings.ToUpper(tt) == strings.ToUpper(t) {
			isCorrect = true
			break
		}
	}

	return isCorrect
}
