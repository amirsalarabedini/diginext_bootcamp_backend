package services

import (
    "errors"
    "following-system/models"
)

var users = make(map[string]*models.User)

func FollowUser(followerID, followeeID string) error {
    follower, ok := users[followerID]
    if !ok {
        follower = &models.User{ID: followerID}
        users[followerID] = follower
    }
    followee, ok := users[followeeID]
    if !ok {
        followee = &models.User{ID: followeeID}
        users[followeeID] = followee
    }
    follower.Following = append(follower.Following, followeeID)
    followee.Followers = append(followee.Followers, followerID)
    return nil
}

func UnfollowUser(followerID, followeeID string) error {
    follower, ok := users[followerID]
    if !ok {
        return errors.New("follower not found")
    }
    followee, ok := users[followeeID]
    if !ok {
        return errors.New("followee not found")
    }
    follower.Following = removeElement(follower.Following, followeeID)
    followee.Followers = removeElement(followee.Followers, followerID)
    return nil
}

func GetFollowerCount(userID string) (int, error) {
    user, ok := users[userID]
    if !ok {
        return 0, errors.New("user not found")
    }
    return len(user.Followers), nil
}

func GetCommonFollowers(userID1, userID2 string) ([]string, error) {
    user1, ok := users[userID1]
    if !ok {
        return nil, errors.New("user1 not found")
    }
    user2, ok := users[userID2]
    if !ok {
        return nil, errors.New("user2 not found")
    }
    return intersection(user1.Followers, user2.Followers), nil
}

func removeElement(slice []string, element string) []string {
    for i, v := range slice {
        if v == element {
            return append(slice[:i], slice[i+1:]...)
        }
    }
    return slice
}

func intersection(slice1, slice2 []string) []string {
    m := make(map[string]bool)
    var inter []string
    for _, v := range slice1 {
        m[v] = true
    }
    for _, v := range slice2 {
        if m[v] {
            inter = append(inter, v)
        }
    }
    return inter
}
