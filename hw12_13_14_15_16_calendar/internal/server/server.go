package server

import (
	"context"
	"time"

	"github.com/cepmap/otus-go-hws/hw12_13_14_15_calendar/internal/models"
	"github.com/google/uuid"
)

//go:generate mockery --name Application

/* other way - mockgen
go:generate mockgen -destination=mocks/Application.go
-package=mocks github.com/cepmap/otus-go-hws/hw12_13_14_15_16_calendar/internal/server Application
*/

type Application interface {
	AddEvent(ctx context.Context, event *models.Event) (uuid.UUID, error)
	GetEvent(ctx context.Context, event uuid.UUID) (*models.Event, error)
	GetEventsForPeriod(ctx context.Context, from, to time.Time) ([]models.Event, error)
	ListEvents(ctx context.Context, limit, low uint64) ([]models.Event, error)
	UpdateEvent(ctx context.Context, event *models.Event) error
	DeleteEvent(ctx context.Context, id uuid.UUID) error
	GetEventsOfDay(ctx context.Context, from time.Time) ([]models.Event, error)
	GetEventsOfWeek(ctx context.Context, from time.Time) ([]models.Event, error)
	GetEventsOfMonth(ctx context.Context, from time.Time) ([]models.Event, error)
}

type Server interface {
	Start(context.Context) error
	Stop(context.Context) error
}
