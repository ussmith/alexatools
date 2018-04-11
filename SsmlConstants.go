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

//SsmlClock describes teh SSML 12 or 24 hour clock formats
type SsmlClock string

//type Whisper string

//This block is the const values for SSML Processing
const (
	breakTimeTempl     string = "<break time=\"{{.1}}ms\"/>"
	breakStrengthTempl string = "<break strength=\"{{.}}\"/>"
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

	emphasisTempl string = "<emphasis level=\"{{.Level}}\">{{.Value}}</emphasis>"
	//ModerateEmphasis
	ModerateEmphasis Emphasis = "moderate"
	//ReducedEmphasis
	ReducedEmphasis Emphasis = "reduced"
	//StrongEmphasis"
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
	sayAsTempl string = "<say-as interpret-as=\"{{.SayAsType}}\">{{.SayAsValue}}</say-as>"
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
	dateTempl string = "<say-as interpret-as=\"date\" format=\"{{.Format}}\">{{.Value}}</say-as>"
	//SSML 12 Hour clock format
	TwelveHourClock SsmlClock = "hms12"
	//SSML 24 Hour clock format
	TwentyFourHourClock SsmlClock = "hms24"
)

const (
	sentenceTempl  string = "<s>{{.}}</s>"
	paragraphTempl string = "<p>{{.}}</p>"
)

const (
	//WhisperString Whisper = "whispers"

	//Whisper is a template for processing alexa specific SSML directives
	//Whisper template.Template = template.Must(template.New("whisper").Parse("<amazon:effect name=\"whispered\">{{.}}</amazon:effect>"))
	whisperTempl string = "<amazon:effect name=\"whispered\">{{.}}</amazon:effect>"
)
