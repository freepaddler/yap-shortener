package store

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"testing"
)

func TestMemStore_Get_Empty(t *testing.T) {
	sm := NewMemStore()
	wantShort := "XXXXXXXX"
	_, getOk := sm.Get(wantShort)
	require.Equal(t, false, getOk)
}

func TestMemStore_Get_NotFound(t *testing.T) {
	sm := NewMemStore()
	wantShort := "XXXXXXXX"
	setUrl := "http://qwe.asd/zxc"
	_ = sm.Put(setUrl)
	_, getOk := sm.Get(wantShort)
	require.Equal(t, false, getOk)
}

func TestMemStore_Put_Get(t *testing.T) {
	sm := NewMemStore()
	setUrl := "http://qwe.asd/zxc"
	getShort := sm.Put(setUrl)
	require.NotEmpty(t, getShort)
	getUrl, getOk := sm.Get(getShort)
	require.Equal(t, true, getOk)
	assert.Equal(t, setUrl, getUrl)
}
