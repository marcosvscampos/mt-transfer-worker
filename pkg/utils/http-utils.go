package utils

import (
	"fmt"
	"io"
	"net/http"
)

func PrintJsonResponse(resp http.Response) {
	json, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("[http-utils#PrintJsonResponse]Um erro ocorreu durante leitura da resposta:", err)
		return
	}

	fmt.Println(string(json))
}
