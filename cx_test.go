package css

import "testing"

func TestNextIsNotEmpty(t *testing.T) {
	if n := Next(0); n == "" {
		t.Fail()
	}
}

func TestShortestPossible(t *testing.T) {
	for i := 0; i < 26; i++ {
		if len(Next(i)) != 1 {
			t.Fail()
		}
	}
}

func TestNextDoesNotRepeat(t *testing.T) {
	check := map[string]bool{}
	for i := 0; i < 100000; i++ {
		name := Next(i)
		if check[name] {
			t.Logf("Repeated %s at inedx %v", name, i)
			t.Fail()
		}
		check[name] = true
	}
}