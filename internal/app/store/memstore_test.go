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
	setURL := "http://qwe.asd/zxc"
	_ = sm.Put(setURL)
	_, getOk := sm.Get(wantShort)
	require.Equal(t, false, getOk)
}

func TestMemStore_Put_Get(t *testing.T) {
	sm := NewMemStore()
	setURL := "http://qwe.asd/zxc"
	getShort := sm.Put(setURL)
	require.NotEmpty(t, getShort)
	getURL, getOk := sm.Get(getShort)
	require.Equal(t, true, getOk)
	assert.Equal(t, setURL, getURL)
}
