-- name: GetVideo :one
SELECT * FROM Videos
WHERE id = $1 LIMIT 1;

-- name: GetAllVideos :many
SELECT * FROM Videos;

--name: DeleteVideo :exec
DELETE  from Videos
where id =$1 ;

-- name: GetVideosByEvent :many
SELECT * FROM videos a, events_videos b
WHERE event_id = $1 and a.event_id = b.event_id ;

-- name: LinkVideoToEvent :exec
INSERT into events_videos (event_id, video_id)
VALUES($1,$2); 

-- name: CreateVideo :one
INSERT INTO Videos ( name, url  )
VALUES ($1,$2) RETURNING *;
