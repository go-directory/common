package common

import (
	"sync"
)

/*
KVStorable-aligned type flags.
*/
const (
	noneKVType       = iota // 0 (nil, not found)
	intKVType               // 1 (int64)
	uintKVType              // 2 (uint64)
	floatKVType             // 3 (float64)
	boolKVType              // 4 (bool)
	bytesKVType             // 5 ([]byte)
	bytesSliceKVType        // 6 ([][]byte)
	presKVType              // 7 (struct{})
)

/*
KVStorable implements a type classification encompassing the following
concrete types:

  - int64
  - uint64
  - float64
  - bool
  - []byte
  - [][]byte
  - struct{}  (ex.: "seen tables")
*/
type KVStorable interface {
	~int64 | ~uint64 | ~float64 | ~bool | ~[]byte | ~[][]byte | ~struct{}
}

/*
KVStore implements a general-use "key/value" storage type with generics support.

This type is suitable for use-cases in which the creation of a hidden ephemeral
DIT would be overkill, such as simple runtime settings for a particular feature
or single operation. It can also serve as a "seen table" for uniqueness control.

Instances of this type MUST be initialized using the [NewKVStore] helper. Any
stored value must implement a [KVStorable] type.

To write a new key/value pair to the receiver instance, see the [KVStore.Set]
method. To write an empty entry -- that is, a key with no value -- see the
[KVStore.SetEmpty] method.

To enable or disable thread safety, see the [KVStore.SetThreadSafe] method. To
check whether thread safety is enabled, see the [KVStore.IsThreadSafe] method.

To control the behavior with respect to case-folding of indices during any of the
various matching operations conducted, see the [KVStore.SetCaseExact] method.

To read a key/value pair from the receiver instance, see the [KVStore.Get]
method. This requires foreknowledge of the underlying type for input as the
generic type parameter. If not known, see the [KVStore.GetAny] method.

To determine the number of stored indices within the receiver instance, see
the [KVStore.Len] method.

To check whether a particular index exists, regardless of its underlying
value type, see the [KVStore.Exists] method.

To remove a key/value pair from the receiver instance entirely, see the
[KVStore.Delete] method.

To ascertain what kind of value is present within a particular index, see the
[KVStore.Kind] method.

To free the receiver from memory -- thereby annihilating it -- see the
[KVStore.Free] method.
*/
type KVStore struct {
	lock     *sync.Mutex
	doLock   func()
	doUnlock func()
	fold     func(string) string
	store    map[string]any
}

/*
NewKVStore initializes and returns a new instance of *[KVStore]. Use of
this constructor is always required.
*/
func NewKVStore() *KVStore {
	return &KVStore{
		fold:     sLC,
		doLock:   func() {},
		doUnlock: func() {},
		store:    make(map[string]any),
	}
}

/*
isZero returns a boolean value indicative of the receiver instance being
zero, in that all underlying components are either nil, or are allocated
but empty.
*/
func (r *KVStore) isZero() bool {
	hl := r.lock == nil
	fl := r.fold == nil
	sl := (r.store == nil || len(r.store) == 0)
	return hl && fl && sl
}

/*
IsNil returns a Boolean value indicative of the receiver instance being
completely nil, or at least zero.
*/
func (r *KVStore) IsNil() bool { return r == nil || r.isZero() }

/*
Len returns the integer length of the underlying key/value store. The
return indicates the number of indices presently stored.
*/
func (r *KVStore) Len() int {
	var L int
	if !r.IsNil() {
		L = len(r.store)
	}

	return L
}

/*
SetThreadSafe enables or disables the underlying *[sync.Mutex] instance
for write locking.

By default, thread-safety is disabled. If the desired state is already
achieved (e.g.: the user tries to enable locking when it has already been
enabled), the request is ignored.

A working lock will not be needlessly replaced with another working lock.
*/
func (r *KVStore) SetThreadSafe(enable bool) {
	if enable && r.lock == nil {
		r.lock = &sync.Mutex{}
		r.doLock = func() { r.lock.Lock() }
		r.doUnlock = func() { r.lock.Unlock() }
	} else if !enable && r.lock != nil {
		r.lock = nil
		r.doLock = func() {}
		r.doUnlock = func() {}
	}
}

/*
IsThreadSafe returns a Boolean value indicative of whether *[sync.Mutex]
has been enabled for the receiver instance.
*/
func (r *KVStore) IsThreadSafe() bool {
	var safe bool
	if !r.IsNil() {
		safe = r.lock != nil
	}

	return safe
}

/*
SetCaseExact enables or disables the behavior towards case-normalization
with respect to index matching.

A value of false, the default, folds all indices to lower-case whenever
they are evaluated.

A value of true matches indices as-is, meaning no case-folding takes place.
*/
func (r *KVStore) SetCaseExact(enable bool) {
	r.doLock()
	defer r.doUnlock()

	if enable {
		r.fold = func(s string) string { return s }
	} else {
		r.fold = sLC
	}
}

/*
Set assigns input [KVStorable] T to the receiver instance in association
with the input string index. Note that this method will not write to the
receiver instance if the index already exists.

To replace a value instead of setting one anew, see the [KVStore.Update]
method.
*/
func (r *KVStore) Set[T KVStorable](index string, val T) {
	r.doLock()
	defer r.doUnlock()

	if !r.Exists(index) {
		r.store[r.fold(index)] = val
	}
}

/*
SetEmpty assigns a valueless index to the receiver instance. This is most
often used when building so-called "seen tables" for uniqueness control.

This is a convenience method only, and is merely an alternative to the act
of manually performing a call of:

	KVStore.Set[struct{}]("the-index", struct{}{})
*/
func (r *KVStore) SetEmpty(index string) {
	r.doLock()
	defer r.doUnlock()

	if !r.Exists(index) {
		r.store[r.fold(index)] = struct{}{}
	}
}

/*
Update replaces the underlying value associated with the input index with
the new [KVStorable] value. Note that this method is capable of not only
changing the value, but the type as well.
*/
func (r *KVStore) Update[T KVStorable](index string, val T) {
	r.doLock()
	defer r.doUnlock()

	r.store[r.fold(index)] = val
}

/*
Exists returns a Boolean value indicative of the specified index name
being present within the receiver instance in association with a value.
*/
func (r *KVStore) Exists(index string) bool {
	_, ok := r.store[r.fold(index)]
	return ok
}

/*
Kind returns an integer value associated with the kind of [KVStorable]
value assigned to the input index. See the [KVStorable] type for a list
of the various flags and their meaning.
*/
func (r *KVStore) Kind(index string) int {
	var typ int = noneKVType

	if val, err := r.GetAny(index); err == nil {
		switch val.(type) {
		case int64: // 1
			typ = intKVType
		case uint64: // 2
			typ = uintKVType
		case float64: // 3
			typ = floatKVType
		case bool: // 4
			typ = boolKVType
		case []byte: // 5
			typ = bytesKVType
		case [][]byte: // 6
			typ = bytesSliceKVType
		case struct{}: // 7
			typ = presKVType
		}
	}

	return typ
}

/*
Delete removes the specified index and its associated value from the
receiver instance.
*/
func (r *KVStore) Delete(index string) {
	r.doLock()
	defer r.doUnlock()

	delete(r.store, r.fold(index))
}

/*
Free frees the receiver instance from memory.

NOTE: This is a purely destructive method and will not be influenced
by the state of the underlying *[sync.Mutex] instance.
*/
func (r *KVStore) Free() { *r = KVStore{} }

/*
Get retrieves the item and implicitly casts it based on the variable
context of [KVStorable] T. A non-nil error is returned if the index is
not found.

See also the [KVStore.Kind] method as a means for ascertaining the
underlying value type, and [KVStore.GetAny] as a means for retrieving
the value regardless of its type.
*/
func (r *KVStore) Get[T KVStorable](index string) (T, error) {
	var (
		val,
		zero T
		err error
		ok  bool
	)

	if val, ok = r.store[r.fold(index)].(T); !ok {
		val = zero
		err = mkerr("KVStore.Get: index not found: " + index)
	}
	return val, err
}

/*
GetAny returns the underlying interface{} value alongside an error.
A non-nil error is returned if the index is not found.

The caller is responsible for type assertion of the return value.

See also [KVStore.Get].
*/
func (r *KVStore) GetAny(index string) (any, error) {
	var (
		ok  bool
		val any
		err error
	)

	if val, ok = r.store[r.fold(index)]; !ok {
		err = mkerr("KVStore.GetAny: index not found: " + index)
	}

	return val, err
}
