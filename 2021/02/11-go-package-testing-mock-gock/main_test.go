package main

import (
	"context"
	"net/http"
	"path"
	"testing"

	"github.com/google/go-cmp/cmp"
	gock "gopkg.in/h2non/gock.v1"
)

func Test_requestWeather(t *testing.T) {
	t.Cleanup(func() {
		gock.EnableNetworking()
		gock.OffAll()
	})

	gock.DisableNetworking()

	type input struct {
		ctx   context.Context
		appID string
		zip   string
		units string
	}

	type output struct {
		res     Result
		withErr bool
	}

	tests := []struct {
		name   string
		setup  func()
		input  input
		output output
	}{
		{
			"200",
			func() {
				gock.New("https://api.openweathermap.org").
					MatchParams(map[string]string{
						"zip":   "90210",
						"appid": "appID",
						"units": "metric",
					}).
					Get("/data/2.5/weather").
					Reply(http.StatusOK).
					File(path.Join("fixtures", "200.json"))
			},
			input{
				context.Background(),
				"appID",
				"90210",
				"metric",
			},
			output{
				res: Result{
					Weather: []Weather{
						{
							Description: "scattered clouds",
						},
					},
					Main: Main{
						Temperature: 12.859999656677246,
						FeelsLike:   12.079999923706055,
					},
				},
			},
		},
		{
			"500",
			func() {
				gock.New("https://api.openweathermap.org").
					MatchParams(map[string]string{
						"zip":   "90210",
						"appid": "appID-500",
						"units": "metric",
					}).
					Get("/data/2.5/weather").
					Reply(http.StatusInternalServerError)
			},
			input{
				context.Background(),
				"appID-500",
				"90210",
				"metric",
			},
			output{
				withErr: true,
			},
		},
		{
			"401",
			func() {
				gock.New("https://api.openweathermap.org").
					MatchParams(map[string]string{
						"zip":   "90210",
						"appid": "appID-401",
						"units": "metric",
					}).
					Get("/data/2.5/weather").
					Reply(http.StatusUnauthorized).
					JSON(map[string]interface{}{
						"cod":     401,
						"message": "Invalid API key. Please see http://openweathermap.org/faq#error401 for more info.",
					})
			},
			input{
				context.Background(),
				"appID-401",
				"90210",
				"metric",
			},
			output{
				withErr: true,
			},
		},
		{
			"Error: invalid JSON",
			func() {
				gock.New("https://api.openweathermap.org").
					MatchParams(map[string]string{
						"zip":   "90210",
						"appid": "appID-error",
						"units": "metric",
					}).
					Get("/data/2.5/weather").
					Reply(http.StatusOK).
					BodyString(`{"broken":`)
			},
			input{
				context.Background(),
				"appID-error",
				"90210",
				"metric",
			},
			output{
				withErr: true,
			},
		},
		{
			"Error: http.Client.Do",
			func() {
				gock.New("https://api.openweathermap.org").
					MatchParams(map[string]string{
						"zip":   "90210",
						"appid": "appID-err-req",
						"units": "metric",
					}).
					Get("/data/2.5/weather").
					Reply(http.StatusMovedPermanently).
					SetHeader("Location", "")
			},
			input{
				context.Background(),
				"appID-err-req",
				"90210",
				"metric",
			},
			output{
				withErr: true,
			},
		},
		{
			"Error: http.NewRequestWithContext",
			func() {},
			input{
				ctx: nil,
			},
			output{
				withErr: true,
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			tt.setup()

			actualRes, actualErr := requestWeather(tt.input.ctx, &http.Client{}, tt.input.appID, tt.input.zip, tt.input.units)

			if (actualErr != nil) != tt.output.withErr {
				t.Fatalf("expected error %t, actual %s", tt.output.withErr, actualErr)
			}

			if !cmp.Equal(tt.output.res, actualRes) {
				t.Fatalf("expected output do not match\n%s", cmp.Diff(tt.output.res, actualRes))
			}
		})
	}
}
