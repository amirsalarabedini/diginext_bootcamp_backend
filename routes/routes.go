package routes

import (
    "following-system/controllers"
    "github.com/gorilla/mux"
)

func InitializeRoutes() *mux.Router {
    router := mux.NewRouter()
    router.HandleFunc("/follow", controllers.FollowUser).Methods("POST")
    router.HandleFunc("/unfollow", controllers.UnfollowUser).Methods("POST")
    router.HandleFunc("/followers", controllers.GetFollowerCount).Methods("GET")
    router.HandleFunc("/common_followers", controllers.GetCommonFollowers).Methods("GET")
    return router
}
