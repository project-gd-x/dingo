package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImports_Add(t *testing.T) {
	imports := Imports{}

	i1 := imports.Add("github.com/core/application/view/signup", "signup")
	i2 := imports.Add("github.com/core/application/signup", "signup")
	i3 := imports.Add("github.com/core/signup", "signup")
	i4 := imports.Add("github.com/core/application/view/signup", "signup")
	i5 := imports.Add("github.com/core/signup", "signup")
	imports.Add("github.com/project-gd-x/api/golang/gdx/events/v1", "v1")
	imports.Add("github.com/project-gd-x/gd-core/internal/core/infrastructure/events", "events")
	imports.Add("github.com/project-gd-x/gd-core/internal/services/analytics/internal/events", "events")

	assert.Equal(t, &Import{Name: "github.com/core/application/view/signup", LocalName: "signup", Renamed: false, OriginalLocalName: ""}, i1)
	assert.Equal(t, &Import{Name: "github.com/core/application/signup", LocalName: "signup_1", Renamed: true, OriginalLocalName: "signup"}, i2)
	assert.Equal(t, &Import{Name: "github.com/core/signup", LocalName: "signup_2", Renamed: true, OriginalLocalName: "signup"}, i3)
	assert.Equal(t, &Import{Name: "github.com/core/application/view/signup", LocalName: "signup", Renamed: false, OriginalLocalName: ""}, i4)
	assert.Equal(t, &Import{Name: "github.com/core/signup", LocalName: "signup_2", Renamed: true, OriginalLocalName: "signup"}, i5)

	expected := Imports{
		Items: []*Import{
			{Name: "github.com/core/application/view/signup", LocalName: "signup", Renamed: false, OriginalLocalName: ""},
			{Name: "github.com/core/application/signup", LocalName: "signup_1", Renamed: true, OriginalLocalName: "signup"},
			{Name: "github.com/core/signup", LocalName: "signup_2", Renamed: true, OriginalLocalName: "signup"},
		},
	}

	assert.Equal(t, expected, imports)
}

func TestServicesImports(t *testing.T) {
	services := map[string]*Service{
		"EventSubscriber": {
			Type: "*github.com/project-gd-x/gd-core/internal/services/analytics/internal/subscriber.EventSubscriber",
			Constructor: &Constructor{
				Name:      "github.com/project-gd-x/gd-core/internal/services/analytics/internal/subscriber.NewEventSubscriber",
				Arguments: []Argument{"@{EventConsumer}", "@{AdwordsSignUpStorage}"},
			},
		},
		"EventConsumer": {
			Interface: "github.com/project-gd-x/api/golang/gdx/events/v1.Consumer",
			Constructor: &Constructor{
				Name:      "github.com/project-gd-x/gd-core/internal/core/infrastructure/events.NewEventsConsumer",
				Arguments: []Argument{"Analytics", "#{version}", "#{cfg.Amqp.URI}"},
				Error:     true,
			},
		},
		"AdwordsSignUpStorage": {
			Interface: "github.com/project-gd-x/gd-core/internal/services/analytics/internal/events.AdwordsSignUpStorage",
			Constructor: &Constructor{
				Name:      "github.com/project-gd-x/gd-core/internal/services/analytics/internal/infrastructure/bigquery.NewAdwordsSignUpStorage",
				Arguments: []Argument{"@{BigQueryClient}", "rawdata", "marketing_signup_adwords"},
			},
		},
		"BigQueryClient": {
			Type: "*cloud.google.com/go/bigquery.Client",
			Constructor: &Constructor{
				Name:      "github.com/project-gd-x/gd-core/pkg/shared.NewBigQueryClient",
				Arguments: []Argument{"#{appContext}", "#{cfg.BigQuery}"},
				Error:     true,
			},
		},
	}

	imports := &Imports{}

	for _, service := range services {
		imports = service.Imports(imports)
	}

	expected := Imports{
		Items: []*Import{
			{Name: "github.com/project-gd-x/gd-core/internal/services/analytics/internal/subscriber", LocalName: "subscriber", Renamed: false, OriginalLocalName: ""},
			{Name: "github.com/project-gd-x/api/golang/gdx/events/v1", LocalName: "v1", Renamed: false, OriginalLocalName: ""},
			{Name: "github.com/project-gd-x/gd-core/internal/core/infrastructure/events", LocalName: "events", Renamed: false, OriginalLocalName: ""},
			{Name: "github.com/project-gd-x/gd-core/internal/services/analytics/internal/events", LocalName: "events_5", Renamed: true, OriginalLocalName: "events"},
			{Name: "github.com/project-gd-x/gd-core/internal/services/analytics/internal/infrastructure/bigquery", LocalName: "bigquery", Renamed: false, OriginalLocalName: ""},
			{Name: "cloud.google.com/go/bigquery", LocalName: "bigquery_1", Renamed: true, OriginalLocalName: "bigquery"},
			{Name: "github.com/project-gd-x/gd-core/pkg/shared", LocalName: "shared", Renamed: false, OriginalLocalName: ""},
		},
	}

	assert.ElementsMatch(t, expected.Items, imports.Items)
}
