package model

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Account struct {
	ID string `json:"id"`
}

func (a *Account) DeserializeResponse(resp http.Response) {

	jsonBody, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Um erro ocorreu ao recuperar a resposta da chamada ao serviço de contas:", err)
		return
	}

	fmt.Println(string(jsonBody))

	var accounts []Account
	err = json.Unmarshal([]byte(jsonBody), &accounts)
	if err != nil {
		fmt.Println("[Account#DeserializeResponse]Erro ao deserializar JSON:", err)
		return
	}

	a.ID = accounts[0].ID
}
