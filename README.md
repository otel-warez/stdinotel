# Stdin-Otel

This program runs on stdin and sends data it receives (logs only) to an OTLP or Splunk HEC endpoint.

You can use it like this:

```shell
$> cat file.txt | stdinotel
```

## Installation

```shell
$> go install github.com/otel-warez/stdinotel/cmd/stdinotel@latest
```

See #Building instead.

## Configuration

`stdinotel` uses environment variables for configuration:

| Name                               | Description                                                      |
|------------------------------------|------------------------------------------------------------------|
| STDINOTEL_PROTOCOL                 | one of `splunk_hec`, `otlp`, `otlphttp`                          |
| STDINOTEL_ENDPOINT                 | the endpoint to address data to, such as `http://localhost:4317` |
| STDINOTEL_TOKEN                    | the authentication token (only for `splunk_hec`)                 |
| STDINOTEL_SPLUNK_INDEX             | the Splunk index to set (only for `splunk_hec`)                  |
| STDINOTEL_TLS_INSECURE_SKIP_VERIFY | whether to check the TLS certificate                             |

## Building

```shell
$> make stdinotel
```

## Standard in
stdinotel consumes data passed in via standard input.

### Piping
If it receives data via pipe, the program consumes all data passed in, blocking until such time it sends it out.
It exits right away.

### Interactive
If stdinotel is run in an interactive CLI, you can exit by entering an empty line, Ctrl+C or Ctrl+D.
