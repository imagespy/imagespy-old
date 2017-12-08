package imagespy

import (
	"fmt"
	"net/http"
	"time"
)

// Image is a Docker image.
type Image struct {
	Created time.Time `json:"created"`
	Digest  string    `json:"digest"`
	Name    string    `json:"name"`
	Tag     string    `json:"tag"`
}

// ImageSpy is a ImageSpy.
type ImageSpy struct {
	CurrentImage *Image `json:"current_image"`
	LatestImage  *Image `json:"latest_image"`
	Name         string `json:"name"`
}

// ImageSpyService handles interactions.
type ImageSpyService struct {
	cacheUnknownImages bool
	registryWhitelist  map[string]struct{}
	requester          *requester
}

// Get retrieves an ImageSpy.
// Creates the ImageSpy if it does not exist.
func (is *ImageSpyService) Get(name string) (*ImageSpy, error) {
	whitelisted, err := isRegistryWhitelisted(name, is.registryWhitelist)
	if err != nil {
		return nil, err
	}

	if whitelisted == false {
		return nil, fmt.Errorf("Registry domain of image %s is not whitelisted", name)
	}

	resp, err := is.requester.readAsJSON("/v1/images/" + name)
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case http.StatusOK:
		imageSpy := &ImageSpy{}
		err = is.requester.parseJSON(resp.Body, imageSpy)
		if err != nil {
			return nil, err
		}

		return imageSpy, nil
	case http.StatusNotFound:
		// TODO: This is only necessary because httpcache only caches when the body is read. Another solution possible?
		is.requester.parseJSON(resp.Body, struct{}{})
		return nil, fmt.Errorf("Error retrieving ImageSpy: API returned status code %d", resp.StatusCode)
	default:
		return nil, fmt.Errorf("Error retrieving ImageSpy: API returned status code %d", resp.StatusCode)
	}
}
