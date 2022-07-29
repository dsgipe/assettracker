package manageincoming

import (
	"io"
	"math"
	"strings"
	"testing"
)

const tolerance = 0.001

func withTolerance(a, b float64) bool {
	return math.Abs(a-b) > tolerance
}

func TestTrackAssets(t *testing.T) {

	input := "BEFORE \n{\"id\":99,\"market\":99,\"price\":47.73686720241104,\"volume\":1253.893286503535,\"is_buy\":true}\n{\"id\":100,\"market\":100,\"price\":48.04464740087032,\"volume\":1114.5008438076445,\"is_buy\":true}\n{\"id\":101,\"market\":99,\"price\":49.140659465620985,\"volume\":3975.4408412282005,\"is_buy\":false}\n"
	var reader io.Reader = strings.NewReader(input)
	totalAssets := TrackAssets(reader)

	//if assets, ok := totalAssets[99]; ok {
	asset1 := totalAssets[0]
	if withTolerance(asset1.TotalVolumePerMarket, 5229.334) {
		t.Errorf("failed to equal based on tolerance %f, %f", asset1.TotalVolumePerMarket, tolerance)
	}
	if asset1.numberOfTrades != 2 {
		t.Error(asset1.numberOfTrades)
	}
	if asset1.Market != 99 {
		t.Error(asset1.Market)
	}
	if withTolerance(asset1.MeanVolumePerMarket, 2614.667) {
		t.Errorf("failed to equal based on tolerance %f, %f", asset1.MeanVolumePerMarket, tolerance)
	}
	if withTolerance(asset1.TotalVolumePerMarket, 5229.334) {
		t.Errorf("failed to equal based on tolerance %f, %f", asset1.TotalVolumePerMarket, tolerance)
	}
	if withTolerance(asset1.PercentageBuyOrdersPerMarket, .5) {
		t.Errorf("failed to equal based on tolerance %f, %f", asset1.PercentageBuyOrdersPerMarket, tolerance)
	}

	asset2 := totalAssets[1]
	if withTolerance(asset2.TotalVolumePerMarket, 1114.500844) {
		t.Errorf("failed to equal based on tolerance %f, %f", asset2.TotalVolumePerMarket, tolerance)
	}
	if asset2.numberOfTrades != 1 {
		t.Error(asset2.numberOfTrades)
	}
	if withTolerance(asset2.MeanVolumePerMarket, 1114.500844) {
		t.Errorf("failed to equal based on tolerance %f, %f", asset2.MeanVolumePerMarket, tolerance)
	}

	PrintAssetsAsJson(totalAssets)
}
