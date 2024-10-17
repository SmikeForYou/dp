package dp

import (
	"golang.org/x/exp/constraints"
	"time"
)

// Sort performs a bubble sort on the provided data.
// Parameters:
// - data: The slice of elements to sort.
// - compare: A function to compare two elements.
// - reverse: A boolean indicating whether to sort in reverse order.
// Returns:
// - A sorted slice of elements.
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

// Filter filters values of a slice using a callback function.
// Parameters:
// - data: The slice of elements to filter.
// - callback: A function to determine if an element should be included.
// Returns:
// - A slice of filtered elements.
func Filter[T any](data []T, callback func(elem T, index int) bool) []T {
	res := make([]T, 0)
	for index, elem := range data {
		if callback(elem, index) {
			res = append(res, elem)
		}
	}
	return res
}

// Map applies a callback function to each element of a slice and returns a new slice.
// Parameters:
// - data: The slice of elements to map.
// - callback: A function to apply to each element.
// Returns:
// - A slice of mapped elements.
func Map[T, K any](data []T, callback func(elem T, index int) K) []K {
	res := make([]K, 0)
	for index, elem := range data {
		res = append(res, callback(elem, index))
	}
	return res
}

// Min returns the minimum value of a slice.
// Parameters:
// - elems: The elements to find the minimum value of.
// Returns:
// - The minimum value.
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

// Max returns the maximum value of a slice.
// Parameters:
// - elems: The elements to find the maximum value of.
// Returns:
// - The maximum value.
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

// Zip aggregates values from several slices.
// Parameters:
// - iterables: The slices to aggregate.
// Returns:
// - A slice of aggregated values.
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

// Grouper is a grouping struct for GroupBy.
type Grouper[T any] struct {
	Key   any
	Group []T
}

// GroupBy returns consecutive keys and groups from the iterable.
// Parameters:
// - iterable: The slice of elements to group.
// - keyExtractor: A function to extract the key for each element.
// Returns:
// - A slice of Grouper structs.
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

// Repeat copies a value to a slice several times.
// Parameters:
// - value: The value to repeat.
// - count: The number of times to repeat the value.
// Returns:
// - A slice of repeated values.
func Repeat[T any](value T, count int) []T {
	res := make([]T, count)
	for i := 0; i < count; i++ {
		res[i] = value
	}
	return res
}

// Chunk splits a slice into chunks.
// Parameters:
// - chunkSize: The size of each chunk.
// - iterable: The slice to split into chunks.
// Returns:
// - A slice of chunks.
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

// Sum sums all values of a slice.
// Parameters:
// - data: The slice of elements to sum.
// Returns:
// - The sum of the elements.
func Sum[T constraints.Ordered](data ...T) T {
	var res T
	for _, i := range data {
		res += i
	}
	return res
}

// SetTimeout executes a function after a timeout.
// Parameters:
// - timeout: The duration to wait before executing the function.
// - fn: The function to execute.
func SetTimeout(timeout time.Duration, fn func()) {
	go func() {
		time.Sleep(timeout)
		fn()
	}()
}

// Compress gets two slices and adds values from the first slice to the result if the value in the second slice with the same index is true.
// If the slice lengths are not equal, the shorter one will be taken.
// Parameters:
// - data: The slice of elements to compress.
// - compressor: The slice of booleans to determine which elements to include.
// Returns:
// - A slice of compressed elements.
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

// In returns the index of the first appearance of an element in a slice.
// Returns -1 if the element does not appear.
// Parameters:
// - elem: The element to search for.
// - data: The slice to search in.
// Returns:
// - The index of the element, or -1 if not found.
func In[T comparable](elem T, data []T) int {
	for i, t := range data {
		if t == elem {
			return i
		}
	}
	return -1
}

// Iterator implements the iterator pattern without using unsafe code.
type Iterator[T any] struct {
	data  []T
	index int
}

// NewIterator creates a new Iterator.
// Parameters:
// - data: The slice of elements to iterate over.
// Returns:
// - A new Iterator.
func NewIterator[T any](data []T) Iterator[T] {
	return Iterator[T]{data: data, index: -1}
}

// Next advances the Iterator to the next element.
// Returns true if there are more elements to iterate over.
// Returns false if the end of the slice is reached.
func (receiver *Iterator[T]) Next() bool {
	receiver.index++
	return (receiver.index) < len(receiver.data)
}

// Elem returns the current element of the iteration.
// Returns the current element.
func (receiver *Iterator[T]) Elem() T {
	return receiver.data[receiver.index]
}

// Items returns the current index of the iteration and the element at that index.
// Returns the current index and element.
func (receiver *Iterator[T]) Items() (int, T) {
	return receiver.index, receiver.data[receiver.index]
}

// Permutations returns all permutations of a slice.
// Parameters:
// - data: The slice of elements to permute.
// Returns:
// - A slice of permutations.
func Permutations[T comparable](data []T) [][]T {
	if len(data) == 2 {
		return [][]T{data, {data[1], data[0]}}
	} else {
		res := make([][]T, 0)
		for i, elem := range data {
			rest := make([]T, 0)
			rest = append(rest, data[:i]...)
			rest = append(rest, data[i+1:]...)
			for _, perm := range Permutations(rest) {
				res = append(res, append([]T{elem}, perm...))
			}
		}
		return res
	}
}
