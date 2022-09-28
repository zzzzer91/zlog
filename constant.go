package zlog

type entityFieldNameType string

func (e entityFieldNameType) String() string {
	return string(e)
}

const (
	EntityFieldNameTraceID    entityFieldNameType = "traceId"
	EntityFieldNameRequestID  entityFieldNameType = "requestId"
	EntityFieldNameError      entityFieldNameType = "error"
	EntityFieldNameErrorStack entityFieldNameType = "errorStack"
)
