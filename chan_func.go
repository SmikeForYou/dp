package dp

import (
	"sync"
	"time"

	"golang.org/x/exp/constraints"
)

//Chanalize converts array of values to channel with this values
func Chanalize[T any](iterable ...T) <-chan T {
	ch := make(chan T)
	go func() {
		for _, iter := range iterable {
			ch <- iter
		}
		close(ch)
	}()
	return ch
}

//ReleaseChanel read all values from chanel and returns array of this values
func ReleaseChanel[T any](chanel <-chan T) []T {
	res := make([]T, 0)
	for i := range chanel {
		res = append(res, i)
	}
	return res
}

//Cycle will cyclical returns values from iterable to channel one by one until chanel not closed
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

//FilterChan filters channel values using callback function
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

//MapChan applies callback function to each value of chan and return to to new chanel
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

//ZipChan aggregates values from several chanels to portions and send them to result chanel
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

//Timer will send some value with some delay to chan until chanel not closed
func Timer[T any](lap time.Duration, val T) (<-chan T, func()) {
	res := make(chan T, 0)
	stop := make(chan byte)
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
		stop <- 1
	}
}

//AccumulateChan accumulated results of other binary functions which is mentioned in func-parameter
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

//Enumerated is wrapper struct for EnumerateChan func
type Enumerated[T any] struct {
	index int
	value T
}

//EnumerateChan wrapper for provided chanel that will add index number to each chanel value
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

//FanIn aggregate several chanels to one. Fan-in pattern implementation
func FanIn[T any](chanels ...<-chan T) <-chan T {
	res := make(chan T, 0)
	go func() {
		defer close(res)
		wg := sync.WaitGroup{}
		wg.Add(len(chanels))
		for _, ch := range chanels {
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