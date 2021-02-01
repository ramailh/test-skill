package transformator

import "strconv"

type Type int

const (
	String Type = iota
	Int
	Interface
	Float
)

type (
	from interface {
		FromInterfaces(data []interface{}) to
		FromStrings(data []interface{}) to
	}

	to interface {
		ToInts() (datas []int)
	}
)

type transformator struct {
	from interface{}
	tipe Type
}

func NewTransformator() from {
	transform := &transformator{}
	return transform
}

func (trans *transformator) FromStrings(data []interface{}) to {
	trans.from = data
	trans.tipe = String

	return trans
}

func (trans *transformator) FromInterfaces(data []interface{}) to {
	trans.from = data
	trans.tipe = Interface

	return trans
}

func (trans *transformator) ToInts() (datas []int) {
	dataInterfaces := trans.from.([]interface{})

	for _, value := range dataInterfaces {
		if value == nil {
			continue
		}

		switch trans.tipe {
		case String:
			valueint, err := strconv.Atoi(value.(string))
			if err != nil {
				continue
			}
			datas = append(datas, valueint)
		case Interface:
			datas = append(datas, value.(int))
		default:
			datas = append(datas, value.(int))
		}
	}

	return
}

func IntervalsPerWeek(from, to, interval int) (intervals []int) {
	if CountDigits(from) == 10 {
		from = from * 1000
	}

	if CountDigits(to) == 10 {
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

func interfacesToInts(datas []interface{}) []int {
	var values []int
	for _, data := range datas {
		if data == nil {
			continue
		}

		dataInt, _ := strconv.Atoi(data.(string))
		values = append(values, dataInt)
	}
	return values
}

func CountDigits(number int) int {
	count := 0
	for number != 0 {
		number /= 10
		count += 1
	}
	return count

}
