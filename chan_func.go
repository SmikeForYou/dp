package dp

import (
	"sync"
	"time"

	"golang.org/x/exp/constraints"
)

// Chanalize converts an array of values to a channel with these values.
// Parameters:
// - values: The values to send to the channel.
// Returns:
// - A channel that will receive the values.
func Chanalize[T any](values ...T) <-chan T {
	ch := make(chan T)
	go func() {
		for _, iter := range values {
			ch <- iter
		}
		close(ch)
	}()
	return ch
}

// ChanalizeCb converts an array of values to a channel with these values and runs a callback function after all values are sent.
// Parameters:
// - callback: The function to call after all values are sent.
// - values: The values to send to the channel.
// Returns:
// - A channel that will receive the values.
func ChanalizeCb[T any](callback func(), values ...T) <-chan T {
	ch := make(chan T)
	go func() {
		for _, iter := range values {
			ch <- iter
		}
		close(ch)
		callback()
	}()
	return ch
}

// ReleaseChanel reads all values from a channel and returns an array of these values.
// Parameters:
// - chanel: The channel to read values from.
// Returns:
// - An array of values read from the channel.
func ReleaseChanel[T any](chanel <-chan T) []T {
	res := make([]T, 0)
	for i := range chanel {
		res = append(res, i)
	}
	return res
}

// Cycle will cyclically return values from an iterable to a channel one by one until the channel is closed.
// Parameters:
// - iterable: The values to cycle through.
// Returns:
// - A channel that will receive the values cyclically.
// - A function to stop the cycling.
func Cycle[T any](iterable ...T) (<-chan T, func()) {
	res := make(chan T)
	stop := make(chan byte, 5)
	go func() {
		for {
			for _, i := range iterable {
				select {
				case <-stop:
					close(res)
					close(stop)
					return
				default:
					res <- i
				}
			}
		}
	}()
	closefunc := func() {
		stop <- 1
	}
	return res, closefunc
}

// FilterChan filters channel values using a callback function.
// Parameters:
// - data: The input channel to filter.
// - callback: The function to use for filtering values.
// Returns:
// - A channel that will receive the filtered values.
func FilterChan[T any](data <-chan T, callback func(elem T) bool) <-chan T {
	res := make(chan T)
	go func() {
		for elem := range data {
			if callback(elem) {
				res <- elem
			}
		}
		close(res)
	}()
	return res
}

// MapChan applies a callback function to each value of a channel and returns a new channel with the results.
// Parameters:
// - data: The input channel to map.
// - callback: The function to apply to each value.
// Returns:
// - A channel that will receive the mapped values.
func MapChan[T, K any](data <-chan T, callback func(elem T) K) <-chan K {
	res := make(chan K)
	go func() {
		for elem := range data {
			res <- callback(elem)
		}
		close(res)
	}()
	return res
}

// ZipChan aggregates values from several channels into portions and sends them to a result channel.
// Parameters:
// - chanels: The input channels to aggregate.
// Returns:
// - A channel that will receive the aggregated values.
func ZipChan[T any](chanels ...<-chan T) <-chan []T {
	res := make(chan []T)
	go func() {
		for {
			item := make([]T, 0)
			var closed bool
			for _, chanel := range chanels {
				val, ok := <-chanel
				if !ok {
					closed = true
					break
				}
				item = append(item, val)
			}
			if closed {
				close(res)
				break
			}
			res <- item
		}
	}()
	return res
}

// Timer will send a value with a delay to a channel until the channel is closed.
// Parameters:
// - lap: The duration to wait between sending values.
// - val: The value to send.
// Returns:
// - A channel that will receive the values.
// - A function to stop the timer.
func Timer[T any](lap time.Duration, val T) (<-chan T, func()) {
	res := make(chan T, 0)
	stop := make(chan struct{})
	go func() {
		defer close(res)
		dataCh, closeCycle := Cycle(val)
		defer closeCycle()
		timer := time.NewTimer(lap)
		for {
			select {
			case <-timer.C:
				res <- <-dataCh
			case <-stop:
				timer.Stop()
				close(stop)
				return
			}
		}
	}()
	return res, func() {
		stop <- struct{}{}
	}
}

// AccumulateChan accumulates results of other binary functions which are mentioned in the function parameter.
// Parameters:
// - data: The input channel to accumulate.
// - fun: The function to use for accumulation.
// - initial: The initial value for accumulation.
// Returns:
// - A channel that will receive the accumulated values.
func AccumulateChan[T constraints.Ordered](data <-chan T, fun func(total, elem T) T, initial T) <-chan T {
	res := make(chan T)
	acc := initial
	go func() {
		for elem := range data {
			acc = fun(acc, elem)
			res <- acc
		}
		close(res)
	}()
	return res
}

// Enumerated is a wrapper struct for the EnumerateChan function.
type Enumerated[T any] struct {
	index int
	value T
}

// EnumerateChan wraps a provided channel and adds an index number to each channel value.
// Parameters:
// - ch: The input channel to enumerate.
// Returns:
// - A channel that will receive the enumerated values.
func EnumerateChan[T any](ch <-chan T) <-chan Enumerated[T] {
	counter := 0
	res := make(chan Enumerated[T], 0)
	go func() {
		for val := range ch {
			res <- Enumerated[T]{value: val, index: counter}
			counter++
		}
		close(res)
	}()
	return res
}

// FanIn aggregates several channels into one. Fan-in pattern implementation.
// Parameters:
// - channels: The input channels to aggregate.
// Returns:
// - A channel that will receive the aggregated values.
func FanIn[T any](channels ...<-chan T) <-chan T {
	res := make(chan T, 0)
	go func() {
		defer close(res)
		wg := sync.WaitGroup{}
		wg.Add(len(channels))
		for _, ch := range channels {
			go func(c <-chan T, w *sync.WaitGroup) {
				for v := range c {
					res <- v
				}
				w.Done()
			}(ch, &wg)
		}
		wg.Wait()
	}()
	return res
}

// FanOut broadcasts messages from a source channel to target channels.
// Parameters:
// - source: The source channel to broadcast from.
// - channels: The target channels to broadcast to.
func FanOut[T any](source <-chan T, channels ...chan T) {
	go func() {
		for sm := range source {
			for _, c := range channels {
				go func(ch chan T, m T) { ch <- m }(c, sm)
			}
		}
	}()
}
