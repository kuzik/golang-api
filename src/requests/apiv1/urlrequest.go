package apiv1

type UrlRequest struct {
	Label       string `form:"label" binding:"required"`
	Destination string `form:"destination" binding:"required"`
	UserId      int    `form:"user_id"`
}
