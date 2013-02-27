/*
   Package regexp implements a simple library for
   regular expressions.

   The syntax of the regular expressions accepted is:

   regexp:
       concatenation { '|' concatenation }
   concatenation:
       { closure }
   closure:
       term [ '*' | '+' | '?' ]
   term:
       '^'
       '$'
       '.'
       character
       '[' [ '^' ] character-ranges ']'
       '(' regexp ')'
*/
package effective

import (
	"regexp"
	"sync"
)

// Compile parses a regular expression and returns, if successful, a Regexp
// object that can be used to match against text.
func Compile(str string) (regexp *Regexp, err error) {
	return nil, nil
}

// Error codes returned by failures to parse an expression.
var (
	ErrInternal      = errors.New("regexp: internal error")
	ErrUnmatchedLpar = errors.New("regexp: unmatched '('")
	ErrUnmatchedRpar = errors.New("regexp: unmatched ')'")
)

var (
	countLock   sync.Mutex
	inputCount  uint32
	outputCount uint32
	errorCount  uint32
)
