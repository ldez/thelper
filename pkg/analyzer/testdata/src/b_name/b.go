// Code generated by generator. DO NOT EDIT.

package b

import (
	"context"
	"testing"
)

func nonTestHelper(b int) {}

func helperWithoutHelper(b *testing.B) {} 

func helperWithHelper(b *testing.B) {
	b.Helper()
}

func helperWithEmptyStringBeforeHelper(b *testing.B) {

	b.Helper()
}

func helperWithHelperAfterAssignment(b *testing.B) { 
	_ = 0
	b.Helper()
}

func helperWithHelperAfterOtherCall(b *testing.B) { 
	f()
	b.Helper()
}

func helperWithHelperAfterOtherSelectionCall(b *testing.B) { 
	b.Fail()
	b.Helper()
}

func helperParamNotFirst(s string, i int, b *testing.B) { 
	b.Helper()
}

func helperParamSecondWithoutContext(s string, b *testing.B, i int) { 
	b.Helper()
}

func helperParamSecondWithContext(ctx context.Context, b *testing.B) {
	b.Helper()
}

func helperWithIncorrectName(o *testing.B) { // want "parameter \\*testing.B should have name b"
	o.Helper()
}

func helperWithAnonymousHelper(b *testing.B) {
	b.Helper()
	func(b *testing.B) {}(b) 
}

func helperWithNoName(_ *testing.B) {
}

func f() {}
