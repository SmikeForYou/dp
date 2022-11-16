package dp

import (
	"time"

	"golang.org/x/exp/constraints"
)

//Sort sorts array
func Sort[T any](data []T, compare func(T, T) bool, reverse bool) []T {
	res := make([]T, len(data))
	copy(res, data)
	for {
		var sorted = true
		for i := 0; i < len(res)-1; i++ {
			if !(compare(res[i], res[i+1]) != reverse) {
				sorted = false
				res[i], res[i+1] = res[i+1], res[i]
			}
		}
		if sorted {
			break
		}
	}
	return res
}

//Filter filters values of array ussing callback func
func Filter[T any](data []T, callback func(elem T, index int) bool) []T {
	res := make([]T, 0)
	for index, elem := range data {
		if callback(elem, index) {
			res = append(res, elem)
		}
	}
	return res
}

//Map applies callback function to each element of array and returns new array
func Map[T, K any](data []T, callback func(elem T, index int) K) []K {
	res := make([]K, 0)
	for index, elem := range data {
		res = append(res, callback(elem, index))
	}
	return res
}

//Min will return minimum value of array
func Min[T constraints.Ordered](elems ...T) T {
	var min T
	for i, elem := range elems {
		if i == 0 {
			min = elem
			continue
		}
		if elem < min {
			min = elem
		}
	}
	return min
}

//Max will return maximum value of array
func Max[T constraints.Ordered](elems ...T) T {
	var max T
	for i, elem := range elems {
		if i == 0 {
			max = elem
			continue
		}
		if elem > max {
			max = elem
		}
	}
	return max
}

//Zip aggregates values from several arrays
func Zip[T any](iterables ...[]T) [][]T {
	minLength := Min(Map(iterables, func(elem []T, index int) int {
		return len(elem)
	})...)
	res := make([][]T, minLength)
	for i := 0; i < minLength; i++ {
		for _, j := range iterables {
			res[i] = append(res[i], j[i])
		}
	}
	return res
}

//Grouper is goruping stuct for GroupBy
type Grouper[T any] struct {
	Key   any
	Group []T
}

//GroupBy returns consecutive keys and groups from the iterable
func GroupBy[T any](iterable []T, keyExtractor func(elem T) any) []Grouper[T] {
	res := make([]Grouper[T], 0)
	for _, elem := range iterable {
		keyExists := false
		keyExtracted := keyExtractor(elem)
		for i, gr := range res {
			if gr.Key == keyExtracted {
				keyExists = true
				res[i].Group = append(res[i].Group, elem)
				break
			}
		}
		if !keyExists {
			res = append(res, Grouper[T]{Key: keyExtracted, Group: []T{elem}})
		}
	}
	return res
}

//Repeat copy values to array several times
func Repeat[T any](value T, count int) []T {
	res := make([]T, count)
	for i := 0; i < count; i++ {
		res[i] = value
	}
	return res
}

//Chunk splits array to chunks
func Chunk[T any](chunkSize int, iterable ...T) [][]T {
	res := make([][]T, 0)
	buf := make([]T, 0)
	for _, i := range iterable {
		buf = append(buf, i)
		if len(buf) == chunkSize {
			res = append(res, buf)
			buf = make([]T, 0)
		}
	}
	return res
}

//Sum summs all value of array
func Sum[T constraints.Ordered](data ...T) T {
	var res T
	for _, i := range data {
		res += i
	}
	return res
}

//SetTimeout will execute fn after some timeout
func SetTimeout(timeout time.Duration, fn func()) {
	go func() {
		time.Sleep(timeout)
		fn()
	}()
}

//Compress get 2 arrays and add values from 1st array to result if value in 2nd array with same index is true
// If array length not equal shorter will be taken
func Compress[T any](data []T, compressor []bool) []T {
	min := Min(len(data), len(compressor))
	res := make([]T, 0)
	for i := 0; i < min; i++ {
		if compressor[i] {
			res = append(res, data[i])
		}
	}
	return res

}

//In will return index of first appearance of elem in data.
//Returns -1 if elem does not appear
func In[T comparable](elem T, data []T) int {
	for i, t := range data {
		if t == elem {
			return i
		}
	}
	return -1
}