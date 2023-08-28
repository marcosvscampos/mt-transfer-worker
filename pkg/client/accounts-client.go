package client

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/marcosvscampos/mt-transfer-worker/pkg/model"
)

const moneyTransferServiceHost = "http://localhost:8081/money-transfer/api"

func FindAccountIdByFilter(userId string, number string) string {
	var callUrl = fmt.Sprintf("%v/accounts?userId=%s&number=%s", moneyTransferServiceHost, userId, number)

	log.Println("[FindAccountByFilter] Chamando serviço de contas: ", callUrl)

	resp, err := http.Get(callUrl)
	if err != nil {
		log.Fatal("Um erro ocorreu ao chamar serviço de contas:", err)
		return ""
	}

	account := model.Account{}
	account.DeserializeResponse(*resp)

	return account.ID
}

func DepositBalanceByAccountId(accountId string, amount float64) {
	var callUrl = fmt.Sprintf("%v/accounts/%s/balances", moneyTransferServiceHost, accountId)
	jsonStr := fmt.Sprintf(`{"balance": %v, "operation": "DEPOSIT"}`, amount)
	payload := []byte(jsonStr)

	req, err := http.NewRequest("PUT", callUrl, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Erro ao criar requisição:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro ao fazer a requisição:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Código de Status:", resp.Status)
}
