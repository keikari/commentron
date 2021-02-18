// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

	"github.com/pkg/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// CreatorSetting is an object representing the database table.
type CreatorSetting struct {
	ID                     uint64      `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatorChannelID       string      `boil:"creator_channel_id" json:"creator_channel_id" toml:"creator_channel_id" yaml:"creator_channel_id"`
	CommentsEnabled        null.Bool   `boil:"comments_enabled" json:"comments_enabled,omitempty" toml:"comments_enabled" yaml:"comments_enabled,omitempty"`
	MinTipAmmountComment   null.Uint64 `boil:"min_tip_ammount_comment" json:"min_tip_ammount_comment,omitempty" toml:"min_tip_ammount_comment" yaml:"min_tip_ammount_comment,omitempty"`
	MinTipAmmountSuperChat null.Uint64 `boil:"min_tip_ammount_super_chat" json:"min_tip_ammount_super_chat,omitempty" toml:"min_tip_ammount_super_chat" yaml:"min_tip_ammount_super_chat,omitempty"`
	MutedWords             null.String `boil:"muted_words" json:"muted_words,omitempty" toml:"muted_words" yaml:"muted_words,omitempty"`
	CreatedAt              time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt              time.Time   `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *creatorSettingR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L creatorSettingL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CreatorSettingColumns = struct {
	ID                     string
	CreatorChannelID       string
	CommentsEnabled        string
	MinTipAmmountComment   string
	MinTipAmmountSuperChat string
	MutedWords             string
	CreatedAt              string
	UpdatedAt              string
}{
	ID:                     "id",
	CreatorChannelID:       "creator_channel_id",
	CommentsEnabled:        "comments_enabled",
	MinTipAmmountComment:   "min_tip_ammount_comment",
	MinTipAmmountSuperChat: "min_tip_ammount_super_chat",
	MutedWords:             "muted_words",
	CreatedAt:              "created_at",
	UpdatedAt:              "updated_at",
}

// Generated where

type whereHelpernull_Uint64 struct{ field string }

func (w whereHelpernull_Uint64) EQ(x null.Uint64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Uint64) NEQ(x null.Uint64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Uint64) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Uint64) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Uint64) LT(x null.Uint64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Uint64) LTE(x null.Uint64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Uint64) GT(x null.Uint64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Uint64) GTE(x null.Uint64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var CreatorSettingWhere = struct {
	ID                     whereHelperuint64
	CreatorChannelID       whereHelperstring
	CommentsEnabled        whereHelpernull_Bool
	MinTipAmmountComment   whereHelpernull_Uint64
	MinTipAmmountSuperChat whereHelpernull_Uint64
	MutedWords             whereHelpernull_String
	CreatedAt              whereHelpertime_Time
	UpdatedAt              whereHelpertime_Time
}{
	ID:                     whereHelperuint64{field: "`creator_setting`.`id`"},
	CreatorChannelID:       whereHelperstring{field: "`creator_setting`.`creator_channel_id`"},
	CommentsEnabled:        whereHelpernull_Bool{field: "`creator_setting`.`comments_enabled`"},
	MinTipAmmountComment:   whereHelpernull_Uint64{field: "`creator_setting`.`min_tip_ammount_comment`"},
	MinTipAmmountSuperChat: whereHelpernull_Uint64{field: "`creator_setting`.`min_tip_ammount_super_chat`"},
	MutedWords:             whereHelpernull_String{field: "`creator_setting`.`muted_words`"},
	CreatedAt:              whereHelpertime_Time{field: "`creator_setting`.`created_at`"},
	UpdatedAt:              whereHelpertime_Time{field: "`creator_setting`.`updated_at`"},
}

// CreatorSettingRels is where relationship names are stored.
var CreatorSettingRels = struct {
	CreatorChannel string
}{
	CreatorChannel: "CreatorChannel",
}

// creatorSettingR is where relationships are stored.
type creatorSettingR struct {
	CreatorChannel *Channel
}

// NewStruct creates a new relationship struct
func (*creatorSettingR) NewStruct() *creatorSettingR {
	return &creatorSettingR{}
}

// creatorSettingL is where Load methods for each relationship are stored.
type creatorSettingL struct{}

var (
	creatorSettingAllColumns            = []string{"id", "creator_channel_id", "comments_enabled", "min_tip_ammount_comment", "min_tip_ammount_super_chat", "muted_words", "created_at", "updated_at"}
	creatorSettingColumnsWithoutDefault = []string{"creator_channel_id", "min_tip_ammount_comment", "min_tip_ammount_super_chat", "muted_words"}
	creatorSettingColumnsWithDefault    = []string{"id", "comments_enabled", "created_at", "updated_at"}
	creatorSettingPrimaryKeyColumns     = []string{"id"}
)

type (
	// CreatorSettingSlice is an alias for a slice of pointers to CreatorSetting.
	// This should generally be used opposed to []CreatorSetting.
	CreatorSettingSlice []*CreatorSetting

	creatorSettingQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	creatorSettingType                 = reflect.TypeOf(&CreatorSetting{})
	creatorSettingMapping              = queries.MakeStructMapping(creatorSettingType)
	creatorSettingPrimaryKeyMapping, _ = queries.BindMapping(creatorSettingType, creatorSettingMapping, creatorSettingPrimaryKeyColumns)
	creatorSettingInsertCacheMut       sync.RWMutex
	creatorSettingInsertCache          = make(map[string]insertCache)
	creatorSettingUpdateCacheMut       sync.RWMutex
	creatorSettingUpdateCache          = make(map[string]updateCache)
	creatorSettingUpsertCacheMut       sync.RWMutex
	creatorSettingUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single creatorSetting record from the query using the global executor.
func (q creatorSettingQuery) OneG() (*CreatorSetting, error) {
	return q.One(boil.GetDB())
}

// OneGP returns a single creatorSetting record from the query using the global executor, and panics on error.
func (q creatorSettingQuery) OneGP() *CreatorSetting {
	o, err := q.One(boil.GetDB())
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// OneP returns a single creatorSetting record from the query, and panics on error.
func (q creatorSettingQuery) OneP(exec boil.Executor) *CreatorSetting {
	o, err := q.One(exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single creatorSetting record from the query.
func (q creatorSettingQuery) One(exec boil.Executor) (*CreatorSetting, error) {
	o := &CreatorSetting{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: failed to execute a one query for creator_setting")
	}

	return o, nil
}

// AllG returns all CreatorSetting records from the query using the global executor.
func (q creatorSettingQuery) AllG() (CreatorSettingSlice, error) {
	return q.All(boil.GetDB())
}

// AllGP returns all CreatorSetting records from the query using the global executor, and panics on error.
func (q creatorSettingQuery) AllGP() CreatorSettingSlice {
	o, err := q.All(boil.GetDB())
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// AllP returns all CreatorSetting records from the query, and panics on error.
func (q creatorSettingQuery) AllP(exec boil.Executor) CreatorSettingSlice {
	o, err := q.All(exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all CreatorSetting records from the query.
func (q creatorSettingQuery) All(exec boil.Executor) (CreatorSettingSlice, error) {
	var o []*CreatorSetting

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "model: failed to assign all query results to CreatorSetting slice")
	}

	return o, nil
}

// CountG returns the count of all CreatorSetting records in the query, and panics on error.
func (q creatorSettingQuery) CountG() (int64, error) {
	return q.Count(boil.GetDB())
}

// CountGP returns the count of all CreatorSetting records in the query using the global executor, and panics on error.
func (q creatorSettingQuery) CountGP() int64 {
	c, err := q.Count(boil.GetDB())
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// CountP returns the count of all CreatorSetting records in the query, and panics on error.
func (q creatorSettingQuery) CountP(exec boil.Executor) int64 {
	c, err := q.Count(exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all CreatorSetting records in the query.
func (q creatorSettingQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to count creator_setting rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q creatorSettingQuery) ExistsG() (bool, error) {
	return q.Exists(boil.GetDB())
}

// ExistsGP checks if the row exists in the table using the global executor, and panics on error.
func (q creatorSettingQuery) ExistsGP() bool {
	e, err := q.Exists(boil.GetDB())
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// ExistsP checks if the row exists in the table, and panics on error.
func (q creatorSettingQuery) ExistsP(exec boil.Executor) bool {
	e, err := q.Exists(exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q creatorSettingQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "model: failed to check if creator_setting exists")
	}

	return count > 0, nil
}

// CreatorChannel pointed to by the foreign key.
func (o *CreatorSetting) CreatorChannel(mods ...qm.QueryMod) channelQuery {
	queryMods := []qm.QueryMod{
		qm.Where("claim_id=?", o.CreatorChannelID),
	}

	queryMods = append(queryMods, mods...)

	query := Channels(queryMods...)
	queries.SetFrom(query.Query, "`channel`")

	return query
}

// LoadCreatorChannel allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (creatorSettingL) LoadCreatorChannel(e boil.Executor, singular bool, maybeCreatorSetting interface{}, mods queries.Applicator) error {
	var slice []*CreatorSetting
	var object *CreatorSetting

	if singular {
		object = maybeCreatorSetting.(*CreatorSetting)
	} else {
		slice = *maybeCreatorSetting.(*[]*CreatorSetting)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &creatorSettingR{}
		}
		args = append(args, object.CreatorChannelID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &creatorSettingR{}
			}

			for _, a := range args {
				if a == obj.CreatorChannelID {
					continue Outer
				}
			}

			args = append(args, obj.CreatorChannelID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`channel`), qm.WhereIn(`claim_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Channel")
	}

	var resultSlice []*Channel
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Channel")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for channel")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for channel")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.CreatorChannel = foreign
		if foreign.R == nil {
			foreign.R = &channelR{}
		}
		foreign.R.CreatorChannelCreatorSettings = append(foreign.R.CreatorChannelCreatorSettings, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.CreatorChannelID == foreign.ClaimID {
				local.R.CreatorChannel = foreign
				if foreign.R == nil {
					foreign.R = &channelR{}
				}
				foreign.R.CreatorChannelCreatorSettings = append(foreign.R.CreatorChannelCreatorSettings, local)
				break
			}
		}
	}

	return nil
}

// SetCreatorChannelG of the creatorSetting to the related item.
// Sets o.R.CreatorChannel to related.
// Adds o to related.R.CreatorChannelCreatorSettings.
// Uses the global database handle.
func (o *CreatorSetting) SetCreatorChannelG(insert bool, related *Channel) error {
	return o.SetCreatorChannel(boil.GetDB(), insert, related)
}

// SetCreatorChannelP of the creatorSetting to the related item.
// Sets o.R.CreatorChannel to related.
// Adds o to related.R.CreatorChannelCreatorSettings.
// Panics on error.
func (o *CreatorSetting) SetCreatorChannelP(exec boil.Executor, insert bool, related *Channel) {
	if err := o.SetCreatorChannel(exec, insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetCreatorChannelGP of the creatorSetting to the related item.
// Sets o.R.CreatorChannel to related.
// Adds o to related.R.CreatorChannelCreatorSettings.
// Uses the global database handle and panics on error.
func (o *CreatorSetting) SetCreatorChannelGP(insert bool, related *Channel) {
	if err := o.SetCreatorChannel(boil.GetDB(), insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetCreatorChannel of the creatorSetting to the related item.
// Sets o.R.CreatorChannel to related.
// Adds o to related.R.CreatorChannelCreatorSettings.
func (o *CreatorSetting) SetCreatorChannel(exec boil.Executor, insert bool, related *Channel) error {
	var err error
	if insert {
		if err = related.Insert(exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `creator_setting` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"creator_channel_id"}),
		strmangle.WhereClause("`", "`", 0, creatorSettingPrimaryKeyColumns),
	)
	values := []interface{}{related.ClaimID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.CreatorChannelID = related.ClaimID
	if o.R == nil {
		o.R = &creatorSettingR{
			CreatorChannel: related,
		}
	} else {
		o.R.CreatorChannel = related
	}

	if related.R == nil {
		related.R = &channelR{
			CreatorChannelCreatorSettings: CreatorSettingSlice{o},
		}
	} else {
		related.R.CreatorChannelCreatorSettings = append(related.R.CreatorChannelCreatorSettings, o)
	}

	return nil
}

// CreatorSettings retrieves all the records using an executor.
func CreatorSettings(mods ...qm.QueryMod) creatorSettingQuery {
	mods = append(mods, qm.From("`creator_setting`"))
	return creatorSettingQuery{NewQuery(mods...)}
}

// FindCreatorSettingG retrieves a single record by ID.
func FindCreatorSettingG(iD uint64, selectCols ...string) (*CreatorSetting, error) {
	return FindCreatorSetting(boil.GetDB(), iD, selectCols...)
}

// FindCreatorSettingP retrieves a single record by ID with an executor, and panics on error.
func FindCreatorSettingP(exec boil.Executor, iD uint64, selectCols ...string) *CreatorSetting {
	retobj, err := FindCreatorSetting(exec, iD, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindCreatorSettingGP retrieves a single record by ID, and panics on error.
func FindCreatorSettingGP(iD uint64, selectCols ...string) *CreatorSetting {
	retobj, err := FindCreatorSetting(boil.GetDB(), iD, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindCreatorSetting retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCreatorSetting(exec boil.Executor, iD uint64, selectCols ...string) (*CreatorSetting, error) {
	creatorSettingObj := &CreatorSetting{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `creator_setting` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, creatorSettingObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: unable to select from creator_setting")
	}

	return creatorSettingObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *CreatorSetting) InsertG(columns boil.Columns) error {
	return o.Insert(boil.GetDB(), columns)
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *CreatorSetting) InsertP(exec boil.Executor, columns boil.Columns) {
	if err := o.Insert(exec, columns); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *CreatorSetting) InsertGP(columns boil.Columns) {
	if err := o.Insert(boil.GetDB(), columns); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CreatorSetting) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("model: no creator_setting provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(creatorSettingColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	creatorSettingInsertCacheMut.RLock()
	cache, cached := creatorSettingInsertCache[key]
	creatorSettingInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			creatorSettingAllColumns,
			creatorSettingColumnsWithDefault,
			creatorSettingColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(creatorSettingType, creatorSettingMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(creatorSettingType, creatorSettingMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `creator_setting` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `creator_setting` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `creator_setting` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, creatorSettingPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "model: unable to insert into creator_setting")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == creatorSettingMapping["ID"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "model: unable to populate default values for creator_setting")
	}

CacheNoHooks:
	if !cached {
		creatorSettingInsertCacheMut.Lock()
		creatorSettingInsertCache[key] = cache
		creatorSettingInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single CreatorSetting record using the global executor.
// See Update for more documentation.
func (o *CreatorSetting) UpdateG(columns boil.Columns) error {
	return o.Update(boil.GetDB(), columns)
}

// UpdateP uses an executor to update the CreatorSetting, and panics on error.
// See Update for more documentation.
func (o *CreatorSetting) UpdateP(exec boil.Executor, columns boil.Columns) {
	err := o.Update(exec, columns)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateGP a single CreatorSetting record using the global executor. Panics on error.
// See Update for more documentation.
func (o *CreatorSetting) UpdateGP(columns boil.Columns) {
	err := o.Update(boil.GetDB(), columns)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the CreatorSetting.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CreatorSetting) Update(exec boil.Executor, columns boil.Columns) error {
	var err error
	key := makeCacheKey(columns, nil)
	creatorSettingUpdateCacheMut.RLock()
	cache, cached := creatorSettingUpdateCache[key]
	creatorSettingUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			creatorSettingAllColumns,
			creatorSettingPrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return errors.New("model: unable to update creator_setting, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `creator_setting` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, creatorSettingPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(creatorSettingType, creatorSettingMapping, append(wl, creatorSettingPrimaryKeyColumns...))
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
		return errors.Wrap(err, "model: unable to update creator_setting row")
	}

	if !cached {
		creatorSettingUpdateCacheMut.Lock()
		creatorSettingUpdateCache[key] = cache
		creatorSettingUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q creatorSettingQuery) UpdateAllP(exec boil.Executor, cols M) {
	err := q.UpdateAll(exec, cols)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllG updates all rows with the specified column values.
func (q creatorSettingQuery) UpdateAllG(cols M) error {
	return q.UpdateAll(boil.GetDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q creatorSettingQuery) UpdateAll(exec boil.Executor, cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec(exec)
	if err != nil {
		return errors.Wrap(err, "model: unable to update all for creator_setting")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o CreatorSettingSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o CreatorSettingSlice) UpdateAllGP(cols M) {
	err := o.UpdateAll(boil.GetDB(), cols)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o CreatorSettingSlice) UpdateAllP(exec boil.Executor, cols M) {
	err := o.UpdateAll(exec, cols)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CreatorSettingSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), creatorSettingPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `creator_setting` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, creatorSettingPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "model: unable to update all in creatorSetting slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *CreatorSetting) UpsertG(updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(boil.GetDB(), updateColumns, insertColumns)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *CreatorSetting) UpsertGP(updateColumns, insertColumns boil.Columns) {
	if err := o.Upsert(boil.GetDB(), updateColumns, insertColumns); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *CreatorSetting) UpsertP(exec boil.Executor, updateColumns, insertColumns boil.Columns) {
	if err := o.Upsert(exec, updateColumns, insertColumns); err != nil {
		panic(boil.WrapErr(err))
	}
}

var mySQLCreatorSettingUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CreatorSetting) Upsert(exec boil.Executor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("model: no creator_setting provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(creatorSettingColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCreatorSettingUniqueColumns, o)

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

	creatorSettingUpsertCacheMut.RLock()
	cache, cached := creatorSettingUpsertCache[key]
	creatorSettingUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			creatorSettingAllColumns,
			creatorSettingColumnsWithDefault,
			creatorSettingColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			creatorSettingAllColumns,
			creatorSettingPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("model: unable to upsert creator_setting, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "creator_setting", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `creator_setting` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(creatorSettingType, creatorSettingMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(creatorSettingType, creatorSettingMapping, ret)
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

	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "model: unable to upsert for creator_setting")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == creatorSettingMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(creatorSettingType, creatorSettingMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "model: unable to retrieve unique values for creator_setting")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRow(cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "model: unable to populate default values for creator_setting")
	}

CacheNoHooks:
	if !cached {
		creatorSettingUpsertCacheMut.Lock()
		creatorSettingUpsertCache[key] = cache
		creatorSettingUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteG deletes a single CreatorSetting record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *CreatorSetting) DeleteG() error {
	return o.Delete(boil.GetDB())
}

// DeleteP deletes a single CreatorSetting record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *CreatorSetting) DeleteP(exec boil.Executor) {
	err := o.Delete(exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteGP deletes a single CreatorSetting record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *CreatorSetting) DeleteGP() {
	err := o.Delete(boil.GetDB())
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single CreatorSetting record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CreatorSetting) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("model: no CreatorSetting provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), creatorSettingPrimaryKeyMapping)
	sql := "DELETE FROM `creator_setting` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "model: unable to delete from creator_setting")
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q creatorSettingQuery) DeleteAllP(exec boil.Executor) {
	err := q.DeleteAll(exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q creatorSettingQuery) DeleteAll(exec boil.Executor) error {
	if q.Query == nil {
		return errors.New("model: no creatorSettingQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec(exec)
	if err != nil {
		return errors.Wrap(err, "model: unable to delete all from creator_setting")
	}

	return nil
}

// DeleteAllG deletes all rows in the slice.
func (o CreatorSettingSlice) DeleteAllG() error {
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o CreatorSettingSlice) DeleteAllP(exec boil.Executor) {
	err := o.DeleteAll(exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o CreatorSettingSlice) DeleteAllGP() {
	err := o.DeleteAll(boil.GetDB())
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CreatorSettingSlice) DeleteAll(exec boil.Executor) error {
	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), creatorSettingPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `creator_setting` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, creatorSettingPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "model: unable to delete all from creatorSetting slice")
	}

	return nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *CreatorSetting) ReloadG() error {
	if o == nil {
		return errors.New("model: no CreatorSetting provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *CreatorSetting) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadGP refetches the object from the database and panics on error.
func (o *CreatorSetting) ReloadGP() {
	if err := o.Reload(boil.GetDB()); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *CreatorSetting) Reload(exec boil.Executor) error {
	ret, err := FindCreatorSetting(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CreatorSettingSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("model: empty CreatorSettingSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *CreatorSettingSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *CreatorSettingSlice) ReloadAllGP() {
	if err := o.ReloadAll(boil.GetDB()); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CreatorSettingSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CreatorSettingSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), creatorSettingPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `creator_setting`.* FROM `creator_setting` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, creatorSettingPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "model: unable to reload all in CreatorSettingSlice")
	}

	*o = slice

	return nil
}

// CreatorSettingExistsG checks if the CreatorSetting row exists.
func CreatorSettingExistsG(iD uint64) (bool, error) {
	return CreatorSettingExists(boil.GetDB(), iD)
}

// CreatorSettingExistsP checks if the CreatorSetting row exists. Panics on error.
func CreatorSettingExistsP(exec boil.Executor, iD uint64) bool {
	e, err := CreatorSettingExists(exec, iD)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// CreatorSettingExistsGP checks if the CreatorSetting row exists. Panics on error.
func CreatorSettingExistsGP(iD uint64) bool {
	e, err := CreatorSettingExists(boil.GetDB(), iD)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// CreatorSettingExists checks if the CreatorSetting row exists.
func CreatorSettingExists(exec boil.Executor, iD uint64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `creator_setting` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "model: unable to check if creator_setting exists")
	}

	return exists, nil
}
