package utils

import "errors"

// Event errors
var ErrEventLoad = errors.New("error loading events")
var ErrEventAppend = errors.New("error appending events")
var ErrEventApply = errors.New("error applying event")

// Command errors
var ErrCommandHandle = errors.New("error handling command")

// Match metadata errors
var ErrMatchMetadataLoad = errors.New("error loading match metadata")
var ErrMatchMetadataCreate = errors.New("error creating match metadata")
var ErrMatchMetadataGenerate = errors.New("error generating match metadata")
