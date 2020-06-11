package errors

import "errors"

var InternalError = errors.New("INTERNAL_ERROR")
var BadRequest = errors.New("BAD_REQUEST")
var InvalidParams = errors.New("INVALID_PARAMS")

var NoteNotFound = errors.New("NOTE_NOT_FOUND")
