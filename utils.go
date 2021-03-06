// Copyright 2015 mparaiso<mparaiso@online.fr>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package expect

import (
	"fmt"
	"strconv"
)

// tries to convert val to a float64, panics if it fails
func toFloat64(val interface{}) float64 {

	switch t := val.(type) {
	case int:
		return float64(t)
	case int8:
		return float64(t)
	case int16:
		return float64(t)
	case int32:
		return float64(t)
	case int64:
		return float64(t)
	case uint:
		return float64(t)
	case uint8:
		return float64(t)
	case uint16:
		return float64(t)
	case uint32:
		return float64(t)
	case uint64:
		return float64(t)
	case float32:
		return float64(t)
	case float64:
		return float64(t)
	case bool:
		if t == true {
			return float64(1)
		}
		return float64(0)
	case string:
		f, _ := strconv.ParseFloat(val.(string), 64)
		return f
	}

	panic(fmt.Sprintf("cannot convert %v to float64", val))

}
