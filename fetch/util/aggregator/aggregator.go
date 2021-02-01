package aggregator

import (
	"github.com/ramailh/backend/fetch/util/helper"
	"github.com/ramailh/backend/fetch/util/transformator"
)

const (
	WeekInterval = 604800000
)

func Metrics(metrics []interface{}) map[string]interface{} {
	values := transformator.NewTransformator().FromStrings(metrics).ToInts()
	minValue, maxValue, sumValue := MinMaxSum(values)

	dataValues := make(map[string]interface{})
	dataValues["max"] = maxValue
	dataValues["min"] = minValue
	dataValues["avg"] = sumValue / len(values)
	dataValues["median"] = ((maxValue - minValue) / 2) + minValue

	return dataValues
}

func DateHistogram(from, to, interval int) (intervals []int) {
	if helper.CountDigits(from) == 10 {
		from = from * 1000
	}

	if helper.CountDigits(to) == 10 {
		to = to * 1000
	}

	for {
		intervals = append(intervals, from+interval)
		from += interval
		if from > to {
			break
		}
	}

	return
}

func GetTimeRange(data []int) (from, to int) {
	for _, value := range data {
		if helper.CountDigits(value) == 10 {
			value = value * 1000
		}

		if to < value {
			to = value
		}

		if from >= value || from == 0 {
			from = value
		}
	}
	return
}

func MinMaxSum(data []int) (min, max, sum int) {
	for _, value := range data {
		if max < value {
			max = value
		}

		if min >= value || min == 0 {
			min = value
		}

		sum += value
	}

	return
}
