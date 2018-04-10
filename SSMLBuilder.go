package alexatools

import (
	"bytes"
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
	Whisper(Whisper) AlexaBuilder
}

type ssmlBuilder struct {
	buffer bytes.Buffer
}

func (builder *ssmlBuilder) Whisper(whisper Whisper) AlexaBuilder {
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
