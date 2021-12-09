package apiv1

type URLRequest struct {
	Label       string `form:"label" binding:"required"`
	Destination string `form:"destination" binding:"required"`
	UserID      int    `form:"user_id"`
}
