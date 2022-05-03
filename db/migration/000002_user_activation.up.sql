ALTER TABLE users
ADD COLUMN "modified_at" timestamp NOT NULL DEFAULT (now()),
ADD COLUMN "activated" boolean, 
ADD COLUMN "activation_code" varchar(200);
