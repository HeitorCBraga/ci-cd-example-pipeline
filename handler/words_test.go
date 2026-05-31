package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jhermesn/ci-cd-example-pipeline/handler"
)

func newWordsRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/words", handler.Words)
	return r
}

func TestWords_HappyPath(t *testing.T) {
	r := newWordsRouter()

	body := `{"text": "hello world"}`
	req := httptest.NewRequest(http.MethodPost, "/words", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", w.Code, w.Body.String())
	}

	var resp wordsResponse
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}

	expected := []string{"hello", "world"}
	if len(resp.Words) != len(expected) {
		t.Fatalf("expected %v, got %v", expected, resp.Words)
	}
	for i, word := range expected {
		if resp.Words[i] != word {
			t.Errorf("word[%d]: expected %q, got %q", i, word, resp.Words[i])
		}
	}
}

func TestWords_MultipleSpaces(t *testing.T) {
	r := newWordsRouter()

	body := `{"text": "  foo   bar  baz  "}`
	req := httptest.NewRequest(http.MethodPost, "/words", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var resp wordsResponse
	json.Unmarshal(w.Body.Bytes(), &resp)
	if len(resp.Words) != 3 {
		t.Errorf("expected 3 words, got %d: %v", len(resp.Words), resp.Words)
	}
}

func TestWords_EmptyText(t *testing.T) {
	r := newWordsRouter()

	body := `{"text": ""}`
	req := httptest.NewRequest(http.MethodPost, "/words", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400 for empty text, got %d", w.Code)
	}
}

func TestWords_MissingBody(t *testing.T) {
	r := newWordsRouter()

	req := httptest.NewRequest(http.MethodPost, "/words", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400 for missing body, got %d", w.Code)
	}
}

type wordsResponse struct {
	Words []string `json:"words"`
}
