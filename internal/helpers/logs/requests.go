package helpers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
)

func PrintMultipartRequest(req *http.Request) {
	mr := multipart.NewReader(req.Body, req.Header.Get("Content-Type")[strings.Index(req.Header.Get("Content-Type"), "boundary=")+9:])
	defer req.Body.Close()

	for {
		part, err := mr.NextPart()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error reading part:", err)
			break
		}

		fmt.Println("--- Start Part ---")
		fmt.Println("Headers:")
		for key, value := range part.Header {
			fmt.Println("  ", key, ":", strings.Join(value, ", "))
		}
		fmt.Println("Content:")

		contentType := part.Header.Get("Content-Type")
		bodyBytes, err := io.ReadAll(part)
		if err != nil {
			fmt.Println("Error reading part content:", err)
			continue
		}

		switch contentType {
		case "text/plain":
			if len(bodyBytes) > 20 {
				fmt.Println(string(bodyBytes[:20]), "...")
			} else {
				fmt.Println(string(bodyBytes))
			}
		case "image/jpeg":
			fmt.Println("Image data (base64 encoded):")
			fmt.Println(string(base64.StdEncoding.EncodeToString(bodyBytes)))
		case "application/json":
			var jsonData interface{}
			err = json.Unmarshal(bodyBytes, &jsonData)
			if err != nil {
				fmt.Printf("Error parsing JSON: %v\n", err)
				fmt.Println(string(bodyBytes[:20]), "...")
			} else {
				fmt.Println(jsonData)
			}
		default:
			if len(bodyBytes) > 20 {
				fmt.Println(string(bodyBytes[:20]), "...")
			} else {
				fmt.Println(string(bodyBytes))
			}
		}
	}
	fmt.Println("--- End Part ---")
}

func PrintRequest(req *http.Request) {
	fmt.Println("--- START: SIMPLE REQUEST ---")

	fmt.Println("Method:", req.Method)
	fmt.Println("URL:", req.URL)
	fmt.Println("Headers:")
	for key, value := range req.Header {
		fmt.Println("  ", key, ":", strings.Join(value, ", "))
	}
	fmt.Println("Body:")
	if req.Body != nil {
		bodyBytes, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Println("Error reading body:", err)
		} else {
			fmt.Println(string(bodyBytes))
		}
		req.Body.Close() // Make sure to close the body
	} else {
		fmt.Println("<empty>")
	}
	fmt.Println("--- END: SIMPLE REQUEST ---")
}

func PrintResponse(resp *http.Response) {
	fmt.Println("--- START: SIMPLE RESPONSE ---")
	fmt.Println("Status:", resp.StatusCode)
	fmt.Println("Headers:")
	for key, value := range resp.Header {
		fmt.Println("  ", key, ":", strings.Join(value, ", "))
	}
	fmt.Println("Body:")
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
	} else {
		fmt.Println(string(bodyBytes))
	}
	fmt.Println("--- END: SIMPLE RESPONSE ---")
}
