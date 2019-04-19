# Pulumi DataDog provider (based on Terraform provider)

## Configuration

```sh
pulumi config set --secret datadog:apiKey YOUR_API_KEY
pulumi config set --secret datadog:appKey YOUR_APP_KEY
```

## Supported resources:

* DatadogMonitor
* DatadogUser
* DatadogMetricMetadata
* DatadogDowntime
* DatadogScreenboard
