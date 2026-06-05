package application

import "errors"

var ErrEventsLoad = errors.New("error loading events")
var ErrEventsAppend = errors.New("error appending events")
var ErrEventApply = errors.New("error applying event")
var ErrCommandHandle = errors.New("error handling command")
