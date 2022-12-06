# Testing Pulumi with Pact

Proof of Concept using [Pact](https://pact.io) to test the different components of Pulumi.

## Requirements

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
