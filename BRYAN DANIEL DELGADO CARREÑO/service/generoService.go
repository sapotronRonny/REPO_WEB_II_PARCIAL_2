package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/danie/APWII_PROYECTO/config"
	"github.com/danie/APWII_PROYECTO/models"
)

// GeneroServiceImplementation implementa GeneroService utilizando HTTP para interactuar con la API.
type GeneroServiceImplementation struct{}

func (gs *GeneroServiceImplementation) GetGeneros() ([]models.GeneroSerie, error) {
	url := fmt.Sprintf("%s/generos", config.URL_GO)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching generos: %v", err)
	}
	defer resp.Body.Close()

	var generos []models.GeneroSerie
	if err := json.NewDecoder(resp.Body).Decode(&generos); err != nil {
		return nil, fmt.Errorf("error decoding response body: %v", err)
	}
	return generos, nil
}

func (gs *GeneroServiceImplementation) GetGeneroById(id string) (*models.GeneroSerie, error) {
	url := fmt.Sprintf("%s/generos/%s", config.URL_GO, id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching genero with ID %s: %v", id, err)
	}
	defer resp.Body.Close()

	var genero models.GeneroSerie
	if err := json.NewDecoder(resp.Body).Decode(&genero); err != nil {
		return nil, fmt.Errorf("error decoding response body: %v", err)
	}
	return &genero, nil
}

func (gs *GeneroServiceImplementation) CreateGenero(input map[string]interface{}) (*models.GeneroSerie, error) {
	url := fmt.Sprintf("%s/generos", config.URL_GO)
	generoJson, err := json.Marshal(input)
	if err != nil {
		return nil, fmt.Errorf("error encoding genero data: %v", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(generoJson))
	if err != nil {
		return nil, fmt.Errorf("error creating genero: %v", err)
	}
	defer resp.Body.Close()

	var createdGenero models.GeneroSerie
	if err := json.NewDecoder(resp.Body).Decode(&createdGenero); err != nil {
		return nil, fmt.Errorf("error decoding response body: %v", err)
	}
	return &createdGenero, nil
}

func (gs *GeneroServiceImplementation) UpdateGenero(id string, input map[string]interface{}) (*models.GeneroSerie, error) {
	url := fmt.Sprintf("%s/generos/%s", config.URL_GO, id)
	generoJson, err := json.Marshal(input)
	if err != nil {
		return nil, fmt.Errorf("error encoding genero data: %v", err)
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(generoJson))
	if err != nil {
		return nil, fmt.Errorf("error creating PUT request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error updating genero with ID %s: %v", id, err)
	}
	defer resp.Body.Close()

	var updatedGenero models.GeneroSerie
	if err := json.NewDecoder(resp.Body).Decode(&updatedGenero); err != nil {
		return nil, fmt.Errorf("error decoding response body: %v", err)
	}
	return &updatedGenero, nil
}

func (gs *GeneroServiceImplementation) DeleteGenero(id string) (bool, error) {
	url := fmt.Sprintf("%s/generos/%s", config.URL_GO, id)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return false, fmt.Errorf("error creating DELETE request: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, fmt.Errorf("error deleting genero with ID %s: %v", id, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	return true, nil
}
