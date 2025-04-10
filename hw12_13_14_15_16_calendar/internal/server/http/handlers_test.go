package internalhttp

import (
	"context"
	"testing"
	"time"

	"github.com/cepmap/otus-go-hws/hw12_13_14_15_calendar/internal/config"
	"github.com/cepmap/otus-go-hws/hw12_13_14_15_calendar/internal/models"
	"github.com/cepmap/otus-go-hws/hw12_13_14_15_calendar/internal/server/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestAppMockGetEvent(t *testing.T) {
	ctx := context.Background()
	eventID := uuid.New()
	existingEvent := models.Event{
		ID:          eventID,
		Title:       "some title",
		DateStart:   time.Now(),
		DateEnd:     time.Now().Add(1 * time.Hour),
		UserID:      uuid.UUID{},
		Description: "some desc for test",
	}

	tests := []struct {
		title      string
		event      *models.Event
		err        error
		GetAppMock func(t *testing.T) *mocks.Application
	}{
		{
			title: "positive case",
			event: &existingEvent,
			GetAppMock: func(t *testing.T) *mocks.Application {
				t.Helper()

				appMock := mocks.NewApplication(t)
				appMock.On("GetEvent", ctx, eventID).Return(&existingEvent, nil)
				return appMock
			},
		},
		{
			title: "event not found",
			err:   models.ErrEventNotFound,
			GetAppMock: func(t *testing.T) *mocks.Application {
				t.Helper()

				appMock := mocks.NewApplication(t)
				appMock.On("GetEvent", ctx, eventID).Return(&models.Event{}, models.ErrEventNotFound)
				return appMock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			appMock := tt.GetAppMock(t)
			server := NewServer(appMock, &config.Server{Host: "0.0.0.0", HTTPPort: "8081"})

			res, err := server.app.GetEvent(ctx, eventID)
			if tt.err != nil {
				require.Equal(t, tt.err, err)
			} else {
				require.Equal(t, tt.event, res)
				require.Equal(t, nil, err)
			}

			appMock.AssertExpectations(t)
		})
	}
}
