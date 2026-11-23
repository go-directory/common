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
	lock    sync.Mutex
	counter uint64
	also    *Counter
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
	r.counter += i

	if r.also != nil {
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
	if i > r.counter {
		r.counter = 0
	} else {
		r.counter -= i
	}

	if r.also != nil {
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

