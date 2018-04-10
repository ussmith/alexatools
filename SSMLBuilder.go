package alexatools

import (
	"bytes"
	"text/template"
)

//SsmlBuilder is a convenience tools for building complex
//SSML based strings
type SsmlBuilder interface {
	Emphasis(Emphasis) SsmlBuilder
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

var whisperTemplate *template.Template

func init() {
	var err error
	whisperTemplate = template.New("whisper")
	whisperTemplate, err = whisperTemplate.Parse(whisperTempl)
	if err != nil {
		panic("Failed to parse whisper template")
	}
}
func (builder *ssmlBuilder) Whisper(whisper string) AlexaBuilder {
	var tpl bytes.Buffer
	whisperTemplate.Execute(&tpl, whisper)
	builder.buffer.WriteString(tpl.String())
	return builder
}

func (builder *ssmlBuilder) Emphasis(emphasis Emphasis) SsmlBuilder {
	builder.buffer.WriteString(string(emphasis))
	return builder
}

func (builder *ssmlBuilder) BreakStrength(bs BreakStrength) SsmlBuilder {
	builder.buffer.WriteString(string(bs))
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
