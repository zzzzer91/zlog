package zlog

type entityFieldNameType string

func (e entityFieldNameType) String() string {
	return string(e)
}

const (
	EntityFieldNameTraceId    entityFieldNameType = "traceID"
	EntityFieldNameRequestId  entityFieldNameType = "requestID"
	EntityFieldNameLogId      entityFieldNameType = "logID"
	EntityFieldNameError      entityFieldNameType = "error"
	EntityFieldNameErrorStack entityFieldNameType = "errorStack"
)

const (
	defaultTimeFormat = "2006-01-02T15:04:05.000Z0700"
)
