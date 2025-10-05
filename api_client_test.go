package gokwork_test

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/1e0cf/gokwork/openapi"
	"github.com/joho/godotenv"
)

func BasicAuthMiddleware() openapi.RequestEditorFn {
	return func(ctx context.Context, req *http.Request) error {
		req.SetBasicAuth("mobile_api", "qFvfRl7w")
		return nil
	}
}
func TestSignIn(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		t.Fatal("Error loading .env file")
	}

	email := os.Getenv("EMAIL")
	password := os.Getenv("PASSWORD")

	if email == "" || password == "" {
		t.Fatal("EMAIL and PASSWORD must be set in .env file")
	}

	client, err := openapi.NewClientWithResponses("https://api.kwork.ru", openapi.WithRequestEditorFn(BasicAuthMiddleware()))
	if err != nil {
		t.Fatalf("Failed to create API client: %v", err)
	}

	testCases := []struct {
		name               string
		email              string
		password           string
		gRecaptchaResponse string
		expectCaptcha      bool
	}{
		{
			name:          "Successful Sign In",
			email:         email,
			password:      password,
			expectCaptcha: false,
		},
		// {
		// 	name:          "Captcha Required",
		// 	email:         email,
		// 	password:      "wrongpassword", // Assuming this will trigger captcha
		// 	expectCaptcha: true,
		// },
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			body := openapi.SignInFormdataRequestBody{
				Login:    tc.email,
				Password: tc.password,
			}
			if tc.gRecaptchaResponse != "" {
				body.GRecaptchaResponse = &tc.gRecaptchaResponse
			}

			resp, err := client.SignInWithFormdataBodyWithResponse(context.Background(), body)
			if err != nil {
				t.Fatalf("Failed to send request: %v", err)
			}
			if resp.StatusCode() != http.StatusOK {
				t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode())
			}
			fmt.Printf("%v", resp.StatusCode())

			// if tc.expectCaptcha {
			// 	if resp.JSON200 == nil || resp.JSON200.ErrorCode == nil || *resp.JSON200.ErroxrCode != 118 {
			// 		t.Errorf("Expected error code 118 for captcha, but got: %v", resp.JSON200)
			// 	}
			// } else {
			// 	if resp.JSON200 == nil || resp.JSON200.Token == nil || *resp.JSON200.Token == "" {
			// 		t.Errorf("Expected a token, but got none. Response: %v", resp.JSON200)
			// 	}
			// }
		})
	}
}
