package storage

import (
	"context"
	"time"

	"github.com/cepmap/otus-go-hws/hw12_13_14_15_calendar/internal/models"
	"github.com/google/uuid"
)

type Storage interface {
	InitStorage()
	AddEvent(ctx context.Context, event *models.Event) error
	GetEvent(ctx context.Context, id uuid.UUID) (*models.Event, error)
	GetEventsForPeriod(ctx context.Context, from, to time.Time) ([]models.Event, error)
	ListEvents(ctx context.Context, limit, low uint64) ([]models.Event, error)
	UpdateEvent(ctx context.Context, event *models.Event) error
	DeleteEvent(ctx context.Context, id uuid.UUID) error
	Close() error
}
