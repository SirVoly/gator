-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, feed_id, user_id)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING *
)
SELECT inserted_feed_follow.*, f.name as feed, u.name as user
FROM inserted_feed_follow
INNER JOIN feeds f ON inserted_feed_follow.feed_id = f.id
INNER JOIN users u ON inserted_feed_follow.user_id = u.id;

-- name: RemoveFeedFollow :exec
DELETE
FROM feed_follows follow
USING feeds
WHERE feeds.id = follow.feed_id AND feeds.url = $1;

-- name: GetFeedFollowsForUser :many
SELECT ff.*, f.name as feed, u.name as user
FROM feed_follows ff
INNER JOIN feeds f ON ff.feed_id = f.id
INNER JOIN users u ON ff.user_id = u.id
WHERE u.name = $1;