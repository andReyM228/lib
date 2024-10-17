package bus

type (
	LoginRequest struct {
		ChatID   int64  `json:"chat_id" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	LoginResponse struct {
		UserID int64 `json:"user_id" validate:"required"`
	}

	GetUserByIDRequest struct {
		ID int64 `json:"id"`
	}

	GetCarByIDRequest struct {
		ID    int64  `json:"id"`
		Token string `json:"token"`
	}

	BuyCarRequest struct {
		ChatID int64  `json:"chat_id"`
		CarID  int64  `json:"car_id"`
		TxHash string `json:"tx_hash"`
	}

	IssueRequest struct {
		ToAddress string `json:"to_address" validate:"required"`
		Amount    int64  `json:"amount" validate:"required"`
		Memo      string `json:"memo" validate:"required"`
	}

	WithdrawRequest struct {
		ToAddress string `json:"to_address" validate:"required"`
		Amount    int64  `json:"amount" validate:"required"`
		Memo      string `json:"memo" validate:"required"`
	}

	TxResponse struct {
		TxHash string `json:"tx_hash"`
	}
)
