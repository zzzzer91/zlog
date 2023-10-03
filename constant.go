package zlog

type entityFieldNameType string

func (e entityFieldNameType) String() string {
	return string(e)
}

const (
	EntityFieldNameTraceID    entityFieldNameType = "traceId"
	EntityFieldNameRequestID  entityFieldNameType = "requestId"
	EntityFieldNameLogID      entityFieldNameType = "logId"
	EntityFieldNameError      entityFieldNameType = "error"
	EntityFieldNameErrorStack entityFieldNameType = "errorStack"
)

const (
	defaultTimeFormat = "2006-01-02T15:04:05.000Z0700"
)
