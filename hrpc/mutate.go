// Copyright (C) 2015  The GoHBase Authors.  All rights reserved.
// This file is part of GoHBase.
// Use of this source code is governed by the Apache License 2.0
// that can be found in the COPYING file.

package hrpc

import (
	"context"
	"encoding/binary"
	"errors"
	"math"
	"time"

	"github.com/LoneWolf38/gohbase/pb"
	"google.golang.org/protobuf/proto"
)

var attributeNameTTL = "_ttl"

// DurabilityType is used to set durability for Durability option
type DurabilityType int32

const (
	// UseDefault is USER_DEFAULT
	UseDefault DurabilityType = iota
	// SkipWal is SKIP_WAL
	SkipWal
	// AsyncWal is ASYNC_WAL
	AsyncWal
	// SyncWal is SYNC_WAL
	SyncWal
	// FsyncWal is FSYNC_WAL
	FsyncWal
)

const (
	putType                 = 4
	deleteType              = 8
	deleteFamilyVersionType = 10
	deleteColumnType        = 12
	deleteFamilyType        = 14
)

var emptyQualifier = map[string][]byte{"": nil}

// Mutate represents a mutation on HBase.
type Mutate struct {
	base

	mutationType pb.MutationProto_MutationType //*int32

	// values is a map of column families to a map of column qualifiers to bytes
	values map[string]map[string][]byte

	ttl              []byte
	timestamp        uint64
	durability       DurabilityType
	deleteOneVersion bool
	skipbatch        bool
}

// TTL sets a time-to-live for mutation queries.
// The value will be in millisecond resolution.
func TTL(t time.Duration) func(Call) error {
	return func(o Call) error {
		m, ok := o.(*Mutate)
		if !ok {
			return errors.New("'TTL' option can only be used with mutation queries")
		}

		buf := make([]byte, 8)
		binary.BigEndian.PutUint64(buf, uint64(t.Nanoseconds()/1e6))
		m.ttl = buf

		return nil
	}
}

// Timestamp sets timestamp for mutation queries.
// The time object passed will be rounded to a millisecond resolution, as by default,
// if no timestamp is provided, HBase sets it to current time in milliseconds.
// In order to have custom time precision, use TimestampUint64 call option for
// mutation requests and corresponding TimeRangeUint64 for retrieval requests.
func Timestamp(ts time.Time) func(Call) error {
	return func(o Call) error {
		m, ok := o.(*Mutate)
		if !ok {
			return errors.New("'Timestamp' option can only be used with mutation queries")
		}
		m.timestamp = uint64(ts.UnixNano() / 1e6)
		return nil
	}
}

// TimestampUint64 sets timestamp for mutation queries.
func TimestampUint64(ts uint64) func(Call) error {
	return func(o Call) error {
		m, ok := o.(*Mutate)
		if !ok {
			return errors.New("'TimestampUint64' option can only be used with mutation queries")
		}
		m.timestamp = ts
		return nil
	}
}

// Durability sets durability for mutation queries.
func Durability(d DurabilityType) func(Call) error {
	return func(o Call) error {
		m, ok := o.(*Mutate)
		if !ok {
			return errors.New("'Durability' option can only be used with mutation queries")
		}
		if d < UseDefault || d > FsyncWal {
			return errors.New("invalid durability value")
		}
		m.durability = d
		return nil
	}
}

// DeleteOneVersion is a delete option that can be passed in order to delete only
// one latest version of the specified qualifiers. Without timestamp specified,
// it will have no effect for delete specific column families request.
// If a Timestamp option is passed along, only the version at that timestamp will be removed
// for delete specific column families and/or qualifier request.
// This option cannot be used for delete entire row request.
func DeleteOneVersion() func(Call) error {
	return func(o Call) error {
		m, ok := o.(*Mutate)
		if !ok {
			return errors.New("'DeleteOneVersion' option can only be used with mutation queries")
		}
		m.deleteOneVersion = true
		return nil
	}
}

// baseMutate returns a Mutate struct without the mutationType filled in.
func baseMutate(ctx context.Context, table, key []byte, values map[string]map[string][]byte,
	options ...func(Call) error,
) (*Mutate, error) {
	m := &Mutate{
		base: base{
			table:    table,
			key:      key,
			ctx:      ctx,
			resultch: make(chan RPCResult, 1),
		},
		values:    values,
		timestamp: MaxTimestamp,
	}
	err := applyOptions(m, options...)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// NewPut creates a new Mutation request to insert the given
// family-column-values in the given row key of the given table.
func NewPut(ctx context.Context, table, key []byte,
	values map[string]map[string][]byte, options ...func(Call) error,
) (*Mutate, error) {
	m, err := baseMutate(ctx, table, key, values, options...)
	if err != nil {
		return nil, err
	}
	m.mutationType = pb.MutationProto_PUT
	return m, nil
}

// NewPutStr is just like NewPut but takes table and key as strings.
func NewPutStr(ctx context.Context, table, key string,
	values map[string]map[string][]byte, options ...func(Call) error,
) (*Mutate, error) {
	return NewPut(ctx, []byte(table), []byte(key), values, options...)
}

// NewDel is used to perform Delete operations on a single row.
// To delete entire row, values should be nil.
//
// To delete specific families, qualifiers map should be nil:
//
//	 map[string]map[string][]byte{
//			"cf1": nil,
//			"cf2": nil,
//	 }
//
// To delete specific qualifiers:
//
//	 map[string]map[string][]byte{
//	     "cf": map[string][]byte{
//				"q1": nil,
//				"q2": nil,
//			},
//	 }
//
// To delete all versions before and at a timestamp, pass hrpc.Timestamp() option.
// By default all versions will be removed.
//
// To delete only a specific version at a timestamp, pass hrpc.DeleteOneVersion() option
// along with a timestamp. For delete specific qualifiers request, if timestamp is not
// passed, only the latest version will be removed. For delete specific families request,
// the timestamp should be passed or it will have no effect as it's an expensive
// operation to perform.
func NewDel(ctx context.Context, table, key []byte,
	values map[string]map[string][]byte, options ...func(Call) error,
) (*Mutate, error) {
	m, err := baseMutate(ctx, table, key, values, options...)
	if err != nil {
		return nil, err
	}

	if len(m.values) == 0 && m.deleteOneVersion {
		return nil, errors.New(
			"'DeleteOneVersion' option cannot be specified for delete entire row request")
	}

	m.mutationType = pb.MutationProto_DELETE
	return m, nil
}

// NewDelStr is just like NewDel but takes table and key as strings.
func NewDelStr(ctx context.Context, table, key string,
	values map[string]map[string][]byte, options ...func(Call) error,
) (*Mutate, error) {
	return NewDel(ctx, []byte(table), []byte(key), values, options...)
}

// NewApp creates a new Mutation request to append the given
// family-column-values into the existing cells in HBase (or create them if
// needed), in given row key of the given table.
func NewApp(ctx context.Context, table, key []byte,
	values map[string]map[string][]byte, options ...func(Call) error,
) (*Mutate, error) {
	m, err := baseMutate(ctx, table, key, values, options...)
	if err != nil {
		return nil, err
	}
	m.mutationType = pb.MutationProto_APPEND
	return m, nil
}

// NewAppStr is just like NewApp but takes table and key as strings.
func NewAppStr(ctx context.Context, table, key string,
	values map[string]map[string][]byte, options ...func(Call) error,
) (*Mutate, error) {
	return NewApp(ctx, []byte(table), []byte(key), values, options...)
}

// NewIncSingle creates a new Mutation request that will increment the given value
// by amount in HBase under the given table, key, family and qualifier.
func NewIncSingle(ctx context.Context, table, key []byte, family, qualifier string,
	amount int64, options ...func(Call) error,
) (*Mutate, error) {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(amount))
	value := map[string]map[string][]byte{family: {qualifier: buf}}
	return NewInc(ctx, table, key, value, options...)
}

// NewIncStrSingle is just like NewIncSingle but takes table and key as strings.
func NewIncStrSingle(ctx context.Context, table, key, family, qualifier string,
	amount int64, options ...func(Call) error,
) (*Mutate, error) {
	return NewIncSingle(ctx, []byte(table), []byte(key), family, qualifier, amount, options...)
}

// NewInc creates a new Mutation request that will increment the given values
// in HBase under the given table and key.
func NewInc(ctx context.Context, table, key []byte,
	values map[string]map[string][]byte, options ...func(Call) error,
) (*Mutate, error) {
	m, err := baseMutate(ctx, table, key, values, options...)
	if err != nil {
		return nil, err
	}
	m.mutationType = pb.MutationProto_INCREMENT
	return m, nil
}

// NewIncStr is just like NewInc but takes table and key as strings.
func NewIncStr(ctx context.Context, table, key string,
	values map[string]map[string][]byte, options ...func(Call) error,
) (*Mutate, error) {
	return NewInc(ctx, []byte(table), []byte(key), values, options...)
}

// Name returns the name of this RPC call.
func (m *Mutate) Name() string {
	return "Mutate"
}

// Description returns the string of type of mutation being performed
func (m *Mutate) Description() string {
	return pb.MutationProto_MutationType_name[int32(m.mutationType)]
}

// SkipBatch returns true if the Mutate request shouldn't be batched,
// but should be sent to Region Server right away.
func (m *Mutate) SkipBatch() bool {
	return m.skipbatch
}

// Values returns the internal values object
// which should be treated as read-only.
// This would typically be used for calculations
// related to metrics and workloads
func (m *Mutate) Values() map[string]map[string][]byte {
	return m.values
}

func (m *Mutate) setSkipBatch(v bool) {
	m.skipbatch = v
}

var (
	MutationProtoDeleteFamilyVersion    = pb.MutationProto_DELETE_FAMILY_VERSION.Enum()
	MutationProtoDeleteFamily           = pb.MutationProto_DELETE_FAMILY.Enum()
	MutationProtoDeleteOneVersion       = pb.MutationProto_DELETE_ONE_VERSION.Enum()
	MutationProtoDeleteMultipleVersions = pb.MutationProto_DELETE_MULTIPLE_VERSIONS.Enum()
)

func (m *Mutate) valuesToProto(ts *uint64) []*pb.MutationProto_ColumnValue {
	cvs := make([]*pb.MutationProto_ColumnValue, len(m.values))
	i := 0
	for k, v := range m.values {
		// And likewise, each item in each column needs to be converted to a
		// protobuf QualifierValue

		// if it's a delete, figure out the type
		var dt *pb.MutationProto_DeleteType
		if m.mutationType == pb.MutationProto_DELETE {
			if len(v) == 0 {
				// delete the whole column family
				if m.deleteOneVersion {
					dt = MutationProtoDeleteFamilyVersion
				} else {
					dt = MutationProtoDeleteFamily
				}
				// add empty qualifier
				if v == nil {
					v = emptyQualifier
				}
			} else {
				// delete specific qualifiers
				if m.deleteOneVersion {
					dt = MutationProtoDeleteOneVersion
				} else {
					dt = MutationProtoDeleteMultipleVersions
				}
			}
		}

		qvs := make([]*pb.MutationProto_ColumnValue_QualifierValue, len(v))
		j := 0
		for k1, v1 := range v {
			qvs[j] = &pb.MutationProto_ColumnValue_QualifierValue{
				Qualifier:  []byte(k1),
				Value:      v1,
				Timestamp:  ts,
				DeleteType: dt,
			}
			j++
		}
		cvs[i] = &pb.MutationProto_ColumnValue{
			Family:         []byte(k),
			QualifierValue: qvs,
		}
		i++
	}
	return cvs
}

func cellblockLen(rowLen, familyLen, qualifierLen, valueLen int) int {
	keyLength := 2 + rowLen + 1 + familyLen + qualifierLen + 8 + 1
	keyValueLength := 4 + 4 + keyLength + valueLen
	return 4 + keyValueLength
}

func appendCellblock(row []byte, family, qualifier string, value []byte, ts uint64, typ byte,
	cbs []byte,
) []byte {
	// cellblock layout:
	//
	// Header:
	// 4 byte length of key + value
	// 4 byte length of key
	// 4 byte length of value
	//
	// Key:
	// 2 byte length of row
	// <row>
	// 1 byte length of row family
	// <family>
	// <qualifier>
	// 8 byte timestamp
	// 1 byte type
	//
	// Value:
	// <value>
	keylength := 2 + len(row) + 1 + len(family) + len(qualifier) + 8 + 1
	valuelength := len(value)

	keyvaluelength := 4 + 4 + keylength + valuelength
	i := len(cbs)
	cbs = append(cbs, make([]byte,
		cellblockLen(len(row), len(family), len(qualifier), len(value)))...)

	// Header:
	binary.BigEndian.PutUint32(cbs[i:], uint32(keyvaluelength))
	i += 4
	binary.BigEndian.PutUint32(cbs[i:], uint32(keylength))
	i += 4
	binary.BigEndian.PutUint32(cbs[i:], uint32(valuelength))
	i += 4

	// Key:
	binary.BigEndian.PutUint16(cbs[i:], uint16(len(row)))
	i += 2
	i += copy(cbs[i:], row)
	cbs[i] = byte(len(family))
	i++
	i += copy(cbs[i:], family)
	i += copy(cbs[i:], qualifier)
	binary.BigEndian.PutUint64(cbs[i:], ts)
	i += 8
	cbs[i] = typ
	i++

	// Value:
	copy(cbs[i:], value)

	return cbs
}

func (m *Mutate) valuesToCellblocks() ([]byte, int32, uint32) {
	if len(m.values) == 0 {
		return nil, 0, 0
	}
	var cbsLen int
	var count int
	for family, v := range m.values {
		if v == nil {
			v = emptyQualifier
		}
		count += len(v)
		for k1, v1 := range v {
			cbsLen += cellblockLen(len(m.key), len(family), len(k1), len(v1))
		}
	}
	cbs := make([]byte, 0, cbsLen)

	var ts uint64
	if m.timestamp == MaxTimestamp {
		ts = math.MaxInt64 // Java's Long.MAX_VALUE use for HBase's LATEST_TIMESTAMP
	} else {
		ts = m.timestamp
	}
	for family, v := range m.values {
		// figure out mutation type
		var mt byte
		if m.mutationType == pb.MutationProto_DELETE {
			if len(v) == 0 {
				// delete the whole column family
				if m.deleteOneVersion {
					mt = deleteFamilyVersionType
				} else {
					mt = deleteFamilyType
				}
				// add empty qualifier
				if v == nil {
					v = emptyQualifier
				}
			} else {
				// delete specific qualifiers
				if m.deleteOneVersion {
					mt = deleteType
				} else {
					mt = deleteColumnType
				}
			}
		} else {
			mt = putType
		}

		for k1, v1 := range v {
			cbs = appendCellblock(m.key, family, k1, v1, ts, mt, cbs)
		}
	}
	if len(cbs) != cbsLen {
		panic("cellblocks len mismatch")
	}
	return cbs, int32(count), uint32(len(cbs))
}

var durabilities = []*pb.MutationProto_Durability{
	pb.MutationProto_Durability(UseDefault).Enum(),
	pb.MutationProto_Durability(SkipWal).Enum(),
	pb.MutationProto_Durability(AsyncWal).Enum(),
	pb.MutationProto_Durability(SyncWal).Enum(),
	pb.MutationProto_Durability(FsyncWal).Enum(),
}

func (m *Mutate) toProto(isCellblocks bool, cbs [][]byte) (*pb.MutateRequest, [][]byte, uint32) {
	var ts *uint64
	if m.timestamp != MaxTimestamp {
		ts = &m.timestamp
	}

	var size uint32
	mProto := &pb.MutationProto{
		Row:        m.key,
		MutateType: &m.mutationType,
		Durability: durabilities[m.durability],
		Timestamp:  ts,
	}

	if isCellblocks {
		// if cellblocks we only add associated cell count as the actual
		// data will be sent after protobuf
		cellblocks, count, sz := m.valuesToCellblocks()
		mProto.AssociatedCellCount = &count
		size = sz
		if size > 0 {
			cbs = append(cbs, cellblocks)
		}
	} else {
		// otherwise, convert the values to protobuf
		mProto.ColumnValue = m.valuesToProto(ts)
	}

	if len(m.ttl) > 0 {
		mProto.Attribute = append(mProto.Attribute, &pb.NameBytesPair{
			Name:  &attributeNameTTL,
			Value: m.ttl,
		})
	}

	return &pb.MutateRequest{
		Region:   m.regionSpecifier(),
		Mutation: mProto,
	}, cbs, size
}

// ToProto converts this mutate RPC into a protobuf message
func (m *Mutate) ToProto() proto.Message {
	p, _, _ := m.toProto(false, nil)
	return p
}

// NewResponse creates an empty protobuf message to read the response of this RPC.
func (m *Mutate) NewResponse() proto.Message {
	return &pb.MutateResponse{}
}

// DeserializeCellBlocks deserializes mutate result from cell blocks
func (m *Mutate) DeserializeCellBlocks(pm proto.Message, b []byte) (uint32, error) {
	resp := pm.(*pb.MutateResponse)
	if resp.Result == nil {
		// TODO: is this possible?
		return 0, nil
	}
	cells, read, err := deserializeCellBlocks(b, uint32(resp.Result.GetAssociatedCellCount()))
	if err != nil {
		return 0, err
	}
	resp.Result.Cell = append(resp.Result.Cell, cells...)
	return read, nil
}

func (m *Mutate) SerializeCellBlocks(cbs [][]byte) (proto.Message, [][]byte, uint32) {
	return m.toProto(true, cbs)
}

func (m *Mutate) CellBlocksEnabled() bool {
	// TODO: maybe have some global client option
	return true
}
