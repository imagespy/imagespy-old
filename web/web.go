package web

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/imagespy/client-go"
	log "github.com/sirupsen/logrus"
	"github.com/imagespy/imagespy/source"
)

type apiError struct {
	Code    int
	Message string
}

type response struct {
	source *source.Result
	spy    *imagespy.ImageSpy
}

type Spy struct {
	Created      time.Time         `json:"created"`
	CurrentImage *imagespy.Image   `json:"current_image"`
	Labels       map[string]string `json:"labels"`
	LatestImage  *imagespy.Image   `json:"latest_image"`
	Name         string            `json:"name"`
}

type SpiesHandler struct {
	collector      *source.Collector
	imageSpyClient *imagespy.ClientV1
	log            *log.Logger
}

// Get returns a list of spies in JSON.
func (sh *SpiesHandler) Get(w http.ResponseWriter, r *http.Request) {
	results := sh.collector.Collect()
	spies := []*Spy{}
	c := make(chan *response, len(results))
	wg := &sync.WaitGroup{}
	wg.Add(len(results))

	for _, result := range results {
		go func(s *source.Result) {
			imageSpy, err := sh.imageSpyClient.ImageSpy.Get(s.Name)
			if err != nil {
				sh.log.Warnf("Error getting ImageSpy with name %s: %s", s.Name, err)
				wg.Done()
				return
			}

			c <- &response{s, imageSpy}
			wg.Done()
		}(result)
	}

	wg.Wait()
	close(c)
	for resp := range c {
		spy := &Spy{
			Created:      resp.source.Created,
			CurrentImage: resp.spy.CurrentImage,
			Labels:       resp.source.Labels,
			LatestImage:  resp.spy.LatestImage,
			Name:         resp.source.Name,
		}

		spies = append(spies, spy)
	}

	payload, err := json.Marshal(spies)
	if err != nil {
		sh.log.Errorf("Error marshalling spies into JSON: %s", err)
		writeError(100, "Error marshalling spies into JSON", 500, w)
		return
	}

	w.WriteHeader(200)
	w.Write(payload)
}

// NewSpiesHandler instantiates a new SpiesHandler.
func NewSpiesHandler(c *source.Collector, isc *imagespy.ClientV1, l *log.Logger) (*SpiesHandler, error) {
	return &SpiesHandler{
		collector:      c,
		imageSpyClient: isc,
		log:            l,
	}, nil
}

func writeError(code int, msg string, status int, w http.ResponseWriter) {
	err := &apiError{code, msg}
	payload, _ := json.Marshal(err)
	w.WriteHeader(status)
	w.Write(payload)
}
