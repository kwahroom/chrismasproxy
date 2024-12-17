package chrismasify

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"text/template"
)

type ResponseData struct {
	Image   string
	OldBody string
}

var presentSVG = `<svg width="200" height="200" viewBox="0 0 200 200">
      <!-- The gift box -->
      <rect x="50" y="50" width="100" height="100" fill="#FFC080" rx="10" />
      
      <!-- The ribbon -->
      <rect x="70" y="30" width="60" height="20" fill="#FF0033" rx="5" />
      <rect x="70" y="130" width="60" height="20" fill="#FF0033" rx="5" />
      <line x1="100" y1="50" x2="100" y2="150" stroke="#FF0033" stroke-width="10" />
      
      <!-- The bow on top -->
      <path d="M 100,40 C 120,60 80,60 100,40 Z" fill="#FF0033" />
      <path d="M 100,40 C 120,20 80,20 100,40 Z" fill="#FF0033" />
    </svg>`

func ModifyResponse(r *http.Response) error {
	responseTemplate, err := template.New("ChrismasResponse").Parse(`<html><body>
{{.Image}}
<code>
{{.OldBody}}
</code>
</body></html>`)

	if err != nil {
		return err
	}

	// Read the response body
	oldBody, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	newBody := new(bytes.Buffer)

	if r.StatusCode == http.StatusNotFound {

		newBody.Write([]byte("Santa lost your present."))

	} else if r.StatusCode == http.StatusOK {
		responseTemplate.Execute(newBody, ResponseData{Image: presentSVG, OldBody: string(oldBody)})
	} else {
		newBody.Write(oldBody)
	}

	r.Header.Set("Content-Length", fmt.Sprintf("%d", newBody.Len()))
	r.Body = io.NopCloser(bytes.NewReader(newBody.Bytes()))

	return nil
}
