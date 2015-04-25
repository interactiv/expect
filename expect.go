// Copyright 2015 mparaiso<mparaiso@online.fr>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package expect

import (
	"regexp"
	"strings"
	"testing"
)

// Expectation is an expectation to be tested
type Expectation struct {
	value interface{}
	test  *testing.T
}

// NegativeExpectation is a negative expectation to be tested
type NegativeExpectation struct {
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

// Expect returns a new Expectation
func (e *expectationBuilder) Expect(val interface{}) *Expectation {
	return Expect(val, e.test)
}

// ToEqual expects 2 values to be equal
func (e *Expectation) ToEqual(val interface{}) {
	if e.value != val {
		e.test.Errorf("%+v should be equal to %+v", e.value, val)
	}
}

func (e *Expectation) ToPanic() {
	defer func() {
		if err := recover(); err == nil {
			e.test.Errorf("%v should not panic", e.value)
		}
	}()
	e.value.(func())()
}

// ToBe expects 2 values to be equal
func (e *Expectation) ToBe(val interface{}) {
	e.ToEqual(val)
}

// ToMatch expects a value to match a regular expression
func (e *Expectation) ToMatch(val string) {
	if match, err := regexp.MatchString(val, e.value.(string)); err != nil {
		e.test.Error(err)
	} else if match == false {
		e.test.Errorf("%+v should match to %+v", e.value, val)
	}
}

// ToBeNil expects a value to be nil
func (e *Expectation) ToBeNil() {
	if e.value != nil {
		e.test.Errorf("%+v should be nil", e.value)
	}
}

// ToBeTrue expects a value to be true
func (e *Expectation) ToBeTrue() {
	if e.value.(bool) != true {
		e.test.Errorf("%+v should be true", e.value)
	}
}

// ToBeFalse expects a value to be false
func (e *Expectation) ToBeFalse() {
	if e.value.(bool) != false {
		e.test.Errorf("%+v should be false", e.value)
	}
}

// ToContain expects a string to be a substring of value
func (e *Expectation) ToContain(word string) {

	if strings.Contains(e.value.(string), word) == false {
		e.test.Errorf("%+v should contain %+v", e.value, word)
	}
}

// toBeLessThan expects value to be less than  number
func (e *Expectation) toBeLessThan(number interface{}) {
	if toFloat64(e.value) >= toFloat64(number) {
		e.test.Errorf("%+v should be less then %+v", e.value, number)
	}
}

// ToBeGreaterThan expects value to be greater than number
func (e *Expectation) ToBeGreaterThan(number interface{}) {
	if toFloat64(e.value) <= toFloat64(number) {
		e.test.Errorf("%+v should greater than %+v", e.value, number)
	}
}

// Not reverse expectations
func (e *Expectation) Not() *NegativeExpectation {
	return &NegativeExpectation{e.value, e.test}
}

func (e *NegativeExpectation) ToEqual(val interface{}) {
	if e.value == val {
		e.test.Errorf("%+v should not be equal to %+v", e.value, val)
	}
}
func (e *NegativeExpectation) ToBe(val interface{}) {
	e.ToEqual(val)
}

func (e *NegativeExpectation) ToMatch(val string) {
	if match, err := regexp.MatchString(val, e.value.(string)); err != nil {
		e.test.Error(err)
	} else if match == true {
		e.test.Errorf("%+v should not match to %+v", e.value, val)
	}
}

func (e *NegativeExpectation) ToBeNil() {
	if e.value == nil {
		e.test.Errorf("%+v should not be nil", e.value)
	}
}

func (e *NegativeExpectation) ToBeTrue() {
	if e.value.(bool) == true {
		e.test.Errorf("%+v should not be true", e.value)
	}
}

func (e *NegativeExpectation) ToBeFalse() {
	if e.value.(bool) == false {
		e.test.Errorf("%+v should not be false", e.value)
	}
}

func (e *NegativeExpectation) ToContain(word string) {

	if strings.Contains(e.value.(string), word) == true {
		e.test.Errorf("%+v should not contain %+v", e.value, word)
	}
}

func (e *NegativeExpectation) toBeLessThan(number float64) {
	if toFloat64(e.value) < toFloat64(number) {
		e.test.Errorf("%+v should not be less than %+v", e.value, number)
	}
}

func (e *NegativeExpectation) ToBeGreaterThan(number interface{}) {
	if toFloat64(e.value) > toFloat64(number) {
		e.test.Errorf("%+v should not be greater than %+v", e.value, number)
	}
}

func (e *NegativeExpectation) ToPanic() {
	defer func() {
		if err := recover(); err != nil {
			e.test.Errorf("%+v should not panic", e.value)
		}
	}()
	e.value.(func())()
}

// Expect returns a new Expectation,
// usefull if a test suit contains only one assertion
func Expect(val interface{}, t *testing.T) *Expectation {
	return &Expectation{val, t}
}
