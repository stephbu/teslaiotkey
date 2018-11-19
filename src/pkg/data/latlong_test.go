package data

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLatLongFromString(t *testing.T) {
	res, err := LatLongFromString("1.00,1.00")
	assert.NotNil(t, res)
	assert.Equal(t, res.Lat, res.Long)
	assert.Nil(t, err)
}

func TestLatLongFromStringFail1(t *testing.T) {
	res, err := LatLongFromString(",1.00")
	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestLatLongFromStringFail2(t *testing.T) {
	res, err := LatLongFromString("a,1.00")
	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestLatLongFromStringFail3(t *testing.T) {
	res, err := LatLongFromString("1.00,")
	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestLatLongFromStringFail4(t *testing.T) {
	res, err := LatLongFromString("1.00,a")
	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestLatLongFromStringFail5(t *testing.T) {
	res, err := LatLongFromString(",")
	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestLatLongFromStringFail6(t *testing.T) {
	res, err := LatLongFromString("a,b")
	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestLatLongFromStringFail7(t *testing.T) {
	res, err := LatLongFromString("1.00")
	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestLatLongFromStringFail8(t *testing.T) {
	res, err := LatLongFromString("a")
	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestLatLongFromStringFail9(t *testing.T) {
	res, err := LatLongFromString("1.00,1.00,1.00")
	assert.Nil(t, res)
	assert.Error(t, err)
}
