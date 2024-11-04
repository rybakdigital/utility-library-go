package despatcher

import (
	mapset "github.com/deckarep/golang-set/v2"
	log "github.com/rybakdigital/utility-library-go/logging/logger"
)

type Despatcher struct {
	Logger   *log.Logger
	Adapters mapset.Set[Adapter]
}
