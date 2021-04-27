package lib

import "time"

// RawReq struct of the raw request
type RawReq struct {
	ID  int64
	Req []byte
}

// RawResp struct of the raw response
type RawResp struct {
	ID     int64
	Resp   []byte
	Err    error
	Elapse time.Duration
}

// RetCode type of the result code
type RetCode int

const (
	RET_CODE_SUCCESS              RetCode = 0    // success
	RET_CODE_WARNING_CALL_TIMEOUT         = 1001 // warning of timeout
	RET_CODE_ERROR_CALL                   = 2001 // error with call(int)
	RET_CODE_ERROR_RESPONSE               = 2002 // error with response
	RET_CODE_ERROR_CALEE                  = 2003 // internal error with the callee
	RET_CODE_FATAL_CALL                   = 3001 // fatal error during the call
)

// GetRetCodePlain returns corresponding text depending on the error code
func GetRetCodePlain(code RetCode) string {
	var codePlain string
	switch code {
	case RET_CODE_SUCCESS:
		codePlain = "Success"
	case RET_CODE_WARNING_CALL_TIMEOUT:
		codePlain = "Call Timeout Warning"
	case RET_CODE_ERROR_CALL:
		codePlain = "Call Error"
	case RET_CODE_ERROR_RESPONSE:
		codePlain = "Response Error"
	case RET_CODE_ERROR_CALEE:
		codePlain = "Callee Error"
	case RET_CODE_FATAL_CALL:
		codePlain = "Call Fatal Error"
	default:
		codePlain = "Unknown Result Code"
	}
	return codePlain
}

// CallResult struct of the result of the call
type CallResult struct {
	ID     int64         // ID
	Req    RawReq        // raw request
	Resp   RawResp       // raw response
	Code   RetCode       // result code
	Msg    string        // brief cause of the result
	Elapse time.Duration // time elapsed
}

// constants, states of the load generator
const (
	// the state of the original
	STATUS_ORIGINAL uint32 = 0
	// the state of the starting
	STATUS_STARTING uint32 = 1
	// the state of being started
	STATUS_STARTED uint32 = 2
	// the state of being stopped
	STATUS_STOPPING uint32 = 3
	// the state of having been stopped
	STATUS_STOPPED uint32 = 4
)

// Generator interface of the load generator
type Generator interface {
	// to start the load generator
	// the result represents whether it has been successfully started
	Start() bool
	// to stop the load generator
	// the result represents whether it has been successfully stopped
	Stop() bool
	// to get the status of the load generator
	Status() uint32
	// to call the counter
	// to reset the counter everytimethe load generator starts
	CallCount() int64
}
