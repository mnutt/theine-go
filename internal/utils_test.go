package internal

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

type Foo struct {
	Bar string
}

func TestStringKey(t *testing.T) {
	hasher := NewHasher[string](nil)
	h, _ := hasher.hash(strconv.Itoa(123456))
	for i := 0; i < 10; i++ {
		ih, _ := hasher.hash(strconv.Itoa(123456))
		require.Equal(t, h, ih)
	}
}

func TestStructStringKey(t *testing.T) {
	hasher1 := NewHasher[Foo](nil)
	hasher2 := NewHasher[Foo](func(k Foo) string {
		return k.Bar
	})
	h1 := uint64(0)
	h2 := uint64(0)
	for i := 0; i < 10; i++ {
		foo := Foo{Bar: strconv.Itoa(123456)}
		if h1 == 0 {
			h1, _ = hasher1.hash(foo)
		} else {
			h1i, _ := hasher1.hash(foo)
			require.NotEqual(t, h1, h1i)
		}
	}
	for i := 0; i < 10; i++ {
		foo := Foo{Bar: strconv.Itoa(123456)}
		if h2 == 0 {
			h2, _ = hasher2.hash(foo)
		} else {
			h2i, _ := hasher2.hash(foo)
			require.Equal(t, h2, h2i)
		}
	}
}
