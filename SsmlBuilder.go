package alexatools

import (
	"bytes"
	"text/template"
)

//SsmlBuilder is a convenience tools for building complex
//SSML based strings
type SsmlBuilder interface {
	Emphasis(Emphasis, string) SsmlBuilder
	BreakStrength(BreakStrength) SsmlBuilder
	Pitch(Pitch) SsmlBuilder
	Rate(Rate) SsmlBuilder
	Volume(Volume) SsmlBuilder
	SayAs(SayAs) SsmlBuilder
	SSMLClock(SSMLClock) SsmlBuilder
	Build() string
}

//AlexaBuilder adds alexa specific capbability whisper
type AlexaBuilder interface {
	Whisper(string) AlexaBuilder
}

type ssmlBuilder struct {
	buffer bytes.Buffer
}

type emphasisLevel struct {
	Level Emphasis
	Value string
}

var whisperTemplate *template.Template
var breakStrengthTemplate *template.Template
var breakTimeTemplate *template.Template
var emphasisTemplate *template.Template

func init() {
	var err error
	whisperTemplate = template.New("whisper")
	whisperTemplate, err = whisperTemplate.Parse(whisperTempl)
	if err != nil {
		panic("Failed to parse whisper template")
	}

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

	emphasisTemplate = template.New("emphasis")
	emphasisTemplate, err = emphasisTemplate.Parse(emphasisTempl)
	if err != nil {
		panic("Failed to parse emphasis template")
	}
}

func (builder *ssmlBuilder) Whisper(whisper string) AlexaBuilder {
	var tpl bytes.Buffer
	whisperTemplate.Execute(&tpl, whisper)
	builder.buffer.WriteString(tpl.String())
	return builder
}

func (builder *ssmlBuilder) Emphasis(emphasis Emphasis, value string) SsmlBuilder {
	var tpl bytes.Buffer
	el := emphasisLevel{Level: emphasis, Value: value}
	emphasisTemplate.Execute(&tpl, el)
	builder.buffer.WriteString(tpl.String())
	return builder
}

func (builder *ssmlBuilder) BreakStrength(bs BreakStrength) SsmlBuilder {
	var tpl bytes.Buffer
	breakStrengthTemplate.Execute(&tpl, string(bs))
	builder.buffer.WriteString(tpl.String())
	return builder
}

func (builder *ssmlBuilder) Break(bt int) SsmlBuilder {
	var tpl bytes.Buffer
	breakTimeTemplate.Execute(&tpl, bt)
	builder.buffer.WriteString(tpl.String())
	return builder
}

func (builder *ssmlBuilder) Pitch(p Pitch) SsmlBuilder {
	return builder
}

func (builder *ssmlBuilder) Rate(r Rate) SsmlBuilder {
	return builder
}

func (builder *ssmlBuilder) Volume(v Volume) SsmlBuilder {
	return builder
}

func (builder *ssmlBuilder) SayAs(sa SayAs) SsmlBuilder {
	return builder
}

func (builder *ssmlBuilder) SSMLClock(clock SSMLClock) SsmlBuilder {
	return builder
}

func (builder *ssmlBuilder) Build() string {
	return builder.buffer.String()
}
