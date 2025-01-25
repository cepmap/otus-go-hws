package app

import (
	"context"
	"github.com/cepmap/otus-go-hws/hw12_13_14_15_calendar/internal/storage"
	"github.com/google/uuid"
	"time"
)

type App struct {
	storage Storage
}

type Logger interface { // TODO
}

type Storage interface {
	InitStorage()
	AddEvent(ctx context.Context, event *storage.Event) error
	GetEvent(ctx context.Context, id uuid.UUID) (*storage.Event, error)
	GetEventsForPeriod(ctx context.Context, from, to time.Time) ([]storage.Event, error)
	ListEvents(ctx context.Context, limit, low uint64) ([]storage.Event, error)
	UpdateEvent(ctx context.Context, event *storage.Event) error
	DeleteEvent(ctx context.Context, id uuid.UUID) error
	Close() error
}

func New(storage Storage) *App {
	return &App{storage: storage}
}

func (a *App) CreateEvent(ctx context.Context, id, title string) error {
	return a.storage.AddEvent(ctx, &storage.Event{})
}

func (a *App) EditEvent(ctx context.Context, id, title string) error         { return nil }
func (a *App) DeleteEvent(ctx context.Context, id string) error              { return nil }
func (a *App) GetEvents(ctx context.Context, id string) (interface{}, error) { return nil, nil }
