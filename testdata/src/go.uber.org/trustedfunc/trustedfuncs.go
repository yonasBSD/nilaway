//  Copyright (c) 2023 Uber Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
This package aims to test any nilaway behavior specific to accomdating tests, such as the `github.com/stretchr/testify` library

<nilaway no inference>
*/
package trustedfunc

import (
	"errors"
	"os/exec"
	"testing"

	"stubs/github.com/stretchr/testify/assert"
	"stubs/github.com/stretchr/testify/require"
	"stubs/github.com/stretchr/testify/suite"
)

type any interface{}

func errs() (any, error) {
	return 0, nil
}

// nilable(param 0)
func consume(any) {}

var dummy bool

// nilable(x)
func testRequire(t *testing.T, x any, z any, m map[any]any) interface{} {
	switch 0.0 {
	case 1.0:
		return x //want "returned"
	case 1.1:
		return z
	case 2.0:
		require.NotNil(t, x)
		return x
	case 2.1:
		require.NotNil(t, z)
		return z
	case 2.2:
		require.Nil(t, x)
		return x //want "returned"
	case 2.3:
		require.Nil(t, z)
		// this is unreachable, so no diagnostics should be reported
		return z
	case 2.4:
		require.NotNilf(t, x, "mymsg: %s", "arg")
		return x
	case 2.5:
		require.NotNilf(t, z, "mymsg: %s", "arg")
		return z
	case 2.6:
		require.Nilf(t, x, "mymsg: %s", "arg")
		return x //want "returned"
	case 2.7:
		require.Nilf(t, z, "mymsg: %s", "arg")
		// this is unreachable, so no diagnostics should be reported
		return z
	case 3:
		y, err := errs()
		consume(err)
		return y //want "returned"
	case 4.0:
		y, err := errs()
		require.NoError(t, err)
		return y
	case 4.1:
		y, err := errs()
		require.Error(t, err)
		return y //want "returned"
	case 4.2:
		y, err := errs()
		require.NoErrorf(t, err, "mymsg: %s", "arg")
		return y
	case 4.3:
		y, err := errs()
		require.Errorf(t, err, "mymsg: %s", "arg")
		return y //want "returned"
	case 5:
		require.True(t, x != nil)
		return x
	case 6:
		require.True(t, nil != x)
		return x
	case 7:
		require.True(t, x == nil)
		return x //want "returned"
	case 8:
		require.True(t, nil == x)
		return x //want "returned"
	case 9:
		require.True(t, x != nil && dummy)
		return x
	case 10:
		require.True(t, x != nil || dummy)
		return x //want "returned"
	case 11:
		require.True(t, dummy && x != nil)
		return x
	case 12:
		require.True(t, dummy || x != nil)
		return x //want "returned"
	case 11.1:
		require.Truef(t, dummy && x != nil, "mymsg: %s", "arg")
		return x
	case 12.1:
		require.Truef(t, dummy || x != nil, "mymsg: %s", "arg")
		return x //want "returned"
	case 13:
		require.False(t, x != nil)
		return x //want "returned"
	case 14:
		require.False(t, nil != x)
		return x //want "returned"
	case 15:
		require.False(t, x == nil)
		return x
	case 16:
		require.False(t, nil == x)
		return x
	case 17:
		require.False(t, x == nil && dummy)
		return x //want "returned"
	case 16.1:
		require.Falsef(t, nil == x, "mymsg: %s", "arg")
		return x
	case 17.1:
		require.Falsef(t, x == nil && dummy, "mymsg: %s", "arg")
		return x //want "returned"
	case 18:
		require.False(t, x == nil || dummy)
		return x
	case 19:
		require.False(t, dummy && x == nil)
		return x //want "returned"
	case 20:
		require.False(t, dummy || x == nil)
		return x
	case 21:
		v, ok := m[0]
		require.True(t, ok)
		return v
	case 22:
		v, ok := m[0]
		require.False(t, ok)
		return v //want "returned"
	case 23:
		_, ok := m[x]
		require.True(t, ok)
		return m[x]
	case 24:
		_, ok := m[x]
		require.False(t, ok)
		return m[x] //want "returned"
	case 25:
		require.NotNil(t, m[x])
		return m[x]
	case 26:
		require.Nil(t, m[x])
		return m[x] //want "returned"
	case 27:
		require.True(t, m[x] != nil)
		return m[x]
	case 28:
		require.True(t, m[x] == nil)
		return m[x] //want "returned"
	case 29:
		mapOfSlice := map[any][]int{}
		require.True(t, len(mapOfSlice[x]) == 1)
		return mapOfSlice[x][0]
	case 30:
		mapOfSlice := map[any][]int{}
		require.True(t, len(mapOfSlice[x]) == 0)
		return mapOfSlice[x][0] //want "sliced into"
	}
	return 0
}

// nilable(a, b, c)
func testMultipleRequires(t *testing.T, a, b, c any) any {
	if dummy {
		return a //want "returned"
	}
	if dummy {
		return b //want "returned"
	}
	if dummy {
		return c //want "returned"
	}

	require.NotNil(t, a)

	if dummy {
		return a
	}
	if dummy {
		return b //want "returned"
	}
	if dummy {
		return c //want "returned"
	}

	require.NotNil(t, b)

	if dummy {
		return a
	}
	if dummy {
		return b
	}
	if dummy {
		return c //want "returned"
	}

	require.NotNil(t, c)

	if dummy {
		return a
	}
	if dummy {
		return b
	}
	if dummy {
		return c
	}
	return 0
}

func takesNonnil(interface{}) {}

func testBackToBack(t *testing.T) {
	var x, y any
	var err, err2 error
	x, err = errs()
	require.NoError(t, err)
	takesNonnil(x)
	x, err = errs()
	require.NoError(t, err)
	takesNonnil(x)
	y, err = errs()
	require.NoError(t, err)
	takesNonnil(y)
	x, err = errs()
	require.NoError(t, err)
	takesNonnil(x)
	y, err = errs()
	require.NoError(t, err)
	takesNonnil(y)
	y, err = errs()
	require.NoError(t, err)
	takesNonnil(y)

	x, err = errs()
	require.NoError(t, err)
	takesNonnil(x)
	x, err = errs()
	require.NoError(t, err)
	takesNonnil(x)
	x, err2 = errs()
	require.NoError(t, err2)
	takesNonnil(x)
	x, err = errs()
	require.NoError(t, err)
	takesNonnil(x)
	x, err2 = errs()
	require.NoError(t, err2)
	takesNonnil(x)
	x, err2 = errs()
	require.NoError(t, err2)
	takesNonnil(x)
	x, err2 = errs()
	require.NoErrorf(t, err2, "mymsg: %s", "arg")
	takesNonnil(x)
}

// test for embedded testify package `suite` at depth 1
type testSetupEmbeddedDepth1 struct {
	suite.Suite
}

func (u *testSetupEmbeddedDepth1) testSuiteDepth1() any {
	response, err := errs()
	u.Nil(err)
	u.NotNil(response)
	return response
}

// nilable(x)
func (u *testSetupEmbeddedDepth1) testAmbiguity(t *testing.T, x *int) *int {
	// We have two kinds of ways to denote a variable as "not nil" in tests:
	// (1) top-level function `assert.NotNil(t *testing.T, x any, msgAndArgs ...any)` where the
	//     first argument is the API object for tests, and the second is the variable that we want to
	//     ensure it is not nil.
	// (2) method `suite.Suite.NotNil(x any, msgAndArgs ...any)` where the first argument is the
	//     variable to be nonnil and the second is an optional format string to report when the
	//     check fails. Note that the API object `T` does not need to be passed since it is passed
	//     in when the suite.Suite struct is constructed.

	// Now, in the following compilable but incorrect code, the developers falsely assumed they are
	// calling the top-level function, where they are actually calling the `suite.Suite` method.
	// NilAway should not be confused and assert that `x` is nonnil.

	// The first error is for passing nilable x to the `msgAndArgs` argument.
	u.NotNil(t, x) //want "passed"
	// The second error is that x is still nilable (u.NotNil does not really do anything).
	return x //want "returned"
}

// test for embedded testify package `suite` at depth 4
type testSetupEmbeddedDepth4 struct {
	embeddedDepth2
	f1 *int
}

type embeddedDepth2 struct {
	embeddedDepth3
	f2 string
}

type embeddedDepth3 struct {
	testSetupEmbeddedDepth1
}

func (u *testSetupEmbeddedDepth4) testSuiteDepth4() any {
	response, err := errs()
	u.NotNil(err)
	u.Nil(response)
	return response //want "returned"
}

// test for field of type testify package `suite` at depth 2
type testSetupFieldDepth2 struct {
	depthField2
}

type depthField2 struct {
	s suite.Suite
}

// nilable(x)
func (u *testSetupFieldDepth2) testSuiteFieldDepth2(x *int) {
	u.s.NotNil(x)
	takesNonnil(x)
}

// test for checking if NilAway can correctly function with checking only one of the two, `nil` or `nonnil`, operations
func (u *testSetupEmbeddedDepth1) testSuiteOnlyNil() any {
	response, err := errs()
	u.Nil(err)
	return response
}

func (u *testSetupEmbeddedDepth1) testSuiteOnlyNonnil() any {
	response, _ := errs()
	u.NotNil(response)
	return response
}

// Below tests check our handling of arbitrarily deeply nested structs (e.g., depth = 5 and depth = 6).

//						A								Depth = 1
//					/		\
//				B				C						Depth = 2
//			/	|	\		/		\
//			D	E	F		G		H					Depth = 3
//								/	/	\	\
//								I	J	K	L			Depth = 4
//										|
//									Suite				Depth = 5

type K struct {
	suite.Suite
}

type H struct {
	I
	j J
	K
	L
}

type C struct {
	g G
	H
}

type A struct {
	b B
	C
}

type B struct {
	d D
	e E
	f F
}

type I struct{}
type J struct{}
type L struct{}
type G struct{}
type D struct{}
type E struct{}
type F struct{}

func (a *A) testMaxDepthOf5() any {
	response, err := errs()
	a.Nil(err)
	a.NotNil(response)
	return response
}

//						Z								Depth = 1
//						|
//						A								Depth = 2
//					/		\
//				B				C						Depth = 3
//			/	|	\		/		\
//			D	E	F		G		H					Depth = 4
//								/	/	\	\
//								I	J	K	L			Depth = 5
//										|
//									Suite				Depth = 6

type Z struct {
	A
}

func (z *Z) testDepthOf6() any {
	response, err := errs()
	z.Nil(err)
	z.NotNil(response)
	return response
}

func (z *Z) testDepthOf6f() any {
	response, err := errs()
	z.Nilf(err, "mymsg: %s", "arg")
	z.NotNilf(response, "mymsg: %s", "arg")
	return response
}

// Similar to `suite.Suite`, `assert` package provides a struct `assert.Assertions` that have
// similar functions (in fact, `suite.Suite` embeds `assert.Assertions` in its implementation). So
// we should test `assert.Assertions` as well.

// Test embedding `assert.Assertions`
type testEmbeddedAssertionStruct struct {
	*assert.Assertions
}

// nilable(x, a)
func (u *testEmbeddedAssertionStruct) testEmbeddedAssertion(x *int, a []int, i int) *int {
	switch i {
	case 0:
		u.Greater(len(a), 0)
		print(a[0])
	case 1:
		u.GreaterOrEqual(len(a), 0)
		print(a[0]) //want "sliced into"
	case 2:
		u.Len(a, 1)
		print(a[0])
	case 3:
		u.Lenf(a, 0, "mymsg: %s", "arg")
		print(a[0]) //want "sliced into"
	case 4:
		u.Less(len(a), 1)
		print(a[0]) //want "sliced into"
	case 5:
		u.Less(1, len(a))
		print(a[0])
	case 6:
		u.LessOrEqualf(1, len(a), "msg", "arg")
		print(a[0])
	}

	u.NotNil(x)
	return x
}

// Test passing `assert.Assertions` as argument.
// nilable(x, s)
func testHelper(a *assert.Assertions, x *int, s []int, i int) *int {
	switch i {
	case 0:
		a.Greater(len(s), 0)
		print(s[0])
	case 1:
		a.GreaterOrEqual(len(s), 0)
		print(s[0]) //want "sliced into"
	case 2:
		a.Len(s, 1)
		print(s[0])
	case 3:
		a.Len(s, 0)
		print(s[0]) //want "sliced into"
	case 4:
		a.Less(len(s), 1)
		print(s[0]) //want "sliced into"
	case 5:
		a.Less(1, len(s))
		print(s[0])
	case 6:
		a.LessOrEqualf(1, len(s), "msg", "arg")
		print(s[0])
	}

	a.NotNil(x)
	return x
}

// To make it more complicated, we shadow the name `assert`.
// nilable(x, s)
func testShadow(assert *assert.Assertions, x *int, s []int, i int) *int {
	switch i {
	case 0:
		assert.Greater(len(s), 0)
		print(s[0])
	case 1:
		assert.GreaterOrEqual(len(s), 0)
		print(s[0]) //want "sliced into"
	case 2:
		assert.Len(s, 1)
		print(s[0])
	case 3:
		assert.Len(s, 0)
		print(s[0]) //want "sliced into"
	case 4:
		assert.Less(len(s), 1)
		print(s[0]) //want "sliced into"
	case 5:
		assert.Less(1, len(s))
		print(s[0])
	case 6:
		assert.LessOrEqualf(1, len(s), "msg", "arg")
		print(s[0])
	}

	// We shouldn't mistake `assert` as the package `assert` here.
	assert.NotNil(x)
	return x
}

// test for checking trusted functions through a call chain
func (s *testSetupEmbeddedDepth1) testCallChain(i int) interface{} {
	v, err := errs()
	switch i {
	case 0:
		s.Require().NoError(err)
		return v
	case 1:
		s.Require().NotNil(v)
		return v
	case 2:
		s.Require().Error(err)
		return v //want "returned"
	case 3:
		s.Require().Nil(v)
		return v //want "returned"
	case 4:
		s.Assert().NoError(err)
		return v
	case 5:
		s.Assert().NotNil(v)
		return v
	case 6:
		s.Assert().Error(err)
		return v //want "returned"
	case 7:
		s.Assert().Nil(v)
		return v //want "returned"
	case 8:
		var a []int
		s.Require().Greater(len(a), 0)
		print(a[0])
	case 9:
		var a []int
		s.Require().GreaterOrEqual(len(a), 0)
		print(a[0]) //want "sliced into"
	case 10:
		var a []int
		s.Assert().Less(len(a), 1)
		print(a[0]) //want "sliced into"
	case 11:
		var a []int
		s.Assert().Less(1, len(a))
		print(a[0])
	}
	return 0
}

// test for checking longer access paths
type W struct {
	y    *Y
	wptr *W
}

func (w *W) x() *W {
	return w.wptr
}

type Y struct {
	suite.Suite
}

func (y *Y) z() *require.Assertions {
	return y.Require()
}

func testLongerAccessPath(w *W) any {
	var a []int
	w.x().y.z().Len(a, 1)
	print(a[0])

	response, err := errs()
	w.x().y.z().NoError(err)
	return response
}

// nilable(a)
func testEqual(t *testing.T, i int, a []int) interface{} {
	switch i {
	case 0:
		require.Equal(t, len(a), 1)
		print(a[0])
	case 1:
		require.Equal(t, len(a), 0)
		print(a[0]) //want "sliced into"

	// Swapping the positions of args should not affect the analysis.
	case 2:
		require.Equal(t, 1, len(a))
		print(a[0])
	case 3:
		require.Equal(t, 0, len(a))
		print(a[0]) //want "sliced into"

	// Using a constant is also OK.
	case 4:
		const zero = 0
		require.Equal(t, zero, len(a))
		print(a[0]) //want "sliced into"

	// We can reason constant value without problems (thanks to constant folding in Go's type checker).
	case 5:
		const zero = 0
		require.Equal(t, zero+1-1, len(a))
		print(a[0]) //want "sliced into"
	case 6:
		const one = 1
		require.Equal(t, one-1+1, len(a))
		print(a[0])

	// The f variant should also be supported.
	case 7:
		require.Equalf(t, 1, len(a), "mymsg: %s", "arg")
		print(a[0])

	// The NotEqual variant should also be supported.
	case 81:
		require.NotEqual(t, len(a), 0)
		print(a[0])

	case 83:
		// Swapping the positions of args should not affect the analysis.
		require.NotEqual(t, 0, len(a))
		print(a[0])

	case 84:
		// Using a constant is also OK.
		const zero = 0
		require.NotEqual(t, zero, len(a))

	case 82:
		// `len(a) != 1` implies that len(a) can be 0, hence we should report an error.
		require.NotEqual(t, len(a), 1)
		print(a[0]) //want "sliced into"

	case 85:
		require.NotEqual(t, len(a), 1)
		print(a[0]) //want "sliced into"

	// Equal/NotEqual with nil should also be supported.
	case 8:
		x, err := errs()
		require.Equal(t, err, nil)
		return x

	case 9:
		x, err := errs()
		require.Equal(t, nil, err)
		return x

	case 10:
		x, err := errs()
		require.NotEqual(t, err, nil)
		return x //want "result 0 of `errs.*` lacking guarding"

	case 11:
		x, err := errs()
		require.NotEqual(t, nil, err)
		return x //want "result 0 of `errs.*` lacking guarding"

	// test with suite.Suite
	case 12:
		x, err := errs()
		s := &testSetupEmbeddedDepth1{}
		s.Equal(nil, err)
		return x

	case 13:
		x, err := errs()
		s := &testSetupEmbeddedDepth1{}
		s.NotEqual(err, nil)
		return x //want "result 0 of `errs.*` lacking guarding"

	case 14:
		var x *int
		require.NotEqual(t, nil, x)
		print(*x)

	case 15:
		var x *int
		require.Equal(t, nil, x)
		print(*x) //want "unassigned variable `x` dereferenced"

	case 16:
		var x *int
		require.Equal(t, x, nil)
		print(*x) //want "unassigned variable `x` dereferenced"

	case 17:
		var x *int
		require.NotEqual(t, x, nil)
		print(*x)
	}
	return 0
}

// nilable(a)
func testGreater(t *testing.T, i int, a []int) {
	switch i {
	case 0:
		require.Greater(t, len(a), 0)
		print(a[0])

	// Swapping the position of args is _not_ OK: `1 > len(a)` does not imply `a != nil`.
	case 1:
		require.Greater(t, 1, len(a))
		print(a[0]) //want "sliced into"

	// Admittedly weird, but you can assert `len(a) > -1`, and that will not imply the nilability of `a`.
	case 2:
		require.Greater(t, len(a), -1)
		print(a[0]) //want "sliced into"

	// The f variant should be supported.
	case 3:
		require.Greaterf(t, len(a), 0, "mymsg: %s", "arg")
		print(a[0])

	// GreaterOrEqual has slightly different semantics, we should capture that.
	case 4:
		// len(a) could be 0, so this does not imply the nilability of `a`.
		require.GreaterOrEqual(t, len(a), 0)
		print(a[0]) //want "sliced into"
	case 5:
		// len(a) >= 1 => len(a) > 0, so it is OK.
		require.GreaterOrEqual(t, len(a), 1)
		print(a[0])

	// Again, swapping the positions of args is _not_ OK.
	case 6:
		require.GreaterOrEqual(t, 1, len(a))
		print(a[0]) //want "sliced into"

	// The f variants should also be supported.
	case 7:
		// len(a) could be 0, so this does not imply the nilability of `a`.
		require.GreaterOrEqualf(t, len(a), 0, "mymsg: %s", "arg")
		print(a[0]) //want "sliced into"
	case 8:
		// len(a) >= 1 => len(a) > 0, so it is OK.
		require.GreaterOrEqualf(t, len(a), 1, "mymsg: %s", "arg")
		print(a[0])
	case 9:
		const zero = 0
		require.Greater(t, len(a), 1+zero-1)
		print(a[0])
	}
}

// nilable(a)
func testLess(t *testing.T, i int, a []int) {
	// This is basically a symmetric test suite to the "greater" one.
	switch i {
	case 0:
		require.Less(t, 1, len(a))
		print(a[0])

	// Swapping the position of args is _not_ OK: `len(a) < 1` does not imply the nilability of `a`.
	case 1:
		require.Less(t, len(a), 1)
		print(a[0]) //want "sliced into"

	// The f variant should be supported.
	case 2:
		require.Lessf(t, 1, len(a), "mymsg: %s", "arg")
		print(a[0])

	// LessOrEqual has slightly different semantics, we should capture that.
	case 3:
		// len(a) could be 0, so this does not imply the nilability of `a`.
		require.LessOrEqual(t, 0, len(a))
		print(a[0]) //want "sliced into"
	case 4:
		// 1 <= len(a) => len(a) > 0, so it is OK.
		require.LessOrEqual(t, 1, len(a))
		print(a[0])

	// Again, swapping the positions of args is _not_ OK.
	case 5:
		require.LessOrEqual(t, len(a), 1)
		print(a[0]) //want "sliced into"

	// The f variants should also be supported.
	case 6:
		// len(a) could be 0, so this does not imply the nilability of `a`.
		require.LessOrEqualf(t, 0, len(a), "mymsg: %s", "arg")
		print(a[0]) //want "sliced into"
	case 7:
		// len(a) >= 1 => len(a) > 0, so it is OK.
		require.LessOrEqualf(t, 1, len(a), "mymsg: %s", "arg")
		print(a[0])

	case 8:
		const zero = 0
		require.Less(t, 1+zero-1, len(a))
		print(a[0])
	}
}

// nilable(a)
func testLen(t *testing.T, i int, a []int) {
	switch i {
	case 0:
		require.Len(t, a, 1)
		print(a[0])
	case 1:
		require.Len(t, a, 0)
		print(a[0]) //want "sliced into"
	case 2:
		const zero = 0
		const one = 1
		require.Len(t, a, 1+zero-1+one)
		print(a[0])

	// The f variant should also be supported.
	case 3:
		const zero = 0
		const one = 1
		require.Lenf(t, a, 1+zero-1+one, "mymsg: %s", "arg")
		print(a[0])
	case 4:
		const zero = 0
		require.Lenf(t, a, 1+zero-1, "mymsg: %s", "arg")
		print(a[0]) //want "sliced into"
	}
}

// nilable(a)
func testEmpty(t *testing.T, i int, a []int, mp map[int]*int) interface{} {
	switch i {
	// zero value with slice len check should be supported
	case 0:
		require.Empty(t, a)
		print(a[0]) //want "sliced into"
	case 1:
		require.NotEmpty(t, a)
		print(a[0])

	// zero value of pointer check should be supported
	case 2:
		var x *int
		require.NotEmpty(t, x)
		print(*x)
	case 3:
		var x *int
		require.Empty(t, x)
		print(*x) //want "unassigned variable `x` dereferenced"
	case 4:
		x, err := errs()
		require.Empty(t, err)
		return x
	case 5:
		x, err := errs()
		require.NotEmpty(t, err)
		return x //want "result 0 of `errs.*` lacking guarding"

	// zero value of boolean check should be supported
	case 6:
		v, ok := mp[0]
		require.NotEmpty(t, ok)
		print(*v)
	case 7:
		v, ok := mp[0]
		require.Empty(t, ok)
		print(*v) //want "deep read from parameter `mp` lacking guarding"

	// The f variants should also be supported.
	case 8:
		require.Emptyf(t, a, "mymsg: %s", "arg")
		print(a[0]) //want "sliced into"
	case 9:
		require.NotEmptyf(t, a, "mymsg: %s", "arg")
		print(a[0])

	// Calls with `suite.Suite` should also be supported.
	case 10:
		x, err := errs()
		s := &testSetupEmbeddedDepth1{}
		s.Empty(err)
		return x
	case 11:
		var x *int
		s := &testSetupEmbeddedDepth1{}
		s.Empty(x)
		print(*x) //want "unassigned variable `x` dereferenced"
	case 12:
		v, ok := mp[0]
		s := &testSetupEmbeddedDepth1{}
		s.Empty(ok)
		print(*v) //want "deep read from parameter `mp` lacking guarding"
	case 13:
		s := &testSetupEmbeddedDepth1{}
		s.Empty(a)
		print(a[0]) //want "sliced into"
	}

	return 0
}

// nilable(err)
func errorsAs(err error, num string, dummy bool) {
	switch num {
	case "simple":
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			print(*exitErr)
		}
		print(*exitErr) //want "unassigned variable `exitErr` dereferenced"
	case "not in if block":
		var exitErr *exec.ExitError
		// Not checking the result of `errors.As` would not guard the variable.
		errors.As(err, &exitErr)
		print(*exitErr) //want "unassigned variable `exitErr` dereferenced"
	case "two errors connected by AND":
		var exitErr, anotherErr *exec.ExitError
		if errors.As(err, &exitErr) && errors.As(err, &anotherErr) {
			print(*exitErr)
			print(*anotherErr)
		}
	case "errors.As with other conditionals connected by AND":
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) && dummy {
			print(*exitErr)
		}
	case "errors.As with other conditionals connected by OR":
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) || dummy {
			print(*exitErr) //want "unassigned variable `exitErr` dereferenced"
		}
	case "two errors connected by OR":
		var exitErr, anotherErr *exec.ExitError
		if errors.As(err, &exitErr) || errors.As(err, &anotherErr) {
			// We do not know the nilability of either.
			print(*exitErr)    //want "unassigned variable `exitErr` dereferenced"
			print(*anotherErr) //want "unassigned variable `anotherErr` dereferenced"
		}
	case "nil dereference in first argument":
		var exitErr *exec.ExitError
		var nilError *error
		if errors.As(*nilError, &exitErr) { //want "unassigned variable `nilError` dereferenced"
			print(*exitErr) // But this is fine!
		}
	case "short assignment in if statement":
		var exitErr *exec.ExitError
		var nilError error
		if ok := errors.As(nilError, &exitErr); ok {
			print(*exitErr)
			print(ok)
		}
	case "short assignment in if statement with OR condition":
		var exitErr *exec.ExitError
		var nilError error
		if ok := errors.As(nilError, &exitErr); ok || dummy {
			print(*exitErr) //want "unassigned variable `exitErr` dereferenced"
			print(ok)
		}
	case "short assignment in if statement with AND condition (ok && dummy)":
		var exitErr *exec.ExitError
		var nilError error
		if ok := errors.As(nilError, &exitErr); ok && dummy {
			print(*exitErr)
			print(ok)
		}
	case "short assignment in if statement with AND condition (dummy && ok)":
		var exitErr *exec.ExitError
		var nilError error
		if ok := errors.As(nilError, &exitErr); dummy && ok {
			// The following is an FP: since the `dummy && ok` is canonicalized to `if dummy { if ok { } }`,
			// our current handling fails to find the `ok := errors.As(...)` call at the end of the block.
			print(*exitErr) //want "unassigned variable `exitErr` dereferenced"
			print(ok)
		}
	}
}

type S struct {
	f *int
}

type myErr struct{}

func (myErr) Error() string { return "myErr message" }

func NewS() (*S, error) {
	if dummy {
		return &S{}, nil
	}
	return nil, myErr{}
}

type SSuite struct {
	suite.Suite
	S *S
}

func (s *SSuite) TestFieldAssignment() {
	var err error
	// TODO: this is a false positive. Fix it by adding support for this case.
	//  Note that this is already being suppressed in the inference mode.
	s.S, err = NewS() //want "lacking guarding"
	s.NoError(err)
	print(s.S.f)
}
