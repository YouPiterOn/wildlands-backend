package utils

import "errors"

// Event errors
var ErrEventLoad = errors.New("error loading events")
var ErrEventAppend = errors.New("error appending events")
var ErrEventApply = errors.New("error applying event")

// Command errors
var ErrCommandHandle = errors.New("error handling command")

// Match errors
var ErrMatchCreate = errors.New("error creating match")

// Player errors
var ErrPlayerCreate = errors.New("error creating player")
var ErrPlayerNotFound = errors.New("player not found")

// Match metadata errors
var ErrMatchMetadataLoad = errors.New("error loading match metadata")
var ErrMatchMetadataCreate = errors.New("error creating match metadata")
var ErrMatchMetadataGenerate = errors.New("error generating match metadata")

// Transport errors
var ErrRequestDecode = errors.New("error decoding request")
var ErrResponseEncode = errors.New("error encoding response")
var ErrUUIDParse = errors.New("error parsing UUID")
