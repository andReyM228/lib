package bus

type (
	LoginRequest struct {
		ChatID   int64  `json:"chat_id"`
		Password string `json:"password"`
	}

	LoginResponse struct {
		UserID int64 `json:"user_id"`
	}

	GetUserByIDRequest struct {
		ID int64 `json:"id"`
	}

	GetCarByIDRequest struct {
		ID    int64  `json:"id"`
		Token string `json:"token"`
	}
)
