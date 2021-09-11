package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
)

func TestPactCrimeServiceCreateCrime(t *testing.T) {
	godotenv.Overload(".env")
	if os.Getenv("WRITE_PACTS") == "" {
		t.Skip("skipping Pact contracts; set WRITE_PACTS environment variable to enable")
	}

	pact := dsl.Pact{
		Consumer: "main",
		Provider: "crimeService",
	}
	defer pact.Teardown()
	body := `{
				"location_name":"region",
				"longitude":55,	
				"latitude":55,
				"description":"smth"
			}`
	pact.AddInteraction().
		UponReceiving("crimeService Create Crime").
		WithRequest(dsl.Request{
			Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json; charset=utf-8")},
			Method:  "POST",
			Path:    dsl.String("/crime"),
			Body:    body,
		}).
		WillRespondWith(dsl.Response{
			Status:  200,
			Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json; charset=utf-8")},
			Body:    `{
						"id": 3,
						"location_name": "region",
						"longitude": 55,
						"latitude": 55,
						"description": "smth",
						"image": "",
						"date": "2021-09-11T23:42:17.973503Z"
					}`,
		})

	if err := pact.Verify(func() error {
		u := fmt.Sprintf("http://localhost:%d/crime", pact.Server.Port)
		req, err := http.NewRequest("POST", u, strings.NewReader(body))
		if err != nil {
			return err
		}
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
		if _, err = http.DefaultClient.Do(req); err != nil {
			return err
		}
		return nil
	}); err != nil {
		t.Fatal(err)
	}

	pact.WritePact()
}