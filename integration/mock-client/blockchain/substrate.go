package blockchain

import (
	"encoding/json"
	"errors"
	"fmt"
)

func handleSubstrateRequest(conn string, msg JsonrpcMessage) ([]JsonrpcMessage, error) {
		switch msg.Method {
		case "state_subscribeStorage":
			return handleSubstrateSubscribe(msg)
		case "state_getMetadata":
			return handleSubstrateMetadata(msg)
		}

	return nil, errors.New(fmt.Sprint("unexpected method: ", msg.Method))
}
func handleSubstrateMetadata(msg JsonrpcMessage) ([]JsonrpcMessage, error) {
	return []JsonrpcMessage{
		{
			Version: msg.Version,
			ID:     msg.ID,
			Method: "state_getMetadata",
		},
	}, nil
}

func handleSubstrateSubscribe(msg JsonrpcMessage) ([]JsonrpcMessage, error) {
	var contents []interface{}
	err := json.Unmarshal(msg.Params, &contents)
	if err != nil {
		return nil, err
	}

	if len(contents) != 2 {
		return nil, errors.New(fmt.Sprint("possibly incorrect length of params array:", len(contents)))
	}

	// Add check for valid substrate address (if address is invalid, return "not valid address" error)

	// Add method = "state_subscribeStorage", params to data JSON object
	return []JsonrpcMessage{
		{
			Version: msg.Version,
			ID:     msg.ID,
			Method: "state_subscribeStorage",
		},
	}, nil
}