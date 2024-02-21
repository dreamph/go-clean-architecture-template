package models

import "fmt"

/*
type MultiLangMessage struct {
	Code     string `json:"code"`
	Messages map[string]*Message
}
*/

type MultiLangMessage struct {
	Code string   `json:"code"`
	Th   *Message `json:"th"`
	En   *Message `json:"en"`
}

type Message struct {
	Code               string `json:"code"`
	Message            string `json:"message"`
	MessageDescription string `json:"messageDescription"`
}

func (m *Message) FormatMessage(args ...interface{}) string {
	msg := m.Message

	if len(args) != 0 {
		return fmt.Sprintf(msg, args...)
	}
	return msg
}

func (m *Message) FormatMessageDescription(args ...interface{}) string {
	msg := m.MessageDescription

	if len(args) != 0 {
		return fmt.Sprintf(msg, args...)
	}
	return msg
}

func (m *Message) Builder() *MessageBuilder {
	return &MessageBuilder{Message: &Message{Code: m.Code, Message: m.Message, MessageDescription: m.MessageDescription}}
}

type MessageBuilder struct {
	Message *Message
}

func (b *MessageBuilder) MessageParams(args ...interface{}) *Message {
	if len(args) != 0 {
		b.Message.Message = fmt.Sprintf(b.Message.Message, args...)
	}
	return b.Message
}

func (b *MessageBuilder) MessageDescriptionParams(args ...interface{}) *Message {
	if len(args) != 0 {
		b.Message.MessageDescription = fmt.Sprintf(b.Message.MessageDescription, args...)
	}
	return b.Message
}
