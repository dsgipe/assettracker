package manageincoming

import (
	"bufio"
	"github.com/mailru/easyjson"
	"io"
)

func TrackAssets(stdin io.Reader) []TotalAssets {
	scanner := bufio.NewScanner(stdin)

	totalAssets := make([]TotalAssets, 0)
	marketToIndex := map[int]int{}

	for scanner.Scan() {
		marketInfo := MarketInfo{}
		err := easyjson.Unmarshal(scanner.Bytes(), &marketInfo)

		if err != nil {
			StdErrLogger.Println("did not scan:", scanner.Text())
			continue
		}

		//Map the market to an array index in order to allow for non index like market values.
		if _, ok := marketToIndex[marketInfo.Market]; !ok {
			marketToIndex[marketInfo.Market] = len(totalAssets)
			totalAssets = append(totalAssets, TotalAssets{})
		}

		totalAssets[marketToIndex[marketInfo.Market]].addCurrent(marketInfo)
	}
	for market := range totalAssets {
		totalAssets[market].processAverages()
	}

	return totalAssets
}

func PrintAssetsAsJson(totalAssets []TotalAssets) {
	for _, asset := range totalAssets {
		asset.print()
	}
}
