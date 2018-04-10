package alexatools

//Emphasis describes the SSML Emphasis directives
type Emphasis string

//BreakStrength describes the SSML BreakStrength directives
type BreakStrength string

//Pitch describes the SSML pitch directives
type Pitch string

//Rate describes the SSML rate directives
type Rate string

//Volume describes the SSML volume directives
type Volume string

//SayAs describes the SSML SayAs directives
type SayAs string

//SSMLClock describes teh SSML 12 or 24 hour clock formats
type SSMLClock string

type Whisper string

//This block is the const values for SSML Processing
const (

	//NoBreakStrength
	NoBreakStrength BreakStrength = "none"
	//ExtraWeakBreakStrength
	ExtraWeakBreakStrength BreakStrength = "x-weak"
	//WeakBreakStrength
	WeakBreakStrength BreakStrength = "weak"
	//MediumBreakStrenght
	MediumBreakStrength BreakStrength = "medium"
	//StrongBreakStrength
	StrongBreakStrength BreakStrength = "strong"
	//ExtraStrongBreakStrength
	ExtraStrongBreakStrength BreakStrength = "x-strong"

	//ModerateEmphasis
	ModerateEmphasis Emphasis = "moderate"
	//ReducedEmphasis
	ReducedEmphasis Emphasis = "reduced"
	//StrongEmphasis
	StrongEmphasis Emphasis = "strong"

	//ExtraLowPitch
	ExtraLowPitch Pitch = "x-low"
	//LowPitch
	LowPitch Pitch = "low"
	//MediumPitch
	MediumPitch Pitch = "medium"
	//HighPitch
	HighPitch Pitch = "high"
	//ExtraHighPitch
	ExtraHighPitch Pitch = "x-high"

	//ExtraSlowRate
	ExtraSlowRate Rate = ("x-slow")
	//SlowRate
	SlowRate Rate = "slow"
	//MediumRate
	MediumRate Rate = "medium"
	//FastRateRate
	Fast Rate = "fast"
	//ExtraFastRate
	ExtraFastRate Rate = "x-fast"

	//Silent
	Silent Volume = "silent"
	//ExtraSoftVolume
	ExtraSoftVolume Volume = "x-soft"
	//SoftVolume
	Soft Volume = "soft"
	//MediumVolume
	MediumVolume Volume = "medium"
	//LoudVolume
	LoudVolume Volume = "loud"
	//ExtraLoudVolume
	ExtraLoudVolume Volume = "x-loud"
)

//These consts describe the SSML say as features
const (
	//Characters
	Characters SayAs = "characters"
	//Spell Out
	SpellOut SayAs = "spell-out"
	//Carinal
	Cardinal SayAs = "cardinal"
	//Number
	Number SayAs = "number"
	//Oridinal
	Ordinal SayAs = "ordinal"
	//Digits
	Digits SayAs = "digits"
	//Fraction
	Fraction SayAs = "fraction"
	//Unit
	Unit SayAs = "unit"
	//Time
	Time SayAs = "time"
	//Telephone
	Telephone SayAs = "telephone"
	//Address
	Address SayAs = "address"
	//Interjection
	Interjection SayAs = "interjection"
	//Expletive
	Expletive SayAs = "expletive"
)

//These consts describe the SSML time format
const (
	//SSML 12 Hour clock format
	TwelveHourClock SSMLClock = "hms12"
	//SSML 24 Hour clock format
	TwentyFourHourClock SSMLClock = "hms24"
)

const (
	//WhisperString is an alexa specific ssml directive
	WhisperString Whisper = "whispers"
)
