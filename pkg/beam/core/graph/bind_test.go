package graph

import (
	"testing"

	"reflect"

	"github.com/apache/beam/sdks/go/pkg/beam/core/funcx"
	"github.com/apache/beam/sdks/go/pkg/beam/core/typex"
	"github.com/apache/beam/sdks/go/pkg/beam/core/util/reflectx"
)

func TestBind(t *testing.T) {
	tests := []struct {
		In  []typex.FullType // Incoming Node type
		Fn  interface{}
		Out []typex.FullType // Outgoing Node type; nil == cannot bind
	}{
		{ // Direct
			[]typex.FullType{typex.NewW(typex.New(reflectx.Int))},
			func(int) int { return 0 },
			[]typex.FullType{typex.NewW(typex.New(reflectx.Int))},
		},
		{ // Direct w/ KV out
			[]typex.FullType{typex.NewW(typex.New(reflectx.Int))},
			func(int) (int, string) { return 0, "" },
			[]typex.FullType{typex.NewWKV(typex.New(reflectx.Int), typex.New(reflectx.String))},
		},
		{ // KV Emitter
			[]typex.FullType{typex.NewW(typex.New(reflectx.Int))},
			func(int, func(int, string)) {},
			[]typex.FullType{typex.NewWKV(typex.New(reflectx.Int), typex.New(reflectx.String))},
		},
		{ // Direct with optionals time/error
			[]typex.FullType{typex.NewW(typex.New(reflectx.Int))},
			func(typex.EventTime, int) (typex.EventTime, int, error) { return typex.EventTime{}, 0, nil },
			[]typex.FullType{typex.NewW(typex.New(reflectx.Int))},
		},
		{ // Emitter w/ optionals
			[]typex.FullType{typex.NewW(typex.New(reflectx.Int))},
			func(typex.EventTime, int, func(typex.EventTime, int)) error { return nil },
			[]typex.FullType{typex.NewW(typex.New(reflectx.Int))},
		},
		{
			[]typex.FullType{typex.NewW(typex.New(reflectx.Int))},
			func(string) int { return 0 },
			nil, // int cannot bind to string
		},
		{
			[]typex.FullType{typex.NewW(typex.New(reflectx.Int))},
			func(int, int) int { return 0 },
			nil, // int cannot bind to int x int
		},
		{ // Generic
			[]typex.FullType{typex.NewW(typex.New(reflectx.Int))},
			func(x typex.X) typex.X { return x },
			[]typex.FullType{typex.NewW(typex.New(reflectx.Int))},
		},
		{
			[]typex.FullType{typex.NewWKV(typex.New(reflectx.Int), typex.New(reflectx.String))},
			func(x typex.X) typex.X { return x },
			nil, // structural mismatch
		},
		{ // Generic swap
			[]typex.FullType{typex.NewWKV(typex.New(reflectx.Int), typex.New(reflectx.String))},
			func(x typex.X, y typex.Y) (typex.Y, typex.X) { return y, x },
			[]typex.FullType{typex.NewWKV(typex.New(reflectx.String), typex.New(reflectx.Int))},
		},
		{ // Side input (as singletons)
			[]typex.FullType{typex.NewW(typex.New(reflectx.Int8)), typex.NewW(typex.New(reflectx.Int16)), typex.NewW(typex.New(reflectx.Int32))},
			func(int8, int16, int32, func(string, int)) {},
			[]typex.FullType{typex.NewWKV(typex.New(reflectx.String), typex.New(reflectx.Int))},
		},
		{ // Side input (as slice and iter)
			[]typex.FullType{typex.NewW(typex.New(reflectx.Int8)), typex.NewW(typex.New(reflectx.Int16)), typex.NewW(typex.New(reflectx.Int32))},
			func(int8, []int16, func(*int32) bool, func(int8, []int16)) {},
			[]typex.FullType{typex.NewWKV(typex.New(reflectx.Int8), typex.New(reflect.SliceOf(reflectx.Int16)))},
		},
		{ // Generic side input (as iter and re-iter)
			[]typex.FullType{typex.NewW(typex.New(reflectx.Int8)), typex.NewW(typex.New(reflectx.Int16)), typex.NewW(typex.New(reflectx.Int32))},
			func(typex.X, func(*typex.Y) bool, func() func(*typex.T) bool, func(typex.X, []typex.Y)) {},
			[]typex.FullType{typex.NewWKV(typex.New(reflectx.Int8), typex.New(reflect.SliceOf(reflectx.Int16)))},
		},
		{ // Generic side output
			[]typex.FullType{typex.NewW(typex.New(reflectx.Int8)), typex.NewW(typex.New(reflectx.Int16)), typex.NewW(typex.New(reflectx.Int32))},
			func(typex.X, typex.Y, typex.Z, func(typex.X, []typex.Y), func(int), func(typex.Z)) {},
			[]typex.FullType{typex.NewWKV(typex.New(reflectx.Int8), typex.New(reflect.SliceOf(reflectx.Int16))), typex.NewW(typex.New(reflectx.Int)), typex.NewW(typex.New(reflectx.Int32))},
		},
		{ // Bind as (K, V) ..
			[]typex.FullType{typex.NewWKV(typex.New(reflectx.Int8), typex.New(reflectx.Int16))},
			func(int8, int16) int { return 0 },
			[]typex.FullType{typex.NewW(typex.New(reflectx.Int))},
		},
		{ // .. bind same input as (V, SI) ..
			[]typex.FullType{typex.NewW(typex.New(reflectx.Int8)), typex.NewW(typex.New(reflectx.Int16))},
			func(int8, int16) int { return 0 },
			[]typex.FullType{typex.NewW(typex.New(reflectx.Int))},
		},
		{ // .. and allow other SI forms ..
			[]typex.FullType{typex.NewW(typex.New(reflectx.Int8)), typex.NewW(typex.New(reflectx.Int16))},
			func(int8, func(*int16) bool) int { return 0 },
			[]typex.FullType{typex.NewW(typex.New(reflectx.Int))},
		},
		{ // .. which won't work as (K, V).
			[]typex.FullType{typex.NewWKV(typex.New(reflectx.Int8), typex.New(reflectx.Int16))},
			func(int8, func(*int16) bool) int { return 0 },
			nil,
		},
	}

	for _, test := range tests {
		fn, err := funcx.New(test.Fn)
		if err != nil {
			t.Errorf("Invalid Fn: %v", err)
			continue
		}
		_, _, _, actual, err := Bind(fn, test.In...)
		if err != nil {
			if test.Out == nil {
				continue // expected
			}
			t.Errorf("Bind(%v, %v) failed: %v", fn, test.In, err)
			continue
		}

		if !typex.IsEqualList(actual, test.Out) {
			t.Errorf("Bind(%v, %v) = %v, want %v", fn, test.In, actual, test.Out)
		}
	}
}