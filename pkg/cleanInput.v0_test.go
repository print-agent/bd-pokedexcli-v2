package pkg

import (
	"fmt"
	"slices"
	"testing"
)

type testCases []struct {
	argInput       string
	returnExpected []string
}

func TestCleanInput(t *testing.T) {
	cases := testCases{
		{
			"Hello, World",
			[]string{"hello", "world"},
		},
		{
			"  Test-Case 123!  ",
			[]string{"testcase"},
		},
		{
			"It's a won'derful day, isn't it?",
			[]string{"it's", "a", "won'derful", "day", "isn't", "it"},
		},
		{
			"!@#$",
			[]string{},
		},
		{
			"12345",
			[]string{},
		},
		{
			"Déjà-vu 2024",
			[]string{"déjàvu"},
		},
		{
			"   ",
			[]string{},
		},
		{
			"'Hello,' said the man.",
			[]string{"'hello'", "said", "the", "man"},
		},
		{
			"asdf1248",
			[]string{"asdf"},
		},
		{
			"124*A)(FST's",
			[]string{"afst's"},
		},
		{
			"Flabébé123!",
			[]string{"flabébé"},
		},
		{
			"''hello!''",
			[]string{"''hello''"},
		},
		{
			"Don't stop believin'",
			[]string{"don't", "stop", "believin'"},
		},
		{
			"They're goin' home.",
			[]string{"they're", "goin'", "home"},
		},
	}

	for _, c := range cases {
		returnActual := cleanInput(c.argInput)

		// TEST: WHEN LENGTH DON'T MATCH
		// NOTE: this is not necessary if you use slices.Equal check
		if len(returnActual) != len(c.returnExpected) {
			t.Errorf("cleanInput(%q) = %v, want %v", c.argInput, returnActual, c.returnExpected)
			t.Errorf("\n")
			t.Errorf("🤬 FAIL\n")
			t.Errorf("=======================================================================\n")
		}

		// TEST: WHEN CONTENT DON'T MATCH V1
		if !slices.Equal(returnActual, c.returnExpected) {
			t.Errorf("cleanInput(%q) = %v, want %v", c.argInput, returnActual, c.returnExpected)
			t.Errorf("\n")
			t.Errorf("😡 FAIL\n")
			t.Errorf("=======================================================================\n")
		} else {
			fmt.Printf("cleanInput(%q) = %v, want %v", c.argInput, returnActual, c.returnExpected)
			fmt.Println()
			fmt.Println("🤤 SUCCESS")
			fmt.Println("=======================================================================")
		}

		// TEST: WHEN CONTENT DON'T MATCH V2
		/*
			if !reflect.DeepEqual(returnActual, c.returnExpected) {
				t.Errorf("cleanInput(%q) = %v, want %v", c.argInput, returnActual, c.returnExpected)
			}
		*/
	}
}
