package pkg

import "sync"

type Result struct {
	TraceCode string
	Error     error
	Data      interface{}
}

// ResultChannel handles sending results to a channel stored in a sync.Map.
type ResultChannel struct {
	ResultMap *sync.Map
}

// NewResultChannel creates a new ResultChannel with an initialized sync.Map.
func NewResultChannel() *ResultChannel {
	return &ResultChannel{
		ResultMap: &sync.Map{},
	}
}

// SendResult sends a result to the channel associated with the traceCode.
func (rc *ResultChannel) SendResult(traceCode string, err error, data ...interface{}) {
	if ch, ok := rc.ResultMap.Load(traceCode); ok {
		var payload interface{}
		if len(data) > 0 {
			payload = data[0]
		}
		ch.(chan Result) <- Result{TraceCode: traceCode, Error: err, Data: payload}
	}
}
