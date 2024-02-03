package gomaith

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_difference(t *testing.T) {
	for _, tt := range []struct {
		name string
		m1   map[int]struct{}
		m2   map[int]struct{}
		want map[int]struct{}
	}{
		{
			name: "both empty",
			m1:   M[int](),
			m2:   M[int](),
			want: M[int](),
		},
		{
			name: "m1 empty",
			m1:   M[int](),
			m2:   M(4, 5, 6),
			want: M(4, 5, 6),
		},
		{
			name: "m2 empty",
			m1:   M(1, 2, 3),
			m2:   M[int](),
			want: M(1, 2, 3),
		},
		{
			name: "disjoint",
			m1:   M(1, 2, 3),
			m2:   M(4, 5, 6),
			want: M(1, 2, 3, 4, 5, 6),
		},
		{
			name: "overlap",
			m1:   M(1, 2, 3),
			m2:   M(2, 3, 4),
			want: M(1, 4),
		},
		{
			name: "no difference",
			m1:   M(1, 2, 3),
			m2:   M(1, 2, 3),
			want: M[int](),
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := difference(tt.m1, tt.m2)
			assert.Equalf(t, tt.want, got, "difference(%v, %v)", tt.want, got)
		})
	}
}

func Test_lastIndex(t *testing.T) {
	t.Parallel()
	testFn := func(i int) func(int) bool {
		return func(j int) bool {
			return i == j
		}
	}
	t.Run("Empty", func(t *testing.T) {
		t.Parallel()
		got := lastIndex(S[int](), testFn(1))
		assert.Equal(t, -1, got)
	})
	t.Run("NotFound", func(t *testing.T) {
		t.Parallel()
		got := lastIndex(S(1, 1, 2, 2, 3, 3), testFn(4))
		assert.Equal(t, -1, got)
	})
	t.Run("FoundAtBeginning", func(t *testing.T) {
		t.Parallel()
		got := lastIndex(S(1, 1, 2, 2, 3, 3), testFn(1))
		assert.Equal(t, 1, got)
	})
	t.Run("FoundInMiddle", func(t *testing.T) {
		t.Parallel()
		got := lastIndex(S(1, 1, 2, 2, 3, 3), testFn(2))
		assert.Equal(t, 3, got)
	})
	t.Run("FoundAtEnd", func(t *testing.T) {
		t.Parallel()
		got := lastIndex(S(1, 1, 2, 2, 3, 3), testFn(3))
		assert.Equal(t, 5, got)
	})
}
