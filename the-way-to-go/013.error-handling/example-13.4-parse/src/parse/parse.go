package parse

import (
    "fmt"
    "strings"
    "strconv"
)

// A ParseError indicates an error in converting a word into an integer.
type ParseError struct {
    Index int       // THe index into the space-separated list of words.
    Word  string    // The word that generated the parse error.
    Error error     // The raw error that precipitated this error, if any
}

// String returns a human-readable error message.
func (e *ParseError) String() string {
    return fmt.Sprintf("parse.Parse ParseError[Index=%d]: error parsing %q as int(Reason: %s)", e.Index, e.Word, e.Error.Error())
}

// Parse parses the space-separated words in in put as integers.
func Parse(input string) (numbers []int, err error) {
    defer func() {
        if r := recover(); r != nil {
            var ok bool
            err, ok = r.(error)
            if !ok {
                err = fmt.Errorf("%s", r)
            }
        }
    }()

    fields := strings.Fields(input)
    numbers = fields2numbers(fields) // here panic can occur
    return
}

func fields2numbers(fields []string) (numbers []int) {
    if len(fields) == 0 {
        panic("no words to parse")
    }
    for idx, field := range fields {
        num, err := strconv.Atoi(field)
        if err != nil {
            panic(&ParseError{idx, field, err})
        }
        numbers = append(numbers, num)
    }
    return
}