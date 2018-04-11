package alexatools

import (
	"fmt"
	"testing"
)

func TestSayAndEmphasis(t *testing.T) {
	builder := New()
	builder.Say("Hello")
	builder.Emphasis(StrongEmphasis, "world")
	fmt.Println(builder.Build())
}

func TestEmphasis(t *testing.T) {
	builder := New()
	builder = builder.Emphasis(ModerateEmphasis, "hello")
	fmt.Println(builder.Build())
}

func TestBreakStrength(t *testing.T) {
	builder := New()
	builder = builder.Pause(StrongBreakStrength)
	fmt.Println(builder.Build())
}

func TestProsody(t *testing.T) {
	builder := New()
	//builder.Pitch(HighPitch).Say("High Pitchin' here boss")
	builder.Pitch(HighPitch).Volume(LoudVolume).Say("High Pitchin' here boss")
	builder.SayAs(Ordinal, "12345")
	fmt.Println(builder.Build())
}
