package models

import (
	"space-game/game/utils"

	"github.com/google/uuid"
)

type Meta struct {
	AbilityPool int    `json:"ability_pool"`
	LastAbility string `json:"last_ability"`
	FocusPool   int    `json:"focus_pool"`
}

type Character struct {
	Experience      int                    `json:"experience"`
	Name            string                 `json:"name"`
	Level           int                    `json:"level"`
	Origin          Origin                 `json:"origin"`
	Background      Background             `json:"background"`
	Talents         map[Talent]int         `json:"talents"`
	Focus           map[Focus]int          `json:"focus"`
	Specializations map[Specialization]int `json:"specializations"`
	SocialClass     SocialClass            `json:"social_class"`
	Profession      Profession             `json:"profession"`
	Drive           Drive                  `json:"drive"`
	Abilities       map[Ability]int        `json:"abilities"`
	Fortune         int                    `json:"fortune"`
	Conditions      []Condition            `json:"conditions"`
	Meta            Meta                   `json:"meta"`
}

type CharacterOpts struct{}

var CharacterOptions CharacterOpts

type CharacterOpt func(g *Character)

func NewCharacter(opts ...CharacterOpt) *Character {
	c := Character{
		Talents:         map[Talent]int{},
		Focus:           map[Focus]int{},
		Specializations: map[Specialization]int{},
		Abilities:       map[Ability]int{},
		Conditions:      make([]Condition, 0),
		Meta:            Meta{},
	}

	for _, opt := range opts {
		opt(&c)
	}
	return &c
}

func (CharacterOpts) Profession(p Profession) CharacterOpt {
	return func(c *Character) {
		c.Profession = p
	}
}

func (CharacterOpts) Background(b Background) CharacterOpt {
	return func(c *Character) {
		c.Background = b
	}
}

func (CharacterOpts) Focus(f map[Focus]int) CharacterOpt {
	return func(c *Character) {
		c.Focus = f
	}
}

func (CharacterOpts) Drive(drive Drive) CharacterOpt {
	return func(c *Character) {
		c.Drive = drive
	}
}

func (CharacterOpts) Talents(talents map[Talent]int) CharacterOpt {
	return func(c *Character) {
		c.Talents = talents
	}
}

//old but functional level 1 generator... 
func GenerateRandomCharacter(name string) Character {
	profession := Profession(utils.Rand(0, 23))
	background := Background(utils.Rand(0, 11))
	focus := map[Focus]int{
		background.BenefitDefinitions().FocusPool[utils.Rand(0, len(background.BenefitDefinitions().FocusPool))]: 1,
		profession.BenefitDefinitions().FocusPool[utils.Rand(0, len(profession.BenefitDefinitions().FocusPool))]: 1,
	}

	drive := Drive(utils.Rand(0, 11))
	talents := map[Talent]int{
		background.BenefitDefinitions().TalentPool[utils.Rand(0, len(background.BenefitDefinitions().TalentPool))]: 1,
		profession.BenefitDefinitions().TalentPool[utils.Rand(0, len(profession.BenefitDefinitions().TalentPool))]: 1,
		drive.Talents()[utils.Rand(0, len(drive.Talents()))]:                                                       1,
	}
	fortune := 15
	if _, ok := talents[Fortune]; ok {
		fortune = 20
	}
	character := Character{
		Name:            uuid.NewString(),
		Origin:          Origin(utils.Rand(0, 3)),
		Background:      background,
		Level:           1,
		Profession:      profession,
		SocialClass:     profession.SocialClass(),
		Drive:           drive,
		Conditions:      []Condition{},
		Specializations: map[Specialization]int{},
		Focus:           focus,
		Fortune:         fortune,
		Abilities: map[Ability]int{
			Accuracy:      utils.Rand(3, 18),
			Constitution:  utils.Rand(3, 18),
			Fighting:      utils.Rand(3, 18),
			Communication: utils.Rand(3, 18),
			Dexterity:     utils.Rand(3, 18),
			Intelligence:  utils.Rand(3, 18),
			Perception:    utils.Rand(3, 18),
			Strength:      utils.Rand(3, 18),
			Willpower:     utils.Rand(3, 18),
		},
		Meta: Meta{},
	}
	character.Abilities[Defense] = 10 + AbilityScoreModifier(character.Abilities[Dexterity])
	character.Abilities[Speed] = 10 + AbilityScoreModifier(character.Abilities[Dexterity])
	character.Abilities[Toughness] = AbilityScoreModifier(character.Abilities[Constitution])
	return character
}
