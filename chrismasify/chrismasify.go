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
	Message string
}

// WriteChrismasResponse Write the html reponse to a http code
func WriteChrismasResponse(newBody io.Writer, oldBody *[]byte, status int) {
	responseTemplate, err := template.New("ChrismasResponse").Parse(`<html><body>
{{.Image}}
<h1>{{.Message}}</h1>
</body></html>`)

	if err != nil {
		panic(err)
	}

	if status == http.StatusGatewayTimeout {
		responseTemplate.Execute(newBody, ResponseData{Image: lostSanta, Message: "504 - Santa got lost and could not find your present on time."})
	} else if status == http.StatusServiceUnavailable {
		responseTemplate.Execute(newBody, ResponseData{Image: drunkSanta, Message: "503 - Drunk Santa passed out and cannot deliver your present."})
	} else if status == http.StatusBadGateway {
		responseTemplate.Execute(newBody, ResponseData{Image: wrongPole, Message: "502 - You landed at the south pole instead of the north pole!"})
	} else if status == http.StatusNotImplemented {
		responseTemplate.Execute(newBody, ResponseData{Image: confusedSanta, Message: "501 - Santa has not yet learned how to do this."})
	} else if status == http.StatusInternalServerError {
		responseTemplate.Execute(newBody, ResponseData{Image: crushedSanta, Message: "500 - Oh No! Santa got crushed by the presents!"})
	} else if status == http.StatusTeapot {
		responseTemplate.Execute(newBody, ResponseData{Image: teaSanta, Message: "418 - Santa is drinking tea."})
	} else if status == http.StatusNotFound {
		responseTemplate.Execute(newBody, ResponseData{Image: noPresentSanta, Message: "404 - Oh no Santa lost your present!"})
	} else if status == http.StatusForbidden {
		responseTemplate.Execute(newBody, ResponseData{Image: fuckOff, Message: "403 - You are on the list. The naughty list. That means no present for you."})
	} else if status == http.StatusPaymentRequired {
		responseTemplate.Execute(newBody, ResponseData{Image: brokeSanta, Message: "402 - Santa is broke cause he spent it all on pepparkakor the fat bastard."})
	} else if status == http.StatusUnauthorized {
		responseTemplate.Execute(newBody, ResponseData{Image: naughtyList, Message: "401 - Sorry you are not on the list."})
	} else if status == http.StatusBadRequest {
		responseTemplate.Execute(newBody, ResponseData{Image: badSanta, Message: "400 - Bad Santa."})
	} else {
		newBody.Write(*oldBody)
	}
}

func ModifyResponse(r *http.Response) error {

	// Read the response body
	oldBody, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	newBody := new(bytes.Buffer)

	WriteChrismasResponse(newBody, &oldBody, r.StatusCode)

	r.Header.Set("Content-Length", fmt.Sprintf("%d", newBody.Len()))
	r.Body = io.NopCloser(bytes.NewReader(newBody.Bytes()))

	return nil
}
