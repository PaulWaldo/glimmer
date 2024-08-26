package ui

import (
	"testing"

	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/test"
	"github.com/google/go-cmp/cmp"
	"github.com/mattn/go-mastodon"
	"github.com/stretchr/testify/assert"
)

func TestNewClientFromPrefs(t *testing.T) {
	type args struct {
		p appPrefs
	}
	server := "server"
	token := "token"
	id := "id"
	secret := "secret"

	tests := []struct {
		name string
		args args
		want *mastodon.Client
	}{
		{
			name: "places prefs into client",
			args: args{
				p: appPrefs{
					MastodonServer: binding.BindString(&server),
					AccessToken:    binding.BindString(&token),
					ClientID:       binding.BindString(&id),
					ClientSecret:   binding.BindString(&secret),
				}},
			want: &mastodon.Client{
				Config: &mastodon.Config{
					Server:       "server",
					ClientID:     "id",
					ClientSecret: "secret",
					AccessToken:  "token",
				}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClientFromPrefs(tt.args.p); !cmp.Equal(got, tt.want) {
				t.Errorf("NewClientFromPrefs(): %s ", cmp.Diff(got, tt.want))
			}
		})
	}
}

func TestClearCredentialsPrefs_RemovesPreferenceValues(t *testing.T) {
	ma := myApp{}
	ma.app = test.NewApp()
	ma.prefs = NewPreferences(ma.app)
	_ = ma.prefs.AccessToken.Set("aaa")
	_ = ma.prefs.ClientID.Set("bbb")
	_ = ma.prefs.ClientSecret.Set("ccc")
	_ = ma.prefs.MastodonServer.Set("ddd")

	ClearCredentialsPrefs()

	var x string
	var err error

	x, err = ma.prefs.AccessToken.Get()
	assert.NoError(t, err)
	assert.Equal(t, "", x, "AccessToken")

	x, err = ma.prefs.ClientID.Get()
	assert.NoError(t, err)
	assert.Equal(t, "", x, "ClientID")

	x, err = ma.prefs.ClientSecret.Get()
	assert.NoError(t, err)
	assert.Equal(t, "", x, "ClientSecret")

	x, err = ma.prefs.MastodonServer.Get()
	assert.NoError(t, err)
	assert.Equal(t, "", x, "MastodonServer")
}
