package xpath

import (
	"testing"
)

func TestElementTest(t *testing.T) {
	t.Run(`elementTest without attribute`, func(t *testing.T) {
		elementTest := newElementTest("a", nil, nil)

		t.Run(`test`, func(t *testing.T) {
			var tests = []struct {
				input interface{}
				want  bool
			}{
				{newText("some text", nil), false},
				{newElement("not-a", nil, nil), false},
				{newElement("not-a", []*attribute{newAttribute("key1", "value1")}, nil), false},
				{newElement("a", nil, nil), true},
				{newElement("a", []*attribute{newAttribute("key1", "value1"), newAttribute("key2", "value2")}, nil), true},
			}

			for _, test := range tests {
				if got := elementTest.test(test.input); got != test.want {
					t.Errorf(`et.test(%v)=%v | want %v`, test.input, got, test.want)
				}
			}
		})
	})

	t.Run(`elementTest with attribute`, func(t *testing.T) {
		elementTest := newElementTest("a", newAttribute("key-want", "value-want"), nil)

		t.Run(`test`, func(t *testing.T) {
			var tests = []struct {
				input interface{}
				want  bool
			}{
				{newText("some text", nil), false},
				{newElement("not-a", nil, nil), false},
				{newElement("not-a", []*attribute{newAttribute("key1", "value1")}, nil), false},
				{newElement("a", []*attribute{newAttribute("key1", "value1"), newAttribute("key2", "value2")}, nil), false},
				{newElement("a", []*attribute{newAttribute("key1", "value1"), newAttribute("key-want", "value-want")}, nil), true},
			}

			for _, test := range tests {
				if got := elementTest.test(test.input); got != test.want {
					t.Errorf(`et.test(%v)=%v | want %v`, test.input, got, test.want)
				}
			}
		})
	})

	t.Run(`elementTest with predicate`, func(t *testing.T) {
		predicate := newPredicate()
		elementTest := newElementTest("a", nil, predicate)

		t.Run(`predicate`, func(t *testing.T) {

			etPredicate := elementTest.predicate()
			if etPredicate == predicate {
				t.Errorf(`et.predicate() = et.pred | want a different predicate`)
			}
		})
	})
}

func TestTextTest(t *testing.T) {
	t.Run(`textTest with empty data`, func(t *testing.T) {
		textTest := newTextTest("")

		t.Run(`test`, func(t *testing.T) {
			var tests = []struct {
				input interface{}
				want  bool
			}{
				{newElement("a", nil, nil), false},
				{newText("some text", nil), true},
			}

			for _, test := range tests {
				if got := textTest.test(test.input); got != test.want {
					t.Errorf(`tt.test(%v)=%v | want %v`, test.input, got, test.want)
				}
			}
		})
	})

	t.Run(`textTest with data`, func(t *testing.T) {
		textTest := newTextTest("some text")

		t.Run(`test`, func(t *testing.T) {

			var tests = []struct {
				input interface{}
				want  bool
			}{
				{newElement("a", nil, nil), false},
				{newText("not some text", nil), false},
				{newText("some text", nil), true},
			}
			for _, test := range tests {
				if got := textTest.test(test.input); got != test.want {
					t.Errorf(`tt.test(%v)=%v | want %v`, test.input, got, test.want)
				}
			}
		})
	})
}
