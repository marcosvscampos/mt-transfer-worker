package repository

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/marcosvscampos/mt-transfer-worker/pkg/database"
	"github.com/marcosvscampos/mt-transfer-worker/pkg/model"
)

func GetTransactionById(id string) *model.Transaction {
	db := database.DB{}
	db.Connect()

	log.Println("Searching Transactions for ID", id)

	instance := db.Instance
	rows, err := instance.Query("SELECT amount, origin_user_id, origin_account_number, destination_user_id, destination_account_number, status FROM transactions WHERE id = ?", id)

	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer rows.Close()

	transaction := model.Transaction{}

	for rows.Next() {
		err := rows.Scan(&transaction.Amount,
			&transaction.OriginUserID,
			&transaction.OriginAccountNumber,
			&transaction.DestinationUserID,
			&transaction.DestinationAccountNumber,
			&transaction.Status)
		if err != nil {
			log.Fatal(err)
			return nil
		}
	}

	log.Println("Transaction Found ->", transaction)

	db.CloseConnection()
	return &transaction
}

func UpdateTransactionStatusById(id string, status string) {
	db := database.DB{}
	db.Connect()

	log.Println("Updating Transaction ID", id, "to Status:", status)

	instance := db.Instance
	result, err := instance.Exec("UPDATE transactions SET status = ? WHERE id = ?", status, id)
	if err != nil {
		log.Fatal(err)
		return
	}
	_, err = result.RowsAffected()
	if err != nil {
		log.Fatal(err)
		return
	}

	db.CloseConnection()
}
