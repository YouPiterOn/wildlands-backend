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
var ErrMatchJoin = errors.New("error joining match")

// Player errors
var ErrPlayerCreate = errors.New("error creating player")
var ErrPlayerNotFound = errors.New("player not found")
var ErrPlayerGetByID = errors.New("error getting player by ID")

// Match metadata errors
var ErrMatchMetadataLoad = errors.New("error loading match metadata")
var ErrMatchMetadataCreate = errors.New("error creating match metadata")
var ErrMatchMetadataGenerate = errors.New("error generating match metadata")

// Transport errors
var ErrRequestDecode = errors.New("error decoding request")
var ErrResponseEncode = errors.New("error encoding response")
var ErrWsCommandRead = errors.New("websocket error reading command")
var ErrWsEventSerialize = errors.New("websocket error serializing events")
var ErrWsEventWrite = errors.New("websocket error writing events")
var ErrWsErrorWrite = errors.New("websocket error writing error")
var ErrWsAccept = errors.New("websocket error accepting connection")

func CustomErrUUIDParse(name string, id string) error {
	return errors.New("error parsing " + name + " UUID: " + id)
}
