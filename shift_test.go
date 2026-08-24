package common

import (
	"fmt"
	"testing"
)

var testMap map[int]Kind

func ExampleNew() {
	bits := New(Uint16)
	fmt.Printf("%T size %d, max %d", bits, bits.Size(), bits.Max())
	// Output: common.BitValue size 16, max 65535
}

func ExampleBitValue_Shift() {
	bits := New(Uint8)
	bits.Shift(2, 4, 32)
	fmt.Printf("Value: %d", bits.Int())
	// Output: Value: 38
}

func ExampleBitValue_Unshift() {
	bits := New(Uint8)
	bits.Shift(2, 4, 32)
	bits.Unshift(32)
	fmt.Printf("Value: %d", bits.Int())
	// Output: Value: 6
}

func ExampleBitValue_Min() {
	var bits BitValue
	fmt.Printf("%d", bits.Min())
	// Output: 0
}

func ExampleBitValue_Max_for8Bit() {
	bits := New(Uint8)
	fmt.Printf("%d", bits.Max())
	// Output: 255
}

func ExampleBitValue_Max_for16Bit() {
	bits := New(Uint16)
	fmt.Printf("%d", bits.Max())
	// Output: 65535
}

func ExampleBitValue_Max_for32bit() {
	bits := New(Uint32)
	fmt.Printf("%d", bits.Max())
	// Output: 4294967295
}

func ExampleBitValue_Int() {
	bits := New(Uint8)
	bits.Shift(2, 4, 32)
	fmt.Printf("%d", bits.Int())
	// Output: 38
}

func ExampleBitValue_Int_mixed() {
	var ints []int
	for i := 0; i < 3; i++ {
		bits := New(Kind(i + 1))
		bits.Shift(bits.Size() << i)
		ints = append(ints, bits.Int())
	}
	fmt.Printf("%v", ints)
	// Output: [8 32 128]
}

func ExampleBitValue_Value() {
	bits := New(Uint32)
	bits.Shift(bits.Max())
	fmt.Printf("%T", bits.Value())
	// Output: *uint32
}

func ExampleBitValue_Kind() {
	bits := New(Uint32)
	fmt.Printf("%s", bits.Kind())
	// Output: uint32
}

func ExampleBitValue_Size() {
	bits := New(Uint32)
	fmt.Printf("%d", bits.Size())
	// Output: 32
}

func ExampleKind_Size() {
	k := Uint32
	fmt.Printf("%d", k.Size())
	// Output: 32
}

func ExampleKind_String() {
	fmt.Printf("%s", Uint32)
	// Output: uint32
}

func ExampleBitValue_Positive_uint8() {
	// user-defined shift values
	type B uint8
	const (
		Bopt1 B = 1 << iota //   1
		Bopt2               //   2
		Bopt3               //   4
		Bopt4               //   8
		Bopt5               //  16
		Bopt6               //  32
		Bopt7               //  64
		Bopt8               // 128	// go no higher (else, overflow uint8)
	)

	bits := New(Uint8)
	bits.Shift(Bopt1, Bopt3, Bopt6)
	fmt.Printf("Value contains B-options #6 (32): %t", bits.Positive(Bopt6))
	// Output: Value contains B-options #6 (32): true
}

func ExampleBitValue_Positive_int32() {
	// user-defined shift values
	type B int32
	const (
		Bopt1  B = 1 << iota // 1
		Bopt2                // 2
		Bopt3                // 4
		Bopt4                // 8
		Bopt5                // 16
		Bopt6                // 32
		Bopt7                // 64
		Bopt8                // 128
		Bopt9                // 256
		Bopt10               // 512
		Bopt11               // 1024
		Bopt12               // 2048
		Bopt13               // 4096
		Bopt14               // 8192
		Bopt15               // 16384
		Bopt16               // 32768
		Bopt17               // 65536
		Bopt18               // 131072
		Bopt19               // 262144
		Bopt20               // 524288
		Bopt21               // 1048576
		Bopt22               // 2097152
		Bopt23               // 4194304
		Bopt24               // 8388608
		Bopt25               // 16777216
		Bopt26               // 33554432
		Bopt27               // 67108864
		Bopt28               // 134217728
		Bopt29               // 268435456
		Bopt30               // 536870912
		Bopt31               // 1073741824 // go no higher (else, overflow int32)
	)

	bits := New(Uint32)
	bits.Shift(Bopt1, Bopt31, Bopt6)
	fmt.Printf("Value contains B-options #31 (1073741824): %t", bits.Positive(Bopt31))
	// Output: Value contains B-options #31 (1073741824): true
}

func ExampleBitValue_SetNamesMap() {
	// user-defined shift values
	bits := New(Uint8)

	type B uint8
	const (
		Bopt1 B = 1 << iota //   1
		Bopt2               //   2
		Bopt3               //   4
		Bopt4               //   8
		Bopt5               //  16
		Bopt6               //  32
		Bopt7               //  64
		Bopt8               // 128      // go no higher (else, overflow uint8)
	)

	// create a const->name map using
	// the above const vals by way of
	// incremental shifts. Note that
	// we could have just as easily
	// built the map manually, though
	// but looping is neater and more
	// succinct.
	var m map[int]string = make(map[int]string, 0)
	for i := 0; i < bits.Size(); i++ {
		str := fmt.Sprintf("bit_option%d", i+1) // inc. label number since we start at zero
		bv := 1 << i
		m[bv] = str
	}

	// assign map to receiver
	bits.SetNamesMap(m)

	// now shift by string name instead
	// of a constant directly :)
	name := `bit_option6`
	bits.Shift(name)

	fmt.Printf("Value contains %s: %t (val:%d)", name, bits.Positive(name), bits.Int())
	// Output: Value contains bit_option6: true (val:32)
}

func ExampleBitValue_None() {
	bits := New(Uint8)
	bits.Shift(8 << 1)
	bits.None() // annihilate any value

	fmt.Printf("%d", bits.Int())
	// Output: 0
}

func ExampleBitValue_All() {
	bits := New(Uint16)
	bits.All() // shift EVERYTHING

	fmt.Printf("%d", bits.Int())
	// Output: 65535
}

func ExampleBitValue_NamesMap() {
	bits := New(Uint8)
	fmt.Printf("%T", bits.NamesMap()) // note this is a nil map
	// Output: map[int]string
}

func TestBitValue_codecov(t *testing.T) {
	var bits BitValue
	bits.Kind()
	bits.Int()
	bits.Value()
	bits.Shift(-1)
	bits.Shift(8 << 8)
	bits.Unshift(-1)
	bits.Positive(-1)
	bits.Unshift(40000000000)
	if i := bits.Int(); i != 0 {
		t.Errorf("%s failed: bogus value set (%d) where none should be",
			t.Name(), i)
	}

	bits = New(Uint8)
	bits.Shift(bits.Max())
	bits.Shift(8 << 8)
	bits.Shift(8 << 1)
	bits.Positive(8 << 2)
	bits.Unshift(8 << 8)
	bits.Kind()
	bits.Int()
	bits.Value()

	for _, kind := range testMap {
		instance := New(kind)
		size := instance.Size()
		_ = kind.String()
		_ = bits.Int()
		for i := 0; i < size; i++ {
			instance.Shift(size << i)
			instance.Positive(size << i)
			instance.Unshift(size << i)
			instance.Shift(instance.Max())
			instance.Unshift(instance.Max())
			switch instance.Value().(type) {
			case *uint8:
				_, _ = toInt(uint8(size))
			case *uint16:
				_, _ = toInt(uint16(size))
			case *uint32:
				_, _ = toInt(uint32(size))
			}
		}
	}
}

func init() {
	testMap = map[int]Kind{
		8:  Uint8,
		16: Uint16,
		32: Uint32,
	}
}
