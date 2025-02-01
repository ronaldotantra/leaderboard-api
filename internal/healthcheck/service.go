package healthcheck

import (
	"context"
	"sync"
)

type component struct {
	Name string
	Fn   HealthCheckFunc
}

type service struct {
	mutex      sync.Mutex
	components []component
}

func New(opts ...Option) Service {
	srv := &service{}

	for _, opt := range opts {
		opt(srv)
	}

	return srv
}

func (s *service) HealthCheck(ctx context.Context) HealthReport {
	failures := map[string]string{}

	wg := &sync.WaitGroup{}
	for _, comp := range s.components {
		wg.Add(1)
		go func(cmp component) {
			defer wg.Done()

			err := cmp.Fn(ctx)
			if err != nil {
				defer s.mutex.Unlock()
				s.mutex.Lock()
				failures[cmp.Name] = err.Error()
			}
		}(comp)
	}
	wg.Wait()

	report := HealthReport{
		Status:     StatusHealthy,
		Components: s.components,
		Failures:   failures,
	}
	if len(failures) > 0 {
		report.Status = StatusUnhealthy
	}

	return report
}
