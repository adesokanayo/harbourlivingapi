--name: GetSponsors: many
SELECT * FROM sponsor
where user_id = $1;

--name: CreateSponsor: one
INSERT INTO sponsor ( user_id )
VALUES ($1) RETURNING *;

--name: DeleteSponsor: exec
DELETE  from sponsor where
user_id = $1;

--name: GetSponsorsForEvents: many
SELECT * FROM events_sponsor
where sponsor_id = $1 and event_id = $2;