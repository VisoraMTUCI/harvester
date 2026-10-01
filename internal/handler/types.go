package handler

import (
	"context"

	"github.com/ppMTUCI/harvester/internal/infra/kafka"
)

type EventWriter interface {
	WriteMessage(ctx context.Context, msg kafka.Message) error
}
