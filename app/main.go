package main

import (
	"flag"
	"net/http"
	"path/filepath"

	imagespy "github.com/imagespy/client-go"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"github.com/wndhydrnt/imagespy/source"
	"github.com/wndhydrnt/imagespy/web"
)

var (
	httpAddress = flag.String("http.address", ":8888", "")
	logLevel    = flag.String("log.level", "ERROR", "")
	uiPath      = flag.String("ui.path", "./ui", "")
)

func main() {
	flag.Parse()
	lvl, err := log.ParseLevel(*logLevel)
	if err != nil {
		log.Fatal(err)
	}

	log.SetLevel(lvl)
	collector := source.NewCollector(log.StandardLogger())
	dockerSource, err := source.NewDockerSource(log.StandardLogger())
	if err != nil {
		log.Fatal(err)
	}

	collector.AddSource("docker", dockerSource)
	imageSpyClient := imagespy.NewClientV1()
	imagespy.Log = log.StandardLogger()
	spiesHandler, _ := web.NewSpiesHandler(collector, imageSpyClient, log.StandardLogger())
	router := httprouter.New()
	router.HandlerFunc("GET", "/spies", spiesHandler.Get)
	static, err := filepath.Abs(*uiPath)
	if err != nil {
		log.Fatal(err)
	}

	router.ServeFiles("/app/*filepath", http.Dir(static))
	router.HandlerFunc("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Location", "/app/")
		w.WriteHeader(http.StatusMovedPermanently)
	})
	log.Fatal(http.ListenAndServe(*httpAddress, router))
}
