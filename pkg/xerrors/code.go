package xerrors

type Code uint

const (
	CodeUnknown Code = iota
	CodeNotFound
	CodeInvalidArgument
	CodeAlreadyExists
	CodeUnauthenticated
	CodeInternal
)
