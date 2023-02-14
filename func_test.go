package dp

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	a string
	b int
	c float64
}

func TestFilter(t *testing.T) {
	unfiltered := []testStruct{
		{"A", 1, 1.0},
		{"B", 2, 2.0},
		{"C", 3, 3.0},
		{"D", 4, 4.0},
		{"E", 5, 5.0},
		{"A", 6, 6.0},
	}
	filtered := Filter(unfiltered, func(elem testStruct, _ int) bool {
		return elem.a == "A"
	})
	assert.Len(t, filtered, 2)
	assert.Equal(t, 1, filtered[0].b)
}

func BenchmarkFilter(b *testing.B) {
	unfiltered := Repeat(testStruct{"A", 1, 1.0}, b.N)
	for i := 0; i < b.N; i++ {
		filtered := Filter(unfiltered, func(elem testStruct, _ int) bool {
			return elem.a == "A"
		})
		assert.Len(b, filtered, b.N)
	}
}

func TestMap(t *testing.T) {
	raw := []testStruct{
		{"A", 1, 1.0},
		{"B", 2, 2.0},
		{"C", 3, 3.0},
		{"D", 4, 4.0},
		{"E", 5, 5.0},
		{"A", 6, 6.0},
	}
	type newStruct struct {
		d string
		e int
		f float64
	}
	mp := Map(raw, func(elem testStruct, index int) newStruct {
		return newStruct{
			d: elem.a,
			e: elem.b * 2,
			f: elem.c,
		}
	})
	assert.Len(t, mp, 6)
	assert.Equal(t, 12, mp[5].e)
}

func TestMin(t *testing.T) {
	data := []int{0, 1, 2, 3, 4, 5, 6}
	assert.Equal(t, 0, Min(data...))
}

func TestMax(t *testing.T) {
	data := []int{0, 1, 2, 3, 4, 5, 7, 6}
	assert.Equal(t, 7, Max(data...))
}

func TestZip(t *testing.T) {
	initial := [][]int{
		{1, 2, 3, 5, 6},
		{4, 5, 6, 7},
		{7, 8, 9},
	}
	zipped := Zip(initial...)
	assert.Len(t, zipped, 3)
	assert.Len(t, zipped[0], 3)
	assert.Len(t, zipped[1], 3)
	assert.Len(t, zipped[2], 3)
	assert.Equal(t, zipped[0], []int{1, 4, 7})
}

func TestGroupBy(t *testing.T) {
	ungrouped := []testStruct{
		{"A", 1, 1.0},
		{"B", 2, 2.0},
		{"A", 6, 6.0},
	}
	grouped := GroupBy(ungrouped, func(elem testStruct) any {
		return elem.a
	})
	assert.Len(t, grouped, 2)
	assert.Equal(t, "A", grouped[0].Key)
	assert.Equal(t, "B", grouped[1].Key)
	assert.Len(t, grouped[0].Group, 2)
	assert.Equal(t, 6, grouped[0].Group[1].b)
}

func TestChunk(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6}
	chunked := Chunk(3, data...)
	assert.Len(t, chunked, 2)
	assert.Equal(t, []int{1, 2, 3}, chunked[0])
	assert.Equal(t, []int{4, 5, 6}, chunked[1])
}

func TestSort(t *testing.T) {
	data := []int{4, 5, 1, 2, 3, 6}
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, Sort(data, func(curr, prev int) bool {
		return curr < prev
	}, false))
	assert.Equal(t, []int{6, 5, 4, 3, 2, 1}, Sort(data, func(curr, prev int) bool {
		return curr < prev
	}, true))
}

func TestRepeat(t *testing.T) {
	assert.Equal(t, []int{1, 1, 1, 1, 1}, Repeat(1, 5))
}

func TestSetTimeout(t *testing.T) {
	start := time.Now()
	var wg sync.WaitGroup
	wg.Add(1)
	fn := func() {
		wg.Done()
	}
	SetTimeout(time.Duration(1000)*time.Millisecond, fn)
	wg.Wait()
	dur := time.Now().Sub(start).Milliseconds()
	assert.GreaterOrEqual(t, uint64(dur), uint64(1000))

}

func TestCompress(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6}
	compresor := []bool{true, true, false, false, true}
	res := Compress(data, compresor)
	assert.Equal(t, []int{1, 2, 5}, res)
}

func TestIn(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6}
	assert.Equal(t, 0, In(1, data))
	assert.Equal(t, -1, In(10, data))
}

func TestIterator(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6}
	res := make([]int, 0)
	iterator := NewIterator(data)
	for iterator.Next() {
		res = append(res, iterator.Elem())
	}
	assert.Equal(t, data, res)
}

func TestPermutations(t *testing.T) {
	data := []int{1, 2, 3}
	res := Permutations(data)
	assert.Len(t, res, 6)
	assert.Equal(t, []int{1, 2, 3}, res[0])
	assert.Equal(t, []int{1, 3, 2}, res[1])
	assert.Equal(t, []int{2, 1, 3}, res[2])
	assert.Equal(t, []int{2, 3, 1}, res[3])
	assert.Equal(t, []int{3, 1, 2}, res[4])
	assert.Equal(t, []int{3, 2, 1}, res[5])
	datastr := []string{"a", "b", "c"}
	resstr := Permutations(datastr)
	assert.Len(t, resstr, 6)
	assert.Equal(t, []string{"a", "b", "c"}, resstr[0])
	assert.Equal(t, []string{"a", "c", "b"}, resstr[1])
	assert.Equal(t, []string{"b", "a", "c"}, resstr[2])
	assert.Equal(t, []string{"b", "c", "a"}, resstr[3])
	assert.Equal(t, []string{"c", "a", "b"}, resstr[4])
	assert.Equal(t, []string{"c", "b", "a"}, resstr[5])

}
