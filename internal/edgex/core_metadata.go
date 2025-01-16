package edgex

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
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

type AllDevicesResponse struct {
	Count   int      `json:"totalCount"`
	Devices []Device `json:"devices"`
}

type DeviceCreateRequest struct {
	ApiVersion string `json:"apiVersion"`
	Device     Device `json:"device"`
}

func (c *CoreMetadataService) DeviceExistsFromName(deviceName string) (bool, error) {
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

func (c *CoreMetadataService) DeviceExistsFromId(deviceId string) (bool, error) {
	allDevices, err := c.GetAllDevices(100)
	if err != nil {
		return false, err
	}

	for _, device := range *allDevices {
		if device.Id == deviceId {
			return true, nil
		}
	}

	return false, nil
}

func (c *CoreMetadataService) GetAllDevices(limit int) (*[]Device, error) {
	path := fmt.Sprintf("%s/device/all?offset=0&limit=%d", c.BaseURL.String(), limit)
	reqUUID := uuid.New()

	req, err := c.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Correlation-ID", reqUUID.String())

	allDevices := AllDevicesResponse{}
	res, err := c.client.Do(req, &allDevices)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %d", res.StatusCode)
	}

	return &allDevices.Devices, nil
}

func (c *CoreMetadataService) GetDeviceFromName(deviceName string) (*Device, error) {
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

func (c *CoreMetadataService) GetDeviceFromId(deviceId string) (*Device, error) {
	allDevices, err := c.GetAllDevices(100)
	if err != nil {
		return nil, err
	}

	for _, device := range *allDevices {
		if device.Id == deviceId {
			return &device, nil
		}
	}

	return nil, fmt.Errorf("device not found")
}

func (c *CoreMetadataService) DeleteDeviceFromName(deviceName string) error {
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

func (c *CoreMetadataService) DeleteDeviceFromId(deviceId string) error {
	allDevices, err := c.GetAllDevices(100)
	if err != nil {
		return err
	}

	for _, device := range *allDevices {
		if device.Id == deviceId {
			return c.DeleteDeviceFromName(device.Name)
		}
	}

	return fmt.Errorf("device not found")
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
		return "", fmt.Errorf("failed to create device: %d - %s", int(creationInfo[0]["statusCode"].(float64)), creationInfo[0]["message"])
	}

	return creationInfo[0]["id"].(string), nil
}
