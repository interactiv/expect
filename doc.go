// Copyright 2015 mparaiso<mparaiso@online.fr>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/*
expect provides expectations for go test suits

Usage

	import (
		"testing"
		"github.com/interactiv/expect"
	)

	func TestNew(t *testing.T) {
		e := New(t)
		e.Expect(1).ToEqual(1)
		e.Expect(2).Not().ToEqual(1)
		e.Expect(1).ToBe(1)
		e.Expect(2).Not().ToEqual(1)
		e.Expect("foo").ToMatch("\\w{1,3}")
		e.Expect("foo").Not().ToMatch("bar")
		type Foo struct{}
		var foo *Foo = nil
		e.Expect(foo).Not().ToBeNil()
		e.Expect(true).ToBeTrue()
		e.Expect(false).ToBeFalse()
		e.Expect(true).Not().ToBeFalse()
		e.Expect(false).Not().ToBeTrue()
		e.Expect("house").ToContain("ou")
		e.Expect("hotel").Not().ToContain("foo")
		e.Expect(1).toBeLessThan(2)
		e.Expect(1).Not().toBeLessThan(0)
		e.Expect(2).toBeGreaterThan(1)
		e.Expect(1).Not().toBeGreaterThan(2)
	}

*/
package expect
