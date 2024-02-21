package errors

import (
	"fmt"
)

type ErrMessageCode struct {
	ErrorCode     string
	MessageCode   string
	MessageStore  func(lang, code string) string
	DefaultParams []interface{}
	Params        []interface{}
}

type ErrMessage struct {
	Code    string
	Message string
}

type ErrMessageMultiLanguage struct {
	Lang           string
	ErrMessageCode *ErrMessageCode
}

func NewErrMessage(code string, msg string) *ErrMessage {
	return &ErrMessage{Code: code, Message: msg}
}

func (m *ErrMessage) WithParam(args ...interface{}) *ErrMessage {
	if len(args) == 0 {
		return &ErrMessage{Code: m.Code, Message: m.Message}
	}
	return &ErrMessage{Code: m.Code, Message: fmt.Sprintf(m.Message, args...)}
}

func (m *ErrMessageCode) WithLang(lang string) *ErrMessageMultiLanguage {
	return &ErrMessageMultiLanguage{
		Lang:           lang,
		ErrMessageCode: m,
	}
}

func (m *ErrMessageMultiLanguage) WithParam(args ...interface{}) *ErrMessageMultiLanguage {
	if len(args) == 0 {
		return m
	}
	m.ErrMessageCode.Params = args
	return m
}

func (m *ErrMessageMultiLanguage) Build() *ErrMessage {
	args := m.ErrMessageCode.Params
	if len(args) == 0 {
		args = m.ErrMessageCode.DefaultParams
	}

	message := m.ErrMessageCode.MessageStore(m.Lang, m.ErrMessageCode.MessageCode)
	if len(message) == 0 {
		return &ErrMessage{Code: m.ErrMessageCode.ErrorCode, Message: message}
	}

	if len(args) == 0 {
		return &ErrMessage{Code: m.ErrMessageCode.ErrorCode, Message: message}
	}

	return &ErrMessage{Code: m.ErrMessageCode.ErrorCode, Message: fmt.Sprintf(message, args...)}
}

/*
func (m *ErrMessage) WithLang(lang string) *ErrMessage {
	if m.MessageStore == nil {
		return &ErrMessage{Code: m.Code, Message: m.Message}
	}
	message := m.MessageStore(lang, m.Code)
	if len(message) != 0 {
		return &ErrMessage{Code: m.Code, Message: message}
	}
	return &ErrMessage{Code: m.Code, Message: m.Message}
}

*/

/*
type ErrMultiLangMessage struct {
	Code     string
	Messages map[string]*ErrMessage
}

func NewErrMultiLangMessage(code string, messages map[string]*ErrMessage) *ErrMultiLangMessage {
	return &ErrMultiLangMessage{Code: code, Messages: messages}
}

func (m *ErrMultiLangMessage) WithLang(lang string) *ErrMessage {
	message, ok := m.Messages[lang]
	if ok {
		return &ErrMessage{Code: m.Code, Message: message.Message}
	}
	message = m.Messages[coreconstants.EN]
	return &ErrMessage{Code: m.Code, Message: message.Message}
}
*/
