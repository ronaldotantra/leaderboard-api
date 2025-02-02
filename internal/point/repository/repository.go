package userrepository

import (
	"context"
	"database/sql"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/leporo/sqlf"
	"github.com/lib/pq"
	"github.com/ronaldotantra/go-atomic"
	"github.com/ronaldotantra/leaderboard-api/internal/constant"
	"github.com/ronaldotantra/leaderboard-api/internal/point"
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

func (r *Repository) BulkInsertPoint(ctx context.Context, input point.InsertPointPayload) (err error) {
	if len(input.Items) == 0 {
		return
	}

	var id int
	matchStmt := r.sql.InsertInto(constant.TableMatch).
		Set("participant_id", pq.Array(input.UserIDs)).
		Set("created_at", time.Now()).
		Set("updated_at", time.Now()).
		Returning("id").To(&id)

	err = matchStmt.QueryRowAndClose(ctx, r.UseTx(ctx))
	if err != nil {
		sentryHub := sentry.GetHubFromContext(ctx)
		if sentryHub != nil {
			sentryHub.CaptureException(err)
		}
		return
	}
	stmt := r.sql.InsertInto(constant.TablePoint)
	for _, item := range input.Items {
		stmt.NewRow().
			Set("match_id", id).
			Set("user_id", item.UserID).
			Set("point", item.Point).
			Set("date", item.Date).
			Set("created_at", time.Now()).
			Set("updated_at", time.Now())
	}

	_, err = stmt.ExecAndClose(ctx, r.UseTx(ctx))
	if err != nil {
		sentryHub := sentry.GetHubFromContext(ctx)
		if sentryHub != nil {
			sentryHub.CaptureException(err)
		}
		return
	}
	return
}

func (r *Repository) GetTotalPoint(ctx context.Context, input point.GetTotalPointInput) (output []point.GetTotalPointOutput, err error) {
	totalPointStmt := r.sql.From(constant.TablePoint).
		Select("user_id").
		Select("sum(point) as total_point").
		Where("date >= ?", input.StartDate).
		Where("date < ?", input.EndDate).
		GroupBy("user_id").To(&output)

	item := point.GetTotalPointOutput{}
	query := r.sql.From(constant.TableUser+" u").
		With("total_point", totalPointStmt).
		LeftJoin("total_point", "u.id = total_point.user_id").
		Select("u.name").To(&item.Name).
		Select("coalesce(total_point.total_point, 0) as total").To(&item.TotalPoint).
		OrderBy("total desc, u.name asc")

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
