// Code generated by generator. DO NOT EDIT.

package t

import (
	"context"
	"testing"
)

func nonTestHelper(t int) {}

func helperWithoutHelper(t *testing.T) {} 

func helperWithHelper(t *testing.T) {
	t.Helper()
}

func helperWithEmptyStringBeforeHelper(t *testing.T) {

	t.Helper()
}

func helperWithHelperAfterAssignment(t *testing.T) { 
	_ = 0
	t.Helper()
}

func helperWithHelperAfterOtherCall(t *testing.T) { 
	f()
	t.Helper()
}

func helperWithHelperAfterOtherSelectionCall(t *testing.T) { 
	t.Fail()
	t.Helper()
}

func helperParamNotFirst(s string, i int, t *testing.T) { // want "parameter \\*testing.T should be the first or after context.Context"
	t.Helper()
}

func helperParamSecondWithoutContext(s string, t *testing.T, i int) { // want "parameter \\*testing.T should be the first or after context.Context"
	t.Helper()
}

func helperParamSecondWithContext(ctx context.Context, t *testing.T) {
	t.Helper()
}

func helperWithIncorrectName(o *testing.T) { 
	o.Helper()
}

func helperWithAnonymousHelper(t *testing.T) {
	t.Helper()
	func(t *testing.T) {}(t) 
}

func helperWithNoName(_ *testing.T) {
}

func f() {}
