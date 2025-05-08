-- name: CreateFeedFollow :one

WITH new_row AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
        VALUES (
            $1,
            $2,
            $3,
            $4,
            $5
        )
    RETURNING *
)
SELECT new_row.*
FROM new_row
INNER JOIN users ON new_row.user_id = users.user_id
INNER JOIN feeds ON new_row.feed_id = feeds.feed_id;

-- name: GetFeedFollows :many
SELECT * FROM feed_follows;

-- name: GetFeedFollowsForUser :many
SELECT feeds.name FROM feed_follows
INNER JOIN users ON feed_follows.user_id = users.user_id
INNER JOIN feeds ON feed_follows.feed_id = feeds.feed_id
WHERE users.name = $1;