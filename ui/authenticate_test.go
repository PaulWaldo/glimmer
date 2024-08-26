package ui

import (
	"testing"

	"fyne.io/fyne/v2/test"
	"github.com/PaulWaldo/glimmer/api"
	"github.com/stretchr/testify/assert"
)

func Test_myApp_LoadSecrets(t *testing.T) {
	app := test.NewApp()
	ma := &myApp{
		app: app,
	}
	expectedSecrets := api.Secrets{
		ApiKey:      "apikey",
		ApiSecret:   "apisecret",
		AccessToken: "accesstoken",
		OAuthToken:  "oauthtoken",
		OAuthSecret: "oauthsecret",
	}

	// Store secrets into preferences
	p := app.Preferences()
	p.SetString(PrefKeyAPIKey, expectedSecrets.ApiKey)
	p.SetString(PrefKeyAPISecret, expectedSecrets.ApiSecret)
	p.SetString(PrefKeyAccessToken, expectedSecrets.AccessToken)
	p.SetString(PrefKeyOauthToken, expectedSecrets.OAuthToken)
	p.SetString(PrefKeyOauthSecret, expectedSecrets.OAuthSecret)

	secrets := ma.LoadSecrets()

	assert.Equal(t, expectedSecrets.ApiKey, secrets.ApiKey)
	assert.Equal(t, expectedSecrets.ApiSecret, secrets.ApiSecret)
	assert.Equal(t, expectedSecrets.AccessToken, secrets.AccessToken)
	assert.Equal(t, expectedSecrets.OAuthSecret, secrets.OAuthSecret)
	assert.Equal(t, expectedSecrets.OAuthToken, secrets.OAuthToken)
}

func Test_apiInfoEntry_InfoHandling(t *testing.T) {
	e := apiInfoEntry{}
	e.makeUI()
	test.Type(e.apiKeyEntry, "abc123")
	test.Type(e.apiSecretEntry, "xyz789")

	e.form.OnSubmit()

	assert.Equal(t, "abc123", e.apiKeyEntry.Text)
}

// func Test_myApp_CreateSecrets(t *testing.T) {
// 	type fields struct {
// 		ma      *myApp
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		want    api.Secrets
// 		wantErr bool
// 	}{

// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ma := &myApp{
// 				app:    tt.fields.app,
// 				window: tt.fields.window,
// 			}
// 			got, err := ma.CreateSecrets()
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("myApp.CreateSecrets() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("myApp.CreateSecrets() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
