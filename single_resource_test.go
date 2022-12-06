package main

import (
	// "fmt"
	"log"
	"os"
	"path/filepath"
	"testing"

	message "github.com/pact-foundation/pact-go/v2/message/v4"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/cmdutil"
	"github.com/stretchr/testify/assert"
	// engine "github.com/ringods/pulumi-pact-poc/github.com/pulumi/pulumi/sdk/v3/proto/go"
)

func TestUserAPIClient(t *testing.T) {
	// Specify the two applications in the integration we are testing
	// NOTE: this can usually be extracted out of the individual test for re-use)
	mockProvider, err := message.NewSynchronousPact(message.Config{
		Consumer: "pulumi-language-host",
		Provider: "pulumi-engine",
	})
	assert.NoError(t, err)

	dir, _ := filepath.Abs("./proto/engine.proto")

	grpcInteraction := `{
		"pact:proto": "` + filepath.ToSlash(dir) + `",
		"pact:proto-service": "Calculator/calculateMulti",
		"pact:content-type": "application/protobuf",
		"request": {
			"shapes": [
				{
					"rectangle": {
						"length": "matching(number, 3)",
						"width": "matching(number, 4)"
					}
				},
				{
					"square": {
						"edge_length": "matching(number, 3)"
					}
				}
			]
		},
		"response": {
			"value": [ "matching(number, 12)", "matching(number, 9)" ]
		}
	}`

	// Defined a new message interaction, and add the plugin config and the contents
	err = mockProvider.
		AddSynchronousMessage("calculate rectangle area request").
		UsingPlugin(message.PluginConfig{
			Plugin: "protobuf",
		}).
		WithContents(grpcInteraction, "application/grpc").
		// Start the gRPC mock server
		StartTransport("grpc", "127.0.0.1", nil).
		// Execute the test
		ExecuteTest(t, func(transport message.TransportConfig, m message.SynchronousMessage) error {
			// Execute the gRPC client against the mock server
			log.Println("Mock server is running on ", transport.Port)
			wd, err := os.Getwd()
			if err != nil {
				return err
			}
			sink := cmdutil.Diag()
			pCtx, err := plugin.NewContext(sink, sink, nil, nil, wd, nil, false, nil)
			if err != nil {
				return err
			}
			_, err = plugin.NewDefaultHost(pCtx, nil, false, nil)

			// Assert: check the result
			assert.NoError(t, err)
			// assert.Equal(t, float32(12.0), area[0])
			// assert.Equal(t, float32(9.0), area[1])

			// return err
			return nil
		})

	assert.NoError(t, err)
}
