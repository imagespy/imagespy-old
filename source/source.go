package source

import (
	"time"

	log "github.com/sirupsen/logrus"
)

type Source interface {
	Get() ([]*Result, error)
}

type Result struct {
	Created time.Time
	Labels  map[string]string
	Name    string
}

type Collector struct {
	log     *log.Logger
	sources []Source
}

// AddSource adds a new source to the collector.
func (c *Collector) AddSource(name string, s Source) {
	c.sources = append(c.sources, s)
}

// Collect collects results from sources.
func (c *Collector) Collect() []*Result {
	results := []*Result{}
	for _, s := range c.sources {
		r, err := s.Get()
		if err != nil {
			log.Errorf("Error collecting results from source: %s", err)
			continue
		}

		results = append(results, r...)
	}

	return results
}

// NewCollector instantiates a new Collector.
func NewCollector(l *log.Logger) *Collector {
	return &Collector{log: l}
}
