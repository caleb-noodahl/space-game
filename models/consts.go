package models

type Origin int
type SocialClass int
type Background int
type Profession int
type Ability int
type Focus int
type Talent int
type Specialization int
type Drive int
type Condition int

type StationState int
type ShipCombatState int
type ShipSize int

const (
	Unknown ShipSize = iota
	Tiny
	Small
	Medium
	Large
	Huge
	Gigantic
	Colossal
	Titanic
)

const (
	Earther Origin = iota
	Martian
	Belter
)

const (
	Outsider SocialClass = iota
	LowerClass
	MiddleClass
	UpperClass
)

const (
	Blinded Condition = iota
	Deafened
	Dying
	Exhausted
	Fatigued
	FreeFalling
	Helpless
	Hindered
	Injured
	Prone
	Restrained
	Unconscious
	Wounded
)

const (
	Bohemian Background = iota
	Exile
	Outcast
	Military
	Laborer
	Urban
	Academic
	Suburban
	Trade
	Aristocratic
	Corporate
	Cosmopolitan
)

const (
	Accuracy Ability = iota
	Constitution
	Fighting
	Communication
	Dexterity
	Intelligence
	Perception
	Strength
	Willpower
	Defense
	Speed
	Toughness
)

const (
	Affluent Talent = iota
	Agility
	Artistry
	Attractive
	Burglary
	Carousing
	Command
	Contacts
	Doctor
	DualWeaponStyle
	Fringer
	Expertise
	GrapplingStyle
	Hacking
	Improvisation
	Inspire
	Intrigue
	KnowItAll
	Knowledge
	Linguistics
	Maker
	Medic
	Misdirection
	Observation
	Oratory
	OverwhelmStyle
	Performance
	PilotTalent
	PinpointAccuracy
	QuickReflexes
	PistolStyle
	RifleStyle
	Protector
	Scouting
	SelfDefenseStyle
	TacticalAwareness
	SingleWeaponStyle
	ThrownWeaponStyle
	TwoHandedStyle
	StrikingStyle
	Fortune
)

const (
	AcademicSpecialization Specialization = iota
	AceSpecialization
	AgentSpecialization
	CommandoSpecialization
	ExecutiveSpecialization
	GunfighterSpecialization
	HackerSpecialization
	InvestigatorSpecialization
	MartialSpecialization
	ArtistSpecialization
	SniperSpecialization
	SocialiteSpecialization
	StarSpecialization
)

type BenefitDefinition struct {
	Background      Background `json:"background"`
	Ability         Ability    `json:"ability"`
	AbilityModifier int        `json:"ability_mod"`
	FocusPool       []Focus    `json:"focus_pool"`
	TalentPool      []Talent   `json:"talent_pool"`
}

type Benefit struct {
	Name   string `json:"name"`
	Focus  Focus  `json:"focus"`
	Talent Talent `json:"talent"`
}

const (
	Brawler Profession = iota
	Survivalist
	Criminal
	Scavenger
	Fixer
	Artist
	Athlete
	Soldier
	Investigator
	Technician
	Clergy
	Negotiator
	Pilot
	Security
	Professional
	Scholar
	Merchant
	Politician
	Commander
	Explorer
	Dilettante
	Expert
	Executive
	Socialite
)

const (
	Bows Focus = iota
	Gunnery
	Pistols
	Rifles
	Throwing
	Bargaining
	Deception
	Disguise
	Etiquette
	Expression
	Gambling
	Investigation
	Leadership
	Performing
	Persuasion
	Seduction
	Running
	Stamina
	Swimming
	Tolerance
	Acrobatics
	Crafting
	Driving
	FreeFall
	Initiative
	Piloting
	SleightOfHand
	Stealth
	Brawling
	Grappling
	HeavyWeapons
	LightWeapons
	Art
	Business
	Cryptography
	CurrentAffairs
	Demolitions
	Engineering
	Evaluation
	Law
	Medicine
	Navigation
	Research
	Science
	SecurityFocus
	Tactics
	Technology
	Empathy
	Hearing
	Intuition
	Searching
	Seeing
	Smelling
	Survival
	Tasting
	Touching
	Tracking
	Climbing
	Intimidation
	Jumping
	Might
	Courage
	Faith
	SelfDiscipline
	History
	Theology
)

const (
	BuilderDrive Drive = iota
	CaregiverDrive
	EcstaticDrive
	JudgeDrive
	LeaderDrive
	NetworkerDrive
	PenitentDrive
	ProtectorDrive
	RebelDrive
	SurvivorDrive
	VisionaryDrive
	FortuneDrive
)

const (
	Target ShipCombatState = iota
	Engage
	Evade
	Disabled
)

const (
	Operational StationState = iota
	Lockdown
	Crisis
)

func (s Specialization) String() string {
	switch s {
	case AcademicSpecialization:
		return "academic"
	case AceSpecialization:
		return "ace"
	case AgentSpecialization:
		return "agent"
	case CommandoSpecialization:
		return "commando"
	case ExecutiveSpecialization:
		return "executive"
	case GunfighterSpecialization:
		return "gunfighter"
	case HackerSpecialization:
		return "hacker"
	case InvestigatorSpecialization:
		return "investor"
	case MartialSpecialization:
		return "martial"
	case ArtistSpecialization:
		return "artistic"
	case SniperSpecialization:
		return "sniper"
	case SocialiteSpecialization:
		return "socialite"
	case StarSpecialization:
		return "star"
	}
	return ""
}

func (c Condition) String() string {
	switch c {
	case Blinded:
		return "blinded"
	case Deafened:
		return "deafened"
	case Dying:
		return "dying"
	case Exhausted:
		return "exhausted"
	case Fatigued:
		return "fatigued"
	case FreeFalling:
		return "free_falling"
	case Helpless:
		return "helpless"
	case Hindered:
		return "hindered"
	case Injured:
		return "injured"
	case Prone:
		return "prone"
	case Restrained:
		return "restrained"
	case Unconscious:
		return "unconscious"
	case Wounded:
		return "wounded"
	}
	return ""
}

func (d Drive) String() string {
	switch d {
	case BuilderDrive:
		return "builder"
	case CaregiverDrive:
		return "caregiver"
	case EcstaticDrive:
		return "ecstatic"
	case JudgeDrive:
		return "judge"
	case LeaderDrive:
		return "leader"
	case NetworkerDrive:
		return "networker"
	case PenitentDrive:
		return "penitent"
	case ProtectorDrive:
		return "protector"
	case RebelDrive:
		return "rebel"
	case SurvivorDrive:
		return "survivor"
	case VisionaryDrive:
		return "visionary"
	case FortuneDrive:
		return "fortune"
	}
	return ""
}

func (d Drive) Talents() []Talent {
	switch d {
	case BuilderDrive:
		return []Talent{Maker, Oratory}
	case CaregiverDrive:
		return []Talent{Inspire, Medic}
	case EcstaticDrive:
		return []Talent{Attractive, Carousing}
	case JudgeDrive:
		return []Talent{Knowledge, Observation}
	case LeaderDrive:
		return []Talent{Command, Inspire}
	case NetworkerDrive:
		return []Talent{Contacts, Intrigue}
	case PenitentDrive:
		return []Talent{Fringer, KnowItAll}
	case ProtectorDrive:
		return []Talent{Misdirection, Protector}
	case RebelDrive:
		return []Talent{Expertise, Improvisation}
	case SurvivorDrive:
		return []Talent{Fringer, TacticalAwareness}
	case VisionaryDrive:
		return []Talent{Artistry, Oratory, Performance}
	case FortuneDrive:
		return []Talent{Fortune}

	}
	return []Talent{}
}

func (s SocialClass) IncomeScore() int {
	switch s {
	case Outsider:
		return 0
	case LowerClass:
		return 2
	case MiddleClass:
		return 4
	case UpperClass:
		return 6
	}
	return 0
}

func AbilityScoreModifier(val int) int {
	if val <= 3 {
		return -2
	}
	if val >= 4 && val <= 5 {
		return -1
	}
	if val >= 6 && val <= 8 {
		return 0
	}
	if val >= 9 && val <= 11 {
		return 1
	}
	if val >= 12 && val <= 14 {
		return 2
	}
	if val >= 15 && val <= 17 {
		return 3
	}
	if val >= 18 && val <= 20 {
		return 4
	}
	if val >= 22 {
		return 5
	}
	return 0
}

func (t Talent) String() string {
	switch t {
	case Affluent:
		return "affluent"
	case Agility:
		return "agility"
	case Artistry:
		return "artistry"
	case Attractive:
		return "attractive"
	case Burglary:
		return "burglary"
	case Carousing:
		return "carousing"
	case Command:
		return "command"
	case Contacts:
		return "contacts"
	case Doctor:
		return "doctor"
	case DualWeaponStyle:
		return "dual_weapon_style"
	case Fringer:
		return "fringer"
	case Expertise:
		return "expertise"
	case GrapplingStyle:
		return "grappling_style"
	case Hacking:
		return "hacking"
	case Improvisation:
		return "improvisation"
	case Inspire:
		return "inspire"
	case Intrigue:
		return "intrigue"
	case KnowItAll:
		return "know_it_all"
	case Knowledge:
		return "knowledge"
	case Linguistics:
		return "linguistics"
	case Maker:
		return "maker"
	case Medic:
		return "medic"
	case Misdirection:
		return "misdirection"
	case Observation:
		return "observation"
	case Oratory:
		return "oratory"
	case OverwhelmStyle:
		return "overwhelm_style"
	case Performance:
		return "performance"
	case PilotTalent:
		return "pilot"
	case PinpointAccuracy:
		return "pinpoint_accuracy"
	case QuickReflexes:
		return "quick_reflexes"
	case PistolStyle:
		return "pistol_style"
	case RifleStyle:
		return "rifle_style"
	case Protector:
		return "protector"
	case Scouting:
		return "scouting"
	case SelfDefenseStyle:
		return "self_defense_style"
	case TacticalAwareness:
		return "tactical_awareness"
	case SingleWeaponStyle:
		return "single_weapon_style"
	case ThrownWeaponStyle:
		return "thrown_weapon_style"
	case TwoHandedStyle:
		return "two_handed_style"
	case StrikingStyle:
		return "striking_style"
	}
	return ""
}

func (f Focus) String() string {
	switch f {
	case Bows:
		return "bows"
	case Gunnery:
		return "gunnery"
	case Pistols:
		return "pistols"
	case Rifles:
		return "rifles"
	case Throwing:
		return "throwing"
	case Bargaining:
		return "bargaining"
	case Deception:
		return "deception"
	case Disguise:
		return "disguise"
	case Etiquette:
		return "etiquette"
	case Expression:
		return "expression"
	case Gambling:
		return "gambling"
	case Investigation:
		return "investigation"
	case Leadership:
		return "leadership"
	case Performing:
		return "performing"
	case Persuasion:
		return "persuasion"
	case Seduction:
		return "seduction"
	case Running:
		return "running"
	case Stamina:
		return "stamina"
	case Swimming:
		return "swimming"
	case Tolerance:
		return "tolerance"
	case Acrobatics:
		return "acrobatics"
	case Crafting:
		return "crafting"
	case Driving:
		return "driving"
	case FreeFall:
		return "free_fall"
	case Initiative:
		return "initiative"
	case Piloting:
		return "piloting"
	case SleightOfHand:
		return "slight_of_hand"
	case Stealth:
		return "stealth"
	case Brawling:
		return "brawling"
	case Grappling:
		return "grappling"
	case HeavyWeapons:
		return "heavy_weapons"
	case LightWeapons:
		return "light_weapons"
	case Art:
		return "art"
	case Business:
		return "business"
	case Cryptography:
		return "cryptography"
	case CurrentAffairs:
		return "current_affairs"
	case Demolitions:
		return "demolitions"
	case Engineering:
		return "engineering"
	case Evaluation:
		return "evaluation"
	case Law:
		return "law"
	case Medicine:
		return "medicine"
	case Navigation:
		return "navigation"
	case Research:
		return "research"
	case Science:
		return "science"
	case SecurityFocus:
		return "security"
	case Tactics:
		return "tactics"
	case Technology:
		return "technology"
	case Empathy:
		return "empathy"
	case Hearing:
		return "hearing"
	case Intuition:
		return "intuition"
	case Searching:
		return "searching"
	case Seeing:
		return "seeing"
	case Smelling:
		return "smelling"
	case Survival:
		return "survival"
	case Tasting:
		return "tasting"
	case Touching:
		return "touching"
	case Tracking:
		return "tracking"
	case Climbing:
		return "climbing"
	case Intimidation:
		return "intimidation"
	case Jumping:
		return "jumping"
	case Might:
		return "might"
	case Courage:
		return "courage"
	case Faith:
		return "faith"
	case SelfDiscipline:
		return "self_discipline"
	case Theology:
		return "theology"
	}
	return ""
}

func (p Profession) String() string {
	switch p {
	case Brawler:
		return "brawler"
	case Survivalist:
		return "survivalist"
	case Criminal:
		return "criminal"
	case Scavenger:
		return "scavenger"
	case Fixer:
		return "fixer"
	case Artist:
		return "artist"
	case Athlete:
		return "athlete"
	case Soldier:
		return "soldier"
	case Investigator:
		return "investigator"
	case Technician:
		return "technician"
	case Clergy:
		return "clergy"
	case Negotiator:
		return "negotiator"
	case Pilot:
		return "pilot"
	case Security:
		return "security"
	case Professional:
		return "professional"
	case Scholar:
		return "scholar"
	case Merchant:
		return "merchant"
	case Politician:
		return "politician"
	case Commander:
		return "commander"
	case Explorer:
		return "explorer"
	case Dilettante:
		return "dilettante"
	case Expert:
		return "expert"
	case Executive:
		return "executive"
	case Socialite:
		return "socialite"
	}
	return ""
}

func (p Profession) SocialClass() SocialClass {
	switch p {
	case Brawler, Survivalist, Criminal, Scavenger, Fixer, Artist:
		return Outsider
	case Athlete, Soldier, Investigator, Technician, Clergy, Negotiator:
		return LowerClass
	case Pilot, Security, Professional, Scholar, Merchant, Politician:
		return MiddleClass
	case Commander, Explorer, Dilettante, Expert, Executive, Socialite:
		return UpperClass
	}
	return Outsider
}

func (a Ability) String() string {
	switch a {
	case Accuracy:
		return "accuracy"
	case Constitution:
		return "constitution"
	case Fighting:
		return "fighting"
	case Communication:
		return "communication"
	case Dexterity:
		return "dexterity"
	case Intelligence:
		return "intelligence"
	case Perception:
		return "perception"
	case Strength:
		return "strength"
	case Willpower:
		return "willpower"
	case Defense:
		return "defense"
	case Toughness:
		return "toughness"
	case Speed:
		return "speed"
	}
	return ""
}

func (s ShipCombatState) String() string {
	switch s {
	case Target:
		return "target"
	case Engage:
		return "engage"
	case Evade:
		return "evade"
	case Disabled:
		return "disabled"
	}
	return ""
}

func (a Ability) Focus() []Focus {
	switch a {
	case Accuracy:
		return []Focus{Bows, Gunnery, Pistols, Rifles, Throwing}
	case Communication:
		return []Focus{
			Bargaining, Deception, Disguise,
			Etiquette, Expression, Gambling,
			Investigation, Leadership, Performing,
			Persuasion, Seduction,
		}
	case Constitution:
		return []Focus{Running, Stamina, Swimming, Tolerance}
	case Dexterity:
		return []Focus{Acrobatics, Crafting, Driving, FreeFall, Initiative, Piloting, SleightOfHand, Stealth}
	case Fighting:
		return []Focus{Brawling, Grappling, HeavyWeapons, LightWeapons}
	case Intelligence:
		return []Focus{
			Art, Business, Cryptography,
			CurrentAffairs, Demolitions, Engineering,
			Evaluation, Law, Medicine,
			Navigation, Research, Science,
			SecurityFocus, Tactics, Technology,
			History,
		}
	case Perception:
		return []Focus{
			Empathy, Hearing, Intuition,
			Searching, Seeing, Smelling,
			Survival, Tasting, Touching,
			Tracking,
		}
	case Strength:
		return []Focus{Climbing, Intimidation, Jumping, Might}
	case Willpower:
		return []Focus{Courage, Faith, SelfDiscipline}
	}
	return []Focus{}
}
func (p Profession) BenefitDefinitions() BenefitDefinition {
	switch p {
	case Artist:
		return BenefitDefinition{
			FocusPool:  []Focus{Expression, Art},
			TalentPool: []Talent{Artistry, Performance},
		}
	case Athlete:
		return BenefitDefinition{
			FocusPool:  []Focus{Running, Swimming, Acrobatics, FreeFall, Climbing, Jumping},
			TalentPool: []Talent{Agility, QuickReflexes},
		}
	case Brawler:
		return BenefitDefinition{
			FocusPool:  []Focus{Brawling, Grappling},
			TalentPool: []Talent{GrapplingStyle, StrikingStyle},
		}
	case Clergy:
		return BenefitDefinition{
			FocusPool:  []Focus{Faith, Theology},
			TalentPool: []Talent{Inspire, Oratory},
		}
	case Commander:
		return BenefitDefinition{
			FocusPool:  []Focus{Leadership, Tactics},
			TalentPool: []Talent{Command, TacticalAwareness},
		}
	case Criminal:
		return BenefitDefinition{
			FocusPool:  []Focus{Deception, SleightOfHand, Stealth},
			TalentPool: []Talent{Burglary, Scouting},
		}
	case Dilettante:
		pool := Perception.Focus()
		pool = append(pool, Research)
		return BenefitDefinition{
			FocusPool:  pool,
			TalentPool: []Talent{Improvisation, KnowItAll},
		}
	case Executive:
		return BenefitDefinition{
			FocusPool:  []Focus{Leadership, Business},
			TalentPool: []Talent{Command, Intrigue},
		}
	case Expert:
		return BenefitDefinition{
			FocusPool:  Intelligence.Focus(),
			TalentPool: []Talent{Expertise, KnowItAll},
		}
	case Explorer:
		pool := Perception.Focus()
		pool = append(pool, Navigation)
		return BenefitDefinition{
			FocusPool:  pool,
			TalentPool: []Talent{PilotTalent, Scouting},
		}
	case Fixer:
		return BenefitDefinition{
			FocusPool:  []Focus{Bargaining, Evaluation},
			TalentPool: []Talent{Fringer, Improvisation},
		}
	case Investigator:
		pool := Perception.Focus()
		pool = append(pool, Investigation)
		return BenefitDefinition{
			FocusPool:  pool,
			TalentPool: []Talent{Intrigue, Observation},
		}
	case Merchant:
		return BenefitDefinition{
			FocusPool:  []Focus{Bargaining, Business},
			TalentPool: []Talent{Affluent, Contacts},
		}
	case Negotiator:
		return BenefitDefinition{
			FocusPool:  []Focus{Bargaining, Persuasion, Empathy},
			TalentPool: []Talent{Intrigue, Oratory},
		}
	case Pilot:
		return BenefitDefinition{
			FocusPool:  []Focus{Piloting},
			TalentPool: []Talent{PilotTalent},
		}
	case Politician:
		return BenefitDefinition{
			FocusPool:  []Focus{Deception, Persuasion, CurrentAffairs, Law},
			TalentPool: []Talent{Contacts, Oratory},
		}
	case Professional:
		return BenefitDefinition{
			FocusPool:  []Focus{Bargaining, Expression, Business, Technology, Research},
			TalentPool: []Talent{Affluent, Expertise},
		}
	case Scavenger:
		return BenefitDefinition{
			FocusPool:  []Focus{Technology, Searching},
			TalentPool: []Talent{Improvisation, Maker},
		}
	case Scholar:
		return BenefitDefinition{
			FocusPool:  Intelligence.Focus(),
			TalentPool: []Talent{Expertise, Knowledge},
		}
	case Security:
		return BenefitDefinition{
			FocusPool: []Focus{SecurityFocus, Empathy, Intuition, Seeing},
			TalentPool: []Talent{
				DualWeaponStyle, GrapplingStyle, OverwhelmStyle,
				PistolStyle, RifleStyle, SelfDefenseStyle,
				SingleWeaponStyle, StrikingStyle, ThrownWeaponStyle,
			},
		}
	case Socialite:
		return BenefitDefinition{
			FocusPool:  []Focus{Etiquette, Seduction},
			TalentPool: []Talent{Attractive, Contacts},
		}
	case Soldier:
		return BenefitDefinition{
			FocusPool: []Focus{Pistols, Rifles, Brawling},
			TalentPool: []Talent{TacticalAwareness, DualWeaponStyle, GrapplingStyle, OverwhelmStyle,
				PistolStyle, RifleStyle, SelfDefenseStyle,
				SingleWeaponStyle, StrikingStyle, ThrownWeaponStyle},
		}
	case Survivalist:
		return BenefitDefinition{
			FocusPool:  []Focus{Bows, Pistols},
			TalentPool: []Talent{Fringer, TacticalAwareness},
		}
	case Technician:
		return BenefitDefinition{
			FocusPool:  []Focus{Engineering, Technology},
			TalentPool: []Talent{Expertise, Hacking, Maker},
		}
	}
	return BenefitDefinition{}
}

func (b Background) BenefitDefinitions() BenefitDefinition {
	switch b {
	case Academic:
		return BenefitDefinition{
			Background:      Academic,
			Ability:         Intelligence,
			AbilityModifier: 1,
			FocusPool:       Intelligence.Focus(),
			TalentPool:      []Talent{Knowledge, Linguistics},
		}
	case Aristocratic:
		return BenefitDefinition{
			Background:      Aristocratic,
			Ability:         Communication,
			AbilityModifier: 1,
			FocusPool:       []Focus{Etiquette, History},
			TalentPool:      []Talent{Affluent, Contacts},
		}
	case Bohemian:
		pool := Intelligence.Focus()
		pool = append(pool, Performing)
		return BenefitDefinition{
			Background:      Bohemian,
			Ability:         Communication,
			AbilityModifier: 1,
			FocusPool:       pool,
			TalentPool:      []Talent{Affluent, Contacts},
		}
	case Corporate:
		return BenefitDefinition{
			Background:      Corporate,
			Ability:         Communication,
			AbilityModifier: 1,
			FocusPool:       []Focus{Bargaining, Business},
			TalentPool:      []Talent{Contacts, Intrigue},
		}
	case Cosmopolitan:
		return BenefitDefinition{
			Background:      Cosmopolitan,
			Ability:         Perception,
			AbilityModifier: 1,
			FocusPool:       []Focus{Etiquette, CurrentAffairs},
			TalentPool:      []Talent{Knowledge, Observation},
		}
	case Exile:
		return BenefitDefinition{
			Background:      Exile,
			Ability:         Constitution,
			AbilityModifier: 1,
			FocusPool:       []Focus{Brawling, SelfDiscipline},
			TalentPool:      []Talent{Affluent, Fringer},
		}
	case Military:
		return BenefitDefinition{
			Background:      Military,
			Ability:         Fighting,
			AbilityModifier: 1,
			FocusPool:       []Focus{Pistols, Tactics},
			TalentPool: []Talent{
				DualWeaponStyle, GrapplingStyle, GrapplingStyle, PistolStyle, RifleStyle, SelfDefenseStyle,
				SingleWeaponStyle, StrikingStyle, ThrownWeaponStyle, TwoHandedStyle, Observation,
			},
		}
	case Outcast:
		return BenefitDefinition{
			Background:      Outcast,
			Ability:         Perception,
			AbilityModifier: 1,
			FocusPool:       []Focus{Deception, Stealth},
			TalentPool:      []Talent{Fringer, Misdirection},
		}
	case Laborer:
		return BenefitDefinition{
			Background:      Laborer,
			Ability:         Constitution,
			AbilityModifier: 1,
			FocusPool:       []Focus{Crafting, Might},
			TalentPool:      []Talent{GrapplingStyle, SelfDefenseStyle, StrikingStyle, TwoHandedStyle, Carousing},
		}
	case Suburban:
		return BenefitDefinition{
			Background:      Suburban,
			Ability:         Communication,
			AbilityModifier: 1,
			FocusPool:       []Focus{Etiquette, CurrentAffairs},
			TalentPool:      []Talent{Affluent, Contacts},
		}
	case Trade:
		return BenefitDefinition{
			Background:      Trade,
			Ability:         Dexterity,
			AbilityModifier: 1,
			FocusPool:       []Focus{Crafting, Engineering},
			TalentPool:      []Talent{Improvisation, Maker},
		}
	case Urban:
		return BenefitDefinition{
			Background:      Urban,
			Ability:         Dexterity,
			AbilityModifier: 1,
			FocusPool:       []Focus{Persuasion, Stamina},
			TalentPool:      []Talent{Agility, Misdirection},
		}
	}
	return BenefitDefinition{}
}

func (b Background) String() string {
	switch b {
	case Bohemian:
		return "bohemian"
	case Exile:
		return "exile"
	case Outcast:
		return "outcast"
	case Military:
		return "military"
	case Laborer:
		return "laborer"
	case Urban:
		return "urban"
	case Academic:
		return "academic"
	case Suburban:
		return "suburban"
	case Trade:
		return "trade"
	case Aristocratic:
		return "aristocratic"
	case Corporate:
		return "corporate"
	case Cosmopolitan:
		return "cosmopolitan"
	}
	return ""
}

func (b Background) SocialClass() SocialClass {
	if (b == Bohemian) || (b == Exile) || (b == Outcast) {
		return Outsider
	}
	if (b == Military) || (b == Laborer) || (b == Urban) {
		return LowerClass
	}
	if (b == Academic) || (b == Suburban) || (b == Trade) {
		return MiddleClass
	}
	if (b == Aristocratic) || (b == Corporate) || (b == Cosmopolitan) {
		return UpperClass
	}
	return Outsider
}

func (o Origin) String() string {
	switch o {
	case Earther:
		return "earther"
	case Martian:
		return "martian"
	case Belter:
		return "belter"
	}
	return ""
}

func (s SocialClass) String() string {
	switch s {
	case Outsider:
		return "outsider"
	case LowerClass:
		return "lower_class"
	case MiddleClass:
		return "middle_class"
	case UpperClass:
		return "upper_class"
	}
	return ""
}
