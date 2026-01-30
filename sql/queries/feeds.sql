-- name: CreateFeeds :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,    
    $2, 
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: ListFeeds :many
SELECT * FROM feeds
ORDER BY name;

-- name: FeedByUrl :one
SELECT * FROM feeds WHERE url = $1;

-- name: FeedById :one
SELECT * FROM feeds WHERE id = $1;
