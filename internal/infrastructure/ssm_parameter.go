package infrastructure

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// Get service account key
func GetServiceAccountKey() ([]byte, error) {
	// Get parameter name from environment variable
	paramName := os.Getenv("SERVICE_ACCOUNT_KEY_PARAMETER")
	if paramName == "" {
		return nil, fmt.Errorf("SERVICE_ACCOUNT_KEY_PARAMETER environment variable is not set")
	}

	result, err := fetchParameterFromExtensionWithRetry(paramName)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Fetch parameter from Lambda Extension (with retry functionality)
func fetchParameterFromExtensionWithRetry(paramName string) ([]byte, error) {
	const maxRetries = 5
	const baseDelay = 100 * time.Millisecond

	for attempt := 0; attempt < maxRetries; attempt++ {
		result, err := fetchParameterFromExtension(paramName)
		if err == nil {
			return result, nil
		}

		// Retry for "not ready to serve traffic" errors
		if strings.Contains(err.Error(), "not ready to serve traffic") ||
			strings.Contains(err.Error(), "status 400") {
			if attempt < maxRetries-1 {
				delay := baseDelay * time.Duration(1<<attempt) // Exponential backoff
				time.Sleep(delay)
				continue
			}
		}

		return nil, err
	}

	return nil, fmt.Errorf("failed to fetch parameter after %d attempts", maxRetries)
}

// Fetch SSM parameter using Lambda Extension
func fetchParameterFromExtension(paramName string) ([]byte, error) {
	// Request to Lambda Extension endpoint (with SecureString decryption enabled)
	url := fmt.Sprintf("http://localhost:2773/systemsmanager/parameters/get?name=%s&withDecryption=true", paramName)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Header settings for Lambda Extension
	// Session token required for AWS Parameter Store Extension
	sessionToken := os.Getenv("AWS_SESSION_TOKEN")
	if sessionToken != "" {
		req.Header.Set("X-Aws-Parameters-Secrets-Token", sessionToken)
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch parameter from extension: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("extension returned status %d, body: %s", resp.StatusCode, string(body))
	}

	// Parse response (Parameter Store Extension response format)
	var response struct {
		Parameter struct {
			ARN              string        `json:"ARN"`
			DataType         string        `json:"DataType"`
			LastModifiedDate string        `json:"LastModifiedDate"`
			Name             string        `json:"Name"`
			Type             string        `json:"Type"`
			Value            string        `json:"Value"`
			Version          int           `json:"Version"`
			Labels           []string      `json:"Labels"`
			Tier             string        `json:"Tier"`
			Policies         []interface{} `json:"Policies"`
		} `json:"Parameter"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return []byte(response.Parameter.Value), nil
}
