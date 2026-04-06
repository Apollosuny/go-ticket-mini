package domain

import "errors"

var (
	ErrTicketNotFound = errors.New("ticket not found")
	ErrTitleRequired = errors.New("title is required")
	ErrRequesterEmailEmpty = errors.New("requester email is required")
	ErrInvalidTicketStatus = errors.New("invalid ticket status")
	ErrCommentMessageEmpty = errors.New("comment message is required")
	ErrAuthorNameEmpty = errors.New("author name is required")
)