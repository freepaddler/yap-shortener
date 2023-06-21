package handlers

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/freepaddler/yap-shortener/internal/app/store"
)

func TestHTTPHandler_Put(t *testing.T) {
	s := store.NewMemStore()
	h := NewHTTPHandler(s)

	tests := []struct {
		name     string
		body     string
		ct       string
		wantCode int
		wantCT   string
	}{
		{
			name:     "put empty body",
			body:     "",
			ct:       "text/plain",
			wantCode: http.StatusBadRequest,
			wantCT:   "",
		},
		{
			name:     "put invalid content-type",
			body:     "qweasdwe",
			ct:       "application/json",
			wantCode: http.StatusBadRequest,
			wantCT:   "",
		},
		{
			name:     "put success",
			body:     "http://qwe.asd/zxc",
			ct:       "text/plain",
			wantCode: http.StatusCreated,
			wantCT:   "text/plain",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqBody := strings.NewReader(tt.body)
			req := httptest.NewRequest(http.MethodPost, "/", reqBody)
			req.Header.Set("Content-Type", tt.ct)
			w := httptest.NewRecorder()
			h.Put(w, req)
			res := w.Result()
			defer res.Body.Close()

			require.Equal(t, tt.wantCode, res.StatusCode)
			if res.StatusCode == http.StatusCreated {
				assert.Contains(t, res.Header.Get("Content-Type"), tt.wantCT)
				resBody, err := io.ReadAll(res.Body)
				require.NoError(t, err)
				assert.NotEmpty(t, string(resBody))
			}
		})
	}
}

func TestHTTPHandler_Get(t *testing.T) {
	s := store.NewMemStore()
	h := NewHTTPHandler(s)
	storedUrl := "http://qwe.asd/zxc"
	storedShort := s.Put(storedUrl)
	fakeShort := "XXXXxxxx"

	tests := []struct {
		name     string
		reqShort string
		wantCode int
		wantUrl  string
	}{
		{
			name:     "Get failed",
			reqShort: fakeShort,
			wantCode: http.StatusBadRequest,
		},
		{
			name:     "Get succeeded",
			reqShort: storedShort,
			wantCode: http.StatusTemporaryRedirect,
			wantUrl:  storedUrl,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/"+tt.reqShort, nil)
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("id", tt.reqShort)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
			w := httptest.NewRecorder()
			h.Get(w, req)
			res := w.Result()
			defer res.Body.Close()

			require.Equal(t, tt.wantCode, res.StatusCode)
			if res.StatusCode == http.StatusTemporaryRedirect {
				assert.Equal(t, res.Header.Get("Location"), tt.wantUrl)
			}
		})
	}
}
