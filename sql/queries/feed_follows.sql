-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
  INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
  VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2
  )
  ON CONFLICT (user_id, feed_id)
  DO UPDATE SET updated_at = NOW()
  RETURNING *
)
SELECT
  inserted_feed_follow.*,
  feeds.name AS feed_name,
  users.name AS user_name
FROM inserted_feed_follow
INNER JOIN users ON users.id = inserted_feed_follow.user_id
INNER JOIN feeds ON feeds.id = inserted_feed_follow.feed_id;

-- name: GetFeedFollowsForUser :many
SELECT feed_follows.id,
  feed_follows.created_at,
  feed_follows.updated_at,
  feed_follows.user_id,
  feed_follows.feed_id,
  feeds.name AS feed_name,
  users.name AS user_name
FROM feed_follows
INNER JOIN users ON feed_follows.user_id = users.id
INNER JOIN feeds ON feed_follows.feed_id = feeds.id
WHERE feed_follows.user_id = $1;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows
USING users, feeds
WHERE feed_follows.user_id = users.id
  AND feed_follows.feed_id = feeds.id
  AND feed_follows.user_id = $1
  AND feed_follows.feed_id = $2;
