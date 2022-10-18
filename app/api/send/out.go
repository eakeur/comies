package send

import (
	"encoding/json"
	"net/http"
)

func encodeJSON(wr http.ResponseWriter, data interface{}) error {
	enc := json.NewEncoder(wr)

	err := enc.Encode(data)
	if err == nil {
		return nil
	}

	jsonErr := enc.Encode(InternalServerError)
	if jsonErr != nil {
		return jsonErr
	}

	return err
}
