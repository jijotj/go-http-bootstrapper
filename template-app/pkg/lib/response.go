package lib

import (
	"encoding/json"
	"io"

	"github.com/rs/zerolog/log"
)

func WriteResponseJSON(w io.Writer, response interface{}) {
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error().Err(err).Msg("Write JSON response")
	}
}
