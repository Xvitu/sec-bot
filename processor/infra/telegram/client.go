package telegram

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
	"xvitu/sec-bot/infra/telegram/request"
	"xvitu/sec-bot/shared/env"
)

type Client struct {
	baseURL string
	client  *http.Client
}

func NewTelegramClient(cfg *env.Config) *Client {
	return &Client{
		baseURL: fmt.Sprintf("%s/bot%s", strings.TrimRight(cfg.TelegramUrl, "/"), cfg.TelegramBotToken),
		client:  &http.Client{Timeout: getTimeout(cfg)},
	}
}

func getTimeout(config *env.Config) time.Duration {
	timeout, err := strconv.Atoi(config.TeleGramClientTimeout)

	if err != nil || timeout <= 0 {
		timeout = 5
	}

	return time.Duration(timeout) * time.Second
}

func (t *Client) Post(request request.TelegramRequest) ([]byte, error) {
	endpoint := t.baseURL + request.Endpoint()

	jsonBody, err := request.Body()
	if err != nil {
		return nil, fmt.Errorf("erro ao serializar body: %w", err)
	}

	resp, err := t.client.Post(endpoint, "application/json", bytes.NewBuffer(jsonBody))

	return handleResponse(resp, err)
}

func handleResponse(resp *http.Response, err error) ([]byte, error) {
	if err != nil {
		return nil, fmt.Errorf("erro na requisição: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erro HTTP %d: %s", resp.StatusCode, string(bodyBytes))
	}

	return bodyBytes, nil
}
