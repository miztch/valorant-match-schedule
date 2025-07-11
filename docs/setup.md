# Setup instruction

## Service account keys for Google Cloud IAM

- See [here](https://cloud.google.com/iam/docs/creating-managing-service-account-keys)
- Save as `service_account_key.json` and create a SSM Parameter.

```bash
aws ssm put-parameter --name "/google/service-account-key" --value "$(cat service_account_key.json | jq -c .)" --type "SecureString"
```

## Google Calendars

- Create calendars and copy their ids for each region: see [here](https://docs.simplecalendar.io/find-google-calendar-id/)
  - 5 calendars are needed for each region.
- Fill in `samconfig.toml`

## mise

- Install [mise](https://mise.jdx.dev) as package manager and task runner.
- Run `mise install` in project root to install golang and AWS SAM CLI

## Provisioning

```bash
mise build
sam deploy --guided --config-env production --config-file samconfig.toml
```
