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

// This package aims to test nilability behavior for `range` in loops.

// <nilaway no inference>
package looprange

import (
	"maps"
	"slices"
)

func dummyConsume(interface{}) {}
func dummyBool() bool          { return true }

// Test for checking range over slices.
// The below tests ensure that all forms of range loops correctly produce their ranging expression
// as non-nil - including limiting the scope of that production to within their loop bodies
func testRangeForSlices(a []*int) *int {
	for range a {
		// here and in following similar cases:
		// check that a can be indexed into (since it's nonnil)
		// validity of the index is outside the scope of NilAway, so we can use any index, e.g. 0
		return a[0]
	}
	if dummyBool() {
		return a[0] //want "sliced into"
	}
	for _ = range a {
		return a[0]
	}
	if dummyBool() {
		return a[0] //want "sliced into"
	}
	for i := range a {
		dummyConsume(i)
		return a[0]
	}
	if dummyBool() {
		return a[0] //want "sliced into"
	}
	for _, _ = range a {
		return a[0]
	}
	if dummyBool() {
		return a[0] //want "sliced into"
	}
	for i, _ := range a {
		dummyConsume(i)
		return a[0]
	}
	if dummyBool() {
		return a[0] //want "sliced into"
	}
	for _, j := range a {
		dummyConsume(j)
		return a[0]
	}
	if dummyBool() {
		return a[0] //want "sliced into"
	}
	for i, j := range a {
		dummyConsume(i)
		dummyConsume(j)
		return a[0]
	}
	if dummyBool() {
		return a[0] //want "sliced into"
	}
	var i2, j2 interface{}
	for i2, _ = range a {
		dummyConsume(i2)
		return a[0]
	}
	if dummyBool() {
		return a[0] //want "sliced into"
	}
	for _, j2 = range a {
		dummyConsume(j2)
		return a[0]
	}
	if dummyBool() {
		return a[0] //want "sliced into"
	}
	for i2, j2 = range a {
		dummyConsume(i2)
		dummyConsume(j2)
		return a[0]
	}
	if dummyBool() {
		return a[0] //want "sliced into"
	}
	i := 0
	return &i
}

// Test for checking range over arrays.
// nonnil(a[])
func testRangeForArrays(a [5]*int) *int {
	for range a {
		return a[0]
	}

	for i := range a {
		a[i] = &i
		if dummyBool() {
			return a[i]
		}
	}

	for _, v := range a {
		if dummyBool() {
			return v
		}
	}
	i := 0
	return &i
}

// Test for checking range over maps.
// nonnil(a, b) nilable(b[], d[])
func testRangeOverMaps(a, b, c, d map[int]*int) *int {
	switch 0 {
	case 1:
		for _, a_elem := range a {
			return a_elem
		}
	case 2:
		for _, b_elem := range b {
			return b_elem //want "returned"
		}
	case 3:
		for _, c_elem := range c {
			return c_elem
		}
	case 4:
		for _, d_elem := range d {
			return d_elem //want "returned"
		}
	}
	i := 0
	return &i
}

// Test for checking range over channels.
// nonnil(a, b) nilable(<-b, <-d)
func testRangeOverChannels(a, b, c, d chan *int) *int {
	switch 0 {
	case 1:
		for a_elem := range a {
			return a_elem
		}
	case 2:
		for b_elem := range b {
			return b_elem //want "returned"
		}
	case 3:
		for c_elem := range c {
			return c_elem
		}
	case 4:
		for d_elem := range d {
			return d_elem //want "returned"
		}
	}
	i := 0
	return &i
}

// Test for checking range over basic types, such as integers and strings.
func testRangeOverBasicTypes(j int) {
	switch j {
	case 0:
		for i := range 10 {
			print(i)
		}
	case 1:
		n := 10
		for i := range n {
			print(i)
		}
	case 2:
		for i := range "hello" {
			print(i)
		}
	case 3:
		s := "hello"
		for i := range s {
			print(i)
		}
	case 4:
		type intType int
		var s intType = 42
		for i := range s {
			print(i)
		}
	case 5:
		type intAlias int
		var s intAlias = 42
		for i := range s {
			print(i)
		}
	}
}

type Set map[string]bool

type MyAlias = Set

//nilable(s)
func testAlias(s MyAlias) {
	for myStr := range s {
		print(myStr)
	}
}

func testIter() {
	i := 42
	for element := range slices.Values([]*int{&i, &i, nil}) {
		print(*element) // FN: we do not really handle iterators for now, the elements from iterators are assumed to be nonnil.
	}
	for k, v := range maps.All(map[string]*int{"abc": &i, "def": nil}) {
		print(k)
		print(*v) // FN: we do not really handle iterators for now, the elements from iterators are assumed to be nonnil.
	}
}

// Custom iterator functions introduced in Go 1.23

func iteratorNoArgs(yield func() bool) {
	for i := 0; i < 3; i++ {
		if !yield() {
			break
		}
	}
}

func iteratorWithKey(yield func(*int) bool) {
	one := 1
	data := []*int{&one, nil}
	for _, i := range data {
		if !yield(i) {
			break
		}
	}
}

func iteratorWithKeyValue(yield func(int, *string) bool) {
	one := "one"
	data := map[int]*string{1: &one, 42: nil}
	for k, v := range data {
		if !yield(k, v) {
			break
		}
	}
}

func useCustomIters() {
	for range iteratorNoArgs {
	}

	for k := range iteratorWithKey {
		print(*k) // FN: we do not really handle iterators for now, the elements from iterators are assumed to be nonnil.
	}

	for k, v := range iteratorWithKeyValue {
		print(k)
		print(*v) // FN: we do not really handle iterators for now, the elements from iterators are assumed to be nonnil.
	}
}
