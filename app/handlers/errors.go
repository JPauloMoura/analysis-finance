package handlers

import "fmt"

var (
	_errorRetrievingFile    = fmt.Errorf("error retrieving the file")
	_errorCreateingTempFile = fmt.Errorf("error creating temporary file")
	_errorReadingFile       = fmt.Errorf("error reading file")
)
