package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"catalina-support/backend/shared/auth"
)

type AIGreeting struct {
	Greeting              string `json:"greeting"`
	Current               string `json:"current"`
	TimeZone              string `json:"timeZone"`
	ModelID               string `json:"modelId"`
	ModelName             string `json:"modelName"`
	AccelerationMode      string `json:"accelerationMode"`
	EffectiveAcceleration string `json:"effectiveAcceleration"`
	DeviceName            string `json:"deviceName"`
	FallbackReason        string `json:"fallbackReason,omitempty"`
}

var (
	ErrAIManagerUnavailable = errors.New("settings.ai.managerUnavailable")
	ErrAILicenseRequired    = errors.New("settings.ai.licenseRequired")
	ErrAINVIDIAUnavailable  = errors.New("settings.ai.nvidiaUnavailable")
)

type AIInsufficientMemoryError struct {
	RequiredBytes  int64
	AvailableBytes int64
}

func (e *AIInsufficientMemoryError) Error() string { return "settings.ai.insufficientMemory" }

type AIModel struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Version         string `json:"version"`
	License         string `json:"license"`
	Source          string `json:"source"`
	Checksum        string `json:"checksum"`
	DownloadBytes   int64  `json:"downloadBytes"`
	RAMBytes        int64  `json:"ramBytes"`
	Installed       bool   `json:"installed"`
	Active          bool   `json:"active"`
	Healthy         bool   `json:"healthy"`
	Downloading     bool   `json:"downloading"`
	DownloadedBytes int64  `json:"downloadedBytes"`
	Phase           string `json:"phase"`
	Partial         bool   `json:"partial"`
	Error           string `json:"error,omitempty"`
}
type AIModelCatalog struct {
	Models                []AIModel `json:"models"`
	DiskAvailableBytes    int64     `json:"diskAvailableBytes"`
	AccelerationMode      string    `json:"accelerationMode"`
	EffectiveAcceleration string    `json:"effectiveAcceleration"`
	DeviceName            string    `json:"deviceName"`
	FallbackReason        string    `json:"fallbackReason,omitempty"`
	Restarting            bool      `json:"restarting"`
}
type aiModelAcceptanceRepository interface {
	AcceptAIModel(string, string, string, *int64, bool) error
}

func (s *Service) SetAIManager(baseURL, token string) {
	s.aiManagerURL = strings.TrimRight(baseURL, "/")
	s.aiManagerToken = token
}
func (s *Service) AIModels() (AIModelCatalog, error) {
	var out AIModelCatalog
	err := s.aiManager(http.MethodGet, "/models", nil, &out)
	return out, err
}
func (s *Service) DownloadAIModel(id string, accepted bool, actor *auth.Identity, setup bool) error {
	if !accepted {
		return ErrAILicenseRequired
	}
	catalog, err := s.AIModels()
	if err != nil {
		return err
	}
	var chosen *AIModel
	for i := range catalog.Models {
		if catalog.Models[i].ID == id {
			chosen = &catalog.Models[i]
			break
		}
	}
	if chosen == nil {
		return ErrAIModelRequired
	}
	var actorIDValue *int64
	if actor != nil {
		actorIDValue = actorID(*actor)
	}
	if repo, ok := s.repo.(aiModelAcceptanceRepository); ok {
		if err := repo.AcceptAIModel(chosen.ID, chosen.Version, chosen.Checksum, actorIDValue, setup); err != nil {
			return traducir(err)
		}
	}
	return s.aiManager(http.MethodPost, "/models/"+url.PathEscape(id)+"/download?acceptLicense=true", bytes.NewReader(nil), nil)
}
func (s *Service) ActivateAIModel(id string) error {
	acceleration := "auto"
	if repo, ok := s.repo.(aiRepository); ok {
		if row, err := repo.AI(); err == nil && row.AccelerationMode != "" {
			acceleration = row.AccelerationMode
		}
	}
	if catalog, err := s.AIModels(); err != nil {
		return err
	} else if catalog.AccelerationMode != acceleration {
		if err := s.setAIAcceleration(acceleration); err != nil {
			return err
		}
	}
	if err := s.aiManager(http.MethodPost, "/models/"+url.PathEscape(id)+"/activate", bytes.NewReader(nil), nil); err != nil {
		return err
	}
	if repo, ok := s.repo.(aiRepository); ok {
		now := s.now()
		if err := repo.UpdateAI(map[string]any{"mode": "local", "provider": "local", "base_url": "http://ai:8080", "model": id, "auth_type": "none", "auth_header": "", "tested_at": now, "acceleration_mode": acceleration}, nil); err != nil {
			return traducir(err)
		}
	}
	return nil
}
func (s *Service) DeactivateAIModel(id string) error {
	if err := s.aiManager(http.MethodPost, "/models/"+url.PathEscape(id)+"/deactivate", bytes.NewReader(nil), nil); err != nil {
		return err
	}
	if repo, ok := s.repo.(aiRepository); ok {
		if err := repo.UpdateAI(map[string]any{"model": "", "tested_at": nil}, nil); err != nil {
			return traducir(err)
		}
	}
	return nil
}
func (s *Service) DeleteAIModel(id string) error {
	return s.aiManager(http.MethodDelete, "/models/"+url.PathEscape(id), nil, nil)
}
func (s *Service) CancelAIModelDownload(id string) error {
	return s.aiManager(http.MethodDelete, "/models/"+url.PathEscape(id)+"/download", nil, nil)
}
func (s *Service) DeleteAIModelPartial(id string) error {
	return s.aiManager(http.MethodDelete, "/models/"+url.PathEscape(id)+"/partial", nil, nil)
}
func (s *Service) TestActiveAIModel() (AIGreeting, error) {
	if s.ia == nil {
		return AIGreeting{}, ErrAIManagerUnavailable
	}
	catalog, err := s.AIModels()
	if err != nil {
		return AIGreeting{}, err
	}
	model := ""
	modelName := ""
	for _, item := range catalog.Models {
		if item.Active && item.Healthy {
			model = item.ID
			modelName = item.Name
			break
		}
	}
	if model == "" {
		return AIGreeting{}, ErrAIManagerUnavailable
	}
	installation, err := s.repo.Installation()
	if err != nil {
		return AIGreeting{}, traducir(err)
	}
	zone := zonaHorariaDe(installation.TimeZone)
	location, err := time.LoadLocation(zone)
	if err != nil {
		location = time.UTC
		zone = "UTC"
	}
	current := s.now().In(location).Format("2006-01-02 15:04:05")
	greeting, err := s.ia.ProbarSaludo("http://ai:8080", model, installation.Language, current)
	if err != nil {
		return AIGreeting{}, ErrAIManagerUnavailable
	}
	return AIGreeting{Greeting: greeting, Current: current, TimeZone: zone, ModelID: model, ModelName: modelName, AccelerationMode: catalog.AccelerationMode, EffectiveAcceleration: catalog.EffectiveAcceleration, DeviceName: catalog.DeviceName, FallbackReason: catalog.FallbackReason}, nil
}

func (s *Service) setAIAcceleration(mode string) error {
	body, _ := json.Marshal(map[string]string{"mode": mode})
	return s.aiManager(http.MethodPut, "/acceleration", bytes.NewReader(body), nil)
}
func (s *Service) aiManager(method, path string, body io.Reader, out any) error {
	if s.aiManagerURL == "" || s.aiManagerToken == "" {
		return ErrAIManagerUnavailable
	}
	req, err := http.NewRequest(method, s.aiManagerURL+path, body)
	if err != nil {
		return ErrAIManagerUnavailable
	}
	req.Header.Set("Authorization", "Bearer "+s.aiManagerToken)
	timeout := 10 * time.Second
	if strings.HasSuffix(path, "/activate") || strings.HasSuffix(path, "/deactivate") || path == "/acceleration" {
		timeout = 310 * time.Second
	}
	client := http.Client{Timeout: timeout}
	resp, err := client.Do(req)
	if err != nil {
		return ErrAIManagerUnavailable
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		var keyed struct {
			Key string `json:"key"`
		}
		payload, _ := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
		_ = json.Unmarshal(payload, &keyed)
		if keyed.Key == "settings.ai.nvidiaUnavailable" {
			return ErrAINVIDIAUnavailable
		}
		if resp.StatusCode == http.StatusUnprocessableEntity {
			var failure struct {
				Key            string `json:"key"`
				RequiredBytes  int64  `json:"requiredBytes"`
				AvailableBytes int64  `json:"availableBytes"`
			}
			if json.Unmarshal(payload, &failure) == nil &&
				failure.Key == "settings.ai.insufficientMemory" && failure.RequiredBytes >= 0 && failure.AvailableBytes >= 0 {
				return &AIInsufficientMemoryError{RequiredBytes: failure.RequiredBytes, AvailableBytes: failure.AvailableBytes}
			}
		}
		return ErrAIManagerUnavailable
	}
	if out != nil && json.NewDecoder(io.LimitReader(resp.Body, 1<<20)).Decode(out) != nil {
		return ErrAIManagerUnavailable
	}
	return nil
}
