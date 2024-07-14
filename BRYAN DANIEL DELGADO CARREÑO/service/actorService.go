package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/danie/APWII_PROYECTO/config"
	"github.com/danie/APWII_PROYECTO/models"
)

// ActorServiceImplementation implementa ActorService utilizando HTTP para interactuar con la API.
type ActorServiceImplementation struct{}

func (as *ActorServiceImplementation) GetActores() ([]models.ActorSerie, error) {
	url := fmt.Sprintf("%s/actores", config.URL_GO)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching actores: %v", err)
	}
	defer resp.Body.Close()

	var actores []models.ActorSerie
	if err := json.NewDecoder(resp.Body).Decode(&actores); err != nil {
		return nil, fmt.Errorf("error decoding response body: %v", err)
	}
	return actores, nil
}

func (as *ActorServiceImplementation) GetActorById(id string) (*models.ActorSerie, error) {
	url := fmt.Sprintf("%s/actores/%s", config.URL_GO, id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching actor with ID %s: %v", id, err)
	}
	defer resp.Body.Close()

	var actor models.ActorSerie
	if err := json.NewDecoder(resp.Body).Decode(&actor); err != nil {
		return nil, fmt.Errorf("error decoding response body: %v", err)
	}
	return &actor, nil
}

func (as *ActorServiceImplementation) CreateActor(input map[string]interface{}) (*models.ActorSerie, error) {
	url := fmt.Sprintf("%s/actores", config.URL_GO)
	actorJson, err := json.Marshal(input)
	if err != nil {
		return nil, fmt.Errorf("error encoding actor data: %v", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(actorJson))
	if err != nil {
		return nil, fmt.Errorf("error creating actor: %v", err)
	}
	defer resp.Body.Close()

	var createdActor models.ActorSerie
	if err := json.NewDecoder(resp.Body).Decode(&createdActor); err != nil {
		return nil, fmt.Errorf("error decoding response body: %v", err)
	}
	return &createdActor, nil
}

func (as *ActorServiceImplementation) UpdateActor(id string, input map[string]interface{}) (*models.ActorSerie, error) {
	url := fmt.Sprintf("%s/actores/%s", config.URL_GO, id)
	actorJson, err := json.Marshal(input)
	if err != nil {
		return nil, fmt.Errorf("error encoding actor data: %v", err)
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(actorJson))
	if err != nil {
		return nil, fmt.Errorf("error creating PUT request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error updating actor with ID %s: %v", id, err)
	}
	defer resp.Body.Close()

	var updatedActor models.ActorSerie
	if err := json.NewDecoder(resp.Body).Decode(&updatedActor); err != nil {
		return nil, fmt.Errorf("error decoding response body: %v", err)
	}
	return &updatedActor, nil
}

func (as *ActorServiceImplementation) DeleteActor(id string) (bool, error) {
	url := fmt.Sprintf("%s/actores/%s", config.URL_GO, id)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return false, fmt.Errorf("error creating DELETE request: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, fmt.Errorf("error deleting actor with ID %s: %v", id, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	return true, nil
}
