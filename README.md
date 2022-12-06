# Testing Pulumi with Pact

Proof of Concept using [Pact](https://pact.io) to test the different components of Pulumi.

## Requirements

* Install `protoc`, `protoc-gen-go` and `protoc-gen-go-grpc` tools
  * Using Homebrew: `brew install protobuf protoc-gen-go protoc-gen-go-grpc`
* Install the [pact-plugin-cli](https://github.com/pact-foundation/pact-plugins/tree/main/cli#installing).
* Download the `install-plugin.sh` script of the [protobuf plugin](https://github.com/pactflow/pact-protobuf-plugin/releases/).
* Execute the `install-plugin.sh`.
* Verify the installation using the command `pact-plugin-cli list`. The output should list the protobuf as installed and enabled:
   ```sh
   ./pact-plugin-cli list
   ┌──────────┬─────────┬───────────────────┬─────────────────────────────────────────────┬─────────┐
   │ Name     ┆ Version ┆ Interface Version ┆ Directory                                   ┆ Status  │
   ╞══════════╪═════════╪═══════════════════╪═════════════════════════════════════════════╪═════════╡
   │ protobuf ┆ 0.2.0   ┆ 1                 ┆ /Users/ringods/.pact/plugins/protobuf-0.2.0 ┆ enabled │
   └──────────┴─────────┴───────────────────┴─────────────────────────────────────────────┴─────────┘
   ```

## Test setup

* Have the protobuf files available. This POC copied the Engine protobuf file from `pulumi/pulumi` over to here.
* Generate Go code from the protobuf file: `protoc --go_out=. --go-grpc_out=. --proto_path ./proto ./proto/engine.proto`
* The `fixtures` folder contains a single Pulumi program creating a `RandomPet` resource:
  * `00-single-resource/yaml`: the handwritten YAML program
  * `00-single-resource/go`: the generated Go program using `pulumi convert`
  * `00-single-resource/ts`: the generated Typescript program using `pulumi convert`


## Reference

* [Pact Go v2](https://github.com/pact-foundation/pact-go/tree/2.x.x)
* [Pact Protobuf Plugin](https://github.com/pactflow/pact-protobuf-plugin) & [examples](https://docs.pact.io/implementation_guides/pact_plugins/examples/protobuf)
