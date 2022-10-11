package main

import "testing"

func TestCountOfAtom(t *testing.T) {
	formula := "K4(ON(SO3)2)2"
	ans := countOfAtoms(formula)

	if ans != "K4N2O14S4" {
		t.Errorf("test failed, formula is:%s\n return is:%s", formula, ans)
	}
}
