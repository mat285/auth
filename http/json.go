package http

import (
	"encoding/json"
	"io"
	"net/http"
)

func ReadJSONBody(r *http.Request, v interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, v)
}

func WriteJSONBody(rw http.ResponseWriter, v interface{}) {
	body, err := json.Marshal(v)
	if err != nil {
		writeError(rw, err)
		return
	}
	rw.WriteHeader(200)
	rw.Write(body)
}
