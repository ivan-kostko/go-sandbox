package MySendBox

import (
	"reflect"
	"testing"
)

func BenchmarkMap(b *testing.B) {
	m := make(map[reflect.Type]string, 5)
	m[reflect.TypeOf("")] = "string"
	//	m[reflect.TypeOf(5)] = "integer"
	//	m[reflect.TypeOf(5.001)] = "float"
	//	m[reflect.TypeOf(true)] = "bit"
	//	m[reflect.TypeOf([]byte(nil))] = "binary"

	fn := func(i interface{}) string { return m[reflect.TypeOf(i)] }

	b.ResetTimer()
	for n := 0; n <= b.N; n++ {
		s := fn("")
		//		s = fn(5)
		//		s = fn(5.001)
		// s = fn(true)
		//		s = fn([]byte(nil))
		s += ""
	}
	b.ReportAllocs()
}

func BenchmarkTypeSwitch(b *testing.B) {

	fn := func(i interface{}) string {
		switch i.(type) {
		case string:
			return "string"
			break
		case int:
			return "integer"
			break
		case float32:
			return "float"
			break
		case bool:
			return "bit"
			break
		case []byte:
			return "binary"
			break

		}
		return ""
	}

	b.ResetTimer()

	for n := 0; n <= b.N; n++ {
		s := fn("")
		//		s = fn(5)
		//		s = fn(5.001)
		//		s = fn(true)
		//		s = fn([]byte(nil))
		s += ""
	}
	b.ReportAllocs()

}
