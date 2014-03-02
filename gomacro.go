// gomacro - a simple textual macro expansion library.  gomacro stores
// keys with text expansions, similar to the C preprocessor. It also
// supports function-like macros and concatenation.
//
// gomacro tokenizes on whitespace, and will recursively expand emitted
// text.
package gomacro

import (
	"strings"
)

type Macro struct {
	macros map[string]string
}

// Return a new, empty macro parsing object.
func NewMacro() *Macro {
	return &Macro{
		macros: make(map[string]string),
	}
}

// Add a new, or overwrites an existing macro definition.
func (m *Macro) Define(key, expansion string) {
	m.macros[key] = expansion
}

// Remove an existing macro definition.
func (m *Macro) Undefine(key string) {
	if _, ok := m.macros[key]; ok {
		delete(m.macros, key)
	}
}

// Return a list of macros currently set.
func (m *Macro) List() []string {
	var keys []string
	for k, _ := range m.macros {
		keys = append(keys, k)
	}
	return keys
}

// Return the macro text for a given key.
func (m *Macro) Macro(key string) string {
	return m.macros[key]
}

// Parse input text with set macros.
func (m *Macro) Parse(input string) string {
	// simple replace, preserves whitespace
	for k, v := range m.macros {
		output := strings.Replace(input, k, v, -1)
		if input == output {
			continue
		} else {
			return m.Parse(output)
		}
	}
	return input
}
