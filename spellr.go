package main

import (
	tb "github.com/nsf/termbox-go"
)

var (
	phrase         = `test phrase`
	typePhrase     = ``
	pi             = 0
	phraseComplete = true
)

func loadNewPhrase(newPhrase string) {
	pi = 0
	tb.Clear(tb.ColorDefault, tb.ColorDefault)

	// Load the phrase on the first line
	for i := 0; i < len(newPhrase); i++ {
		tb.SetCell(i, 0, rune(newPhrase[i]), tb.ColorDefault, tb.ColorDefault)
	}

	tb.Flush()
	phraseComplete = false
}

// This is called everytime a key is pressed
func updatePhrase(ev tb.Event) {
	typeRune := ev.Ch
	if pi >= len(phrase) {
		tb.SetCell(pi, 1, typeRune, tb.ColorDefault, tb.ColorRed)
	} else {
		CurrentPhraseRune := rune(phrase[pi])
		// ev.Ch does not get populated when space is pressed
		if ev.Key == tb.KeySpace {
			typeRune = rune(32)
		}

		fgColor := tb.ColorRed
		if typeRune == CurrentPhraseRune {
			fgColor = tb.ColorGreen
		}

		tb.SetCell(pi, 1, typeRune, tb.ColorDefault, fgColor)
	}

	tb.Flush()
	typePhrase += string(typeRune)
	pi++
}

func removeLastKey() {
	if pi == 0 {
		return
	}
	pi--
	tb.SetCell(pi, 1, rune(32), tb.ColorDefault, tb.ColorDefault)
	tb.Flush()
	typePhrase = typePhrase[:len(typePhrase)-1]
}

func main() {
	err := tb.Init()
	if err != nil {
		panic(err)
	}
	defer tb.Close()

	tb.SetInputMode(tb.InputEsc)

eventLoop:
	for {
		if phraseComplete {
			loadNewPhrase(phrase)
		}

		switch ev := tb.PollEvent(); ev.Type {
		case tb.EventKey:
			// handle standard event key presses
			switch ev.Key {
			case tb.KeyCtrlC:
				break eventLoop
			case tb.KeyDelete, tb.KeyBackspace, tb.KeyBackspace2:
				removeLastKey()
			default:
				updatePhrase(ev)
			}

			if phrase == typePhrase {
				typePhrase = ``
				phraseComplete = true
			}
		case tb.EventResize:
		case tb.EventMouse:
		case tb.EventError:
			panic(ev.Err)
		default:

		}
		// printTypePhrase()
	}
}

func printTypePhrase() {
	for i := 0; i < 50; i++ {
		tb.SetCell(i, 3, rune(32), tb.ColorDefault, tb.ColorDefault)
	}
	for i := 0; i < 50; i++ {
		tb.SetCell(i, 4, rune(32), tb.ColorDefault, tb.ColorDefault)
	}
	for i := 0; i < len(phrase); i++ {
		tb.SetCell(i, 3, rune(phrase[i]), tb.ColorDefault, tb.ColorDefault)
	}
	for i := 0; i < len(typePhrase); i++ {
		tb.SetCell(i, 4, rune(typePhrase[i]), tb.ColorDefault, tb.ColorDefault)
	}
	tb.Flush()
}
