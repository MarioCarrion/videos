package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	var zip, appID, units string

	flag.StringVar(&appID, "appID", "", "OpenWeather token / appID")
	flag.StringVar(&zip, "zip", "90210", "Zip code to use")
	flag.StringVar(&units, "units", "imperial", "Units use: imperial/metric/standard")

	flag.Parse()

	// XXX Validations for arguments omitted

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	res, err := requestWeather(ctx, &http.Client{}, appID, zip, units)
	if err != nil {
		log.Fatalf("couldn't request weather: %s\n", err)
	}

	fmt.Printf("%+v\n", res)
}

type Weather struct {
	Description string `json:"description"`
}
type Main struct {
	Temperature float32 `json:"temp"`
	FeelsLike   float32 `json:"feels_like"`
}

type Result struct {
	Weather []Weather `json:"weather"`
	Main    Main      `json:"main"`
}

func requestWeather(ctx context.Context, client *http.Client, appID, zip, units string) (Result, error) {
	req, err := http.NewRequestWithContext(ctx,
		http.MethodGet,
		"https://api.openweathermap.org/data/2.5/weather",
		nil)
	if err != nil {
		return Result{}, err
	}

	url := req.URL.Query()
	url.Add("zip", zip)
	url.Add("appid", appID)
	url.Add("units", units)

	req.URL.RawQuery = url.Encode()

	res, err := client.Do(req)
	if err != nil {
		return Result{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		var errJSON struct {
			Message string `json:"message"`
		}

		if err := json.NewDecoder(res.Body).Decode(&errJSON); err != nil {
			return Result{}, err
		}

		return Result{}, fmt.Errorf(errJSON.Message)
	}

	var result Result

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return Result{}, err
	}

	return result, nil
}
