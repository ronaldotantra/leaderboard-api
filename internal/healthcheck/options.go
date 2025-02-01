package healthcheck

import (
	"context"
	"database/sql"
)

type HealthCheckFunc func(ctx context.Context) error

type Option func(s *service)

func DBClientOption(name string, db *sql.DB) Option {
	return func(s *service) {
		s.components = append(s.components, component{
			Name: name,
			Fn: func(ctx context.Context) error {
				err := db.PingContext(ctx)
				if err != nil {
					return err
				}

				return nil
			},
		})
	}
}
