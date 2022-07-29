package manageincoming

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func (asset *TotalAssets) addCurrent(info MarketInfo) {
	asset.TotalVolumePerMarket += info.Volume
	asset.totalPricePerMarket += info.Price
	asset.numberOfTrades++
	if info.IsBuy {
		asset.totalBought++
	}
	asset.volumeWeightedPrice += info.Price * info.Volume
	asset.Market = info.Market
}

func (asset *TotalAssets) processAverages() {
	asset.MeanVolumePerMarket = asset.TotalVolumePerMarket / float64(asset.numberOfTrades)
	asset.MeanPricePerMarket = asset.totalPricePerMarket / float64(asset.numberOfTrades)
	asset.PercentageBuyOrdersPerMarket = float64(asset.totalBought) / float64(asset.numberOfTrades)
	asset.VolumeWeightedAveragePricePerMarket = asset.volumeWeightedPrice / asset.TotalVolumePerMarket

}

func (asset *TotalAssets) print() {
	marshal, err := json.Marshal(asset)
	if err != nil {
		return
	}
	fmt.Println(string(marshal))
}

type TotalAssets struct {
	//TODO Evaluate size of total value to insure float64 is sufficient
	Market                              int     `json:"market"`
	TotalVolumePerMarket                float64 `json:"total_volume"`
	MeanPricePerMarket                  float64 `json:"mean_price"`
	MeanVolumePerMarket                 float64 `json:"mean_volume"`
	VolumeWeightedAveragePricePerMarket float64 `json:"volume_weighted_average_price"`
	PercentageBuyOrdersPerMarket        float64 `json:"percentage_buy"`
	totalPricePerMarket                 float64
	numberOfTrades                      int
	totalBought                         int
	volumeWeightedPrice                 float64
}

//easyjson:json
type MarketInfo struct {
	Id     int     `json:"id"`
	Market int     `json:"market"`
	Price  float64 `json:"price"`
	Volume float64 `json:"volume"`
	IsBuy  bool    `json:"is_buy"`
}

type Logger interface {
	Println(v ...interface{})
}

var StdErrLogger = Logger(log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile))
