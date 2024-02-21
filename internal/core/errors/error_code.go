package errors

type ErrMessageCodeOptionsFunc func(*ErrMessageCode)

func WithDefaultParams(v ...interface{}) ErrMessageCodeOptionsFunc {
	return func(o *ErrMessageCode) {
		o.DefaultParams = v
	}
}

type ErrorCodeBuilder struct {
	messageStore func(lang, code string) string
}

func NewErrorCodeBuilder(messageStore func(lang, code string) string) *ErrorCodeBuilder {
	return &ErrorCodeBuilder{
		messageStore: messageStore,
	}
}

func (e *ErrorCodeBuilder) NewErrorCode(errorCode string, messageCode string, optFns ...func(*ErrMessageCode)) *ErrMessageCode {
	errMessage := &ErrMessageCode{
		ErrorCode:    errorCode,
		MessageCode:  messageCode,
		MessageStore: e.messageStore,
	}

	for _, optFn := range optFns {
		optFn(errMessage)
	}

	return errMessage
}
