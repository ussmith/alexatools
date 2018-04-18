package alexatools

import (
	"bytes"
	"text/template"
)

//AlexaBuilder adds alexa specific capbability whisper
type AlexaBuilder interface {
	Whisper(string) AlexaBuilder
	Emphasis(Emphasis, string) AlexaBuilder
	Pause(BreakStrength) AlexaBuilder
	DurationPause(int) AlexaBuilder
	Pitch(Pitch) AlexaBuilder
	Rate(Rate) AlexaBuilder
	Volume(Volume) AlexaBuilder
	SayAs(SayAs, string) AlexaBuilder
	Sentence(string) AlexaBuilder
	Paragraph(string) AlexaBuilder
	Date(SsmlClock, string) AlexaBuilder
	Say(string) AlexaBuilder
	Build() string
}

type alexaBuilder struct {
	builder ssmlBuilder
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

//NewAlexaBuilder creates a new AlexaBuilder
func NewAlexaBuilder() AlexaBuilder {
	return &alexaBuilder{}
}

func (alexaBuilder *alexaBuilder) Whisper(whisper string) AlexaBuilder {
	var tpl bytes.Buffer
	whisperTemplate.Execute(&tpl, whisper)
	alexaBuilder.builder.addElement(tpl.String(), true)
	return alexaBuilder
}

func (alexaBuilder *alexaBuilder) Phomeme(raw, phomeme string, cs CharacterSet) AlexaBuilder {
	return alexaBuilder
}

func (alexaBuilder *alexaBuilder) Sentence(val string) AlexaBuilder {
	alexaBuilder.builder.Sentence(val)
	return alexaBuilder
}

func (alexaBuilder *alexaBuilder) Paragraph(val string) AlexaBuilder {
	alexaBuilder.builder.Paragraph(val)
	return alexaBuilder
}

func (alexaBuilder *alexaBuilder) Emphasis(emphasis Emphasis, value string) AlexaBuilder {
	alexaBuilder.builder.Emphasis(emphasis, value)
	return alexaBuilder
}

func (alexaBuilder *alexaBuilder) Pause(bs BreakStrength) AlexaBuilder {
	alexaBuilder.builder.Pause(bs)
	return alexaBuilder
}

func (alexaBuilder *alexaBuilder) DurationPause(millis int) AlexaBuilder {
	alexaBuilder.builder.DurationPause(millis)
	return alexaBuilder
}

func (alexaBuilder *alexaBuilder) Break(bt int) AlexaBuilder {
	alexaBuilder.Break(bt)
	return alexaBuilder
}

func (alexaBuilder *alexaBuilder) Date(format SsmlClock, value string) AlexaBuilder {
	alexaBuilder.builder.Date(format, value)
	return alexaBuilder
}

func (alexaBuilder *alexaBuilder) Pitch(p Pitch) AlexaBuilder {
	alexaBuilder.builder.Pitch(p)
	return alexaBuilder
}

func (alexaBuilder *alexaBuilder) Rate(r Rate) AlexaBuilder {
	alexaBuilder.builder.Rate(r)
	return alexaBuilder
}

func (alexaBuilder *alexaBuilder) Volume(v Volume) AlexaBuilder {
	alexaBuilder.builder.Volume(v)
	return alexaBuilder
}

func (alexaBuilder *alexaBuilder) SayAs(sa SayAs, value string) AlexaBuilder {
	alexaBuilder.builder.SayAs(sa, value)
	return alexaBuilder
}

func (alexaBuilder *alexaBuilder) Say(s string) AlexaBuilder {
	alexaBuilder.builder.Say(s)
	return alexaBuilder
}

func (alexaBuilder *alexaBuilder) Build() string {
	return alexaBuilder.builder.Build()
}
