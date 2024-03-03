package controllers_test

import (
	"bytes"
	"elewa/pkg/controllers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCreateBook_Success(t *testing.T) {
	r := gin.Default()
	r.POST("/create", controllers.CreateBook)
	recorder := httptest.NewRecorder()

	jsonStr := []byte(`{
		"name":"Test Book",
		"price": 100, 
		"author":{
			"first_name": "Test",
			"last_name": "Author"}}
		`)

	req, err := http.NewRequest("POST", "/create", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	r.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"success":`
	if !bytes.Contains(recorder.Body.Bytes(), []byte(expected)) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			recorder.Body.String(), expected)
	}
}
