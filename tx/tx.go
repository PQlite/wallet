// Пакет tx, ймовірно, буде обробляти створення транзакцій, але підписання буде в crypto
package tx

type UnsignTransaction struct {
	From      string  `json:"from"`
	To        string  `json:"to"`
	Amount    float32 `json:"amount"`
	Timestamp int64   `json:"timestamp"`
	Nonce     int     `json:"nonce"`
}

type Transaction struct {
	From      string  `json:"from"`
	To        string  `json:"to"`
	Amount    float32 `json:"amount"`
	Timestamp int64   `json:"timestamp"`
	Nonce     int     `json:"nonce"`
	Signature string  `json:"signature"`
	PubKey    string  `json:"pubkey"`
}

func (t Transaction) GetUnsignTransaction() UnsignTransaction {
	return UnsignTransaction{
		From:      t.From,
		To:        t.To,
		Amount:    t.Amount,
		Timestamp: t.Timestamp,
		Nonce:     t.Nonce,
	}
}
