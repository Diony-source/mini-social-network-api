package follow

type FollowRequest struct {
	FolloweeID int64 `json:"followee_id" validate:"required,gt=0"`
}
