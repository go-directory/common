package common

import "fmt"

func ExampleKVStore() {
	// Initialize a new KVStore instance (REQUIRED!)
	kvs := NewKVStore()

	// OPTIONAL: Protect with mutex locking. Use of
	// false entirely disables locking.
	kvs.SetThreadSafe(true)

	// OPTIONAL: control index matching with respect
	// to case-folding. A value of false means fold
	// indices, while true means DO NOT fold indices.
	//
	// For the purpose of this example, we'll stick
	// with the default behavior.
	kvs.SetCaseExact(false)

	// Add some values. Each of the seven (7) allowed
	// types are demonstrated below.

	// bool
	kvs.Set[bool]("ReadOnly", true)

	// int64
	kvs.Set[int64]("Gauge", int64(-4775))

	// uint64
	kvs.Set[uint64]("Counter", uint64(137744))

	// float64
	kvs.Set[float64]("Pi", float64(3.14159))

	// []byte
	kvs.Set[[]byte]("Value", []byte(`a value`))

	// [][]byte
	kvs.Set[[][]byte]("Values", [][]byte{
		[]byte(`a value`),
		[]byte(`another value`),
		[]byte(`...`),
	})

	// Special case: this is an example of building
	// a "seen table", which is a construct used in
	// building unique lists of strings. For this
	// kind of entry, there is no functional value,
	// just an index (e.g.: "unique").
	kvs.SetEmpty("unique")

	// Determine the kind of type for an index.
	kind := kvs.Kind("readonly")

	var (
		expected uint64
		whatever any
		err      error
	)

	// Call a value we already know about using a string index.
	if expected, err = kvs.Get[uint64]("counter"); err != nil {
		fmt.Printf("Unable to retrieve counter: %v", err)
		return
	}

	// Call a value we have no knowledge of beyond the index.
	if whatever, err = kvs.GetAny("GAUGE"); err != nil {
		fmt.Printf("Unable to retrieve gauger: %v", err)
		return
	}

	// See how many entries reside within the receiver instance.
	indices := kvs.Len()

	// Print out some facts
	fmt.Printf("Is thread safe:  %t\n", kvs.IsThreadSafe())
	fmt.Printf("ReadOnly kind:   %d\n", kind)
	fmt.Printf("Expected value:  %T (%d)\n", expected, expected)
	fmt.Printf("Whatever value:  %T (%#v)\n", whatever, whatever)
	fmt.Printf("Num. indices:    %d\n", indices)
	fmt.Printf("'unique' exists: %t\n", kvs.Exists("unique"))

	fmt.Println("Freeing KVStore from memory")
	kvs.Free()
	fmt.Printf("KVStore freed:   %t\n", kvs.IsNil())

	// Output:
	// Is thread safe:  true
	// ReadOnly kind:   4
	// Expected value:  uint64 (137744)
	// Whatever value:  int64 (-4775)
	// Num. indices:    7
	// 'unique' exists: true
	// Freeing KVStore from memory
	// KVStore freed:   true
}
