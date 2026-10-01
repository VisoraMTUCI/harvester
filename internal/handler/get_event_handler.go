package handler

import (
	"io"
	"log/slog"
	"net/http"

	jsoniter "github.com/json-iterator/go"

	"github.com/ppMTUCI/harvester/internal/infra/kafka"
)

const (
	eventKeyQueryParam = "mouse-click"
)

func NewGetEventHandler(logger *slog.Logger, writer EventWriter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)

			return
		}

		body, err := io.ReadAll(http.MaxBytesReader(w, r.Body, 1<<20))
		if err != nil {
			logger.WarnContext(ctx, "Cannot read event body", "error", err.Error())
			w.WriteHeader(http.StatusBadRequest)

			return
		}

		if len(body) == 0 {
			logger.WarnContext(ctx, "Empty event body")
			w.WriteHeader(http.StatusBadRequest)

			return
		}

		if !jsoniter.Valid(body) {
			logger.WarnContext(ctx, "Event body is not valid json", "size", len(body))
			w.WriteHeader(http.StatusBadRequest)

			return
		}

		var key []byte
		if rawKey := r.URL.Query().Get(eventKeyQueryParam); rawKey != "" {
			key = []byte(rawKey)
		}

		err = writer.WriteMessage(ctx, kafka.Message{Key: key, Value: body})
		if err != nil {
			logger.ErrorContext(ctx, "Cannot write event to kafka", "error", err.Error())
			w.WriteHeader(http.StatusInternalServerError)

			return
		}

		w.WriteHeader(http.StatusAccepted)
	}
}
