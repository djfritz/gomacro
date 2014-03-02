// gomacro - a simple textual macro expansion library.  gomacro stores
// keys with text expansions, similar to the C preprocessor. It also
// supports function-like macros and concatenation.
//
// gomacro tokenizes on whitespace, and will recursively expand emitted
// text.
package gomacro

import (
	"regexp"
)

type Macro struct {
	macros map[string]*macro
}

type macro struct {
	expansion string
	re        *regexp.Regexp
}

// Return a new, empty macro parsing object.
func NewMacro() *Macro {
	return &Macro{
		macros: make(map[string]*macro),
	}
}

// Add a new, or overwrites an existing macro definition.
func (m *Macro) Define(key, expansion string) error {
	re, err := regexp.Compile(key)
	if err != nil {
		return err
	}
	m.macros[key] = &macro{
		expansion: expansion,
		re:        re,
	}
	return nil
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
	if v, ok := m.macros[key]; ok {
		return v.expansion
	}
	return ""
}

// Parse input text with set macros.
func (m *Macro) Parse(input string) string {
	for _, v := range m.macros {
		output := v.re.ReplaceAllString(input, v.expansion)
		if input != output {
			return m.Parse(output)
		}
	}
	return input
}
