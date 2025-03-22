package service

import (
	"context"
	"log/slog"
	"time"

	"github.com/ExonegeS/REST-API-001/internal/domain"
)

type LoggingService struct {
	logger *slog.Logger
	next   domain.Service
}

func NewLoggingService(logger *slog.Logger, next domain.Service) domain.Service {
	return &LoggingService{
		logger: logger,
		next:   next,
	}
}

func (s *LoggingService) GetUsersMany(ctx context.Context, input domain.GetUsersInput) (response *domain.GetUsersResponse, err error) {
	start := time.Now()
	defer func() {
		logger := s.logger
		if response != nil {
			logger = s.logger.With(slog.Any("user total", response.Total))
		} else {
			logger = s.logger.With(slog.Any("err", err))
		}
		logger.Info(
			"GetUsersMany",
			"took", time.Since(start).String(),
		)
	}()
	return s.next.GetUsersMany(ctx, input)
}

func (s *LoggingService) GetUsersOne(ctx context.Context, input domain.GetUserInput) (response *domain.GetUserResponse, err error) {
	start := time.Now()
	defer func() {
		logger := s.logger
		if response != nil {
			logger = s.logger.With(slog.Any("user", response.User))
		} else {
			logger = s.logger.With(slog.Any("err", err))
		}
		logger.Info(
			"GetUsersOne",
			"took", time.Since(start).String(),
		)
	}()
	return s.next.GetUsersOne(ctx, input)
}
