package api

import (
	"context"
	"fmt"
	"github.com/cepmap/otus-go-hws/hw12_13_14_15_calendar/api/pbapp"
	"github.com/cepmap/otus-go-hws/hw12_13_14_15_calendar/internal/app"
	"github.com/cepmap/otus-go-hws/hw12_13_14_15_calendar/internal/mapper"
)

type api struct {
	pbapp.UnimplementedAppServer
	app *app.App
}

func NewAPI(app *app.App) pbapp.AppServer {
	return &api{app: app}
}

func (a *api) AddEvent(ctx context.Context, req *pbapp.AddEventRequest) (*pbapp.AddEventResponse, error) {
	cmd, err := mapper.AddEventCommand(req)
	if err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}
	eventId, err := a.app.AddEvent(ctx, cmd)
	if err != nil {
		return nil, fmt.Errorf("failed to add event: %w", err)
	}
	return &pbapp.AddEventResponse{EventId: eventId.String()}, nil
}
func (a *api) UpdateEvent(ctx context.Context, req *pbapp.UpdateEventRequest) (*pbapp.UpdateEventResponse, error) {
	event, err := mapper.Event(req)
	if err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}
	if err = a.app.UpdateEvent(ctx, event); err != nil {
		return nil, fmt.Errorf("update event: %w", err)
	}
	return &pbapp.UpdateEventResponse{}, nil
}

func (a *api) DeleteEvent(ctx context.Context, req *pbapp.DeleteEventRequest) (*pbapp.DeleteEventResponse, error) {
	eventID, err := mapper.EventID(req)
	if err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}
	if err = a.app.DeleteEvent(ctx, eventID); err != nil {
		return nil, fmt.Errorf("delete event: %w", err)
	}
	return &pbapp.DeleteEventResponse{}, nil
}

func (a *api) GetEventsOfDay(ctx context.Context, req *pbapp.GetEventsRequest) (*pbapp.GetEventsResponse, error) {
	events, err := a.app.GetEventsOfDay(ctx, mapper.BeginOfDay(req))
	if err != nil {
		return nil, fmt.Errorf("get events: %w", err)
	}
	return mapper.GetEventsResponse(events), nil
}

func (a *api) GetEventsOfWeek(ctx context.Context, req *pbapp.GetEventsRequest) (*pbapp.GetEventsResponse, error) {
	events, err := a.app.GetEventsOfWeek(ctx, mapper.BeginOfDay(req))
	if err != nil {
		return nil, fmt.Errorf("get events: %w", err)
	}
	return mapper.GetEventsResponse(events), nil
}

func (a *api) GetEventsOfMonth(ctx context.Context, req *pbapp.GetEventsRequest) (*pbapp.GetEventsResponse, error) {
	events, err := a.app.GetEventsOfMonth(ctx, mapper.BeginOfDay(req))
	if err != nil {
		return nil, fmt.Errorf("get events: %w", err)
	}
	return mapper.GetEventsResponse(events), nil
}
