# valorant-match-schedule

You can check worldwide VALORANT matches schedule in Google Calendar.

## Public URL

https://valorant-calendar.mizt.ch/

> [!TIP]
> - All Calendars except `International` contain tier-1 and tier-2 matches played in the region.

- [International](https://calendar.google.com/calendar/u/0?cid=ZjJiMjY1MTkwZmNkMDY5NjY5M2FiZjgwODE1OGRkZDQ1MTQwNjY1ZGRkMGM1YjM3YWFhYTA0ZGJjZDYxZWFhOEBncm91cC5jYWxlbmRhci5nb29nbGUuY29t) : contains all VCT international matches and some community events.
- [Americas](https://calendar.google.com/calendar/u/0?cid=YmRkZTBmNWZhYTQ2ZTQxZGRmYWJkZTJmYjhlZDZhZTcxOTdiNjQ2MTQyYjhlZDAzOTdmMGE0ZjBiYmQxMDlmNEBncm91cC5jYWxlbmRhci5nb29nbGUuY29t)
- [EMEA](https://calendar.google.com/calendar/u/0?cid=OWNjOGI0YmIwMjFiMDE5NGIyNDEzOWRiMjU1NjFmYTc0ZGVhYTc0MzU0YWY1NDkzMDFiNDRhZWNjOTJiMzJiOUBncm91cC5jYWxlbmRhci5nb29nbGUuY29t)
- [Pacific](https://calendar.google.com/calendar/u/0?cid=N2EzN2YxZmYzMzdiODBmOWRiOWFhOGVhNWQ2ZDBiYWY5ODE4MDM0YmIyMDhiNDAyNWUzNDFjNGE3NTMyMTE5OUBncm91cC5jYWxlbmRhci5nb29nbGUuY29t)
- [China](https://calendar.google.com/calendar/u/0?cid=ODFhYjAxZTQ1ZmNiNGY5MjViNzZmOTdmYjA0MjA3Mjc3MjUxZjhhM2FhMDUyMGYxMDA3MmUzYzBhMjQyMTdjNEBncm91cC5jYWxlbmRhci5nb29nbGUuY29t)

> [!WARNING]
> - Below are scheduled to be de-updated or removed in the near future
> - I recommend that you replace each of them by subscribing to a new calendar that will replace it.

- [NA](https://calendar.google.com/calendar/u/0?cid=NDg2OTg5YjJjYTEwZTk5OGVlODU4YTcyMDRiMTNjZDkzNWQ0MTFmMjlkYzA4MzA1ZGJiZDFjNDIzN2YyYTIxZEBncm91cC5jYWxlbmRhci5nb29nbGUuY29t) : All events in this calendar are also included in `Americas`.
- [Brazil / LATAM](https://calendar.google.com/calendar/u/0?cid=NTNjYTg1NDZhNGQzM2IzMTkzNmVkMjA0NjE2NmY2OTA2ODc1YzhiMzM3N2JhOThiOWFkMWJhZjNhMDQwZjA3NUBncm91cC5jYWxlbmRhci5nb29nbGUuY29t) : All events in this calendar are also included in `Americas`.
- [APAC](https://calendar.google.com/calendar/u/0?cid=YTZiOWM3ZGI5YzdmM2U2YjFmNjNmMjc4MjhkMzM0ZjQ0NTFhZDI0ZGU1NzA0YjI4YTk5YzU1ZGEwNjk5YjdiZkBncm91cC5jYWxlbmRhci5nb29nbGUuY29t) : All events in this calendar are also included in `Pacific`.
- [East Asia](https://calendar.google.com/calendar/u/0?cid=MWEyMWYzYWFlMDRjYWJiODQxNDFlOGE1ZjRhYTRjZTA4NWNmMzZlNDA0YjQyOGIzYjZjNzYxMWE4MTllZjczMkBncm91cC5jYWxlbmRhci5nb29nbGUuY29t) : All events in this calendar are also included in `Pacific`.

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
  - 9 calendars are needed for each regions.

## Provisioning

You can use [AWS SAM](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html) to provision this application.
- Templates are nested, so you have to execute `sam deploy` only once

```bash
cd valorant-match-schedule/
sam build
sam deploy --guided --capabilities CAPABILITY_IAM CAPABILITY_AUTO_EXPAND
```
