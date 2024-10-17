dp - Go Documentation Server  window.initFuncs = \[\]; var goVersion = "go1.23.2";

...

[Go Documentation Server](http://localhost:6060/pkg/)

[GoDoc](http://localhost:6060/pkg/)

[▽](index.html#)

submit search

Package dp
==========

document.ANALYSIS\_DATA = null; document.CALLGRAPH = null;

`import "github.com/SmikeForYou/dp"`

[Overview](index.html#pkg-overview)

[Index](index.html#pkg-index)

Overview ▹
----------

Overview ▾
----------

Index ▹
-------

Index ▾
-------

[func AccumulateChan\[T constraints.Ordered\](data <-chan T, fun func(total, elem T) T, initial T) <-chan T](index.html#AccumulateChan)

[func Chanalize\[T any\](values ...T) <-chan T](index.html#Chanalize)

[func ChanalizeCb\[T any\](callback func(), values ...T) <-chan T](index.html#ChanalizeCb)

[func Chunk\[T any\](chunkSize int, iterable ...T) \[\]\[\]T](index.html#Chunk)

[func Compress\[T any\](data \[\]T, compressor \[\]bool) \[\]T](index.html#Compress)

[func Cycle\[T any\](iterable ...T) (<-chan T, func())](index.html#Cycle)

[func DerefArray\[T any\](arr \[\]\*T) \[\]T](index.html#DerefArray)

[func EnumerateChan\[T any\](ch <-chan T) <-chan Enumerated\[T\]](index.html#EnumerateChan)

[func FanIn\[T any\](channels ...<-chan T) <-chan T](index.html#FanIn)

[func FanOut\[T any\](source <-chan T, channels ...chan T)](index.html#FanOut)

[func Filter\[T any\](data \[\]T, callback func(elem T, index int) bool) \[\]T](index.html#Filter)

[func FilterChan\[T any\](data <-chan T, callback func(elem T) bool) <-chan T](index.html#FilterChan)

[func GetFieldByTag(in any, tag string, tagValue string) (reflect.StructField, reflect.Value, error)](index.html#GetFieldByTag)

[func In\[T comparable\](elem T, data \[\]T) int](index.html#In)

[func Map\[T, K any\](data \[\]T, callback func(elem T, index int) K) \[\]K](index.html#Map)

[func MapChan\[T, K any\](data <-chan T, callback func(elem T) K) <-chan K](index.html#MapChan)

[func Max\[T constraints.Ordered\](elems ...T) T](index.html#Max)

[func Min\[T constraints.Ordered\](elems ...T) T](index.html#Min)

[func Permutations\[T comparable\](data \[\]T) \[\]\[\]T](index.html#Permutations)

[func RefArray\[T any\](arr \[\]T) \[\]\*T](index.html#RefArray)

[func ReleaseChanel\[T any\](chanel <-chan T) \[\]T](index.html#ReleaseChanel)

[func Repeat\[T any\](value T, count int) \[\]T](index.html#Repeat)

[func SetTimeout(timeout time.Duration, fn func())](index.html#SetTimeout)

[func Sort\[T any\](data \[\]T, compare func(T, T) bool, reverse bool) \[\]T](index.html#Sort)

[func StructToArr(in any, tag string) (\[\]any, error)](index.html#StructToArr)

[func StructToMap(in any, tag string) (map\[string\]any, error)](index.html#StructToMap)

[func Sum\[T constraints.Ordered\](data ...T) T](index.html#Sum)

[func Timer\[T any\](lap time.Duration, val T) (<-chan T, func())](index.html#Timer)

[func Zip\[T any\](iterables ...\[\]T) \[\]\[\]T](index.html#Zip)

[func ZipChan\[T any\](chanels ...<-chan T) <-chan \[\]T](index.html#ZipChan)

[type Cache](index.html#Cache)

[func NewCache\[K comparable, V any\]() Cache\[K, V\]](index.html#NewCache)

[func NewCacheFromArr\[K comparable, V any\](data \[\]V, visitor func(elem V) (K, V)) Cache\[K, V\]](index.html#NewCacheFromArr)

[func (c Cache\[K, V\]) Exists(key K) bool](index.html#Cache.Exists)

[func (c Cache\[K, V\]) Get(key K) (V, bool)](index.html#Cache.Get)

[func (c Cache\[K, V\]) GetWithCallback(key K, onFail func(key K) (V, bool)) (V, bool)](index.html#Cache.GetWithCallback)

[func (c Cache\[K, V\]) Push(key K, val V)](index.html#Cache.Push)

[func (c Cache\[K, V\]) Remove(key K)](index.html#Cache.Remove)

[type Enumerated](index.html#Enumerated)

[type Grouper](index.html#Grouper)

[func GroupBy\[T any\](iterable \[\]T, keyExtractor func(elem T) any) \[\]Grouper\[T\]](index.html#GroupBy)

[type Iterator](index.html#Iterator)

[func NewIterator\[T any\](data \[\]T) Iterator\[T\]](index.html#NewIterator)

[func (receiver \*Iterator\[T\]) Elem() T](index.html#Iterator.Elem)

[func (receiver \*Iterator\[T\]) Items() (int, T)](index.html#Iterator.Items)

[func (receiver \*Iterator\[T\]) Next() bool](index.html#Iterator.Next)

[type TempCache](index.html#TempCache)

[func NewTempCache\[K comparable, V any\]() TempCache\[K, V\]](index.html#NewTempCache)

[func (tc TempCache\[K, V\]) Get(key K) (V, bool)](index.html#TempCache.Get)

[func (tc TempCache\[K, V\]) Push(key K, val V, ttl time.Duration)](index.html#TempCache.Push)

[func (tc TempCache\[K, V\]) Remove(key K)](index.html#TempCache.Remove)

### Package files

[caches.go](http://localhost:6060/src/github.com/SmikeForYou/dp/caches.go) [chan\_func.go](http://localhost:6060/src/github.com/SmikeForYou/dp/chan_func.go) [func.go](http://localhost:6060/src/github.com/SmikeForYou/dp/func.go) [struct\_func.go](http://localhost:6060/src/github.com/SmikeForYou/dp/struct_func.go) [utils.go](http://localhost:6060/src/github.com/SmikeForYou/dp/utils.go)

func [AccumulateChan](http://localhost:6060/src/github.com/SmikeForYou/dp/chan_func.go?s=4452:4556#L179) [¶](index.html#AccumulateChan)
---------------------------------------------------------------------------------------------------------------------------------------

func AccumulateChan\[T [constraints](http://localhost:6060/pkg/golang.org/x/exp/constraints/).[Ordered](http://localhost:6060/pkg/golang.org/x/exp/constraints/#Ordered)\](data <-chan T, fun func(total, elem T) T, initial T) <-chan T

AccumulateChan accumulates results of other binary functions which are mentioned in the function parameter. Parameters: - data: The input channel to accumulate. - fun: The function to use for accumulation. - initial: The initial value for accumulation. Returns: - A channel that will receive the accumulated values.

func [Chanalize](http://localhost:6060/src/github.com/SmikeForYou/dp/chan_func.go?s=266:309#L5) [¶](index.html#Chanalize)
-------------------------------------------------------------------------------------------------------------------------

func Chanalize\[T [any](http://localhost:6060/pkg/builtin/#any)\](values ...T) <-chan T

Chanalize converts an array of values to a channel with these values. Parameters: - values: The values to send to the channel. Returns: - A channel that will receive the values.

func [ChanalizeCb](http://localhost:6060/src/github.com/SmikeForYou/dp/chan_func.go?s=739:801#L22) [¶](index.html#ChanalizeCb)
------------------------------------------------------------------------------------------------------------------------------

func ChanalizeCb\[T [any](http://localhost:6060/pkg/builtin/#any)\](callback func(), values ...T) <-chan T

ChanalizeCb converts an array of values to a channel with these values and runs a callback function after all values are sent. Parameters: - callback: The function to call after all values are sent. - values: The values to send to the channel. Returns: - A channel that will receive the values.

func [Chunk](http://localhost:6060/src/github.com/SmikeForYou/dp/func.go?s=3966:4019#L160) [¶](index.html#Chunk)
----------------------------------------------------------------------------------------------------------------

func Chunk\[T [any](http://localhost:6060/pkg/builtin/#any)\](chunkSize [int](http://localhost:6060/pkg/builtin/#int), iterable ...T) \[\]\[\]T

Chunk splits a slice into chunks. Parameters: - chunkSize: The size of each chunk. - iterable: The slice to split into chunks. Returns: - A slice of chunks.

func [Compress](http://localhost:6060/src/github.com/SmikeForYou/dp/func.go?s=5132:5185#L204) [¶](index.html#Compress)
----------------------------------------------------------------------------------------------------------------------

func Compress\[T [any](http://localhost:6060/pkg/builtin/#any)\](data \[\]T, compressor \[\][bool](http://localhost:6060/pkg/builtin/#bool)) \[\]T

Compress gets two slices and adds values from the first slice to the result if the value in the second slice with the same index is true. If the slice lengths are not equal, the shorter one will be taken. Parameters: - data: The slice of elements to compress. - compressor: The slice of booleans to determine which elements to include. Returns: - A slice of compressed elements.

func [Cycle](http://localhost:6060/src/github.com/SmikeForYou/dp/chan_func.go?s=1547:1598#L53) [¶](index.html#Cycle)
--------------------------------------------------------------------------------------------------------------------

func Cycle\[T [any](http://localhost:6060/pkg/builtin/#any)\](iterable ...T) (<-chan T, func())

Cycle will cyclically return values from an iterable to a channel one by one until the channel is closed. Parameters: - iterable: The values to cycle through. Returns: - A channel that will receive the values cyclically. - A function to stop the cycling.

func [DerefArray](http://localhost:6060/src/github.com/SmikeForYou/dp/struct_func.go?s=3129:3165#L98) [¶](index.html#DerefArray)
--------------------------------------------------------------------------------------------------------------------------------

func DerefArray\[T [any](http://localhost:6060/pkg/builtin/#any)\](arr \[\]\*T) \[\]T

DerefArray converts an array of pointers to an array of values. Parameters: - arr: The input array of pointers. Returns: - An array of values.

func [EnumerateChan](http://localhost:6060/src/github.com/SmikeForYou/dp/chan_func.go?s=5046:5105#L203) [¶](index.html#EnumerateChan)
-------------------------------------------------------------------------------------------------------------------------------------

func EnumerateChan\[T [any](http://localhost:6060/pkg/builtin/#any)\](ch <-chan T) <-chan [Enumerated](index.html#Enumerated)\[T\]

EnumerateChan wraps a provided channel and adds an index number to each channel value. Parameters: - ch: The input channel to enumerate. Returns: - A channel that will receive the enumerated values.

func [FanIn](http://localhost:6060/src/github.com/SmikeForYou/dp/chan_func.go?s=5506:5554#L221) [¶](index.html#FanIn)
---------------------------------------------------------------------------------------------------------------------

func FanIn\[T [any](http://localhost:6060/pkg/builtin/#any)\](channels ...<-chan T) <-chan T

FanIn aggregates several channels into one. Fan-in pattern implementation. Parameters: - channels: The input channels to aggregate. Returns: - A channel that will receive the aggregated values.

func [FanOut](http://localhost:6060/src/github.com/SmikeForYou/dp/chan_func.go?s=6034:6089#L244) [¶](index.html#FanOut)
-----------------------------------------------------------------------------------------------------------------------

func FanOut\[T [any](http://localhost:6060/pkg/builtin/#any)\](source <-chan T, channels ...chan T)

FanOut broadcasts messages from a source channel to target channels. Parameters: - source: The source channel to broadcast from. - channels: The target channels to broadcast to.

func [Filter](http://localhost:6060/src/github.com/SmikeForYou/dp/func.go?s=922:993#L29) [¶](index.html#Filter)
---------------------------------------------------------------------------------------------------------------

func Filter\[T [any](http://localhost:6060/pkg/builtin/#any)\](data \[\]T, callback func(elem T, index [int](http://localhost:6060/pkg/builtin/#int)) [bool](http://localhost:6060/pkg/builtin/#bool)) \[\]T

Filter filters values of a slice using a callback function. Parameters: - data: The slice of elements to filter. - callback: A function to determine if an element should be included. Returns: - A slice of filtered elements.

func [FilterChan](http://localhost:6060/src/github.com/SmikeForYou/dp/chan_func.go?s=2131:2205#L82) [¶](index.html#FilterChan)
------------------------------------------------------------------------------------------------------------------------------

func FilterChan\[T [any](http://localhost:6060/pkg/builtin/#any)\](data <-chan T, callback func(elem T) [bool](http://localhost:6060/pkg/builtin/#bool)) <-chan T

FilterChan filters channel values using a callback function. Parameters: - data: The input channel to filter. - callback: The function to use for filtering values. Returns: - A channel that will receive the filtered values.

func [GetFieldByTag](http://localhost:6060/src/github.com/SmikeForYou/dp/struct_func.go?s=2296:2395#L69) [¶](index.html#GetFieldByTag)
--------------------------------------------------------------------------------------------------------------------------------------

func GetFieldByTag(in [any](http://localhost:6060/pkg/builtin/#any), tag [string](http://localhost:6060/pkg/builtin/#string), tagValue [string](http://localhost:6060/pkg/builtin/#string)) ([reflect](http://localhost:6060/pkg/reflect/).[StructField](http://localhost:6060/pkg/reflect/#StructField), [reflect](http://localhost:6060/pkg/reflect/).[Value](http://localhost:6060/pkg/reflect/#Value), [error](http://localhost:6060/pkg/builtin/#error))

GetFieldByTag returns the struct field and value of the field with the given tag. If tagValue is empty, the first field with the given tag is returned. Parameters: - in: The input struct to be searched. - tag: The tag used to filter the struct fields. - tagValue: The specific tag value to search for. Returns: - The struct field and its value. - An error if the input is not a struct or the tag is not found.

func [In](http://localhost:6060/src/github.com/SmikeForYou/dp/func.go?s=5621:5664#L222) [¶](index.html#In)
----------------------------------------------------------------------------------------------------------

func In\[T [comparable](http://localhost:6060/pkg/builtin/#comparable)\](elem T, data \[\]T) [int](http://localhost:6060/pkg/builtin/#int)

In returns the index of the first appearance of an element in a slice. Returns -1 if the element does not appear. Parameters: - elem: The element to search for. - data: The slice to search in. Returns: - The index of the element, or -1 if not found.

func [Map](http://localhost:6060/src/github.com/SmikeForYou/dp/func.go?s=1368:1436#L45) [¶](index.html#Map)
-----------------------------------------------------------------------------------------------------------

func Map\[T, K [any](http://localhost:6060/pkg/builtin/#any)\](data \[\]T, callback func(elem T, index [int](http://localhost:6060/pkg/builtin/#int)) K) \[\]K

Map applies a callback function to each element of a slice and returns a new slice. Parameters: - data: The slice of elements to map. - callback: A function to apply to each element. Returns: - A slice of mapped elements.

func [MapChan](http://localhost:6060/src/github.com/SmikeForYou/dp/chan_func.go?s=2628:2699#L101) [¶](index.html#MapChan)
-------------------------------------------------------------------------------------------------------------------------

func MapChan\[T, K [any](http://localhost:6060/pkg/builtin/#any)\](data <-chan T, callback func(elem T) K) <-chan K

MapChan applies a callback function to each value of a channel and returns a new channel with the results. Parameters: - data: The input channel to map. - callback: The function to apply to each value. Returns: - A channel that will receive the mapped values.

func [Max](http://localhost:6060/src/github.com/SmikeForYou/dp/func.go?s=2043:2088#L77) [¶](index.html#Max)
-----------------------------------------------------------------------------------------------------------

func Max\[T [constraints](http://localhost:6060/pkg/golang.org/x/exp/constraints/).[Ordered](http://localhost:6060/pkg/golang.org/x/exp/constraints/#Ordered)\](elems ...T) T

Max returns the maximum value of a slice. Parameters: - elems: The elements to find the maximum value of. Returns: - The maximum value.

func [Min](http://localhost:6060/src/github.com/SmikeForYou/dp/func.go?s=1705:1750#L58) [¶](index.html#Min)
-----------------------------------------------------------------------------------------------------------

func Min\[T [constraints](http://localhost:6060/pkg/golang.org/x/exp/constraints/).[Ordered](http://localhost:6060/pkg/golang.org/x/exp/constraints/#Ordered)\](elems ...T) T

Min returns the minimum value of a slice. Parameters: - elems: The elements to find the minimum value of. Returns: - The minimum value.

func [Permutations](http://localhost:6060/src/github.com/SmikeForYou/dp/func.go?s=6926:6973#L271) [¶](index.html#Permutations)
------------------------------------------------------------------------------------------------------------------------------

func Permutations\[T [comparable](http://localhost:6060/pkg/builtin/#comparable)\](data \[\]T) \[\]\[\]T

Permutations returns all permutations of a slice. Parameters: - data: The slice of elements to permute. Returns: - A slice of permutations.

func [RefArray](http://localhost:6060/src/github.com/SmikeForYou/dp/struct_func.go?s=3412:3446#L111) [¶](index.html#RefArray)
-----------------------------------------------------------------------------------------------------------------------------

func RefArray\[T [any](http://localhost:6060/pkg/builtin/#any)\](arr \[\]T) \[\]\*T

RefArray converts an array of values to an array of pointers. Parameters: - arr: The input array of values. Returns: - An array of pointers.

func [ReleaseChanel](http://localhost:6060/src/github.com/SmikeForYou/dp/chan_func.go?s=1138:1184#L39) [¶](index.html#ReleaseChanel)
------------------------------------------------------------------------------------------------------------------------------------

func ReleaseChanel\[T [any](http://localhost:6060/pkg/builtin/#any)\](chanel <-chan T) \[\]T

ReleaseChanel reads all values from a channel and returns an array of these values. Parameters: - chanel: The channel to read values from. Returns: - An array of values read from the channel.

func [Repeat](http://localhost:6060/src/github.com/SmikeForYou/dp/func.go?s=3656:3698#L146) [¶](index.html#Repeat)
------------------------------------------------------------------------------------------------------------------

func Repeat\[T [any](http://localhost:6060/pkg/builtin/#any)\](value T, count [int](http://localhost:6060/pkg/builtin/#int)) \[\]T

Repeat copies a value to a slice several times. Parameters: - value: The value to repeat. - count: The number of times to repeat the value. Returns: - A slice of repeated values.

func [SetTimeout](http://localhost:6060/src/github.com/SmikeForYou/dp/func.go?s=4630:4679#L190) [¶](index.html#SetTimeout)
--------------------------------------------------------------------------------------------------------------------------

func SetTimeout(timeout [time](http://localhost:6060/pkg/time/).[Duration](http://localhost:6060/pkg/time/#Duration), fn func())

SetTimeout executes a function after a timeout. Parameters: - timeout: The duration to wait before executing the function. - fn: The function to execute.

func [Sort](http://localhost:6060/src/github.com/SmikeForYou/dp/func.go?s=338:407#L5) [¶](index.html#Sort)
----------------------------------------------------------------------------------------------------------

func Sort\[T [any](http://localhost:6060/pkg/builtin/#any)\](data \[\]T, compare func(T, T) [bool](http://localhost:6060/pkg/builtin/#bool), reverse [bool](http://localhost:6060/pkg/builtin/#bool)) \[\]T

Sort performs a bubble sort on the provided data. Parameters: - data: The slice of elements to sort. - compare: A function to compare two elements. - reverse: A boolean indicating whether to sort in reverse order. Returns: - A sorted slice of elements.

func [StructToArr](http://localhost:6060/src/github.com/SmikeForYou/dp/struct_func.go?s=1343:1394#L39) [¶](index.html#StructToArr)
----------------------------------------------------------------------------------------------------------------------------------

func StructToArr(in [any](http://localhost:6060/pkg/builtin/#any), tag [string](http://localhost:6060/pkg/builtin/#string)) (\[\][any](http://localhost:6060/pkg/builtin/#any), [error](http://localhost:6060/pkg/builtin/#error))

StructToArr converts a struct to a slice of values of struct field. It uses tags on struct fields to decide which fields to add to the returned slice. Parameters: - in: The input struct to be converted. - tag: The tag used to filter the struct fields. Returns: - A slice of struct field values. - An error if the input is not a struct.

func [StructToMap](http://localhost:6060/src/github.com/SmikeForYou/dp/struct_func.go?s=429:489#L6) [¶](index.html#StructToMap)
-------------------------------------------------------------------------------------------------------------------------------

func StructToMap(in [any](http://localhost:6060/pkg/builtin/#any), tag [string](http://localhost:6060/pkg/builtin/#string)) (map\[[string](http://localhost:6060/pkg/builtin/#string)\][any](http://localhost:6060/pkg/builtin/#any), [error](http://localhost:6060/pkg/builtin/#error))

StructToMap converts a struct to a map using the struct's tags. It uses tags on struct fields to decide which fields to add to the returned map. Parameters: - in: The input struct to be converted. - tag: The tag used to filter the struct fields. Returns: - A map with keys as tag values and values as struct field values. - An error if the input is not a struct.

func [Sum](http://localhost:6060/src/github.com/SmikeForYou/dp/func.go?s=4351:4395#L178) [¶](index.html#Sum)
------------------------------------------------------------------------------------------------------------

func Sum\[T [constraints](http://localhost:6060/pkg/golang.org/x/exp/constraints/).[Ordered](http://localhost:6060/pkg/golang.org/x/exp/constraints/#Ordered)\](data ...T) T

Sum sums all values of a slice. Parameters: - data: The slice of elements to sum. Returns: - The sum of the elements.

func [Timer](http://localhost:6060/src/github.com/SmikeForYou/dp/chan_func.go?s=3697:3759#L148) [¶](index.html#Timer)
---------------------------------------------------------------------------------------------------------------------

func Timer\[T [any](http://localhost:6060/pkg/builtin/#any)\](lap [time](http://localhost:6060/pkg/time/).[Duration](http://localhost:6060/pkg/time/#Duration), val T) (<-chan T, func())

Timer will send a value with a delay to a channel until the channel is closed. Parameters: - lap: The duration to wait between sending values. - val: The value to send. Returns: - A channel that will receive the values. - A function to stop the timer.

func [Zip](http://localhost:6060/src/github.com/SmikeForYou/dp/func.go?s=2379:2418#L96) [¶](index.html#Zip)
-----------------------------------------------------------------------------------------------------------

func Zip\[T [any](http://localhost:6060/pkg/builtin/#any)\](iterables ...\[\]T) \[\]\[\]T

Zip aggregates values from several slices. Parameters: - iterables: The slices to aggregate. Returns: - A slice of aggregated values.

func [ZipChan](http://localhost:6060/src/github.com/SmikeForYou/dp/chan_func.go?s=3056:3107#L117) [¶](index.html#ZipChan)
-------------------------------------------------------------------------------------------------------------------------

func ZipChan\[T [any](http://localhost:6060/pkg/builtin/#any)\](chanels ...<-chan T) <-chan \[\]T

ZipChan aggregates values from several channels into portions and sends them to a result channel. Parameters: - chanels: The input channels to aggregate. Returns: - A channel that will receive the aggregated values.

type [Cache](http://localhost:6060/src/github.com/SmikeForYou/dp/caches.go?s=129:168#L1) [¶](index.html#Cache)
--------------------------------------------------------------------------------------------------------------

Cache is a generic map-based cache. K is the type of the keys, and V is the type of the values.

type Cache\[K [comparable](http://localhost:6060/pkg/builtin/#comparable), V [any](http://localhost:6060/pkg/builtin/#any)\] map\[K\]V

### func [NewCache](http://localhost:6060/src/github.com/SmikeForYou/dp/caches.go?s=209:257#L1) [¶](index.html#NewCache)

func NewCache\[K [comparable](http://localhost:6060/pkg/builtin/#comparable), V [any](http://localhost:6060/pkg/builtin/#any)\]() [Cache](index.html#Cache)\[K, V\]

NewCache creates a new empty Cache.

### func [NewCacheFromArr](http://localhost:6060/src/github.com/SmikeForYou/dp/caches.go?s=427:519#L6) [¶](index.html#NewCacheFromArr)

func NewCacheFromArr\[K [comparable](http://localhost:6060/pkg/builtin/#comparable), V [any](http://localhost:6060/pkg/builtin/#any)\](data \[\]V, visitor func(elem V) (K, V)) [Cache](index.html#Cache)\[K, V\]

NewCacheFromArr creates a new Cache from a slice of values. The visitor function is used to generate the key-value pairs for the cache.

### func (Cache\[K, V\]) [Exists](http://localhost:6060/src/github.com/SmikeForYou/dp/caches.go?s=770:809#L20) [¶](index.html#Cache.Exists)

func (c [Cache](index.html#Cache)\[K, V\]) Exists(key K) [bool](http://localhost:6060/pkg/builtin/#bool)

Exists checks if a key exists in the Cache.

### func (Cache\[K, V\]) [Get](http://localhost:6060/src/github.com/SmikeForYou/dp/caches.go?s=967:1008#L27) [¶](index.html#Cache.Get)

func (c [Cache](index.html#Cache)\[K, V\]) Get(key K) (V, [bool](http://localhost:6060/pkg/builtin/#bool))

Get retrieves a value from the Cache by key. It returns the value and a boolean indicating whether the key was found.

### func (Cache\[K, V\]) [GetWithCallback](http://localhost:6060/src/github.com/SmikeForYou/dp/caches.go?s=1183:1266#L34) [¶](index.html#Cache.GetWithCallback)

func (c [Cache](index.html#Cache)\[K, V\]) GetWithCallback(key K, onFail func(key K) (V, [bool](http://localhost:6060/pkg/builtin/#bool))) (V, [bool](http://localhost:6060/pkg/builtin/#bool))

GetWithCallback retrieves a value from the Cache by key. If the key is not found, the onFail callback is called to provide a value.

### func (Cache\[K, V\]) [Push](http://localhost:6060/src/github.com/SmikeForYou/dp/caches.go?s=664:703#L15) [¶](index.html#Cache.Push)

func (c [Cache](index.html#Cache)\[K, V\]) Push(key K, val V)

Push adds a key-value pair to the Cache.

### func (Cache\[K, V\]) [Remove](http://localhost:6060/src/github.com/SmikeForYou/dp/caches.go?s=1394:1428#L43) [¶](index.html#Cache.Remove)

func (c [Cache](index.html#Cache)\[K, V\]) Remove(key K)

Remove deletes a key-value pair from the Cache by key.

type [Enumerated](http://localhost:6060/src/github.com/SmikeForYou/dp/chan_func.go?s=4777:4830#L193) [¶](index.html#Enumerated)
-------------------------------------------------------------------------------------------------------------------------------

Enumerated is a wrapper struct for the EnumerateChan function.

type Enumerated\[T [any](http://localhost:6060/pkg/builtin/#any)\] struct {
// contains filtered or unexported fields
}

type [Grouper](http://localhost:6060/src/github.com/SmikeForYou/dp/func.go?s=2710:2762#L110) [¶](index.html#Grouper)
--------------------------------------------------------------------------------------------------------------------

Grouper is a grouping struct for GroupBy.

type Grouper\[T [any](http://localhost:6060/pkg/builtin/#any)\] struct {
Key   [any](http://localhost:6060/pkg/builtin/#any)
Group \[\]T
}

### func [GroupBy](http://localhost:6060/src/github.com/SmikeForYou/dp/func.go?s=3004:3081#L121) [¶](index.html#GroupBy)

func GroupBy\[T [any](http://localhost:6060/pkg/builtin/#any)\](iterable \[\]T, keyExtractor func(elem T) [any](http://localhost:6060/pkg/builtin/#any)) \[\][Grouper](index.html#Grouper)\[T\]

GroupBy returns consecutive keys and groups from the iterable. Parameters: - iterable: The slice of elements to group. - keyExtractor: A function to extract the key for each element. Returns: - A slice of Grouper structs.

type [Iterator](http://localhost:6060/src/github.com/SmikeForYou/dp/func.go?s=5814:5867#L232) [¶](index.html#Iterator)
----------------------------------------------------------------------------------------------------------------------

Iterator implements the iterator pattern without using unsafe code.

type Iterator\[T [any](http://localhost:6060/pkg/builtin/#any)\] struct {
// contains filtered or unexported fields
}

### func [NewIterator](http://localhost:6060/src/github.com/SmikeForYou/dp/func.go?s=6006:6051#L242) [¶](index.html#NewIterator)

func NewIterator\[T [any](http://localhost:6060/pkg/builtin/#any)\](data \[\]T) [Iterator](index.html#Iterator)\[T\]

NewIterator creates a new Iterator. Parameters: - data: The slice of elements to iterate over. Returns: - A new Iterator.

### func (\*Iterator\[T\]) [Elem](http://localhost:6060/src/github.com/SmikeForYou/dp/func.go?s=6460:6497#L256) [¶](index.html#Iterator.Elem)

func (receiver \*[Iterator](index.html#Iterator)\[T\]) Elem() T

Elem returns the current element of the iteration. Returns the current element.

### func (\*Iterator\[T\]) [Items](http://localhost:6060/src/github.com/SmikeForYou/dp/func.go?s=6666:6711#L262) [¶](index.html#Iterator.Items)

func (receiver \*[Iterator](index.html#Iterator)\[T\]) Items() ([int](http://localhost:6060/pkg/builtin/#int), T)

Items returns the current index of the iteration and the element at that index. Returns the current index and element.

### func (\*Iterator\[T\]) [Next](http://localhost:6060/src/github.com/SmikeForYou/dp/func.go?s=6264:6304#L249) [¶](index.html#Iterator.Next)

func (receiver \*[Iterator](index.html#Iterator)\[T\]) Next() [bool](http://localhost:6060/pkg/builtin/#bool)

Next advances the Iterator to the next element. Returns true if there are more elements to iterate over. Returns false if the end of the slice is reached.

type [TempCache](http://localhost:6060/src/github.com/SmikeForYou/dp/caches.go?s=1709:1785#L55) [¶](index.html#TempCache)
-------------------------------------------------------------------------------------------------------------------------

TempCache is a cache with time-based expiration for its values. K is the type of the keys, and V is the type of the values.

type TempCache\[K [comparable](http://localhost:6060/pkg/builtin/#comparable), V [any](http://localhost:6060/pkg/builtin/#any)\] struct {
// contains filtered or unexported fields
}

### func [NewTempCache](http://localhost:6060/src/github.com/SmikeForYou/dp/caches.go?s=1834:1890#L60) [¶](index.html#NewTempCache)

func NewTempCache\[K [comparable](http://localhost:6060/pkg/builtin/#comparable), V [any](http://localhost:6060/pkg/builtin/#any)\]() [TempCache](index.html#TempCache)\[K, V\]

NewTempCache creates a new empty TempCache.

### func (TempCache\[K, V\]) [Get](http://localhost:6060/src/github.com/SmikeForYou/dp/caches.go?s=2326:2372#L74) [¶](index.html#TempCache.Get)

func (tc [TempCache](index.html#TempCache)\[K, V\]) Get(key K) (V, [bool](http://localhost:6060/pkg/builtin/#bool))

Get retrieves a value from the TempCache by key. If the value has expired, it is removed from the cache and a zero value is returned.

### func (TempCache\[K, V\]) [Push](http://localhost:6060/src/github.com/SmikeForYou/dp/caches.go?s=2049:2112#L68) [¶](index.html#TempCache.Push)

func (tc [TempCache](index.html#TempCache)\[K, V\]) Push(key K, val V, ttl [time](http://localhost:6060/pkg/time/).[Duration](http://localhost:6060/pkg/time/#Duration))

Push adds a key-value pair to the TempCache with a time-to-live (TTL).

### func (TempCache\[K, V\]) [Remove](http://localhost:6060/src/github.com/SmikeForYou/dp/caches.go?s=2576:2615#L85) [¶](index.html#TempCache.Remove)

func (tc [TempCache](index.html#TempCache)\[K, V\]) Remove(key K)

Remove deletes a key-value pair from the TempCache by key.

Build version go1.23.2.  
Except as [noted](https://developers.google.com/site-policies#restrictions), the content of this page is licensed under the Creative Commons Attribution 3.0 License, and code is licensed under a [BSD license](http://localhost:6060/LICENSE).  
[Terms of Service](https://golang.org/doc/tos.html) | [Privacy Policy](https://www.google.com/intl/en/policies/privacy/)