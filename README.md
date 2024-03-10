# cosmos-sdk-gridnode

# new proto files
Since we don't have an app we cannot use ignite to generate new files from the proto files.

Instead use this command from within the proto directory.

```bash
buf generate --template buf.gen.gogo.yaml
buf generate --template buf.gen.pulsar.yaml
```