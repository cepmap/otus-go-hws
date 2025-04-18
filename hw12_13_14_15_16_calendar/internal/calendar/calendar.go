package calendar

import (
	"context"
	"time"

	"github.com/cepmap/otus-go-hws/hw12_13_14_15_calendar/internal/models"
	"github.com/cepmap/otus-go-hws/hw12_13_14_15_calendar/internal/storage"
	"github.com/google/uuid"
)

type App struct {
	storage storage.Storage
}

func New(storage storage.Storage) *App {
	return &App{storage: storage}
}

func (a *App) AddEvent(ctx context.Context, event *models.Event) (uuid.UUID, error) {
	if event.ID == uuid.Nil {
		event.ID = uuid.New()
	}
	return event.ID, a.storage.AddEvent(ctx, event)
}

func (a *App) GetEvent(ctx context.Context, id uuid.UUID) (*models.Event, error) {
	return a.storage.GetEvent(ctx, id)
}

func (a *App) GetEventsForPeriod(ctx context.Context, since, dateTo time.Time) ([]models.Event, error) {
	return a.storage.GetEventsForPeriod(ctx, since, dateTo)
}

func (a *App) ListEvents(ctx context.Context, limit, low uint64) ([]models.Event, error) {
	return a.storage.ListEvents(ctx, limit, low)
}

func (a *App) UpdateEvent(ctx context.Context, event *models.Event) error {
	return a.storage.UpdateEvent(ctx, event)
}

func (a *App) DeleteEvent(ctx context.Context, id uuid.UUID) error {
	return a.storage.DeleteEvent(ctx, id)
}

func (a *App) GetEventsOfDay(ctx context.Context, since time.Time) ([]models.Event, error) {
	return a.storage.GetEventsForPeriod(ctx, since, since.AddDate(0, 0, 1))
}

func (a *App) GetEventsOfWeek(ctx context.Context, since time.Time) ([]models.Event, error) {
	return a.storage.GetEventsForPeriod(ctx, since, since.AddDate(0, 0, 7))
}

func (a *App) GetEventsOfMonth(ctx context.Context, since time.Time) ([]models.Event, error) {
	return a.storage.GetEventsForPeriod(ctx, since, since.AddDate(0, 1, 0))
}
