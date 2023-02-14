package dp

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestChanalize(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	c := Chanalize(data...)
	res := make([]int, 0)
	for i := range c {
		res = append(res, i)
	}
	assert.Equal(t, data, res)
}

func TestReleaseChanel(t *testing.T) {
	c := make(chan int)
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	go func() {
		for _, i := range data {
			c <- i
		}
		close(c)
	}()
	released := ReleaseChanel(c)
	assert.Equal(t, data, released)
}

func TestCycle(t *testing.T) {
	data := []int{1, 2, 3}
	cycled, cl := Cycle(data...)
	counter := 0
	res := make([]int, 0)
	for i := range cycled {
		res = append(res, i)
		counter++
		if counter == 5 {
			cl()
		}
	}
	assert.Len(t, res, 6)
	assert.Equal(t, []int{1, 2, 3, 1, 2, 3}, res)
}

func TestFilterChan(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	c := Chanalize(data...)
	filtered := FilterChan(c, func(i int) bool {
		return i%2 == 0
	})
	released := ReleaseChanel(filtered)
	assert.Len(t, released, 4)
	assert.Equal(t, []int{2, 4, 6, 8}, released)
}

func TestMapChan(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	c := Chanalize(data...)
	mapped := MapChan(c, func(i int) float32 {
		return float32(i * 2.0)
	})
	assert.Equal(t, Sum(data...)*2, int(Sum(ReleaseChanel(mapped)...)))
}

func TestZipChan(t *testing.T) {
	ch1 := Chanalize(1, 2, 3, 4)
	ch2 := Chanalize(5, 6, 7)
	ch3 := Chanalize(8, 9)
	res := ZipChan(ch1, ch2, ch3)
	rel := ReleaseChanel(res)
	assert.Len(t, rel, 2)
	assert.Equal(t, []int{1, 5, 8}, rel[0])
	assert.Equal(t, []int{2, 6, 9}, rel[1])
}

func TestAccumulateChan(t *testing.T) {
	ch := Chanalize(1, 2, 3, 4, 5)
	res := make([]int, 0)
	for accumulated := range AccumulateChan(ch, func(total, elem int) int {
		total += elem
		return total
	}, 0) {
		res = append(res, accumulated)
	}
	assert.Equal(t, []int{1, 3, 6, 10, 15}, res)
}

func TestTimer(t *testing.T) {
	timeout := 2000
	timerChan, cls := Timer(time.Duration(timeout)*time.Millisecond, 1)
	startTime := time.Now()
	var delta uint64
	for range timerChan {
		delta = uint64(time.Now().Sub(startTime).Milliseconds())
		cls()
	}
	assert.GreaterOrEqual(t, uint64(delta), uint64(timeout))
}

func TestEnumerateChan(t *testing.T) {
	data := []int{0, 0, 0, 0, 0, 0}
	res := make([]Enumerated[int], 0)
	for v := range EnumerateChan(Chanalize(data...)) {
		res = append(res, v)
	}
	assert.Equal(t, Enumerated[int]{index: 3, value: 0}, res[3])
}

func TestFanIn(t *testing.T) {
	ch1 := Chanalize([]int{1, 2, 3}...)
	ch2 := Chanalize([]int{4, 5, 6}...)
	ch3 := Chanalize([]int{7, 8, 9}...)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, Sort(ReleaseChanel(FanIn(ch1, ch2, ch3)), func(a, b int) bool {
		return a < b
	}, false))

}

func TestFunOut(t *testing.T) {
	ch := make(chan int)
	t1 := make(chan int)
	t2 := make(chan int)
	r1, r2 := make([]int, 0), make([]int, 0)
	FanOut(ch, t1, t2)
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		for _, i := range []int{1, 2, 3} {
			ch <- i
		}
		close(ch)
	}()

	go func() {
		defer wg.Done()
		for m := range t1 {
			r1 = append(r1, m)
			if len(r1) == 3 {
				close(t1)
			}
		}

	}()

	go func() {
		defer wg.Done()
		for m := range t2 {
			r2 = append(r2, m)
			if len(r2) == 3 {
				close(t2)
			}
		}
	}()
	wg.Wait()
	r1 = Sort(r1, func(i, j int) bool {
		return j > i
	}, false)
	r2 = Sort(r2, func(i, j int) bool {
		return j > i
	}, false)
	assert.Equal(t, []int{1, 2, 3}, r1)
	assert.Equal(t, r1, r2)

}
