// Copyright 2015 mparaiso<mparaiso@online.fr>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package expect

import (
	"regexp"
	"strings"
	"testing"
)

type expectation struct {
	value interface{}
	test  *testing.T
}

type negativeExpectation struct {
	value interface{}
	test  *testing.T
}

type expectationBuilder struct {
	test *testing.T
}

// New returns a new expectationBuilder
// usefull if test suits contains many assertions
func New(t *testing.T) *expectationBuilder {
	return &expectationBuilder{t}
}

// Expect returns a new expectation
func (e *expectationBuilder) Expect(val interface{}) *expectation {
	return Expect(val, e.test)
}

// ToEqual expect 2 values to be equal
func (e *expectation) ToEqual(val interface{}) {
	if e.value != val {
		e.test.Errorf("%+v should be equal to %+v", e.value, val)
	}
}

// ToEqual expect 2 values to be equal
func (e *expectation) ToBe(val interface{}) {
	e.ToEqual(val)
}

// ToMatch expect a value to match a regular expression
func (e *expectation) ToMatch(val string) {
	if match, err := regexp.MatchString(val, e.value.(string)); err != nil {
		e.test.Error(err)
	} else if match == false {
		e.test.Errorf("%+v should match to %+v", e.value, val)
	}
}

// ToBeNil expect a value to be nil
func (e *expectation) ToBeNil() {
	if e.value != nil {
		e.test.Errorf("%+v should be nil", e.value)
	}
}

// ToBeTrue expect a value to be true
func (e *expectation) ToBeTrue() {
	if e.value.(bool) != true {
		e.test.Errorf("%+v should be true", e.value)
	}
}

// ToBeTrue expect a value to be false
func (e *expectation) ToBeFalse() {
	if e.value.(bool) != false {
		e.test.Errorf("%+v should be false", e.value)
	}
}

// ToContain expects a string to be a substring of value
func (e *expectation) ToContain(word string) {

	if strings.Contains(e.value.(string), word) == false {
		e.test.Errorf("%+v should contain %+v", e.value, word)
	}
}

// toBeLessThan expects value to be less than  number
func (e *expectation) toBeLessThan(number interface{}) {
	if toFloat64(e.value) >= toFloat64(number) {
		e.test.Errorf("%+v should be less then %+v", e.value, number)
	}
}

// ToBeGreaterThan expects value to be greater than number
func (e *expectation) ToBeGreaterThan(number interface{}) {
	if toFloat64(e.value) <= toFloat64(number) {
		e.test.Errorf("%+v should greater than %+v", e.value, number)
	}
}

// Not reverse expectations
func (e *expectation) Not() *negativeExpectation {
	return &negativeExpectation{e.value, e.test}
}

func (e *negativeExpectation) ToEqual(val interface{}) {
	if e.value == val {
		e.test.Errorf("%+v should not be equal to %+v", e.value, val)
	}
}
func (e *negativeExpectation) ToBe(val interface{}) {
	e.ToEqual(val)
}

func (e *negativeExpectation) ToMatch(val string) {
	if match, err := regexp.MatchString(val, e.value.(string)); err != nil {
		e.test.Error(err)
	} else if match == true {
		e.test.Errorf("%+v should not match to %+v", e.value, val)
	}
}

func (e *negativeExpectation) ToBeNil() {
	if e.value == nil {
		e.test.Errorf("%+v should not be nil", e.value)
	}
}

func (e *negativeExpectation) ToBeTrue() {
	if e.value.(bool) == true {
		e.test.Errorf("%+v should not be true", e.value)
	}
}

func (e *negativeExpectation) ToBeFalse() {
	if e.value.(bool) == false {
		e.test.Errorf("%+v should not be false", e.value)
	}
}

func (e *negativeExpectation) ToContain(word string) {

	if strings.Contains(e.value.(string), word) == true {
		e.test.Errorf("%+v should not contain %+v", e.value, word)
	}
}

func (e *negativeExpectation) toBeLessThan(number float64) {
	if toFloat64(e.value) < toFloat64(number) {
		e.test.Errorf("%+v should not be less than %+v", e.value, number)
	}
}

func (e *negativeExpectation) ToBeGreaterThan(number interface{}) {
	if toFloat64(e.value) > toFloat64(number) {
		e.test.Errorf("%+v should not be greater than %+v", e.value, number)
	}
}

// Expect returns a new expectation,
// usefull if a test suit contains only one assertion
func Expect(val interface{}, t *testing.T) *expectation {
	return &expectation{val, t}
}
