package battlenet

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

// API exposes wows realms api
type API interface {
	FindConnectedRealm(realm string) (ConnectedRealm, error)
	ListAuctions(connectedRealmID int) ([]Auction, error)
}

const (
	tokenURL = "battle.net/oauth/token"
	apiURL   = "api.blizzard.com"
)

// API comment
type api struct {
	region string
	http   *http.Client
}

// NewAPI returns Realms API
func NewAPI(region string) API {
	authConfig := &clientcredentials.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		TokenURL:     fmt.Sprintf("https://%s.%s", region, tokenURL),
		AuthStyle:    oauth2.AuthStyleInHeader,
	}

	api := &api{
		region: region,
		http:   authConfig.Client(context.Background()),
	}

	return api
}

// FindConnectedRealms returns connected realm id for given realm
func (a *api) FindConnectedRealm(realm string) (ConnectedRealm, error) {
	connectedRealmsEndpoint := fmt.Sprintf("https://%s.%s/data/wow/realm/%s?locale=en_GB", a.region, apiURL, realm)
	var connectedRealm ConnectedRealm
	req, err := http.NewRequest("GET", connectedRealmsEndpoint, nil)
	if err != nil {
		return connectedRealm, err
	}

	req.Header.Add("Battlenet-Namespace", fmt.Sprintf("dynamic-%s", a.region))

	response, err := a.http.Do(req)
	if err != nil {
		return ConnectedRealm{}, err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return connectedRealm, errors.New("not a 200 response")
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return connectedRealm, err
	}

	err = json.Unmarshal(bytes, &connectedRealm)

	if err != nil {
		return connectedRealm, err
	}

	return connectedRealm, nil
}

// ListAuctions lists all current auctions of the connected realms
func (a *api) ListAuctions(connectedRealmID int) ([]Auction, error) {
	return make([]Auction, 0), nil
}
