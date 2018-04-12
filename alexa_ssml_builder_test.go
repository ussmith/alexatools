package alexatools

import (
	"fmt"
	"testing"
)

func TestWhisper(t *testing.T) {
	builder := NewAlexaBuilder()
	builder.Whisper("quietly")
	fmt.Println(builder.Build())
}
