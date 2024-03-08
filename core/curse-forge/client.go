package curseforge

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"iceline-hosting.com/core/models"
)

const (
	baseURL = "https://api.curseforge.com"
)

var (
	ErrBadStatusCode = fmt.Errorf("bad status code")
)

type client struct {
	token string
}

func NewClient(token string) (client, error) {
	if token == "" {
		return client{}, errors.New("token cannot be empty")
	}
	return client{token}, nil
}

// Implement the modeDownloader interface from ../mods/mods.go. All methods should return an error that is not implemented yet.
func (c client) GetModByModID(ctx context.Context, id int) (models.Mod, error) {
	hc := http.Client{}

	endpoint := fmt.Sprintf("%s/v1/mods/%d", baseURL, id)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return models.Mod{}, err
	}

	req.Header.Set("x-api-key", c.token)

	resp, err := hc.Do(req)
	if err != nil {
		return models.Mod{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.Mod{}, ErrBadStatusCode
	}

	var data struct {
		D models.Mod `json:"data"`
	}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return models.Mod{}, err
	}

	return data.D, nil
}

func (c client) GetFeaturedMods(ctx context.Context, gameID int) (map[string]interface{}, error) {
	body := []byte(fmt.Sprintf(`{"gameId":%d}`, gameID))
	jsonBody := bytes.NewBuffer(body)
	hc := http.Client{}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, baseURL+"/v1/mods/featured", jsonBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("x-api-key", c.token)
	// req.Header.Set("x-api-key", c.token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	log.Println(req.Header)
	log.Println(req.URL)
	resp, err := hc.Do(req)
	if err != nil {
		return nil, err
	}
	log.Println(resp.Body)

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w - %s", ErrBadStatusCode, resp.Status)
	}
	// responseBytes, _ := ioutil.ReadAll(resp.Body)
	// log.Printf("Response Body: %s", responseBytes)
	// var data struct {
	// 	Featured []models.Mod `json:"featured"`
	// }
	data:=map[string]interface{}{}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	log.Println(data)
	return data, nil
}
func (c client) GetModsByGameID(ctx context.Context, gameID int) ([]models.Mod, error) {
	hc := http.Client{}
	finalEndpoint := fmt.Sprintf("%s/v1/mods/search?gameId=%d", baseURL, gameID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, finalEndpoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("x-api-key", c.token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	log.Println(req.URL)
	resp, err := hc.Do(req)
	if err != nil {
		return nil, err
	}
	log.Println(resp.Body)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w - %s", ErrBadStatusCode, resp.Status)
	}
	// responseBytes, _ := ioutil.ReadAll(resp.Body)
	// log.Printf("Response Body: %s", responseBytes)
	var data2 struct {
		D []models.Mod `json:"data"`
	}

	err = json.NewDecoder(resp.Body).Decode(&data2)
	if err != nil {
		return nil, err
	}
	log.Println(data2)
	return data2.D, nil
}

func (c client) GetGames(ctx context.Context) ([]models.CfGame, error) {
	client := http.Client{}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, baseURL+"/v1/games", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("x-api-key", c.token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w - %s", ErrBadStatusCode, resp.Status)
	}

	var data struct {
		D []models.CfGame `json:"data"`
	}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data.D, nil
}
