package acceptance_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"

	"github.com/quii/go-specs-greet/acceptance-tests/drivers"
	"github.com/quii/go-specs-greet/acceptance-tests/specifications"
)

func TestSpecsForGreet(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	exposedPort := runAsContainer(t)

	client := http.Client{Timeout: time.Second}
	baseURL := fmt.Sprintf("http://localhost:%s", exposedPort)
	greeterDriver := drivers.HTTPGreeter{Client: &client, BaseURL: baseURL}

	specifications.SpecsForGreet(t, greeterDriver)
	specifications.SpecsForCurse(t, greeterDriver)
}

func runAsContainer(t *testing.T) string {
	t.Helper()

	ctx := context.Background()
	containerRequest := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:       "../.",
			Dockerfile:    "./Dockerfile",
			PrintBuildLog: true,
		},
		ExposedPorts: []string{"8080"},
	}
	container, err := testcontainers.GenericContainer(
		ctx,
		testcontainers.GenericContainerRequest{
			ContainerRequest: containerRequest,
			Started:          true,
		},
	)
	assert.NoError(t, err)

	t.Cleanup(func() {
		assert.NoError(t, container.Terminate(ctx))
	})

	mappedPort, err := container.MappedPort(ctx, "8080")
	assert.NoError(t, err)
	return mappedPort.Port()
}
