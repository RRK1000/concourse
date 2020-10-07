package buildserver

import (
	"encoding/json"
	"net/http"

	"github.com/concourse/concourse/atc/api/present"
	"github.com/concourse/concourse/atc/db"
)

func (s *Server) GetBuild(build db.Build) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.logger.Session("get-build")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		err := json.NewEncoder(w).Encode(present.Build(build, s.router))
		if err != nil {
			logger.Error("failed-to-encode-build", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	})
}
