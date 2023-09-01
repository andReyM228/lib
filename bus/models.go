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
)
