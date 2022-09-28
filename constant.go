package zlog

type entityFieldNameType string

func (e entityFieldNameType) String() string {
	return string(e)
}

const (
	EntityFieldNameTraceID    entityFieldNameType = "traceID"
	EntityFieldNameRequestID  entityFieldNameType = "requestID"
	EntityFieldNameError      entityFieldNameType = "error"
	EntityFieldNameErrorStack entityFieldNameType = "errorStack"
)
