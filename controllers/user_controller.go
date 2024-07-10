package controllers

import (
    "encoding/json"
    "following-system/models"
    "following-system/services"
    "following-system/utils"
    "net/http"
)

func FollowUser(w http.ResponseWriter, r *http.Request) {
    var data models.FollowRequest
    _ = json.NewDecoder(r.Body).Decode(&data)
    err := services.FollowUser(data.FollowerID, data.FolloweeID)
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }
    utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func UnfollowUser(w http.ResponseWriter, r *http.Request) {
    var data models.FollowRequest
    _ = json.NewDecoder(r.Body).Decode(&data)
    err := services.UnfollowUser(data.FollowerID, data.FolloweeID)
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }
    utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func GetFollowerCount(w http.ResponseWriter, r *http.Request) {
    userID := r.URL.Query().Get("user_id")
    count, err := services.GetFollowerCount(userID)
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }
    utils.RespondWithJSON(w, http.StatusOK, map[string]int{"follower_count": count})
}

func GetCommonFollowers(w http.ResponseWriter, r *http.Request) {
    userID1 := r.URL.Query().Get("user_id1")
    userID2 := r.URL.Query().Get("user_id2")
    commonFollowers, err := services.GetCommonFollowers(userID1, userID2)
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }
    utils.RespondWithJSON(w, http.StatusOK, commonFollowers)
}
