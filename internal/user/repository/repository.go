package userrepository

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/leporo/sqlf"
	"github.com/ronaldotantra/go-atomic"
	"github.com/ronaldotantra/leaderboard-api/internal/constant"
	"github.com/ronaldotantra/leaderboard-api/internal/user"
)

type Repository struct {
	atomic.Executor
	sql *sqlf.Dialect
}

func New(atomicExecutor atomic.Executor) *Repository {
	return &Repository{
		Executor: atomicExecutor,
		sql:      sqlf.PostgreSQL,
	}
}

func (r *Repository) SelectOneUser(ctx context.Context, input user.SelectOneUserInput) (output user.User, err error) {
	query := r.sql.From(constant.TableUser + " u").
		Select("u.id").To(&output.ID).
		Select("u.name").To(&output.Name).
		Select("u.email").To(&output.Email).
		Select("u.password").To(&output.Password).
		Select("u.created_at").To(&output.CreatedAt).
		Select("u.updated_at").To(&output.UpdatedAt)

	if input.UserID != nil {
		query = query.Where("u.id = ?", *input.UserID)
	}

	if input.Email != nil {
		query = query.Where("u.email = ?", strings.ToLower(*input.Email))
	}

	err = query.QueryRowAndClose(ctx, r.UseTx(ctx))
	if err != nil {
		sentryHub := sentry.GetHubFromContext(ctx)
		if sentryHub != nil {
			sentryHub.CaptureException(err)
		}
		return
	}

	return
}

func (r *Repository) SelectUsers(ctx context.Context) (output []user.User, err error) {
	var item user.User
	query := r.sql.From(constant.TableUser + " u").
		Select("u.id").To(&item.ID).
		Select("u.name").To(&item.Name).
		Select("u.email").To(&item.Email).
		Select("u.created_at").To(&item.CreatedAt).
		Select("u.updated_at").To(&item.UpdatedAt)

	err = query.QueryAndClose(ctx, r.UseTx(ctx), func(rows *sql.Rows) {
		output = append(output, item)
	})
	if err != nil {
		sentryHub := sentry.GetHubFromContext(ctx)
		if sentryHub != nil {
			sentryHub.CaptureException(err)
		}
		return
	}
	return
}

func (r *Repository) InsertUser(ctx context.Context, input user.RegisterPayload) error {
	query := r.sql.InsertInto(constant.TableUser).
		Set("name", input.Name).
		Set("email", input.Email).
		Set("password", input.Password).
		Set("created_at", time.Now()).
		Set("updated_at", time.Now())

	_, err := query.ExecAndClose(ctx, r.UseTx(ctx))
	if err != nil {
		sentryHub := sentry.GetHubFromContext(ctx)
		if sentryHub != nil {
			sentryHub.CaptureException(err)
		}
		return err
	}

	return nil
}
