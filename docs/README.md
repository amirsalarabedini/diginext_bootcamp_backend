# Following System API

## Endpoints

### Follow User

- **URL:** `/follow`
- **Method:** `POST`
- **Payload:**
  ```json
  {
    "follower_id": "user1",
    "followee_id": "user2"
  }
  ```

### Unfollow User

- **URL:** `/unfollow`
- **Method:** `POST`
- **Payload:**
  ```json
  {
    "follower_id": "user1",
    "followee_id": "user2"
  }
  ```

### Get Follower Count

- **URL:** `/followers?user_id=user1`
- **Method:** `GET`

### Get Common Followers

- **URL:** `/common_followers?user_id1=user1&user_id2=user2`
- **Method:** `GET`
