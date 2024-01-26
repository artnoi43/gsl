package gsl

import (
	"fmt"

	"github.com/soyart/gsl/data/list"
)

const (
	paren           = '('
	parenClose      = ')'
	curlyBrace      = '{'
	curlyBraceClose = '}'
	sqBracket       = '['
	sqBracketClose  = ']'
	chevron         = '<'
	chevronClose    = '>'
)

func IsOpenChar(openChar rune) bool {
	switch openChar {
	case
		paren,
		curlyBrace,
		sqBracket,
		chevron:
		return true
	}

	return false
}

func IsCloseChar(closeChar rune) bool {
	switch closeChar {
	case
		parenClose,
		curlyBraceClose,
		sqBracketClose,
		chevronClose:
		return true
	}

	return false
}

func CloseChar(openChar rune) (rune, error) {
	switch openChar {
	case paren:
		return parenClose, nil

	case curlyBrace:
		return curlyBraceClose, nil

	case sqBracket:
		return sqBracketClose, nil

	case chevron:
		return chevronClose, nil
	}

	return openChar, fmt.Errorf("unknown open symbol '%c'", openChar)
}

func OpenChar(closeChar rune) (rune, error) {
	switch closeChar {
	case parenClose:
		return paren, nil

	case curlyBraceClose:
		return curlyBrace, nil

	case sqBracketClose:
		return sqBracket, nil

	case chevronClose:
		return chevron, nil
	}

	return closeChar, fmt.Errorf("unknown close symbol '%c'", closeChar)
}

// IsWellClosed validates if all surrounds are properly closed
//
// # Examples (see unit test)
//
// 1. ([bar(foo)])  -> ok
//
// 2. [()[{}]]      -> ok
//
// 3. [((]          -> error
//
// 4. [)            -> error
//
// 5. [             -> error
//
// Notes: Only openning/closing characters are concerned.
func IsWellClosed(s string) error {
	stack := list.NewStackSafe[rune]()

	for _, char := range s {
		if IsOpenChar(char) {
			stack.Push(char)
			continue
		}

		if IsCloseChar(char) {
			charOpen, err := OpenChar(char)
			if err != nil {
				panic(fmt.Sprintf("unexpected gsl bug: missing open char for '%c'", char))
			}

			prev := stack.Pop()
			if prev == nil {
				return fmt.Errorf("unexpected close character '%c'", char)
			}

			if prevOpen := *prev; prevOpen != charOpen {
				expectedClose, err := CloseChar(prevOpen)
				if err != nil {
					panic(fmt.Sprintf("unexpected gsl bug: missing close char for '%c'", char))
				}

				return fmt.Errorf("unexpected close character '%c' - expecting '%c'", char, expectedClose)
			}
		}
	}

	if !stack.IsEmpty() {
		remaining := stack.Pop()

		return fmt.Errorf("unclosed character '%c'", *remaining)
	}

	return nil
}
