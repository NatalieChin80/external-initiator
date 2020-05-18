package blockchain

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/smartcontractkit/chainlink/core/logger"
)

func handleSubstrateRequest(conn string, msg JsonrpcMessage) ([]JsonrpcMessage, error) {
	logger.Infow("entered handle substrate request")
	if conn == "ws" {
		switch msg.Method {
		case "state_subscribeStorage":
			logger.Infow("entered state/subscribeStorage")
			return handleSubstrateSubscribeStorage(msg)
		case "state_getMetadata":
			return handleSubstrateMetadata(msg)
		}
	} else {
		return nil, errors.New(fmt.Sprint("invalid connection type: ", msg.Method))
	}
	return nil, errors.New(fmt.Sprint("unexpected method: ", msg.Method))
}
func handleSubstrateMetadata(msg JsonrpcMessage) ([]JsonrpcMessage, error) {
	logger.Infow("handle substrate metadata")
	// var contents []interface{}
	// err := json.Unmarshal(msg.Params, &contents)
	// if err != nil {
	// 	return nil, err
	// }

	// if len(contents) != 3 {
	// 	return nil, errors.New(fmt.Sprint("possibly incorrect length of params array:", len(contents)))
	// }
	return []JsonrpcMessage{
		{
			Version: "2.0",
			ID:      msg.ID,
			Method:  "state_getMetadata",
		},
	}, nil
}

func handleSubstrateSubscribeStorage(msg JsonrpcMessage) ([]JsonrpcMessage, error) {
	var contents []interface{}
	err := json.Unmarshal(msg.Params, &contents)
	if err != nil {
		return nil, err
	}

	// if len(contents) != 4 {
	// 	return nil, errors.New(fmt.Sprint("possibly incorrect length of params array:", len(contents)))
	// }

	// Add check for valid substrate address (if address is invalid, return "not valid address" error)

	return []JsonrpcMessage{
		{
			Version: "2.0",
			ID:      msg.ID,
			Method:  "state_subscribeStorage",
			Params: msg.Params,
		},
	}, nil
}