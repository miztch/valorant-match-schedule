# valorant-match-schedule

You can check worldwide VALORANT matches schedule in Google Calendar.

## Public URL

- TBA **(Roadmap: after integrated with [dima](https://github.com/miztch/dima) and [sasha](https://github.com/miztch/sasha), it comes!)**

## Architecture overview
![architecture](image/valorant-match-schedule.drawio.svg)

1. Lambda function: `get-match-list` is periodically invoked from EventBridge
2. `get-match-list` decides which day to fetch match information, and publishes records including specified dates strings to SQS queue: `FetchDailyMatches`
3. `fetch-daily-matches` is triggered by `FetchDailyMatches` queue. fetches match data for a day from API service and put into DynamoDB table: `MatchList`
4. `MatchList` table streams captures item-level data modification to Lambda function: `add-gcal-event`
5. `add-gcal-event` creates/modifies Google Calendar events. `Outbox` table manages their state and information.

## Prerequisites

If you want to do it yourself, you should be ready with:
- service account keys for Google Cloud IAM ( `service_account_key.json` ): see [here](https://cloud.google.com/iam/docs/creating-managing-service-account-keys)

    ```bash
    git clone https://github.com/miztch/valorant-match-schedule
    cp service_account_key.json valorant-match-schedule/add-gcal-event/function/service_account_key.json
    ```
- calendars and their ids for each regions: see [here](https://docs.simplecalendar.io/find-google-calendar-id/)
  - 6 calendars are needed. (for `EMEA`, `NA`, `BR/LATAM`, `APAC`, `EAST_ASIA`, `INTERNATIONAL`)

## Provisioning

You can use [AWS SAM](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html) to provision this application.
- Templates are nested, so you have to execute `sam deploy` only once

```bash
cd valorant-match-schedule/
sam build
sam deploy --guided --capabilities CAPABILITY_IAM CAPABILITY_AUTO_EXPAND
```
