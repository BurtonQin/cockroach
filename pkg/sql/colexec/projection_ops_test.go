// Copyright 2018 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package colexec

import (
	"context"
	"fmt"
	"math/rand"
	"reflect"
	"testing"

	"github.com/cockroachdb/cockroach/pkg/col/coldata"
	"github.com/cockroachdb/cockroach/pkg/col/coldatatestutils"
	"github.com/cockroachdb/cockroach/pkg/col/typeconv"
	"github.com/cockroachdb/cockroach/pkg/settings/cluster"
	"github.com/cockroachdb/cockroach/pkg/sql/colexecbase"
	"github.com/cockroachdb/cockroach/pkg/sql/execinfra"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
	"github.com/cockroachdb/cockroach/pkg/sql/sqlbase"
	"github.com/cockroachdb/cockroach/pkg/sql/types"
	"github.com/cockroachdb/cockroach/pkg/util/leaktest"
	"github.com/cockroachdb/cockroach/pkg/util/randutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProjPlusInt64Int64ConstOp(t *testing.T) {
	defer leaktest.AfterTest(t)()
	ctx := context.Background()
	st := cluster.MakeTestingClusterSettings()
	evalCtx := tree.MakeTestingEvalContext(st)
	defer evalCtx.Stop(ctx)
	flowCtx := &execinfra.FlowCtx{
		EvalCtx: &evalCtx,
		Cfg: &execinfra.ServerConfig{
			Settings: st,
		},
	}
	runTests(t, []tuples{{{1}, {2}, {nil}}}, tuples{{1, 2}, {2, 3}, {nil, nil}}, orderedVerifier,
		func(input []colexecbase.Operator) (colexecbase.Operator, error) {
			return createTestProjectingOperator(
				ctx, flowCtx, input[0], []*types.T{types.Int},
				"@1 + 1" /* projectingExpr */, false, /* canFallbackToRowexec */
			)
		})
}

func TestProjPlusInt64Int64Op(t *testing.T) {
	defer leaktest.AfterTest(t)()
	ctx := context.Background()
	st := cluster.MakeTestingClusterSettings()
	evalCtx := tree.MakeTestingEvalContext(st)
	defer evalCtx.Stop(ctx)
	flowCtx := &execinfra.FlowCtx{
		EvalCtx: &evalCtx,
		Cfg: &execinfra.ServerConfig{
			Settings: st,
		},
	}
	runTests(t, []tuples{{{1, 2}, {3, 4}, {5, nil}}}, tuples{{1, 2, 3}, {3, 4, 7}, {5, nil, nil}},
		orderedVerifier,
		func(input []colexecbase.Operator) (colexecbase.Operator, error) {
			return createTestProjectingOperator(
				ctx, flowCtx, input[0], []*types.T{types.Int, types.Int},
				"@1 + @2" /* projectingExpr */, false, /* canFallbackToRowexec */
			)
		})
}

func TestProjDivFloat64Float64Op(t *testing.T) {
	defer leaktest.AfterTest(t)()
	ctx := context.Background()
	st := cluster.MakeTestingClusterSettings()
	evalCtx := tree.MakeTestingEvalContext(st)
	defer evalCtx.Stop(ctx)
	flowCtx := &execinfra.FlowCtx{
		EvalCtx: &evalCtx,
		Cfg: &execinfra.ServerConfig{
			Settings: st,
		},
	}
	runTests(t, []tuples{{{1.0, 2.0}, {3.0, 4.0}, {5.0, nil}}}, tuples{{1.0, 2.0, 0.5}, {3.0, 4.0, 0.75}, {5.0, nil, nil}},
		orderedVerifier,
		func(input []colexecbase.Operator) (colexecbase.Operator, error) {
			return createTestProjectingOperator(
				ctx, flowCtx, input[0], []*types.T{types.Float, types.Float},
				"@1 / @2" /* projectingExpr */, false, /* canFallbackToRowexec */
			)
		})
}

func benchmarkProjPlusInt64Int64ConstOp(b *testing.B, useSelectionVector bool, hasNulls bool) {
	ctx := context.Background()
	st := cluster.MakeTestingClusterSettings()
	evalCtx := tree.MakeTestingEvalContext(st)
	defer evalCtx.Stop(ctx)
	flowCtx := &execinfra.FlowCtx{
		EvalCtx: &evalCtx,
		Cfg: &execinfra.ServerConfig{
			Settings: st,
		},
	}
	typs := []*types.T{types.Int, types.Int}
	batch := testAllocator.NewMemBatch(typs)
	col := batch.ColVec(0).Int64()
	for i := 0; i < coldata.BatchSize(); i++ {
		col[i] = 1
	}
	if hasNulls {
		for i := 0; i < coldata.BatchSize(); i++ {
			if rand.Float64() < nullProbability {
				batch.ColVec(0).Nulls().SetNull(i)
			}
		}
	}
	batch.SetLength(coldata.BatchSize())
	if useSelectionVector {
		batch.SetSelection(true)
		sel := batch.Selection()
		for i := 0; i < coldata.BatchSize(); i++ {
			sel[i] = i
		}
	}
	source := colexecbase.NewRepeatableBatchSource(testAllocator, batch, typs)
	plusOp, err := createTestProjectingOperator(
		ctx, flowCtx, source, []*types.T{types.Int},
		"@1 + 1" /* projectingExpr */, false, /* canFallbackToRowexec */
	)
	require.NoError(b, err)
	plusOp.Init()

	b.SetBytes(int64(8 * coldata.BatchSize()))
	for i := 0; i < b.N; i++ {
		plusOp.Next(ctx)
	}
}

func BenchmarkProjPlusInt64Int64ConstOp(b *testing.B) {
	for _, useSel := range []bool{true, false} {
		for _, hasNulls := range []bool{true, false} {
			b.Run(fmt.Sprintf("useSel=%t,hasNulls=%t", useSel, hasNulls), func(b *testing.B) {
				benchmarkProjPlusInt64Int64ConstOp(b, useSel, hasNulls)
			})
		}
	}
}

func TestGetProjectionConstOperator(t *testing.T) {
	defer leaktest.AfterTest(t)()
	binOp := tree.Mult
	var input colexecbase.Operator
	colIdx := 3
	constVal := 31.37
	constArg := tree.NewDFloat(tree.DFloat(constVal))
	outputIdx := 5
	op, err := GetProjectionRConstOperator(
		testAllocator, types.Float, types.Float, types.Float,
		binOp, input, colIdx, constArg, outputIdx,
	)
	if err != nil {
		t.Error(err)
	}
	expected := &projMultFloat64Float64ConstOp{
		projConstOpBase: projConstOpBase{
			OneInputNode: NewOneInputNode(op.(*projMultFloat64Float64ConstOp).input),
			allocator:    testAllocator,
			colIdx:       colIdx,
			outputIdx:    outputIdx,
		},
		constArg: constVal,
	}
	if !reflect.DeepEqual(op, expected) {
		t.Errorf("got %+v,\nexpected %+v", op, expected)
	}
}

func TestGetProjectionConstMixedTypeOperator(t *testing.T) {
	defer leaktest.AfterTest(t)()
	binOp := tree.GE
	var input colexecbase.Operator
	colIdx := 3
	constVal := int16(31)
	constArg := tree.NewDInt(tree.DInt(constVal))
	outputIdx := 5
	op, err := GetProjectionRConstOperator(
		testAllocator, types.Int, types.Int2, types.Int,
		binOp, input, colIdx, constArg, outputIdx,
	)
	if err != nil {
		t.Error(err)
	}
	expected := &projGEInt64Int16ConstOp{
		projConstOpBase: projConstOpBase{
			OneInputNode: NewOneInputNode(op.(*projGEInt64Int16ConstOp).input),
			allocator:    testAllocator,
			colIdx:       colIdx,
			outputIdx:    outputIdx,
		},
		constArg: constVal,
	}
	if !reflect.DeepEqual(op, expected) {
		t.Errorf("got %+v,\nexpected %+v", op, expected)
	}
}

// TestRandomComparisons runs binary comparisons against all scalar types
// (supported by the vectorized engine) with random non-null data verifying
// that the result of Datum.Compare matches the result of the exec projection.
func TestRandomComparisons(t *testing.T) {
	defer leaktest.AfterTest(t)()
	ctx := context.Background()
	st := cluster.MakeTestingClusterSettings()
	evalCtx := tree.MakeTestingEvalContext(st)
	defer evalCtx.Stop(ctx)
	flowCtx := &execinfra.FlowCtx{
		EvalCtx: &evalCtx,
		Cfg: &execinfra.ServerConfig{
			Settings: st,
		},
	}
	const numTuples = 2048
	rng, _ := randutil.NewPseudoRand()

	expected := make([]bool, numTuples)
	var da sqlbase.DatumAlloc
	lDatums := make([]tree.Datum, numTuples)
	rDatums := make([]tree.Datum, numTuples)
	for _, typ := range types.Scalar {
		if typ.Family() == types.DateFamily {
			// TODO(jordan): #40354 tracks failure to compare infinite dates.
			continue
		}
		if !typeconv.IsTypeSupported(typ) {
			continue
		}
		typs := []*types.T{typ, typ, types.Bool}
		bytesFixedLength := 0
		if typ.Family() == types.UuidFamily {
			bytesFixedLength = 16
		}
		b := testAllocator.NewMemBatchWithSize(typs, numTuples)
		lVec := b.ColVec(0)
		rVec := b.ColVec(1)
		ret := b.ColVec(2)
		coldatatestutils.RandomVec(rng, bytesFixedLength, lVec, numTuples, 0)
		coldatatestutils.RandomVec(rng, bytesFixedLength, rVec, numTuples, 0)
		for i := range lDatums {
			lDatums[i] = PhysicalTypeColElemToDatum(lVec, i, da, typ)
			rDatums[i] = PhysicalTypeColElemToDatum(rVec, i, da, typ)
		}
		for _, cmpOp := range []tree.ComparisonOperator{tree.EQ, tree.NE, tree.LT, tree.LE, tree.GT, tree.GE} {
			for i := range lDatums {
				cmp := lDatums[i].Compare(&evalCtx, rDatums[i])
				var b bool
				switch cmpOp {
				case tree.EQ:
					b = cmp == 0
				case tree.NE:
					b = cmp != 0
				case tree.LT:
					b = cmp < 0
				case tree.LE:
					b = cmp <= 0
				case tree.GT:
					b = cmp > 0
				case tree.GE:
					b = cmp >= 0
				}
				expected[i] = b
			}
			input := newChunkingBatchSource(typs, []coldata.Vec{lVec, rVec, ret}, numTuples)
			op, err := createTestProjectingOperator(
				ctx, flowCtx, input, []*types.T{typ, typ},
				fmt.Sprintf("@1 %s @2", cmpOp), false, /* canFallbackToRowexec */
			)
			require.NoError(t, err)
			if err != nil {
				t.Fatal(err)
			}
			op.Init()
			var idx int
			for batch := op.Next(ctx); batch.Length() > 0; batch = op.Next(ctx) {
				for i := 0; i < batch.Length(); i++ {
					absIdx := idx + i
					assert.Equal(t, expected[absIdx], batch.ColVec(2).Bool()[i],
						"expected %s %s %s (%s[%d]) to be %t found %t", lDatums[absIdx], cmpOp, rDatums[absIdx], typ, absIdx,
						expected[absIdx], ret.Bool()[i])
				}
				idx += batch.Length()
			}
		}
	}
}

func TestGetProjectionOperator(t *testing.T) {
	defer leaktest.AfterTest(t)()
	typ := types.Int2
	binOp := tree.Mult
	var input colexecbase.Operator
	col1Idx := 5
	col2Idx := 7
	outputIdx := 9
	op, err := GetProjectionOperator(
		testAllocator, typ, typ, types.Int2,
		binOp, input, col1Idx, col2Idx, outputIdx,
	)
	if err != nil {
		t.Error(err)
	}
	expected := &projMultInt16Int16Op{
		projOpBase: projOpBase{
			OneInputNode: NewOneInputNode(op.(*projMultInt16Int16Op).input),
			allocator:    testAllocator,
			col1Idx:      col1Idx,
			col2Idx:      col2Idx,
			outputIdx:    outputIdx,
		},
	}
	if !reflect.DeepEqual(op, expected) {
		t.Errorf("got %+v,\nexpected %+v", op, expected)
	}
}

func benchmarkProjOp(
	b *testing.B,
	makeProjOp func(source *colexecbase.RepeatableBatchSource, intWidth int32) (colexecbase.Operator, error),
	useSelectionVector bool,
	hasNulls bool,
	intType *types.T,
) {
	ctx := context.Background()

	typs := []*types.T{intType, intType}
	batch := testAllocator.NewMemBatch(typs)
	switch intType.Width() {
	case 64:
		col1 := batch.ColVec(0).Int64()
		col2 := batch.ColVec(1).Int64()
		for i := 0; i < coldata.BatchSize(); i++ {
			col1[i] = 1
			col2[i] = 1
		}
	case 32:
		col1 := batch.ColVec(0).Int32()
		col2 := batch.ColVec(1).Int32()
		for i := 0; i < coldata.BatchSize(); i++ {
			col1[i] = 1
			col2[i] = 1
		}
	default:
		b.Fatalf("unsupported type: %s", intType)
	}
	if hasNulls {
		for i := 0; i < coldata.BatchSize(); i++ {
			if rand.Float64() < nullProbability {
				batch.ColVec(0).Nulls().SetNull(i)
			}
			if rand.Float64() < nullProbability {
				batch.ColVec(1).Nulls().SetNull(i)
			}
		}
	}
	batch.SetLength(coldata.BatchSize())
	if useSelectionVector {
		batch.SetSelection(true)
		sel := batch.Selection()
		for i := 0; i < coldata.BatchSize(); i++ {
			sel[i] = i
		}
	}
	source := colexecbase.NewRepeatableBatchSource(testAllocator, batch, typs)
	op, err := makeProjOp(source, intType.Width())
	require.NoError(b, err)
	op.Init()

	b.SetBytes(int64(8 * coldata.BatchSize() * 2))
	for i := 0; i < b.N; i++ {
		op.Next(ctx)
	}
}

func BenchmarkProjOp(b *testing.B) {
	ctx := context.Background()
	st := cluster.MakeTestingClusterSettings()
	evalCtx := tree.MakeTestingEvalContext(st)
	defer evalCtx.Stop(ctx)
	flowCtx := &execinfra.FlowCtx{
		EvalCtx: &evalCtx,
		Cfg: &execinfra.ServerConfig{
			Settings: st,
		},
	}
	getInputTypesForIntWidth := func(width int32) []*types.T {
		switch width {
		case 0, 64:
			return []*types.T{types.Int, types.Int}
		case 32:
			return []*types.T{types.Int4, types.Int4}
		default:
			b.Fatalf("unsupported int width: %d", width)
			return nil
		}
	}
	projOpMap := map[string]func(*colexecbase.RepeatableBatchSource, int32) (colexecbase.Operator, error){
		"projPlusIntIntOp": func(source *colexecbase.RepeatableBatchSource, width int32) (colexecbase.Operator, error) {
			return createTestProjectingOperator(
				ctx, flowCtx, source, getInputTypesForIntWidth(width),
				"@1 + @2" /* projectingExpr */, false, /* canFallbackToRowexec */
			)
		},
		"projMinusIntIntOp": func(source *colexecbase.RepeatableBatchSource, width int32) (colexecbase.Operator, error) {
			return createTestProjectingOperator(
				ctx, flowCtx, source, getInputTypesForIntWidth(width),
				"@1 - @2" /* projectingExpr */, false, /* canFallbackToRowexec */
			)
		},
		"projMultIntIntOp": func(source *colexecbase.RepeatableBatchSource, width int32) (colexecbase.Operator, error) {
			return createTestProjectingOperator(
				ctx, flowCtx, source, getInputTypesForIntWidth(width),
				"@1 * @2" /* projectingExpr */, false, /* canFallbackToRowexec */
			)
		},
		"projDivIntIntOp": func(source *colexecbase.RepeatableBatchSource, width int32) (colexecbase.Operator, error) {
			return createTestProjectingOperator(
				ctx, flowCtx, source, getInputTypesForIntWidth(width),
				"@1 / @2" /* projectingExpr */, false, /* canFallbackToRowexec */
			)
		},
	}

	for projOp, makeProjOp := range projOpMap {
		for _, intType := range []*types.T{types.Int, types.Int4} {
			for _, useSel := range []bool{true, false} {
				for _, hasNulls := range []bool{true, false} {
					b.Run(fmt.Sprintf("op=%s/type=%s/useSel=%t/hasNulls=%t",
						projOp, intType, useSel, hasNulls), func(b *testing.B) {
						benchmarkProjOp(b, makeProjOp, useSel, hasNulls, intType)
					})
				}
			}
		}
	}
}
