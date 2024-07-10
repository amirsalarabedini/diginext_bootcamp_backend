package models

type User struct {
    ID             string   `json:"id"`
    Followers      []string `json:"followers"`
    Following      []string `json:"following"`
}

type FollowRequest struct {
    FollowerID string `json:"follower_id"`
    FolloweeID string `json:"followee_id"`
}
