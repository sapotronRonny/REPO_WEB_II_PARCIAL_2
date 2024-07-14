package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/danie/APWII_PROYECTO/config"
	"github.com/danie/APWII_PROYECTO/models"
)

type SerieServiceImplementation struct{}

const (
	BaseURLSeries = config.URL_GO
)

func (s *SerieServiceImplementation) GetSeries() ([]models.Serie, error) {
	url := fmt.Sprintf("%s/series", BaseURLSeries)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching series: %v", err)
	}
	defer resp.Body.Close()

	var series []models.Serie
	if err := json.NewDecoder(resp.Body).Decode(&series); err != nil {
		return nil, fmt.Errorf("error decoding response body: %v", err)
	}
	return series, nil
}

func (s *SerieServiceImplementation) GetSerieById(id string) (*models.Serie, error) {
	url := fmt.Sprintf("%s/series/%s", BaseURLSeries, id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching serie with ID %s: %v", id, err)
	}
	defer resp.Body.Close()

	var serie models.Serie
	if err := json.NewDecoder(resp.Body).Decode(&serie); err != nil {
		return nil, fmt.Errorf("error decoding response body: %v", err)
	}
	return &serie, nil
}

func (s *SerieServiceImplementation) CreateSerie(input map[string]interface{}) (*models.Serie, error) {
	url := fmt.Sprintf("%s/series", BaseURLSeries)
	serieJson, err := json.Marshal(input)
	if err != nil {
		return nil, fmt.Errorf("error encoding serie data: %v", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(serieJson))
	if err != nil {
		return nil, fmt.Errorf("error creating serie: %v", err)
	}
	defer resp.Body.Close()

	var createdSerie models.Serie
	if err := json.NewDecoder(resp.Body).Decode(&createdSerie); err != nil {
		return nil, fmt.Errorf("error decoding response body: %v", err)
	}
	return &createdSerie, nil
}

func (s *SerieServiceImplementation) UpdateSerie(id string, input map[string]interface{}) (*models.Serie, error) {
	url := fmt.Sprintf("%s/series/%s", BaseURLSeries, id)
	serieJson, err := json.Marshal(input)
	if err != nil {
		return nil, fmt.Errorf("error encoding serie data: %v", err)
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(serieJson))
	if err != nil {
		return nil, fmt.Errorf("error creating PUT request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error updating serie with ID %s: %v", id, err)
	}
	defer resp.Body.Close()

	var updatedSerie models.Serie
	if err := json.NewDecoder(resp.Body).Decode(&updatedSerie); err != nil {
		return nil, fmt.Errorf("error decoding response body: %v", err)
	}
	return &updatedSerie, nil
}

func (s *SerieServiceImplementation) DeleteSerie(id string) (bool, error) {
	url := fmt.Sprintf("%s/series/%s", BaseURLSeries, id)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return false, fmt.Errorf("error creating DELETE request: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, fmt.Errorf("error deleting serie with ID %s: %v", id, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	return true, nil
}
