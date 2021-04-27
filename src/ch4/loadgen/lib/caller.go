package lib

import "time"

// Caller interface of the caller
type Caller interface {
	// to construct requests
	BuildReq() RawReq
	// to call
	Call(req []byte, timeoutNS time.Duration) ([]byte, error)
	// to check responses
	CheckResp(rawReq RawReq, rawResp RawResp) *CallResult
}
