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
			"Hello World",
			[]string{"Hello", "World"},
		},
	}

	for _, c := range cases {
		returnActual := cleanInput(c.argInput)

		// TEST: WHEN LENGTH DON'T MATCH
		// NOTE: this is not necessary if you use slices.Equal check
		if len(returnActual) != len(c.returnExpected) {
			t.Errorf("cleanInput(%q) = %v, want %v", c.argInput, returnActual, c.returnExpected)
			fmt.Println()
			fmt.Println("😡 FAIL")
			fmt.Println("=======================================================================")
		}

		// TEST: WHEN CONTENT DON'T MATCH V1
		if !slices.Equal(returnActual, c.returnExpected) {
			t.Errorf("cleanInput(%q) = %v, want %v", c.argInput, returnActual, c.returnExpected)
			fmt.Println()
			fmt.Println("😡 FAIL")
			fmt.Println("=======================================================================")
		}

		// TEST: WHEN CONTENT DON'T MATCH V2
		/*
			if !reflect.DeepEqual(returnActual, c.returnExpected) {
				t.Errorf("cleanInput(%q) = %v, want %v", c.argInput, returnActual, c.returnExpected)
			}
		*/

		fmt.Printf("cleanInput(%q) = %v, want %v", c.argInput, returnActual, c.returnExpected)
		fmt.Println()
		fmt.Println("🤤 SUCCESS")
		fmt.Println("=======================================================================")

	}
}
