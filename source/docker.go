package source

import (
	"context"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	log "github.com/sirupsen/logrus"
)

const (
	ErrorContainerList = iota + 200
	ErrorMarshallJSON
)

type DockerSource struct {
	client *client.Client
	log    *log.Logger
}

func (d *DockerSource) Get() ([]*Result, error) {
	d.log.Debug("Listing Docker containers...")
	containers, err := d.client.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		return nil, err
	}

	d.log.Debugf("Docker API returned %d containers", len(containers))
	results := []*Result{}
	for _, container := range containers {
		result := &Result{Labels: map[string]string{}}
		result.Created = time.Unix(container.Created, 0).UTC()
		result.Name = container.Image
		result.Labels["container"] = strings.TrimLeft(container.Names[0], "/")
		results = append(results, result)
	}

	return results, nil
}

func NewDockerSource(l *log.Logger) (*DockerSource, error) {
	dockerClient, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}
	return &DockerSource{
		client: dockerClient,
		log:    l,
	}, nil
}
