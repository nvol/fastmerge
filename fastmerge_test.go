package fastmerge

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

const N = 20000000

type testDataModel struct {
	given    []int
	expected []int
}

func slicesAreEqual(v1, v2 []int) bool {
	if len(v1) != len(v2) {
		return false
	}
	for i := range v1 {
		if v1[i] != v2[i] {
			return false
		}
	}
	return true
}
func TestMergeSortIntSlice(t *testing.T) {

	var testData testDataModel
	testData.given = make([]int, N)
	testData.expected = make([]int, N)
	rand.Seed(time.Now().UnixNano())
	for i := range testData.given {
		val := rand.Int()
		testData.given[i] = val
		testData.expected[i] = val
	}
	ts0 := time.Now().UnixNano()
	sort.Ints(testData.expected)
	dt0 := time.Now().UnixNano() - ts0
	t.Logf("quick-sort time: %v ns\n", dt0)

	ts1 := time.Now().UnixNano()
	MergeSortIntSlice(testData.given)
	dt1 := time.Now().UnixNano() - ts1
	t.Logf("merge-sort time: %v ns\n", dt1)

	if !slicesAreEqual(testData.given, testData.expected) {
		t.Errorf("Oh, no!\n%#v\n%#v\n", testData.given, testData.expected)
	}

	// other cases
	testDataSet := []testDataModel{
		testDataModel{
			given:    []int{2, 6, 4},
			expected: []int{2, 4, 6},
		},
		testDataModel{
			given:    []int{},
			expected: []int{},
		},
		testDataModel{
			given:    []int{42},
			expected: []int{42},
		},
		testDataModel{
			given:    []int{-1, 0, 1, 2},
			expected: []int{-1, 0, 1, 2},
		},
		testDataModel{
			given:    []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for ix := range testDataSet {
		MergeSortIntSlice(testDataSet[ix].given)
		if !slicesAreEqual(testDataSet[ix].given, testDataSet[ix].expected) {
			t.Errorf("expected: %v\n     got:%v\n",
				testDataSet[ix].expected,
				testDataSet[ix].given,
			)
		}
	}
}
