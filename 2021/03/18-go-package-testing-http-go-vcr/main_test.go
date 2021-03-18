package main

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"path"
	"testing"

	"github.com/dnaeon/go-vcr/cassette"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/google/go-cmp/cmp"
	gock "gopkg.in/h2non/gock.v1"
)

func Test_requestWeather(t *testing.T) {
	t.Parallel()

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
		setup  func(t *testing.T) (*http.Client, func() error)
		input  input
		output output
	}{
		{
			"200",
			func(t *testing.T) (*http.Client, func() error) {
				r, err := recorder.New(path.Join("fixtures", "weather_200"))
				if err != nil {
					t.Fatalf("failed creating recorder %s", err)
				}

				cleanURL := func(u *url.URL) *url.URL {
					q := u.Query()
					q.Del("appid")

					u.RawQuery = q.Encode()

					return u
				}

				r.AddFilter(func(i *cassette.Interaction) error {
					u, err := url.Parse(i.Request.URL)
					if err != nil {
						return err
					}

					i.URL = cleanURL(u).String()

					return nil
				})

				r.SetMatcher(func(r *http.Request, i cassette.Request) bool {
					r.URL = cleanURL(r.URL)

					return cassette.DefaultMatcher(r, i)
				})

				return &http.Client{Transport: r}, r.Stop
			},
			input{
				context.Background(),
				"98fec8cb3815fe82d71cadd8acfb4dfa",
				"90210",
				"metric",
			},
			output{
				res: Result{
					Weather: []Weather{
						{
							Description: "clear sky",
						},
					},
					Main: Main{
						Temperature: 12.199999809265137,
						FeelsLike:   9.960000038146973,
					},
				},
			},
		},
		{
			"401",
			func(t *testing.T) (*http.Client, func() error) {
				r, err := recorder.New(path.Join("fixtures", "weather_401"))
				if err != nil {
					t.Fatalf("failed creating recorder %s", err)
				}

				return &http.Client{Transport: r}, r.Stop
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
			"500",
			func(_ *testing.T) (*http.Client, func() error) {
				gock.New("https://api.openweathermap.org").
					MatchParams(map[string]string{
						"zip":   "90210",
						"appid": "appID-500",
						"units": "metric",
					}).
					Get("/data/2.5/weather").
					Reply(http.StatusInternalServerError)

				client := &http.Client{Transport: &http.Transport{}}
				gock.InterceptClient(client)

				return client, func() error {
					if gock.IsDone() == false {
						return errors.New("gock has pending mocks")
					}

					return nil
				}
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
			"Error: invalid JSON",
			func(_ *testing.T) (*http.Client, func() error) {
				gock.New("https://api.openweathermap.org").
					MatchParams(map[string]string{
						"zip":   "90210",
						"appid": "appID-error",
						"units": "metric",
					}).
					Get("/data/2.5/weather").
					Reply(http.StatusOK).
					BodyString(`{"broken":`)

				client := &http.Client{Transport: &http.Transport{}}
				gock.InterceptClient(client)

				return client, func() error {
					if gock.IsDone() == false {
						return errors.New("gock has pending mocks")
					}

					return nil
				}
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
			func(_ *testing.T) (*http.Client, func() error) {
				gock.New("https://api.openweathermap.org").
					MatchParams(map[string]string{
						"zip":   "90210",
						"appid": "appID-err-req",
						"units": "metric",
					}).
					Get("/data/2.5/weather").
					Reply(http.StatusMovedPermanently).
					SetHeader("Location", "")

				client := &http.Client{Transport: &http.Transport{}}
				gock.InterceptClient(client)

				return client, func() error {
					if gock.IsDone() == false {
						return errors.New("gock has pending mocks")
					}

					return nil
				}
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
			func(_ *testing.T) (*http.Client, func() error) {
				return nil, func() error { return nil }
			},
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
			client, teardown := tt.setup(t)
			defer func() {
				if err := teardown(); err != nil {
					t.Fatalf("teardown failed %s", err)
				}
			}()

			actualRes, actualErr := requestWeather(tt.input.ctx, client, tt.input.appID, tt.input.zip, tt.input.units)

			if (actualErr != nil) != tt.output.withErr {
				t.Fatalf("expected error %t, actual %s", tt.output.withErr, actualErr)
			}

			if !cmp.Equal(tt.output.res, actualRes) {
				t.Fatalf("expected output do not match\n%s", cmp.Diff(tt.output.res, actualRes))
			}
		})
	}
}
