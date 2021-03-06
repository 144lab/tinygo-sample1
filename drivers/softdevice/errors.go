package softdevice

// Error is an error from within the SoftDevice.
type Error uint32

func (e Error) Error() string {
	switch {
	case e < 0x1000:
		// Global errors.
		switch e {
		case 0:
			return "no error"
		case 1:
			return "SVC handler is missing"
		case 2:
			return "SoftDevice has not been enabled"
		case 3:
			return "internal error"
		case 4:
			return "no memory for operation"
		case 5:
			return "not found"
		case 6:
			return "not supported"
		case 7:
			return "invalid parameter"
		case 8:
			return "invalid state, operation disallowed in this state"
		case 9:
			return "invalid Length"
		case 10:
			return "invalid flags"
		case 11:
			return "invalid data"
		case 12:
			return "invalid data size"
		case 13:
			return "operation timed out"
		case 14:
			return "null pointer"
		case 15:
			return "forbidden operation"
		case 16:
			return "bad memory address"
		case 17:
			return "busy"
		case 18:
			return "maximum connection count exceeded"
		case 19:
			return "not enough resources for operation"
		default:
			return "other global error"
		}
	case e < 0x2000:
		// SDM errors.
		return "other SDM error"
	case e < 0x3000:
		// SoC errors.
		return "other SoC error"
	case e < 0x4000:
		// STK errors.
		return "other STK error"
	default:
		// Other errors.
		return "other error"
	}
}

// NrfError returns an error (using the Error type) if the error code is
// non-zero, otherwise it returns nil. It is used with internal API calls.
func NrfError(code uint32) error {
	if code != 0 {
		return Error(code)
	}
	return nil
}
