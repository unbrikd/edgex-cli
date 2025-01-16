package edgex

import (
	"fmt"
	"net/http"
)

type CoreMetadataService service

type Device struct {
	ApiVersion     string                 `json:"apiVersion"`
	Id             string                 `json:"id,omitempty"`
	Name           string                 `json:"name"`
	ServiceName    string                 `json:"serviceName"`
	ProfileName    string                 `json:"profileName"`
	Labels         []string               `json:"labels"`
	Description    string                 `json:"description"`
	AdminState     string                 `json:"adminState"`
	OperatingState string                 `json:"operatingState"`
	AutoEvents     []AutoEvent            `json:"autoEvents"`
	Protocols      map[string]interface{} `json:"protocols"`
}

type AutoEvent struct {
	Interval   string `json:"interval"`
	OnChange   bool   `json:"onChange"`
	SourceName string `json:"sourceName"`
}

type DeviceResponse struct {
	Device Device `json:"device"`
}

type DeviceCreateRequest struct {
	ApiVersion string `json:"apiVersion"`
	Device     Device `json:"device"`
}

func (c *CoreMetadataService) DeviceExists(deviceName string) (bool, error) {
	path := fmt.Sprintf("%s/device/check/name/%s", c.BaseURL.String(), deviceName)

	req, err := c.client.NewRequest("GET", path, nil)
	if err != nil {
		return false, err
	}

	res, err := c.client.Do(req, nil)
	if err != nil {
		return false, err
	}

	return res.StatusCode == http.StatusOK, nil
}

func (c *CoreMetadataService) GetDevice(deviceName string) (*Device, error) {
	path := fmt.Sprintf("%s/device/name/%s", c.BaseURL.String(), deviceName)

	req, err := c.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	deviceResponse := &DeviceResponse{}
	res, err := c.client.Do(req, deviceResponse)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %d", res.StatusCode)
	}

	return &deviceResponse.Device, nil
}

func (c *CoreMetadataService) DeleteDevice(deviceName string) error {
	path := fmt.Sprintf("%s/device/name/%s", c.BaseURL.String(), deviceName)

	req, err := c.client.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	res, err := c.client.Do(req, nil)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("status code: %d", res.StatusCode)
	}

	return nil
}

func (c *CoreMetadataService) CreateDevice(d Device) (string, error) {
	path := fmt.Sprintf("%s/device", c.BaseURL.String())

	payload := []DeviceCreateRequest{
		{
			ApiVersion: c.ApiVersion,
			Device:     d,
		},
	}

	req, err := c.client.NewRequest("POST", path, payload)
	if err != nil {
		return "", err
	}

	creationInfo := []map[string]interface{}{}
	res, err := c.client.Do(req, &creationInfo)
	if err != nil {
		return "", err
	}

	if res.StatusCode != http.StatusMultiStatus {
		return "", fmt.Errorf("status code: %d", res.StatusCode)
	}

	if int(creationInfo[0]["statusCode"].(float64)) != http.StatusCreated {
		return "", fmt.Errorf("failed to create device: %d - %s", creationInfo[0]["statusCode"], creationInfo[0]["message"])
	}

	return creationInfo[0]["id"].(string), nil
}
