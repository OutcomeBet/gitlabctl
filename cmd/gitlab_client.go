package cmd

import (
	"fmt"

	"github.com/xanzy/go-gitlab"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func newBasicAuthClient(username, password, basehttpURL string) (*gitlab.Client, error) {
	gitlabClient, err := gitlab.NewBasicAuthClient(nil,
		basehttpURL,
		username,
		password,
	)
	if err != nil {
		return nil, err
	}
	return gitlabClient, nil
}

func newClient(privateToken, apihttpURL string) (*gitlab.Client, error) {
	gitlabClient := gitlab.NewClient(nil, privateToken)
	if err := gitlabClient.SetBaseURL(apihttpURL); err != nil {
		return nil, err
	}
	return gitlabClient, nil
}

func newOAuthClient(oAuthToken, apihttpURL string) (*gitlab.Client, error) {
	gitlabClient := gitlab.NewOAuthClient(nil, oAuthToken)
	if err := gitlabClient.SetBaseURL(apihttpURL); err != nil {
		return nil, err
	}
	return gitlabClient, nil
}

func newGitlabClient() (*gitlab.Client, error) {
	if get("USERNAME") != "" && get("PASSWORD") != "" && get("HTTP_URL") != "" {
		log.Println("basic auth login")
		gitlab, err := newBasicAuthClient(get("USERNAME"),
			get("PASSWORD"),
			get("HTTP_URL"),
		)
		if err != nil {
			return nil, err
		}
		return gitlab, nil
	}
	if get("PRIVATE_TOKEN") != "" && get("API_HTTP_URL") != "" {
		log.Println("private token login")
		gitlab, err := newClient(get("PRIVATE_TOKEN"), get("API_HTTP_URL"))
		if err != nil {
			return nil, err
		}
		return gitlab, nil
	}
	if get("OAUTH_TOKEN") != "" && get("API_HTTP_URL") != "" {
		log.Println("oauth login")
		gitlab, err := newOAuthClient(get("OAUTH_TOKEN"), get("API_HTTP_URL"))
		if err != nil {
			return nil, err
		}
		return gitlab, nil
	}
	return nil, fmt.Errorf("No clients were created. GITLAB variables may not be set properly.")
}

func get(e string) string {
	viper.SetEnvPrefix("gitlab") // will be automatically uppercased
	viper.BindEnv(e)             // will be automatically uppercased
	return viper.GetString(e)
}
