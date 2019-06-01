// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// TwitterFeed is an object representing the database table.
type TwitterFeed struct {
	ID              int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	GuildID         int64     `boil:"guild_id" json:"guild_id" toml:"guild_id" yaml:"guild_id"`
	CreatedAt       time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	TwitterUsername string    `boil:"twitter_username" json:"twitter_username" toml:"twitter_username" yaml:"twitter_username"`
	TwitterUserID   int64     `boil:"twitter_user_id" json:"twitter_user_id" toml:"twitter_user_id" yaml:"twitter_user_id"`
	ChannelID       int64     `boil:"channel_id" json:"channel_id" toml:"channel_id" yaml:"channel_id"`
	Enabled         bool      `boil:"enabled" json:"enabled" toml:"enabled" yaml:"enabled"`

	R *twitterFeedR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L twitterFeedL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TwitterFeedColumns = struct {
	ID              string
	GuildID         string
	CreatedAt       string
	TwitterUsername string
	TwitterUserID   string
	ChannelID       string
	Enabled         string
}{
	ID:              "id",
	GuildID:         "guild_id",
	CreatedAt:       "created_at",
	TwitterUsername: "twitter_username",
	TwitterUserID:   "twitter_user_id",
	ChannelID:       "channel_id",
	Enabled:         "enabled",
}

// Generated where

type whereHelperint64 struct{ field string }

func (w whereHelperint64) EQ(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint64) NEQ(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint64) LT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint64) LTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint64) GT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint64) GTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelperbool struct{ field string }

func (w whereHelperbool) EQ(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperbool) NEQ(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperbool) LT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperbool) LTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperbool) GT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperbool) GTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

var TwitterFeedWhere = struct {
	ID              whereHelperint64
	GuildID         whereHelperint64
	CreatedAt       whereHelpertime_Time
	TwitterUsername whereHelperstring
	TwitterUserID   whereHelperint64
	ChannelID       whereHelperint64
	Enabled         whereHelperbool
}{
	ID:              whereHelperint64{field: `id`},
	GuildID:         whereHelperint64{field: `guild_id`},
	CreatedAt:       whereHelpertime_Time{field: `created_at`},
	TwitterUsername: whereHelperstring{field: `twitter_username`},
	TwitterUserID:   whereHelperint64{field: `twitter_user_id`},
	ChannelID:       whereHelperint64{field: `channel_id`},
	Enabled:         whereHelperbool{field: `enabled`},
}

// TwitterFeedRels is where relationship names are stored.
var TwitterFeedRels = struct {
}{}

// twitterFeedR is where relationships are stored.
type twitterFeedR struct {
}

// NewStruct creates a new relationship struct
func (*twitterFeedR) NewStruct() *twitterFeedR {
	return &twitterFeedR{}
}

// twitterFeedL is where Load methods for each relationship are stored.
type twitterFeedL struct{}

var (
	twitterFeedColumns               = []string{"id", "guild_id", "created_at", "twitter_username", "twitter_user_id", "channel_id", "enabled"}
	twitterFeedColumnsWithoutDefault = []string{"guild_id", "created_at", "twitter_username", "twitter_user_id", "channel_id", "enabled"}
	twitterFeedColumnsWithDefault    = []string{"id"}
	twitterFeedPrimaryKeyColumns     = []string{"id"}
)

type (
	// TwitterFeedSlice is an alias for a slice of pointers to TwitterFeed.
	// This should generally be used opposed to []TwitterFeed.
	TwitterFeedSlice []*TwitterFeed

	twitterFeedQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	twitterFeedType                 = reflect.TypeOf(&TwitterFeed{})
	twitterFeedMapping              = queries.MakeStructMapping(twitterFeedType)
	twitterFeedPrimaryKeyMapping, _ = queries.BindMapping(twitterFeedType, twitterFeedMapping, twitterFeedPrimaryKeyColumns)
	twitterFeedInsertCacheMut       sync.RWMutex
	twitterFeedInsertCache          = make(map[string]insertCache)
	twitterFeedUpdateCacheMut       sync.RWMutex
	twitterFeedUpdateCache          = make(map[string]updateCache)
	twitterFeedUpsertCacheMut       sync.RWMutex
	twitterFeedUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single twitterFeed record from the query using the global executor.
func (q twitterFeedQuery) OneG(ctx context.Context) (*TwitterFeed, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single twitterFeed record from the query.
func (q twitterFeedQuery) One(ctx context.Context, exec boil.ContextExecutor) (*TwitterFeed, error) {
	o := &TwitterFeed{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for twitter_feeds")
	}

	return o, nil
}

// AllG returns all TwitterFeed records from the query using the global executor.
func (q twitterFeedQuery) AllG(ctx context.Context) (TwitterFeedSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all TwitterFeed records from the query.
func (q twitterFeedQuery) All(ctx context.Context, exec boil.ContextExecutor) (TwitterFeedSlice, error) {
	var o []*TwitterFeed

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to TwitterFeed slice")
	}

	return o, nil
}

// CountG returns the count of all TwitterFeed records in the query, and panics on error.
func (q twitterFeedQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all TwitterFeed records in the query.
func (q twitterFeedQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count twitter_feeds rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q twitterFeedQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q twitterFeedQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if twitter_feeds exists")
	}

	return count > 0, nil
}

// TwitterFeeds retrieves all the records using an executor.
func TwitterFeeds(mods ...qm.QueryMod) twitterFeedQuery {
	mods = append(mods, qm.From("\"twitter_feeds\""))
	return twitterFeedQuery{NewQuery(mods...)}
}

// FindTwitterFeedG retrieves a single record by ID.
func FindTwitterFeedG(ctx context.Context, iD int64, selectCols ...string) (*TwitterFeed, error) {
	return FindTwitterFeed(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindTwitterFeed retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTwitterFeed(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*TwitterFeed, error) {
	twitterFeedObj := &TwitterFeed{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"twitter_feeds\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, twitterFeedObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from twitter_feeds")
	}

	return twitterFeedObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *TwitterFeed) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TwitterFeed) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no twitter_feeds provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(twitterFeedColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	twitterFeedInsertCacheMut.RLock()
	cache, cached := twitterFeedInsertCache[key]
	twitterFeedInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			twitterFeedColumns,
			twitterFeedColumnsWithDefault,
			twitterFeedColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(twitterFeedType, twitterFeedMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(twitterFeedType, twitterFeedMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"twitter_feeds\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"twitter_feeds\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into twitter_feeds")
	}

	if !cached {
		twitterFeedInsertCacheMut.Lock()
		twitterFeedInsertCache[key] = cache
		twitterFeedInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single TwitterFeed record using the global executor.
// See Update for more documentation.
func (o *TwitterFeed) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the TwitterFeed.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TwitterFeed) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	twitterFeedUpdateCacheMut.RLock()
	cache, cached := twitterFeedUpdateCache[key]
	twitterFeedUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			twitterFeedColumns,
			twitterFeedPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update twitter_feeds, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"twitter_feeds\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, twitterFeedPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(twitterFeedType, twitterFeedMapping, append(wl, twitterFeedPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update twitter_feeds row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for twitter_feeds")
	}

	if !cached {
		twitterFeedUpdateCacheMut.Lock()
		twitterFeedUpdateCache[key] = cache
		twitterFeedUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (q twitterFeedQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q twitterFeedQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for twitter_feeds")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for twitter_feeds")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o TwitterFeedSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TwitterFeedSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), twitterFeedPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"twitter_feeds\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, twitterFeedPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in twitterFeed slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all twitterFeed")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *TwitterFeed) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *TwitterFeed) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no twitter_feeds provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(twitterFeedColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	twitterFeedUpsertCacheMut.RLock()
	cache, cached := twitterFeedUpsertCache[key]
	twitterFeedUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			twitterFeedColumns,
			twitterFeedColumnsWithDefault,
			twitterFeedColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			twitterFeedColumns,
			twitterFeedPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert twitter_feeds, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(twitterFeedPrimaryKeyColumns))
			copy(conflict, twitterFeedPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"twitter_feeds\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(twitterFeedType, twitterFeedMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(twitterFeedType, twitterFeedMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert twitter_feeds")
	}

	if !cached {
		twitterFeedUpsertCacheMut.Lock()
		twitterFeedUpsertCache[key] = cache
		twitterFeedUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteG deletes a single TwitterFeed record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *TwitterFeed) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single TwitterFeed record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TwitterFeed) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TwitterFeed provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), twitterFeedPrimaryKeyMapping)
	sql := "DELETE FROM \"twitter_feeds\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from twitter_feeds")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for twitter_feeds")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q twitterFeedQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no twitterFeedQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from twitter_feeds")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for twitter_feeds")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o TwitterFeedSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TwitterFeedSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TwitterFeed slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), twitterFeedPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"twitter_feeds\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, twitterFeedPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from twitterFeed slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for twitter_feeds")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *TwitterFeed) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no TwitterFeed provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *TwitterFeed) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTwitterFeed(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TwitterFeedSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty TwitterFeedSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TwitterFeedSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TwitterFeedSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), twitterFeedPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"twitter_feeds\".* FROM \"twitter_feeds\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, twitterFeedPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TwitterFeedSlice")
	}

	*o = slice

	return nil
}

// TwitterFeedExistsG checks if the TwitterFeed row exists.
func TwitterFeedExistsG(ctx context.Context, iD int64) (bool, error) {
	return TwitterFeedExists(ctx, boil.GetContextDB(), iD)
}

// TwitterFeedExists checks if the TwitterFeed row exists.
func TwitterFeedExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"twitter_feeds\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if twitter_feeds exists")
	}

	return exists, nil
}
