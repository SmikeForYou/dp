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
