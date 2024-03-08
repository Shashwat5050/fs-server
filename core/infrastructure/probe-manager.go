package infra

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap"
	"iceline-hosting.com/core/logger"
)

type probeManager struct {
	log      logger.Logger
	mu       sync.Mutex
	checkers []checker
	timeout  time.Duration

	healthEndpnt string
	readyEndpnt  string
}

const (
	defaultHealthEndpnt = "/health"
	defaultReadyEndpnt  = "/ready"
)

type checker interface {
	Name() string
	CheckHealth(ctx context.Context) error
	CheckReadiness(ctx context.Context) error
}

func (pb *probeManager) SetTimeout(timeout time.Duration) {
	pb.timeout = timeout
}

func (pb *probeManager) SetHealthEndpnt(endpnt string) {
	pb.healthEndpnt = endpnt
}

func (pb *probeManager) SetReadyEndpnt(endpnt string) {
	pb.readyEndpnt = endpnt
}

func NewProbeManager(log logger.Logger, chs ...checker) *probeManager {
	return &probeManager{
		log:          log,
		timeout:      30 * time.Second,
		checkers:     chs,
		healthEndpnt: defaultHealthEndpnt,
		readyEndpnt:  defaultReadyEndpnt,
	}
}

func (pm *probeManager) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case pm.healthEndpnt:
		pm.HandleHealthCheck(w, r)
	case pm.readyEndpnt:
		pm.HandleReadyCheck(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func (pm *probeManager) HandleHealthCheck(w http.ResponseWriter, r *http.Request) {
	unavailableServices := make([]string, 0)

	ctx, cancel := context.WithTimeout(r.Context(), pm.timeout)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(len(pm.checkers))

	for _, ch := range pm.checkers {
		go func(ch checker) {
			defer wg.Done()

			if err := ch.CheckHealth(ctx); err != nil {
				pm.mu.Lock()
				unavailableServices = append(unavailableServices, ch.Name())
				pm.mu.Unlock()

				return
			}
		}(ch)
	}

	wg.Wait()
	if len(unavailableServices) > 0 {
		w.WriteHeader(http.StatusServiceUnavailable)
		_, err := w.Write([]byte(fmt.Sprintf("Unavailable services: %s", strings.Join(unavailableServices, ", "))))
		if err != nil {
			pm.log.Error("failed to write response", zap.Error(err))
		}

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (pm *probeManager) HandleReadyCheck(w http.ResponseWriter, r *http.Request) {
	notReadyServices := make([]string, 0)

	ctx, cancel := context.WithTimeout(r.Context(), pm.timeout)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(len(pm.checkers))

	for _, ch := range pm.checkers {
		go func(ch checker) {
			defer wg.Done()

			if err := ch.CheckReadiness(ctx); err != nil {
				pm.mu.Lock()
				notReadyServices = append(notReadyServices, ch.Name())
				pm.mu.Unlock()

				return
			}
		}(ch)
	}

	wg.Wait()
	if len(notReadyServices) > 0 {
		w.WriteHeader(http.StatusServiceUnavailable)
		_, err := w.Write([]byte(fmt.Sprintf("Not ready services: %s", strings.Join(notReadyServices, ", "))))
		if err != nil {
			pm.log.Error("failed to write response", zap.Error(err))
		}

		return
	}

	w.WriteHeader(http.StatusOK)
}
