package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStructToMap(t *testing.T) {
	type TestStruct struct {
		A int    `tag:"a"`
		B string `tag:"b"`
	}
	v := TestStruct{1, "1"}
	mp, _ := StructToMap(v, "tag")
	assert.Equal(t, 1, mp["a"])
	assert.Equal(t, "1", mp["b"])

	mpp, _ := StructToMap(&v, "tag")
	assert.Equal(t, 1, mpp["a"])
	assert.Equal(t, "1", mpp["b"])

	_, err := StructToMap([]int{}, "tag")
	assert.NotNil(t, err)
}

func TestStructToArr(t *testing.T) {
	type TestStruct struct {
		A int    `tag:"a"`
		B string `tag:"b"`
	}
	v := TestStruct{1, "1"}
	mp, _ := StructToArr(v, "tag")
	assert.Equal(t, 1, mp[0])
	assert.Equal(t, "1", mp[1])

	mpp, _ := StructToArr(&v, "tag")
	assert.Equal(t, 1, mpp[0])
	assert.Equal(t, "1", mpp[1])

	_, err := StructToArr([]int{}, "tag")
	assert.NotNil(t, err)
}

func TestGetFieldByTag(t *testing.T) {
	type TestStruct struct {
		A int    `tag:"a"`
		B string `tag:"b"`
	}
	v := TestStruct{1, "1"}
	_, val, _ := GetFieldByTag(v, "tag", "a")
	assert.Equal(t, 1, int(val.Int()))

	_, val, _ = GetFieldByTag(&v, "tag", "b")
	assert.Equal(t, "1", val.String())

	_, _, err := GetFieldByTag([]int{}, "tag", "a")
	assert.NotNil(t, err)

	_, _, err = GetFieldByTag(v, "tag", "c")
	assert.NotNil(t, err)
}

func TestRefArray(t *testing.T) {
	arr := []int{1, 2, 3}
	res := RefArray(arr)
	assert.Equal(t, 1, *res[0])
	assert.Equal(t, 2, *res[1])
	assert.Equal(t, 3, *res[2])
}

func TestDerefArray(t *testing.T) {
	arr := []*int{new(int), new(int), new(int)}
	*arr[0] = 1
	*arr[1] = 2
	*arr[2] = 3
	res := DerefArray(arr)
	assert.Equal(t, 1, res[0])
	assert.Equal(t, 2, res[1])
	assert.Equal(t, 3, res[2])
}
