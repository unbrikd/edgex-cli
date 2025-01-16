package utils

import (
	"fmt"
	"os/exec"
)

func GetDeviceId() (string, error) {
	ret, err := exec.Command("cdscmd", "-gdg", "1", "-pp").Output()
	if err != nil {
		return "", err
	}

	if string(ret) == "" {
		return "", fmt.Errorf("device ID is an empty string")
	}

	return string(ret), nil
}
