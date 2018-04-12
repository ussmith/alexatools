package alexatools

import (
	"bytes"
	"log"
	"strings"
	"text/template"
)

//AlexaBuilder adds alexa specific capbability whisper
type AlexaBuilder interface {
	Whisper(string) AlexaBuilder
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

//New creates a new SsmlBuilder
func NewAlexaBuilder() AlexaBuilder {
	return &alexaBuilder{}
}

func (builder *ssmlBuilder) Whisper(whisper string) AlexaBuilder {
	var tpl bytes.Buffer
	whisperTemplate.Execute(&tpl, whisper)
	builder.buffer.WriteString(tpl.String() + " ")
	return builder
}
