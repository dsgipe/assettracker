package main

import (
	"asset/tracker/manageincoming"
	"os"
)

func main() {
	assets := manageincoming.TrackAssets(os.Stdin)
	manageincoming.PrintAssetsAsJson(assets)
}
