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

func TestBreakDuration(t *testing.T) {
	builder := New()
	builder.Say("Scooby")
	builder = builder.DurationPause(1000000)
	builder.Say("Dooby Doo")
	fmt.Println(builder.Build())
}

func TestProsody(t *testing.T) {
	builder := New()
	//builder.Pitch(HighPitch).Say("High Pitchin' here boss")
	builder.Pitch(HighPitch).Volume(LoudVolume).Say("High Pitchin' here boss")
	builder.SayAs(Ordinal, "12345")
	fmt.Println(builder.Build())
}

func TestEscape(t *testing.T) {
	builder := New()

	builder.Say("Yes & No")
	fmt.Println(builder.Build())
}

func TestDate(t *testing.T) {
	builder := New()

	builder.Date(YearMonthDay, "2019-12-31")
	t.Log(builder.Build())
}

func TestSsml(t *testing.T) {
	builder := New()

	ssml := builder.Say(" is happening on ").Date(YearMonthDay, "2019-12-31").Build()
	t.Log(ssml)
}