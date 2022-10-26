package test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAllSuccess(t *testing.T) {
	db := setupTestDB()
	defer CloseTestDB(db)
	router := SetupRouter()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/project", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "Success", responseBody["status"])
	assert.Equal(t, "Berhasil find all", responseBody["data"].([]interface{})[0].(map[string]interface{})["name"])
	fmt.Println(responseBody)
}
