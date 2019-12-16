package verrors

// ErrorType is an enum encapsulating the spectrum of all possible types of errors raised by the application
type ErrorType int64

const (
	// NotFound corresponds to errors caused by missing entities
	NotFound ErrorType = 1 + iota
	// InvalidArgument corresponds to errors caused by missing or malformed arguments supplied by a client
	InvalidArgument
	// AlreadyExists corresponds to errors caused by an entity already existing
	AlreadyExists
	// PermissionDenied corresponds to a user not having permission to access a resource.
	PermissionDenied
	// Unauthenticated indicates the request does not have valid authentication credentials for the operation.
	Unauthenticated
	// Unimplemented corresponds to a function that is unimplemented
	Unimplemented
	// Unknown Error occurred
	Unknown
	// Internal Error
	Internal
	// Unavailable error occurred
	Unavailable
	// FailedPrecondition indicates operation was rejected because the
	// system is not in a state required for the operation's execution.
	// For example, directory to be deleted may be non-empty, an rmdir
	// operation is applied to a non-directory, etc.
	//
	// A litmus test that may help a service implementor in deciding
	// between FailedPrecondition, Aborted, and Unavailable:
	//  (a) Use Unavailable if the client can retry just the failing call.
	//  (b) Use Aborted if the client should retry at a higher-level
	//      (e.g., restarting a read-modify-write sequence).
	//  (c) Use FailedPrecondition if the client should not retry until
	//      the system state has been explicitly fixed.  E.g., if an "rmdir"
	//      fails because the directory is non-empty, FailedPrecondition
	//      should be returned since the client should not retry unless
	//      they have first fixed up the directory by deleting files from it.
	//  (d) Use FailedPrecondition if the client performs conditional
	//      REST Get/Update/Delete on a resource and the resource on the
	//      server does not match the condition. E.g., conflicting
	//      read-modify-write on the same resource.
	FailedPrecondition
	// DeadlineExceeded means operation expired before completion.
	// For operations that change the state of the system, this error may be
	// returned even if the operation has completed successfully. For
	// example, a successful response from a server could have been delayed
	// long enough for the deadline to expire.
	DeadlineExceeded
	// ResourceExhausted indicates some resource has been exhausted, perhaps
	// a per-user quota, or perhaps the entire file system is out of space.
	ResourceExhausted
	// Aborted indicates the operation was aborted, typically due to a
	// concurrency issue like sequencer check failures, transaction aborts,
	// etc.
	//
	// See litmus test above for deciding between FailedPrecondition,
	// Aborted, and Unavailable.
	Aborted
	// Canceled indicates whether the downstream client cancels a request
	Canceled
	// BadGateway - The server, while acting as a gateway or proxy, received an invalid response from the upstream server.
	BadGateway
)

// StatusClientClosedRequest represents `request cancel` errors.
// The GRPC Specification states that this is the translation to http code for codes.Canceled.
// https://github.com/grpc/grpc/blob/master/doc/statuscodes.md
// Until it is included in the http library itself, we'll just hardcode it here.
const StatusClientClosedRequest = 499

// String converts an error type to a string
func (errType ErrorType) String() string {
	switch errType {
	case NotFound:
		return "NotFound"
	case InvalidArgument:
		return "InvalidArgument"
	case AlreadyExists:
		return "AlreadyExists"
	case PermissionDenied:
		return "PermissionDenied"
	case Unauthenticated:
		return "Unauthenticated"
	case Unimplemented:
		return "Unimplemented"
	case Unknown:
		return "Unknown"
	case Internal:
		return "Internal"
	case Unavailable:
		return "Unavailable"
	case FailedPrecondition:
		return "FailedPrecondition"
	case DeadlineExceeded:
		return "DeadlineExceeded"
	case ResourceExhausted:
		return "ResourceExhausted"
	case Aborted:
		return "Aborted"
	case Canceled:
		return "Canceled"
	case BadGateway:
		return "BadGateway"
	default:
		return "Unknown"
	}
}
