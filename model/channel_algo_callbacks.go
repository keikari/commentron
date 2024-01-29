// Code generated by SQLBoiler 4.16.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package model

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// ChannelAlgoCallback is an object representing the database table.
type ChannelAlgoCallback struct {
	ChannelID string    `boil:"channel_id" json:"channel_id" toml:"channel_id" yaml:"channel_id"`
	WatcherID uint      `boil:"watcher_id" json:"watcher_id" toml:"watcher_id" yaml:"watcher_id"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *channelAlgoCallbackR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L channelAlgoCallbackL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ChannelAlgoCallbackColumns = struct {
	ChannelID string
	WatcherID string
	CreatedAt string
	UpdatedAt string
}{
	ChannelID: "channel_id",
	WatcherID: "watcher_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var ChannelAlgoCallbackTableColumns = struct {
	ChannelID string
	WatcherID string
	CreatedAt string
	UpdatedAt string
}{
	ChannelID: "channel_algo_callbacks.channel_id",
	WatcherID: "channel_algo_callbacks.watcher_id",
	CreatedAt: "channel_algo_callbacks.created_at",
	UpdatedAt: "channel_algo_callbacks.updated_at",
}

// Generated where

type whereHelperuint struct{ field string }

func (w whereHelperuint) EQ(x uint) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperuint) NEQ(x uint) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperuint) LT(x uint) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperuint) LTE(x uint) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperuint) GT(x uint) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperuint) GTE(x uint) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperuint) IN(slice []uint) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperuint) NIN(slice []uint) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var ChannelAlgoCallbackWhere = struct {
	ChannelID whereHelperstring
	WatcherID whereHelperuint
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	ChannelID: whereHelperstring{field: "`channel_algo_callbacks`.`channel_id`"},
	WatcherID: whereHelperuint{field: "`channel_algo_callbacks`.`watcher_id`"},
	CreatedAt: whereHelpertime_Time{field: "`channel_algo_callbacks`.`created_at`"},
	UpdatedAt: whereHelpertime_Time{field: "`channel_algo_callbacks`.`updated_at`"},
}

// ChannelAlgoCallbackRels is where relationship names are stored.
var ChannelAlgoCallbackRels = struct {
}{}

// channelAlgoCallbackR is where relationships are stored.
type channelAlgoCallbackR struct {
}

// NewStruct creates a new relationship struct
func (*channelAlgoCallbackR) NewStruct() *channelAlgoCallbackR {
	return &channelAlgoCallbackR{}
}

// channelAlgoCallbackL is where Load methods for each relationship are stored.
type channelAlgoCallbackL struct{}

var (
	channelAlgoCallbackAllColumns            = []string{"channel_id", "watcher_id", "created_at", "updated_at"}
	channelAlgoCallbackColumnsWithoutDefault = []string{"channel_id", "watcher_id", "created_at", "updated_at"}
	channelAlgoCallbackColumnsWithDefault    = []string{}
	channelAlgoCallbackPrimaryKeyColumns     = []string{"channel_id", "watcher_id"}
	channelAlgoCallbackGeneratedColumns      = []string{}
)

type (
	// ChannelAlgoCallbackSlice is an alias for a slice of pointers to ChannelAlgoCallback.
	// This should almost always be used instead of []ChannelAlgoCallback.
	ChannelAlgoCallbackSlice []*ChannelAlgoCallback

	channelAlgoCallbackQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	channelAlgoCallbackType                 = reflect.TypeOf(&ChannelAlgoCallback{})
	channelAlgoCallbackMapping              = queries.MakeStructMapping(channelAlgoCallbackType)
	channelAlgoCallbackPrimaryKeyMapping, _ = queries.BindMapping(channelAlgoCallbackType, channelAlgoCallbackMapping, channelAlgoCallbackPrimaryKeyColumns)
	channelAlgoCallbackInsertCacheMut       sync.RWMutex
	channelAlgoCallbackInsertCache          = make(map[string]insertCache)
	channelAlgoCallbackUpdateCacheMut       sync.RWMutex
	channelAlgoCallbackUpdateCache          = make(map[string]updateCache)
	channelAlgoCallbackUpsertCacheMut       sync.RWMutex
	channelAlgoCallbackUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single channelAlgoCallback record from the query.
func (q channelAlgoCallbackQuery) One(exec boil.Executor) (*ChannelAlgoCallback, error) {
	o := &ChannelAlgoCallback{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: failed to execute a one query for channel_algo_callbacks")
	}

	return o, nil
}

// All returns all ChannelAlgoCallback records from the query.
func (q channelAlgoCallbackQuery) All(exec boil.Executor) (ChannelAlgoCallbackSlice, error) {
	var o []*ChannelAlgoCallback

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "model: failed to assign all query results to ChannelAlgoCallback slice")
	}

	return o, nil
}

// Count returns the count of all ChannelAlgoCallback records in the query.
func (q channelAlgoCallbackQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to count channel_algo_callbacks rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q channelAlgoCallbackQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "model: failed to check if channel_algo_callbacks exists")
	}

	return count > 0, nil
}

// ChannelAlgoCallbacks retrieves all the records using an executor.
func ChannelAlgoCallbacks(mods ...qm.QueryMod) channelAlgoCallbackQuery {
	mods = append(mods, qm.From("`channel_algo_callbacks`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`channel_algo_callbacks`.*"})
	}

	return channelAlgoCallbackQuery{q}
}

// FindChannelAlgoCallback retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindChannelAlgoCallback(exec boil.Executor, channelID string, watcherID uint, selectCols ...string) (*ChannelAlgoCallback, error) {
	channelAlgoCallbackObj := &ChannelAlgoCallback{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `channel_algo_callbacks` where `channel_id`=? AND `watcher_id`=?", sel,
	)

	q := queries.Raw(query, channelID, watcherID)

	err := q.Bind(nil, exec, channelAlgoCallbackObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: unable to select from channel_algo_callbacks")
	}

	return channelAlgoCallbackObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ChannelAlgoCallback) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("model: no channel_algo_callbacks provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(channelAlgoCallbackColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	channelAlgoCallbackInsertCacheMut.RLock()
	cache, cached := channelAlgoCallbackInsertCache[key]
	channelAlgoCallbackInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			channelAlgoCallbackAllColumns,
			channelAlgoCallbackColumnsWithDefault,
			channelAlgoCallbackColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(channelAlgoCallbackType, channelAlgoCallbackMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(channelAlgoCallbackType, channelAlgoCallbackMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `channel_algo_callbacks` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `channel_algo_callbacks` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `channel_algo_callbacks` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, channelAlgoCallbackPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}
	_, err = exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "model: unable to insert into channel_algo_callbacks")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ChannelID,
		o.WatcherID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}
	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "model: unable to populate default values for channel_algo_callbacks")
	}

CacheNoHooks:
	if !cached {
		channelAlgoCallbackInsertCacheMut.Lock()
		channelAlgoCallbackInsertCache[key] = cache
		channelAlgoCallbackInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the ChannelAlgoCallback.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ChannelAlgoCallback) Update(exec boil.Executor, columns boil.Columns) error {
	var err error
	key := makeCacheKey(columns, nil)
	channelAlgoCallbackUpdateCacheMut.RLock()
	cache, cached := channelAlgoCallbackUpdateCache[key]
	channelAlgoCallbackUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			channelAlgoCallbackAllColumns,
			channelAlgoCallbackPrimaryKeyColumns,
		)
		if len(wl) == 0 {
			return errors.New("model: unable to update channel_algo_callbacks, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `channel_algo_callbacks` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, channelAlgoCallbackPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(channelAlgoCallbackType, channelAlgoCallbackMapping, append(wl, channelAlgoCallbackPrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	_, err = exec.Exec(cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "model: unable to update channel_algo_callbacks row")
	}

	if !cached {
		channelAlgoCallbackUpdateCacheMut.Lock()
		channelAlgoCallbackUpdateCache[key] = cache
		channelAlgoCallbackUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAll updates all rows with the specified column values.
func (q channelAlgoCallbackQuery) UpdateAll(exec boil.Executor, cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec(exec)
	if err != nil {
		return errors.Wrap(err, "model: unable to update all for channel_algo_callbacks")
	}

	return nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ChannelAlgoCallbackSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("model: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), channelAlgoCallbackPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `channel_algo_callbacks` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, channelAlgoCallbackPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "model: unable to update all in channelAlgoCallback slice")
	}

	return nil
}

var mySQLChannelAlgoCallbackUniqueColumns = []string{}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ChannelAlgoCallback) Upsert(exec boil.Executor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("model: no channel_algo_callbacks provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(channelAlgoCallbackColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLChannelAlgoCallbackUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
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
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	channelAlgoCallbackUpsertCacheMut.RLock()
	cache, cached := channelAlgoCallbackUpsertCache[key]
	channelAlgoCallbackUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			channelAlgoCallbackAllColumns,
			channelAlgoCallbackColumnsWithDefault,
			channelAlgoCallbackColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			channelAlgoCallbackAllColumns,
			channelAlgoCallbackPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("model: unable to upsert channel_algo_callbacks, could not build update column list")
		}

		ret := strmangle.SetComplement(channelAlgoCallbackAllColumns, strmangle.SetIntersect(insert, update))

		cache.query = buildUpsertQueryMySQL(dialect, "`channel_algo_callbacks`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `channel_algo_callbacks` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(channelAlgoCallbackType, channelAlgoCallbackMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(channelAlgoCallbackType, channelAlgoCallbackMapping, ret)
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
	_, err = exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "model: unable to upsert for channel_algo_callbacks")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(channelAlgoCallbackType, channelAlgoCallbackMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "model: unable to retrieve unique values for channel_algo_callbacks")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}
	err = exec.QueryRow(cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "model: unable to populate default values for channel_algo_callbacks")
	}

CacheNoHooks:
	if !cached {
		channelAlgoCallbackUpsertCacheMut.Lock()
		channelAlgoCallbackUpsertCache[key] = cache
		channelAlgoCallbackUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single ChannelAlgoCallback record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ChannelAlgoCallback) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("model: no ChannelAlgoCallback provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), channelAlgoCallbackPrimaryKeyMapping)
	sql := "DELETE FROM `channel_algo_callbacks` WHERE `channel_id`=? AND `watcher_id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "model: unable to delete from channel_algo_callbacks")
	}

	return nil
}

// DeleteAll deletes all matching rows.
func (q channelAlgoCallbackQuery) DeleteAll(exec boil.Executor) error {
	if q.Query == nil {
		return errors.New("model: no channelAlgoCallbackQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec(exec)
	if err != nil {
		return errors.Wrap(err, "model: unable to delete all from channel_algo_callbacks")
	}

	return nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ChannelAlgoCallbackSlice) DeleteAll(exec boil.Executor) error {
	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), channelAlgoCallbackPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `channel_algo_callbacks` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, channelAlgoCallbackPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}
	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "model: unable to delete all from channelAlgoCallback slice")
	}

	return nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *ChannelAlgoCallback) Reload(exec boil.Executor) error {
	ret, err := FindChannelAlgoCallback(exec, o.ChannelID, o.WatcherID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ChannelAlgoCallbackSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ChannelAlgoCallbackSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), channelAlgoCallbackPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `channel_algo_callbacks`.* FROM `channel_algo_callbacks` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, channelAlgoCallbackPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "model: unable to reload all in ChannelAlgoCallbackSlice")
	}

	*o = slice

	return nil
}

// ChannelAlgoCallbackExists checks if the ChannelAlgoCallback row exists.
func ChannelAlgoCallbackExists(exec boil.Executor, channelID string, watcherID uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `channel_algo_callbacks` where `channel_id`=? AND `watcher_id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, channelID, watcherID)
	}
	row := exec.QueryRow(sql, channelID, watcherID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "model: unable to check if channel_algo_callbacks exists")
	}

	return exists, nil
}

// Exists checks if the ChannelAlgoCallback row exists.
func (o *ChannelAlgoCallback) Exists(exec boil.Executor) (bool, error) {
	return ChannelAlgoCallbackExists(exec, o.ChannelID, o.WatcherID)
}
