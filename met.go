package common

/*
met.go implements basic metrics-related types, such as a counter.
*/

import (
	"sync"
)

/*
Counter implements a single, thread-safe general-use counter with
repeater capabilities.
*/
type Counter struct {
	name    []byte
	lock    sync.Mutex
	counter uint64
	also    *Counter
}

/*
Subordinate creates and returns a new instance of *[Counter] that
has automatically been linked to the receiver. If the return value
is modified, the change is chained to the receiver instance.
*/
func (r *Counter) Subordinate(name ...[]byte) *Counter {
	var n []byte
	if len(name) > 0 && len(name[0]) > 0 {
		n = name[0]
	}
	emitT("%T.Subordinate -> %q", *r, n)
	return &Counter{name: n, also: r}
}

/*
RepeatOn allows a [Counter] instance to replay updates upon another
[Counter]. This is useful if a particular instance represents a so
called "grand total" and should be updated any time another specific
[Counter] is incremented or decremented.

Locks do not overlap. A subsequent [Counter] will not be modified
until the current [Counter] mutex is unlocked.
*/
func (r *Counter) RepeatOn(c *Counter) {
	r.lock.Lock()
	defer r.lock.Unlock()

	if r.also != c {
		r.also = c
	}
}

/*
Increment increments the underlying counter value within the receiver
instance.

If the input variadic n is supplied, the counter is incremented according
to the magnitude of n. Otherwise, the counter is incremented by one (1).
*/
func (r *Counter) Increment(n ...uint64) {
	r.lock.Lock()
	defer r.lock.Unlock()

	var i uint64 = 1
	if len(n) > 0 {
		i = n[0]
	}

	repeat := r.also != nil
	r.counter += i
	emitT("%T.Increment %q += %d [now:%d, repeat:%t]",
		*r, r.name, i, r.counter, repeat)

	if repeat {
		r.also.Increment(n...)
	}
}

/*
Decrement decrements the underlying counter value within the receiver
instance.

If the input variadic n is supplied, the counter is decremented according
to the magnitude of n. Otherwise, the counter is decremented by one (1).
*/
func (r *Counter) Decrement(n ...uint64) {
	if r.counter == 0 {
		return
	}

	r.lock.Lock()
	defer r.lock.Unlock()

	var i uint64 = 1
	if len(n) > 0 {
		i = n[0]
	}
	repeat := r.also != nil
	if i > r.counter {
		r.counter = 0
		emitT("%T.Decrement reset %q [now:%d, repeat:%t]",
			*r, r.name, r.counter, repeat)
	} else {
		r.counter -= i
		emitT("%T.Decrement %q -= %d [now:%d, repeat:%t]",
			*r, r.name, i, r.counter, repeat)
	}

	if repeat {
		r.also.Decrement(n...)
	}
}

/*
Count returns the current counter value of the receiver instance.
*/
func (r *Counter) Count() uint64 {
	r.lock.Lock()
	defer r.lock.Unlock()
	return r.counter
}

func (r *Counter) SetName(name []byte) {
	if len(name) > 0 && len(r.name) == 0 {
		r.name = name
	}
}
