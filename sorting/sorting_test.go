package sorting

import (
	"testing"
)

var dataToSort = [...][]int32 {
    {5, 2, 3, 3, 1, -10},
    {5},
    {},
}

var expectedResults = [...][]int32 {
    {-10, 1, 2, 3, 3, 5},
    {5},
    {},
}

func TestInsertionSortInt(t *testing.T) {
    for i, d := range dataToSort {
        data := make([]int32, len(d))
        copy(data, d)
        InsertionSortInt(data, true)
        for j, v := range data {
            if v != expectedResults[i][j] {
                t.Errorf("Test set %d. Expected value %d but value %d received\n",
                        i + 1, expectedResults[i][j], v)
            }
        }
    }
    for i, d := range dataToSort {
        data := make([]int32, len(d))
        copy(data, d)
        InsertionSortInt(data, false)
        for j, v := range data {
            if v != expectedResults[i][len(expectedResults[i]) - (j + 1)] {
                t.Errorf("Test set %d. Expected value %d but value %d received\n",
                        i + 1, expectedResults[i][j], v)
            }
        }
    }


}


