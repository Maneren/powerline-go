package main

import (
	pwl "github.com/justjanne/powerline-go/powerline"
)

func segmentClock(p *powerline) []pwl.Segment {
	return []pwl.Segment{{
		Name:       "clock",
		Content:    "%D{%H:%M:%S}",
		Foreground: p.theme.TimeFg,
		Background: p.theme.TimeBg,
	}}
}
