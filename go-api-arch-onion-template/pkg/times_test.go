package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAdjustedReleaseDay(t *testing.T) {
	// non-leap to non-leap (old -> new)
	oldNonLeapYear := Str2time("2022-01-01")
	assert.Equal(t, 1, GetAdjustedReleaseDay(Str2time("2022-01-01"), oldNonLeapYear))
	assert.Equal(t, 60, GetAdjustedReleaseDay(Str2time("2022-03-01"), oldNonLeapYear))
	assert.Equal(t, 365, GetAdjustedReleaseDay(Str2time("2022-12-31"), oldNonLeapYear))
	// non-leap to leap (old -> new)
	assert.Equal(t, 1, GetAdjustedReleaseDay(Str2time("2024-01-01"), oldNonLeapYear))
	assert.Equal(t, 60, GetAdjustedReleaseDay(Str2time("2024-03-01"), oldNonLeapYear))
	assert.Equal(t, 365, GetAdjustedReleaseDay(Str2time("2024-12-31"), oldNonLeapYear))
	// leap to non-leap (old -> new)
	oldLeapYear := Str2time("2020-01-01")
	assert.Equal(t, 1, GetAdjustedReleaseDay(Str2time("2022-01-01"), oldLeapYear))
	assert.Equal(t, 60, GetAdjustedReleaseDay(Str2time("2022-03-01"), oldLeapYear))
	assert.Equal(t, 365, GetAdjustedReleaseDay(Str2time("2022-12-31"), oldLeapYear))
	// leap to leap (old -> new)
	assert.Equal(t, 1, GetAdjustedReleaseDay(Str2time("2024-01-01"), oldLeapYear))
	assert.Equal(t, 61, GetAdjustedReleaseDay(Str2time("2024-03-01"), oldLeapYear))
	assert.Equal(t, 366, GetAdjustedReleaseDay(Str2time("2024-12-31"), oldLeapYear))

	// non-leap to non-leap (new -> old)
	newNonLeapYear := Str2time("2026-01-01")
	assert.Equal(t, 1, GetAdjustedReleaseDay(Str2time("2022-01-01"), newNonLeapYear))
	assert.Equal(t, 60, GetAdjustedReleaseDay(Str2time("2022-03-01"), newNonLeapYear))
	assert.Equal(t, 365, GetAdjustedReleaseDay(Str2time("2022-12-31"), newNonLeapYear))
	// non-leap to leap (new -> old)
	assert.Equal(t, 1, GetAdjustedReleaseDay(Str2time("2024-01-01"), newNonLeapYear))
	assert.Equal(t, 60, GetAdjustedReleaseDay(Str2time("2024-03-01"), newNonLeapYear))
	assert.Equal(t, 365, GetAdjustedReleaseDay(Str2time("2024-12-31"), newNonLeapYear))
	// leap to non-leap (new -> old)
	newLeapYear := Str2time("2028-01-01")
	assert.Equal(t, 1, GetAdjustedReleaseDay(Str2time("2022-01-01"), newLeapYear))
	assert.Equal(t, 60, GetAdjustedReleaseDay(Str2time("2022-03-01"), newLeapYear))
	assert.Equal(t, 365, GetAdjustedReleaseDay(Str2time("2022-12-31"), newLeapYear))
	// leap to leap (new -> old)
	assert.Equal(t, 1, GetAdjustedReleaseDay(Str2time("2024-01-01"), newLeapYear))
	assert.Equal(t, 61, GetAdjustedReleaseDay(Str2time("2024-03-01"), newLeapYear))
	assert.Equal(t, 366, GetAdjustedReleaseDay(Str2time("2024-12-31"), newLeapYear))
}
