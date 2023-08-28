package services

import (
	"github.com/marcosvscampos/mt-transfer-worker/pkg/client"
	"github.com/marcosvscampos/mt-transfer-worker/pkg/model"
	"github.com/marcosvscampos/mt-transfer-worker/pkg/repository"
)

func EffectTransfer(tm *model.TransactionMessage) {
	transaction := repository.GetTransactionById(tm.TransactionId)
	accountId := client.FindAccountIdByFilter(transaction.DestinationUserID, transaction.DestinationAccountNumber)

	client.CallAuthorizer()
	client.DepositBalanceByAccountId(accountId, transaction.Amount)
	repository.UpdateTransactionStatusById(tm.TransactionId, "COMPLETADO")
}
