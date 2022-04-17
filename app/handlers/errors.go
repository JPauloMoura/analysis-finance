package handlers

import (
	"errors"
)

var (
	_errRetrievingFile    = errors.New("error retrieving the file")
	_errCreateingTempFile = errors.New("error creating temporary file")
	_errReadingFile       = errors.New("error reading file")
)

const (
	ErrorMsgInvalidType = "invalid file type"
)

var ErrorMsgs = map[string]string{
	ErrorMsgInvalidType: "Tipo de arquivo inválido. Realize o upload de um .csv",
}
