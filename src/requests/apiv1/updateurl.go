package apiv1

type UpdateUrlRequest struct {
	Label       string `form:"label" binding:"required"`
	Destination string `form:"destination" binding:"required"`
}
