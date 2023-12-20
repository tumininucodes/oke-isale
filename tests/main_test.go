package tests

import (
	"bytes"
	"file-upload-service/api/files/download"
	"file-upload-service/api/files/upload"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestFileUpload(t *testing.T) {
	r := gin.Default()
	r.POST("files/upload", upload.Upload)

	testRouter := httptest.NewServer(r)
	defer testRouter.Close()

	body := &bytes.Buffer{}
    writer := multipart.NewWriter(body)
    part, _ := writer.CreateFormFile("file", "test_file.txt")
    part.Write(make([]byte, 7 << 20))
    writer.Close()

	// Prepare a file upload request
	request, err := http.NewRequest("POST", testRouter.URL+"/files/upload", body)
	if err != nil {
		t.Fatalf("Failed to create request: %s", err.Error())
	}

	request.Header.Set("Content-Type", writer.FormDataContentType())

	// Perform the request
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		t.Fatalf("Failed to perform request: %s", err.Error())
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status %d; got %d", http.StatusOK, response.StatusCode)
		return
	}
}


func TestDownloadFile(t *testing.T) {
	r := gin.Default()
	r.GET("files/download/:filename", download.Download)

	// Create a temporary file for testing
	testFile, err := os.Create("storage/testfile.txt")
	if err != nil {
		t.Fatal("Error creating test file:", err)
	}
	testFile.Close()
	defer os.Remove("testfile.txt")

	testRouter := httptest.NewServer(r)
	defer testRouter.Close()

	// Perform a download request for the test file
	req, err := http.NewRequest("GET", testRouter.URL+"/files/download/testfile.txt", nil)
	if err != nil {
		t.Fatal("Error creating request:", err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d; got %d", http.StatusOK, w.Code)
	}

	// Check the response headers and content type
	expectedContentType := "application/octet-stream"
	if w.Header().Get("Content-Type") != expectedContentType {
		t.Errorf("Expected Content-Type %s; got %s", expectedContentType, w.Header().Get("Content-Type"))
	}

}
 