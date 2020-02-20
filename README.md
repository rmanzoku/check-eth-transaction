# check-eth-transaction
## Description
Check Ethereum transaction failure for addresses.

The source from Google BigQuery, its confirmation number is probably 20.

## Setting for mackerel-agent
```
[plugin.checks.eth-tx]
command = ["check-eth-transaction", "-p", "project-id", "-a", "0x0000000000000000000000000000000000000001,0x0000000000000000000000000000000000000002"]
notification_interval = 10
max_check_attempts = 1
check_interval = 5
timeout_seconds = 45
prevent_alert_auto_close = true
env = { GOOGLE_APPLICATION_CREDENTIALS = "/path/to/credential.json" }
```

## Usage
### Options

```
  -a string
        addresses (default "0x0000000000000000000000000000000000000000")
  -p string
        projectID
```
