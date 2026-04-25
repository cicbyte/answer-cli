package client

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cicbyte/answer-cli/internal/log"
	"go.uber.org/zap"
)

func init() {
	log.Logger = zap.NewNop()
}

// helper: 创建模拟 Answer API 服务器
func newTestServer(t *testing.T, code int, data any) *httptest.Server {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		body := map[string]any{"code": code, "msg": "Success.", "data": data}
		json.NewEncoder(w).Encode(body)
	}))
	return srv
}

func newTestErrorServer(t *testing.T, code int, msg string) *httptest.Server {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		body := map[string]any{"code": code, "msg": msg}
		json.NewEncoder(w).Encode(body)
	}))
	return srv
}

// --- unwrap ---

func TestUnwrap_Success(t *testing.T) {
	srv := newTestServer(t, 200, map[string]string{"key": "value"})
	defer srv.Close()

	c := NewClient(&Config{BaseURL: srv.URL})
	var result map[string]string
	err := c.GetJSON(t.Context(), "/test", nil, &result)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result["key"] != "value" {
		t.Fatalf("expected value, got %s", result["key"])
	}
}

func TestUnwrap_Code0(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{"code": 0, "msg": "ok", "data": map[string]string{"k": "v"}})
	}))
	defer srv.Close()

	c := NewClient(&Config{BaseURL: srv.URL})
	var result map[string]string
	err := c.GetJSON(t.Context(), "/test", nil, &result)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result["k"] != "v" {
		t.Fatalf("expected v, got %s", result["k"])
	}
}

func TestUnwrap_APIError(t *testing.T) {
	srv := newTestErrorServer(t, 404, "not found")
	defer srv.Close()

	c := NewClient(&Config{BaseURL: srv.URL})
	err := c.GetJSON(t.Context(), "/test", nil, nil)
	if err == nil {
		t.Fatal("expected error")
	}
	apiErr, ok := err.(*APIError)
	if !ok {
		t.Fatalf("expected *APIError, got %T", err)
	}
	if apiErr.Code != 404 {
		t.Fatalf("expected code 404, got %d", apiErr.Code)
	}
	if apiErr.Message != "not found" {
		t.Fatalf("expected 'not found', got %s", apiErr.Message)
	}
}

func TestUnwrap_Unauthorized(t *testing.T) {
	srv := newTestErrorServer(t, 401, "unauthorized")
	defer srv.Close()

	c := NewClient(&Config{BaseURL: srv.URL})
	err := c.GetJSON(t.Context(), "/test", nil, nil)
	if err == nil {
		t.Fatal("expected error")
	}
	if !IsUnauthorizedError(err) {
		t.Fatal("expected IsUnauthorizedError true")
	}
	if IsNotFoundError(err) {
		t.Fatal("expected IsNotFoundError false")
	}
}

func TestUnwrap_NotFound(t *testing.T) {
	srv := newTestErrorServer(t, 404, "not found")
	defer srv.Close()

	c := NewClient(&Config{BaseURL: srv.URL})
	err := c.GetJSON(t.Context(), "/test", nil, nil)
	if !IsNotFoundError(err) {
		t.Fatal("expected IsNotFoundError true")
	}
}

// --- IsUnauthorizedError / IsNotFoundError ---

func TestIsUnauthorizedError_Nil(t *testing.T) {
	if IsUnauthorizedError(nil) {
		t.Fatal("expected false for nil")
	}
}

func TestIsUnauthorizedError_NonAPIError(t *testing.T) {
	if IsUnauthorizedError(http.ErrHandlerTimeout) {
		t.Fatal("expected false for non-APIError")
	}
}

func TestIsNotFoundError_Nil(t *testing.T) {
	if IsNotFoundError(nil) {
		t.Fatal("expected false for nil")
	}
}

func TestIsNetworkError(t *testing.T) {
	if IsNetworkError(nil) {
		t.Fatal("expected false for nil")
	}
}

// --- APIError.Error ---

func TestAPIError_Error(t *testing.T) {
	e := &APIError{Code: 500, Message: "internal error"}
	got := e.Error()
	expected := "API error [500]: internal error"
	if got != expected {
		t.Fatalf("expected %q, got %q", expected, got)
	}
}

func TestAPIError_Error_ReasonOnly(t *testing.T) {
	e := &APIError{Code: 403, Message: ""}
	got := e.Error()
	expected := "API error [403]: "
	if got != expected {
		t.Fatalf("expected %q, got %q", expected, got)
	}
}

// --- NewClient ---

func TestNewClient_Defaults(t *testing.T) {
	c := NewClient(&Config{BaseURL: "http://localhost:8080"})
	if c.baseURL != "http://localhost:8080" {
		t.Fatalf("expected baseURL")
	}
	if c.token != "" {
		t.Fatal("expected empty token")
	}
	if c.Question == nil {
		t.Fatal("expected Question service")
	}
	if c.Answer == nil {
		t.Fatal("expected Answer service")
	}
	if c.Comment == nil {
		t.Fatal("expected Comment service")
	}
}

func TestNewClient_WithToken(t *testing.T) {
	c := NewClient(&Config{BaseURL: "http://localhost", Token: "tok"})
	if c.GetToken() != "tok" {
		t.Fatalf("expected tok, got %s", c.GetToken())
	}
}

func TestSetToken(t *testing.T) {
	c := NewClient(&Config{BaseURL: "http://localhost"})
	c.SetToken("new-tok")
	if c.GetToken() != "new-tok" {
		t.Fatalf("expected new-tok, got %s", c.GetToken())
	}
}

// --- QueryParams ---

func TestGetJSON_WithParams(t *testing.T) {
	var receivedPage string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		receivedPage = r.URL.Query().Get("page")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{"code": 200, "msg": "", "data": map[string]any{"count": 1}})
	}))
	defer srv.Close()

	c := NewClient(&Config{BaseURL: srv.URL})
	var result map[string]any
	err := c.GetJSON(t.Context(), "/test", map[string]string{"page": "3"}, &result)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if receivedPage != "3" {
		t.Fatalf("expected page=3, got %s", receivedPage)
	}
}

// --- Empty data field ---

func TestUnwrap_EmptyData(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{"code": 200, "msg": "", "data": nil})
	}))
	defer srv.Close()

	c := NewClient(&Config{BaseURL: srv.URL})
	var result map[string]string
	err := c.GetJSON(t.Context(), "/test", nil, &result)
	if err != nil {
		t.Fatalf("expected no error for nil data, got %v", err)
	}
}

func TestUnwrap_MissingDataField(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{"code": 200, "msg": ""})
	}))
	defer srv.Close()

	c := NewClient(&Config{BaseURL: srv.URL})
	var result map[string]string
	err := c.GetJSON(t.Context(), "/test", nil, &result)
	if err != nil {
		t.Fatalf("expected no error for missing data field, got %v", err)
	}
}
