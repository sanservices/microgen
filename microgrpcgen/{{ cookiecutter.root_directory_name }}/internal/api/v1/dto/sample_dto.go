package dto

type GetUserRequestValidation struct {
	UserID uint32 `validate:"required"`
}
