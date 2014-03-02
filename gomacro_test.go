package gomacro_test

import (
	. "gomacro"
	"testing"
)

func TestDefine(t *testing.T) {
	m := NewMacro()
	m.Define("key", "value")
	got := m.Macro("key")
	want := "value"
	if got != want {
		t.Errorf("define got %v, wanted %v", got, want)
	}

	// also test getting something that shouldn't exist
	got = m.Macro("foo")
	want = ""
	if got != want {
		t.Errorf("Define got %v, wanted %v", got, want)
	}
}

func TestUndefine(t *testing.T) {
	m := NewMacro()
	m.Define("key", "value")
	m.Undefine("key")
	got := m.Macro("key")
	want := ""
	if got != want {
		t.Errorf("Undefine got %v, wanted %v", got, want)
	}
}

func TestList(t *testing.T) {
	m := NewMacro()
	m.Define("key", "value")
	m.Define("foo", "value")
	m.Define("bar", "value")
	got := m.List()
	want := []string{"key", "foo", "bar"}
	for i, v := range got {
		if v != want[i] {
			t.Errorf("List got %v, wanted %v", got, want)
		}
	}
}

func TestMacro(t *testing.T) {
	m := NewMacro()
	m.Define("key", "value")
	got := m.Macro("key")
	want := "value"
	if got != want {
		t.Errorf("Macro got %v, wanted %v", got, want)
	}
}

func TestParseReplace(t *testing.T) {
	m := NewMacro()
	m.Define("key", "value")
	got := m.Parse("this is my key, my key is this")
	want := "this is my value, my value is this"
	if got != want {
		t.Errorf("ParseReplace got %v, wanted %v", got, want)
	}
}

func TestParseReplaceRecursive(t *testing.T) {
	m := NewMacro()
	m.Define("key", "value")
	m.Define("value", "foo")
	got := m.Parse("this is my key, my value is this")
	want := "this is my foo, my foo is this"
	if got != want {
		t.Errorf("ParseReplace got %v, wanted %v", got, want)
	}
}
