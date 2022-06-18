package main

import (
	"math/rand"
	"time"

	"github.com/kanztu/goblog/pkg/wasm"
)

const (
	MIN_JITTER_INPUT = 50  // ms
	MAX_JITTER_INPUT = 200 // ms

	MIN_JITTER_DEL = 50  // ms
	MAX_JITTER_DEL = 100 // ms

	WORD_DELAY      = 1000
	WORD_DELAY_POST = 500
)

var (
	prefix = "I am "
	items  = []string{
		"a Bassist",
		"an Explorer",
		"a Programmer",
	}
)

func Animation() {
	dom := wasm.GetElementById("description")
	dom.DeleteInnerHTML()
	for {
		for _, s := range items {
			// INPUT
			len := len(s)
			for i := range s {
				dom.SetInnerHTML(prefix + s[0:i+1])
				delay := rand.Intn(MAX_JITTER_INPUT-MIN_JITTER_INPUT) + MIN_JITTER_INPUT
				time.Sleep(time.Duration(delay) * time.Millisecond)
			}
			time.Sleep(time.Duration(WORD_DELAY) * time.Millisecond)
			// DELETE
			for i := range s {
				dom.SetInnerHTML(prefix + s[0:len-i-1])
				delay := rand.Intn(MAX_JITTER_DEL-MIN_JITTER_DEL) + MIN_JITTER_DEL
				time.Sleep(time.Duration(delay) * time.Millisecond)
			}
			time.Sleep(time.Duration(WORD_DELAY_POST) * time.Millisecond)

		}
	}
}
