package apiv1

type CreateUrlRequest struct {
	Label       string `form:"label" binding:"required"`
	Destination string `form:"destination" binding:"required"`
}
