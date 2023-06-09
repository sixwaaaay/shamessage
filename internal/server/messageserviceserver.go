// Code generated by gen. DO NOT EDIT.
// Source: message.proto

package server

import (
	"context"
	"github.com/sixwaaaay/shamessage/internal/logic"
	"github.com/sixwaaaay/shamessage/message"
)

// MessageServiceServer is the implementation of message.MessageServiceServer
type MessageServiceServer struct {
	ListLogic *logic.ListLogic
	PutLogic  *logic.PutLogic
	message.UnimplementedMessageServiceServer
}

// ServerOption is the parameters for MessageServiceServer
type ServerOption struct {
	ListLogic *logic.ListLogic
	PutLogic  *logic.PutLogic
}

// NewMessageServiceServer creates a new MessageServiceServer
func NewMessageServiceServer(opt ServerOption) *MessageServiceServer {
	return &MessageServiceServer{
		ListLogic: opt.ListLogic,
		PutLogic:  opt.PutLogic,
	}
}

// fetch message list
func (s *MessageServiceServer) List(ctx context.Context, in *message.MessageListRequest) (*message.MessageListResponse, error) {
	return s.ListLogic.List(ctx, in)
}

// send message
func (s *MessageServiceServer) Put(ctx context.Context, in *message.MessageActionRequest) (*message.MessageActionResponse, error) {
	return s.PutLogic.Put(ctx, in)
}
