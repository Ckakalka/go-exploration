package sorting

func InsertionSortInt(data []int32, ascending bool) {
    for i, curVal := range data {
        for ; (i > 0) && lessInt(ascending, curVal, data[i -1]); i-- {
            data[i] = data[i - 1]
        }
        data[i] = curVal
    }
}

func lessInt(ascending bool, first int32, second int32) bool {
    if ascending {
        return first < second
    } else {
        return first > second
    }
}

