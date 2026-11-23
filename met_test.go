package common

import (
	"fmt"
)

func ExampleCounter() {
	// Initialize ALL needed counters.
	var (
		totalPDUs      = new(Counter)
		requestPDUs    = new(Counter)
		responsePDUs   = new(Counter)
		addRequest     = new(Counter)
		addResponse    = new(Counter)
		modifyRequest  = new(Counter)
		modifyResponse = new(Counter)
		bindRequest    = new(Counter)
		unbindRequest  = new(Counter)
		bindResponse   = new(Counter)
	)

	// Next, we setup our counter "hierarchy" like so:
	//
	//              TOTAL
	//            /       \
	//       REQUESTS   RESPONSES
	//          |           |
	//   <ADD,MOD,...> <ADD,MOD,...>
	//
	// We want ANY PDU -- whether a Request or Response --
	// to ALWAYS increment the appropriate superior counter.

	// We'll begin with the top-level "grand total" ...
	requestPDUs.RepeatOn(totalPDUs)
	responsePDUs.RepeatOn(totalPDUs)

	// Next, setup our Request type counters to increment
	// the Request total counter.
	bindRequest.RepeatOn(requestPDUs)
	addRequest.RepeatOn(requestPDUs)
	modifyRequest.RepeatOn(requestPDUs)
	unbindRequest.RepeatOn(requestPDUs)

	// Repeat for Response type counters using the Response
	// total counter.
	bindResponse.RepeatOn(responsePDUs)
	addResponse.RepeatOn(responsePDUs)
	modifyResponse.RepeatOn(responsePDUs)

	// Setup is complete. Begin incrementing our type-specific
	// counters. In real life, this would be driven by routine
	// activity on the DSA.
	bindRequest.Increment()
	bindResponse.Increment()
	addRequest.Increment()
	addResponse.Increment()
	modifyRequest.Increment()
	modifyResponse.Increment()
	unbindRequest.Increment()

	fmt.Printf("Total/Req/Resp: %d/%d/%d",
		totalPDUs.Count(),
		requestPDUs.Count(),
		responsePDUs.Count())
	// Output:
	// Total/Req/Resp: 7/4/3
}
