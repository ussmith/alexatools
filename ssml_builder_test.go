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
	builder = builder.BreakStrength(StrongBreakStrength)
	fmt.Println(builder.Build())
}

func TestProsody(t *testing.T) {
	builder := New()
	builder.Pitch(HighPitch).Say("High Pitchin' here boss")
	fmt.Println(builder.Build())
}
