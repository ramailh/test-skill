package services

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ramailh/backend/fetch/props"
	"github.com/ramailh/backend/fetch/repository/redist"
	"github.com/ramailh/backend/fetch/util/aggregator"
	"github.com/ramailh/backend/fetch/util/helper"
	"github.com/ramailh/backend/fetch/util/request"
	"github.com/ramailh/backend/fetch/util/transformator"

	"github.com/thedevsaddam/gojsonq/v2"
)

func GetDataWithUSD() (interface{}, error) {
	resp, err := request.Get("https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list")
	if err != nil {
		return nil, err
	}
	datas, err := resp.ArrayOfMap()
	if err != nil {
		return nil, err
	}

	rdsClient := redist.NewRedisClient()

	year, month, date := time.Now().Date()
	key := fmt.Sprintf("USDtoID-%d%s%d", year, month.String(), date)

	var usdToIDR float64
	usdToIDR, err = rdsClient.GetFloat(key)
	if err != nil || usdToIDR == 0 {
		resp, err = request.Get("https://free.currconv.com/api/v7/convert?q=USD_IDR&compact=ultra&apiKey=" + props.ApiKey)
		if err != nil {
			return nil, err
		}
		dataConverter, err := resp.Map()
		if err != nil {
			return nil, err
		}

		usdToIDR = dataConverter["USD_IDR"].(float64)
		rdsClient.Set(key, usdToIDR, 24*time.Hour)
	}

	for index, data := range datas {
		price, ok := data["price"].(string)
		if !ok {
			continue
		}
		priceInt, _ := strconv.Atoi(price)
		datas[index]["price_in_usd"] = fmt.Sprintf("%.3f", float64(priceInt)/usdToIDR)
	}

	return datas, nil
}

func GetAggregateData() (interface{}, error) {
	resp, err := request.Get("https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list")
	if err != nil {
		return nil, err
	}
	respDatas, err := resp.ArrayOfInterface()
	if err != nil {
		return nil, err
	}

	listOfProvince := gojsonq.New().FromInterface(respDatas).Distinct("area_provinsi").Pluck("area_provinsi").([]interface{})
	aggregateByProvince := gojsonq.New().FromInterface(respDatas).GroupBy("area_provinsi").Get().(map[string][]interface{})

	var datas []map[string]interface{}
	for _, province := range listOfProvince {
		if province == nil {
			continue
		}

		innerData := make(map[string]interface{})

		dataProvince := gojsonq.New().FromInterface(aggregateByProvince[province.(string)])

		timestamps := dataProvince.Pluck("timestamp").([]interface{})
		intTimestamps := transformator.NewTransformator().FromStrings(timestamps).ToInts()
		minRange, maxRange := aggregator.GetTimeRange(intTimestamps)
		intervals := aggregator.DateHistogram(minRange, maxRange, aggregator.WeekInterval)
		dataTimeRange := dataProvince.Get()

		var datasIntervals []map[string]interface{}
		var from int

		for _, interval := range intervals {

			dataPerInterval := make(map[string]interface{})

			var totalPrices []interface{}
			var totalSizes []interface{}
			var mapAggregator = make(map[string]interface{})

			for _, innerData := range dataTimeRange.([]interface{}) {
				innerMap := innerData.(map[string]interface{})

				tmp, ok := innerMap["timestamp"].(string)
				if !ok {
					continue
				}

				timestamp, _ := strconv.Atoi(tmp)
				if helper.CountDigits(timestamp) == 10 {
					timestamp = timestamp * 1000
				}

				if timestamp <= interval && timestamp >= from {

					prices := dataProvince.Pluck("price").([]interface{})
					totalPrices = append(totalPrices, prices...)

					sizes := dataProvince.Pluck("size").([]interface{})
					totalSizes = append(totalSizes, sizes...)

				}
			}

			if len(totalPrices) != 0 && len(totalSizes) != 0 {
				mapAggregator["prices"] = aggregator.Metrics(totalPrices)
				mapAggregator["sizes"] = aggregator.Metrics(totalSizes)
			}

			dataPerInterval["interval"] = interval
			dataPerInterval["interval_as_string"] = time.Unix(0, int64(interval)*int64(1000000)).Format("2006-01-02")
			dataPerInterval["data"] = mapAggregator

			datasIntervals = append(datasIntervals, dataPerInterval)

			from = interval
		}

		innerData["province"] = province
		innerData["series"] = datasIntervals

		datas = append(datas, innerData)
	}

	return datas, nil
}
