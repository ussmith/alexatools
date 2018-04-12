package alexatools

import (
	"bytes"
	"log"
	"strings"
	"text/template"
)

//SsmlBuilder is a convenience tools for building complex
//SSML based strings
type SsmlBuilder interface {
	Emphasis(Emphasis, string) SsmlBuilder
	Pause(BreakStrength) SsmlBuilder
	DurationPause(int) SsmlBuilder
	Pitch(Pitch) SsmlBuilder
	Rate(Rate) SsmlBuilder
	Volume(Volume) SsmlBuilder
	SayAs(SayAs, string) SsmlBuilder
	Sentence(string) SsmlBuilder
	Paragraph(string) SsmlBuilder
	Date(SsmlClock, string) SsmlBuilder
	Say(string) SsmlBuilder
	Build() string
}

type ssmlBuilder struct {
	buffer    bytes.Buffer
	PitchVal  Pitch
	RateVal   Rate
	VolumeVal Volume
}

type emphasisLevel struct {
	Level Emphasis
	Value string
}

type sayAs struct {
	SayAsType  SayAs
	SayAsValue string
}

type ssmlDate struct {
	Format SsmlClock
	Value  string
}

var replacer = strings.NewReplacer("&", " and ", "<", "", ">", "", "'", "", "\"", "")

var breakStrengthTemplate *template.Template
var breakTimeTemplate *template.Template
var emphasisTemplate *template.Template
var sayAsTemplate *template.Template
var sentenceTemplate *template.Template
var paragraphTemplate *template.Template
var dateTemplate *template.Template

func init() {

	breakStrengthTemplate = template.New("breakStrength")
	breakStrengthTemplate, err = breakStrengthTemplate.Parse(breakStrengthTempl)
	if err != nil {
		panic("Failed to parse break strength template")
	}

	breakTimeTemplate = template.New("breakTime")
	breakTimeTemplate, err = breakTimeTemplate.Parse(breakTimeTempl)
	if err != nil {
		panic("Failed to parse break time template")
	}

	sayAsTemplate = template.New("sayAs")
	sayAsTemplate, err = sayAsTemplate.Parse(sayAsTempl)
	if err != nil {
		panic("Failed to parse sayAs template")
	}

	emphasisTemplate = template.New("emphasis")
	emphasisTemplate, err = emphasisTemplate.Parse(emphasisTempl)
	if err != nil {
		panic("Failed to parse emphasis template")
	}

	sentenceTemplate = template.New("sentence")
	sentenceTemplate, err = sentenceTemplate.Parse(sentenceTempl)
	if err != nil {
		panic("Failed to parse sentence template")
	}

	paragraphTemplate = template.New("paragraph")
	paragraphTemplate, err = paragraphTemplate.Parse(paragraphTempl)
	if err != nil {
		panic("Failed to parse paragraph template")
	}

	dateTemplate = template.New("date")
	dateTemplate, err = dateTemplate.Parse(dateTempl)
	if err != nil {
		panic("Failed to parse date template")
	}
}

//New creates a new SsmlBuilder
func New() SsmlBuilder {
	return &ssmlBuilder{}
}

func (builder *ssmlBuilder) Phomeme(raw, phomeme string, cs CharacterSet) SsmlBuilder {
	return builder
}

func (builder *ssmlBuilder) Sentence(val string) SsmlBuilder {
	var tpl bytes.Buffer
	sentenceTemplate.Execute(&tpl, val)
	builder.addElement(tpl.String())
	return builder
}

func (builder *ssmlBuilder) Paragraph(val string) SsmlBuilder {
	var tpl bytes.Buffer
	paragraphTemplate.Execute(&tpl, val)
	builder.addElement(tpl.String())
	return builder
}

func (builder *ssmlBuilder) Emphasis(emphasis Emphasis, value string) SsmlBuilder {
	var tpl bytes.Buffer
	el := emphasisLevel{Level: emphasis, Value: value}
	emphasisTemplate.Execute(&tpl, el)
	builder.addElement(tpl.String())
	return builder
}

func (builder *ssmlBuilder) Pause(bs BreakStrength) SsmlBuilder {
	var tpl bytes.Buffer
	breakStrengthTemplate.Execute(&tpl, string(bs))
	builder.buffer.WriteString(tpl.String() + " ")
	return builder
}

func (builder *ssmlBuilder) DurationPause(millis int) SsmlBuilder {
	var tpl bytes.Buffer
	breakTimeTemplate.Execute(&tpl, millis)
	builder.buffer.WriteString(tpl.String() + " ")
	return builder
}

func (builder *ssmlBuilder) Break(bt int) SsmlBuilder {
	var tpl bytes.Buffer
	breakTimeTemplate.Execute(&tpl, bt)
	builder.buffer.WriteString(tpl.String() + " ")
	return builder
}

func (builder *ssmlBuilder) Date(format SsmlClock, value string) SsmlBuilder {
	var tpl bytes.Buffer
	dv := ssmlDate{Format: format, Value: value}
	dateTemplate.Execute(&tpl, dv)
	builder.addElement(tpl.String())
	return builder
}

func (builder *ssmlBuilder) Pitch(p Pitch) SsmlBuilder {
	builder.PitchVal = p
	return builder
}

func (builder *ssmlBuilder) Rate(r Rate) SsmlBuilder {
	builder.RateVal = r
	return builder
}

func (builder *ssmlBuilder) Volume(v Volume) SsmlBuilder {
	builder.VolumeVal = v
	return builder
}

func (builder *ssmlBuilder) SayAs(sa SayAs, value string) SsmlBuilder {
	var tpl bytes.Buffer
	sas := sayAs{SayAsType: sa, SayAsValue: value}
	sayAsTemplate.Execute(&tpl, sas)
	builder.addElement(tpl.String())
	return builder
}

func (builder *ssmlBuilder) Say(s string) SsmlBuilder {
	builder.addElement(s)
	return builder
}

func (builder *ssmlBuilder) Build() string {
	var tpl bytes.Buffer
	tpl.WriteString("<speak>")
	tpl.WriteString(builder.buffer.String())
	tpl.WriteString("</speak>")
	return tpl.String()
}

func (builder *ssmlBuilder) addElement(value string) {
	if builder.VolumeVal == "" && builder.RateVal == "" && builder.PitchVal == "" {
		builder.buffer.WriteString(value)
		return
	}

	var count int

	if builder.VolumeVal != "" {
		count++
	}
	if builder.PitchVal != "" {
		count++
	}
	if builder.RateVal != "" {
		count++
	}

	log.Printf("Count = %d\n", count)

	var tpl bytes.Buffer
	tpl.WriteString("<prosody ")

	if builder.PitchVal != "" {
		tpl.WriteString("pitch=\"")
		tpl.WriteString(string(builder.PitchVal))
		tpl.WriteString("\"")
		count--
		if count >= 1 {
			tpl.WriteString(" ")
		}
	}

	if builder.RateVal != "" {
		tpl.WriteString("rate=\"")
		tpl.WriteString(string(builder.RateVal))
		tpl.WriteString("\"")
		count--
		if count >= 1 {
			tpl.WriteString(" ")
		}
	}

	if builder.VolumeVal != "" {
		tpl.WriteString("volume=\"")
		tpl.WriteString(string(builder.VolumeVal))
		tpl.WriteString("\"")
	}

	tpl.WriteString(">")
	tpl.WriteString(value)
	tpl.WriteString("</prosody>")

	//Reset the prosody values
	builder.VolumeVal = ""
	builder.RateVal = ""
	builder.PitchVal = ""

	//Add the prosody section
	builder.buffer.WriteString(tpl.String())
}

func escape(in string) (out string) {
	out = replacer.Replace(in)
	return
}
