package transaction

import (
	"github.com/google/uuid"
)

type Transaction struct {
	ID                uuid.UUID `json:"id"`
	SenderAccountID   uuid.UUID `json:"senderaccountid" binding:"required"`
	ReceiverAccountID uuid.UUID `json:"receiveraccountid" binding:"required"`
	Amount            int64     `json:"amount" binding:"required"`
}
