package imagespy

import "github.com/docker/distribution/reference"

var DefaultRegistryWhitelist = map[string]struct{}{
	"index.docker.io": struct{}{},
	"docker.io":       struct{}{},
	"quay.io":         struct{}{},
}

func isRegistryWhitelisted(name string, whitelist map[string]struct{}) (bool, error) {
	ref, err := reference.ParseNormalizedNamed(name)
	if err != nil {
		return false, err
	}

	domain := reference.Domain(ref)
	_, ok := whitelist[domain]
	return ok, nil
}
