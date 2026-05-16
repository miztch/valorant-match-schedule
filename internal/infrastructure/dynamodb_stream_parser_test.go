package infrastructure

import (
	"reflect"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/miztch/valorant-match-schedule/internal/dto"
)

func TestUnmarshalDynamoDBStreamImageToMatchForStream(t *testing.T) {
	intPtr := func(v int) *int { return &v }
	strPtr := func(v string) *string { return &v }
	int64Ptr := func(v int64) *int64 { return &v }

	tests := []struct {
		name      string
		image     map[string]events.DynamoDBAttributeValue
		want      dto.MatchForStream
		wantError bool
	}{
		{
			name: "full match with teams as list of maps",
			image: map[string]events.DynamoDBAttributeValue{
				"match_id":     events.NewNumberAttribute("378829"),
				"best_of":      events.NewNumberAttribute("5"),
				"start_time":   events.NewStringAttribute("2024-08-25T07:00:00+0000"),
				"end_time":     events.NewStringAttribute("2024-08-25T12:00:00+0000"),
				"event_name":   events.NewStringAttribute("Valorant Champions 2024"),
				"event_detail": events.NewStringAttribute("Playoffs: Grand Final"),
				"match_uri":    events.NewStringAttribute("https://vlr.gg/378829/edward-gaming-vs-team-heretics-valorant-champions-2024-gf"),
				"region":       events.NewStringAttribute("PACIFIC#INTERNATIONAL"),
				"ttl":          events.NewNumberAttribute("1724630400"),
				"teams": events.NewListAttribute([]events.DynamoDBAttributeValue{
					events.NewMapAttribute(map[string]events.DynamoDBAttributeValue{
						"id":   events.NewNumberAttribute("2001"),
						"name": events.NewStringAttribute("EDward Gaming"),
					}),
					events.NewMapAttribute(map[string]events.DynamoDBAttributeValue{
						"id":   events.NewNumberAttribute("2002"),
						"name": events.NewStringAttribute("Team Heretics"),
					}),
				}),
			},
			want: dto.MatchForStream{
				MatchID:     intPtr(378829),
				BestOf:      intPtr(5),
				StartTime:   strPtr("2024-08-25T07:00:00+0000"),
				EndTime:     strPtr("2024-08-25T12:00:00+0000"),
				EventName:   strPtr("Valorant Champions 2024"),
				EventDetail: strPtr("Playoffs: Grand Final"),
				MatchURI:    strPtr("https://vlr.gg/378829/edward-gaming-vs-team-heretics-valorant-champions-2024-gf"),
				Region:      strPtr("PACIFIC#INTERNATIONAL"),
				TTL:         int64Ptr(1724630400),
				Teams: []dto.Team{
					{Id: 2001, Name: "EDward Gaming"},
					{Id: 2002, Name: "Team Heretics"},
				},
			},
		},
		{
			name: "teams are TBD but other fields are set",
			image: map[string]events.DynamoDBAttributeValue{
				"match_id":     events.NewNumberAttribute("999999"),
				"best_of":      events.NewNumberAttribute("3"),
				"start_time":   events.NewStringAttribute("2024-01-01T00:00:00+0000"),
				"end_time":     events.NewStringAttribute("2024-01-01T03:00:00+0000"),
				"event_name":   events.NewStringAttribute("VCT 2024: Pacific Stage 1"),
				"event_detail": events.NewStringAttribute("Group Stage"),
				"match_uri":    events.NewStringAttribute("https://vlr.gg/999999/tbd-vs-tbd-vct-2024-pacific-stage-1-gs"),
				"region":       events.NewStringAttribute("PACIFIC#INTERNATIONAL"),
				"ttl":          events.NewNumberAttribute("1704121200"),
				"teams": events.NewListAttribute([]events.DynamoDBAttributeValue{
					events.NewMapAttribute(map[string]events.DynamoDBAttributeValue{
						"id":   events.NewNumberAttribute("0"),
						"name": events.NewStringAttribute("TBD"),
					}),
					events.NewMapAttribute(map[string]events.DynamoDBAttributeValue{
						"id":   events.NewNumberAttribute("0"),
						"name": events.NewStringAttribute("TBD"),
					}),
				}),
			},
			want: dto.MatchForStream{
				MatchID:     intPtr(999999),
				BestOf:      intPtr(3),
				StartTime:   strPtr("2024-01-01T00:00:00+0000"),
				EndTime:     strPtr("2024-01-01T03:00:00+0000"),
				EventName:   strPtr("VCT 2024: Pacific Stage 1"),
				EventDetail: strPtr("Group Stage"),
				MatchURI:    strPtr("https://vlr.gg/999999/tbd-vs-tbd-vct-2024-pacific-stage-1-gs"),
				Region:      strPtr("PACIFIC#INTERNATIONAL"),
				TTL:         int64Ptr(1704121200),
				Teams: []dto.Team{
					{Id: 0, Name: "TBD"},
					{Id: 0, Name: "TBD"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnmarshalDynamoDBStreamImageToMatchForStream(tt.image)
			if (err != nil) != tt.wantError {
				t.Fatalf("UnmarshalDynamoDBStreamImageToMatchForStream() error = %v, wantError %v", err, tt.wantError)
			}
			if !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("UnmarshalDynamoDBStreamImageToMatchForStream() = %+v, want %+v", *got, tt.want)
			}
		})
	}
}
