CREATE TABLE "users" (
                         "id" SERIAL PRIMARY KEY,
                         "phone" varchar,
                         "first_name" varchar NOT NULL,
                         "last_name" varchar NOT NULL,
                         "email" varchar NOT NULL,
                         "username" varchar UNIQUE NOT NULL,
                         "password" varchar NOT NULL,
                         "password_changed_at" timestamp,
                         "usertype" int NOT NULL,
                         "avatar_url" varchar,
                         "date_of_birth" timestamp NOT NULL,
                         "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "users_type" (
                            "id" SERIAL PRIMARY KEY,
                            "description" varchar NOT NULL,
                            "status" int NOT NULL
);

CREATE TABLE "events_type" (
                              "id" SERIAL PRIMARY KEY,
                              "description" varchar NOT NULL,
                              "status" int NOT NULL
);

CREATE TABLE "categories" (
                            "id" SERIAL PRIMARY KEY,
                            "description" varchar NOT NULL,
                            "image" varchar,
                            "status" int NOT NULL,
                            "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "images" (
                              "id" SERIAL PRIMARY KEY,
                              "name" varchar,
                              "url" varchar NOT NULL
);

CREATE TABLE "videos" (
                              "id" SERIAL PRIMARY KEY,
                              "name" varchar,
                              "url" varchar NOT NULL
);

CREATE TABLE "events" (
                          "id" SERIAL PRIMARY KEY,
                          "title" varchar NOT NULL,
                          "description" varchar NOT NULL,
                          "banner_image" varchar NOT NULL,
                          "start_date" timestamptz NOT NULL,
                          "end_date" timestamptz NOT NULL,
                          "venue" int NOT NULL,
                          "type" int NOT NULL,
                          "user_id" int NOT NULL,
                          "category" int NOT NULL,
                          "ticket_id" int,
                          "recurring" boolean,
                          "status" int NOT NULL,
                          "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "venues" (
                         "id" SERIAL PRIMARY KEY,
                         "name" varchar NOT NULL,
                         "address" varchar,
                         "postal_code" varchar,
                         "city" varchar,
                         "province" varchar,
                         "country_code" varchar,
                         "venue_owner" int NOT NULL,
                         "banner_image" varchar,
                         "rating" float DEFAULT (0.00),
                         "longitude" float, 
                         "latitude" float, 
                         "status" int NOT NULL,
                         "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "tickets" (
                          "id" SERIAL PRIMARY KEY,
                          "name" varchar NOT NULL,
                          "event_id" int NOT NULL,
                          "price" float NOT NULL DEFAULT (0.00),
                          "currency" varchar NOT NULL,
                          "description" varchar
);


CREATE TABLE "users_tickets" (
                                "id" SERIAL PRIMARY KEY,
                                "user_id" int NOT NULL,
                                "ticket_id" int NOT NULL,
                                "quantity" int,
                                "total_cost" float,
                                "paid" boolean,
                                "payment_ref" varchar,
                                "payment_method" varchar,
                                "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "sponsors" (
                           "id" SERIAL PRIMARY KEY,
                           "user_id" int NOT NULL,
                           "display_name" varchar,
                           "avatar_url" varchar,
                           "short_bio" varchar,
                           "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "events_sponsors" (
                                  "id" SERIAL PRIMARY KEY,
                                  "event_id" int NOT NULL,
                                  "sponsor_id" int NOT NULL,
                                  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "events_videos" (
                                  "id" SERIAL PRIMARY KEY,
                                  "event_id" int NOT NULL,
                                  "video_id" int NOT NULL,
                                  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "events_images" (
                                  "id" SERIAL PRIMARY KEY,
                                  "event_id" int NOT NULL,
                                  "image_id" int NOT NULL,
                                  "created_at" timestamp NOT NULL DEFAULT (now())
);
CREATE TABLE "hosts" (
                        "id" SERIAL PRIMARY KEY,
                        "user_id" int NOT NULL,
                        "display_name" varchar,
                        "avatar_url" varchar,
                        "short_bio" varchar,
                        "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "events_hosts" (
                               "id" SERIAL PRIMARY KEY,
                               "event_id" int NOT NULL,
                               "host_id" int NOT NULL,
                               "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "events_status" (
                                 "id" SERIAL PRIMARY KEY,
                                 "description" varchar,
                                 "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "venues_status" (
                                 "id" SERIAL PRIMARY KEY,
                                 "description" varchar,
                                 "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "artists" (
                           "id" SERIAL PRIMARY KEY,
                           "user_id" int NOT NULL,
                           "display_name" varchar,
                           "avatar_url" varchar,
                           "short_bio" varchar,
                           "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "events_artists" (
                                  "id" SERIAL PRIMARY KEY,
                                  "event_id" int NOT NULL,
                                  "artist_id" int NOT NULL,
                                  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "events_favorites" (
                                  "id" SERIAL PRIMARY KEY,
                                  "event_id" int NOT NULL,
                                  "user_id" int NOT NULL,
                                  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "venues_favorites" (
                                  "id" SERIAL PRIMARY KEY,
                                  "venue_id" int NOT NULL,
                                  "user_id" int NOT NULL,
                                  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "promotions" (
                                  "id" SERIAL PRIMARY KEY,
                                  "event_id" int NOT NULL,
                                  "user_id" int NOT NULL,
                                  "plan_id" int NOT NULL,
                                  "start_date" timestamp NOT NULL, 
                                  "end_date" timestamp NOT NULL,
                                  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "plans" (
                                  "id" SERIAL PRIMARY KEY,
                                  "name" varchar NOT NULL,
                                  "description" varchar NOT NULL,
                                  "price" float NOT NULL,
                                  "no_of_days" int NOT NULL,
                                  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "news" (
                                  "id" SERIAL PRIMARY KEY,
                                  "title" varchar NOT NULL,
                                  "description" varchar NOT NULL,
                                  "feature_image" varchar ,
                                  "body" text NOT NULL,
                                  "user_id" int NOT NULL,
                                  "publish_date" timestamp NOT NULL,
                                  "tags" text ,
                                  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "events_views" (
                                  "id" SERIAL PRIMARY KEY,
                                  "event_id" int NOT NULL,
                                  "user_id" int NOT NULL,
                                  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "schedules" (
                                  "id" SERIAL PRIMARY KEY,
                                  "event_id" int NOT NULL,
                                  "date" timestamp NOT NULL,
                                  "start_time" time  NOT NULL,
                                  "end_time" time NOT NULL,
                                  "created_at" timestamp NOT NULL DEFAULT (now())
);


CREATE TABLE "dayplans" (
                                  "id" SERIAL PRIMARY KEY,
                                  "start_time" time NOT NULL,
                                  "end_time" time NOT NULL,
                                  "schedule_id" int NOT NULL,
                                  "title" varchar,
                                  "description" varchar,
                                  "performer_name" varchar,
                                  "created_at" timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE "users" ADD FOREIGN KEY ("usertype") REFERENCES "users_type" ("id");

ALTER TABLE "events" ADD FOREIGN KEY ("venue") REFERENCES "venues" ("id")
ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE "events" ADD FOREIGN KEY ("type") REFERENCES "events_type" ("id");

ALTER TABLE "events" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "events" ADD FOREIGN KEY ("category") REFERENCES "categories" ("id");

ALTER TABLE "events" ADD FOREIGN KEY ("status") REFERENCES "events_status" ("id");

ALTER TABLE "tickets" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id");

ALTER TABLE "users_tickets" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id")
ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE "users_tickets" ADD FOREIGN KEY ("ticket_id") REFERENCES "tickets" ("id")
ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE "sponsors" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "events_sponsors" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id")
ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE "events_sponsors" ADD FOREIGN KEY ("sponsor_id") REFERENCES "sponsors" ("id")
ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE "events_images" ADD FOREIGN KEY ("image_id") REFERENCES "images" ("id")
ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE "events_images" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id")
ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE "events_videos" ADD FOREIGN KEY ("video_id") REFERENCES "videos" ("id")
ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE "events_videos" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id")
ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE "hosts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "events_hosts" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id")
ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE "events_hosts" ADD FOREIGN KEY ("host_id") REFERENCES "hosts" ("id")
ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE "artists" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "events_artists" ADD FOREIGN KEY ("artist_id") REFERENCES "artists" ("id")
ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE "events_favorites" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id")
ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE "events_favorites" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id")
ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE "venues_favorites" ADD FOREIGN KEY ("venue_id") REFERENCES "venues" ("id")
ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE "venues_favorites" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id")
ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE "venues" ADD FOREIGN KEY ("status") REFERENCES "venues_status" ("id");

ALTER TABLE "promotions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "promotions" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id");

ALTER TABLE "promotions" ADD FOREIGN KEY ("plan_id") REFERENCES "plans" ("id");

ALTER TABLE "news" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "events_views" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id")
ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE "events_views" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id")
ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE "schedules" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id")
ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE "dayplans" ADD FOREIGN KEY ("schedule_id") REFERENCES "schedules" ("id")
ON UPDATE CASCADE ON DELETE CASCADE;


CREATE INDEX ON "users" ("id");

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "users" ("usertype");

CREATE INDEX ON "events" ("id");

CREATE INDEX ON "events" ("venue");

CREATE INDEX ON "events" ("start_date");

CREATE INDEX ON "events" ("end_date");

CREATE INDEX ON "events" ("type");

---======Seed Data=======-------

INSERT  INTO users_type ( "description", "status")
VALUES
    ('Attendee',1),
    ('Host',1),
    ('Sponsor',1),
    ('Admin',1);

INSERT  INTO events_type ("description", "status")
VALUES
    ( 'Free',1),
    ( 'Paid',1),
    ( 'Special',1);

INSERT  INTO categories ( "description", "status")
VALUES
    ('Education',1),
    ('Food ',1),
    ('Sport',1),
    ('Music',1),
    ('Business',1);

INSERT  INTO venues_status ("description" )
VALUES 
    ( 'Draft'),
    ( 'Published'),
    ( 'Approved'),
    ( 'Rejected'),
    ( 'Deleted');

INSERT  INTO venues ("name", "address", "postal_code","city","province","country_code","venue_owner","status")
VALUES
    ('Eko Hotels','34 TempleBy Way,54532 ','T2A6YG','Calgary','AB','CAN',1,1 ),
    ('Eko Hotels','34 TempleBy Way,54532 ','T2A6YG','Calgary','AB','CAN',1,1 ),
    ('Eko Hotels','34 TempleBy Way,54532 ','T2A6YG','Calgary','AB','CAN',1,1),
    ('Eko Hotels','34 TempleBy Way 54532','T2A6YG','Calgary','AB','CAN',1,1),
    ('Eko Hotels','34 TempleBy Way,54532','T2A6YG','Calgary','AB','CAN',1,1);



INSERT  INTO events_status ("description" )
VALUES 
    ( 'Draft'),
    ( 'Published'),
    ('Approved'),
    ( 'Rejected'),
    ( 'Completed');



INSERT  INTO images ("name","url" )
VALUES
    ( 'face', 'https://www.gravatar.com/avatar/205e460b479e2e5b48aec07710c08d50');
   
INSERT  INTO videos ("name","url" )
VALUES
    ( 'face', 'https://www.gravatar.com/avatar/205e460b479e2e5b48aec07710c08d50');

INSERT INTO users ("phone", "first_name", "last_name","username","password", "email","usertype","date_of_birth")
VALUES
(
  '08067648635','demofirstname','demolastname','demouser','$2a$10$cutZ4yVSagcNMjD8ofqP3eTT4iaOFu1MlZjM977E4ZJLdSrXkV48q','demouser@gmail.com','1','2022-01-01T14:00:12-00:00');

INSERT INTO events ("title", "description", "banner_image","start_date","end_date", "venue","type","user_id","category","status")
VALUES
(
  'demoTitle','demo description','https://demobannerurl','2022-01-01T14:00:12-00:00','2022-01-01T14:00:12-00:00','1','1','1','1','1');

INSERT INTO plans ("name", "description", "price", "no_of_days")
VALUES(
  'Standard Plan','Bess plan for low budget','20.00','10');


INSERT INTO promotions ("event_id", "user_id","plan_id","start_date","end_date")
VALUES(
  '1','1','1','2022-01-01T14:00:12-00:00','2022-03-01T14:00:12-00:00');
