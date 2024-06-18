package main

import (
	souin_skipper "github.com/HobMartin/souin/plugins/skipper"
	"github.com/zalando/skipper"
	"github.com/zalando/skipper/filters"
)

func main() {
	skipper.Run(skipper.Options{
		Address:       ":80",
		RoutesFile:    "examples/example.yaml",
		CustomFilters: []filters.Spec{souin_skipper.NewHTTPCacheFilter()}},
	)
}
