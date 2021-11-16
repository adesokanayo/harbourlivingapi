CREATE TABLE "users" (
                         "id" SERIAL PRIMARY KEY,
                         "title" varchar NOT NULL,
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
                            "desc" varchar NOT NULL,
                            "status" int
);

CREATE TABLE "events_type" (
                              "id" SERIAL PRIMARY KEY,
                              "desc" varchar NOT NULL,
                              "status" int
);

CREATE TABLE "categories" (
                            "id" SERIAL PRIMARY KEY,
                            "desc" varchar NOT NULL,
                            "image" varchar,
                            "status" int
);

CREATE TABLE "subcategories" (
                               "id" SERIAL PRIMARY KEY,
                               "category_id" int NOT NULL,
                               "desc" varchar NOT NULL,
                               "status" int
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
                          "start_date" timestamp NOT NULL,
                          "end_date" timestamp NOT NULL,
                          "venue" int NOT NULL,
                          "type" int NOT NULL,
                          "user_id" int NOT NULL,
                          "category" int NOT NULL,
                          "subcategory" int NOT NULL,
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
                         "url" varchar,
                         "virtual" boolean NOT NULL,
                         "rating" float DEFAULT (0.00),
                         "longitude" float, 
                         "latitude" float 
);

CREATE TABLE "tickets" (
                          "id" SERIAL PRIMARY KEY,
                          "name" varchar NOT NULL,
                          "event_id" int NOT NULL,
                          "price" float NOT NULL DEFAULT (0.00),
                          "quantity" int NOT NULL DEFAULT (0),
                          "status" int NOT NULL
);

CREATE TABLE "tickets_status" (
                                 "id" SERIAL PRIMARY KEY,
                                 "desc" varchar,
                                 "status" int NOT NULL
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
                                 "desc" varchar,
                                 "created_at" timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE "users" ADD FOREIGN KEY ("usertype") REFERENCES "users_type" ("id");

ALTER TABLE "subcategories" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "events" ADD FOREIGN KEY ("venue") REFERENCES "venues" ("id");

ALTER TABLE "events" ADD FOREIGN KEY ("type") REFERENCES "events_type" ("id");

ALTER TABLE "events" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "events" ADD FOREIGN KEY ("category") REFERENCES "categories" ("id");

ALTER TABLE "events" ADD FOREIGN KEY ("subcategory") REFERENCES "subcategories" ("id");

ALTER TABLE "events" ADD FOREIGN KEY ("status") REFERENCES "events_status" ("id");

ALTER TABLE "tickets" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id");

ALTER TABLE "tickets_status" ADD FOREIGN KEY ("status") REFERENCES "tickets_status" ("id");

ALTER TABLE "users_tickets" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "users_tickets" ADD FOREIGN KEY ("ticket_id") REFERENCES "tickets" ("id");

ALTER TABLE "sponsors" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "events_sponsors" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id");

ALTER TABLE "events_sponsors" ADD FOREIGN KEY ("sponsor_id") REFERENCES "sponsors" ("id");
ALTER TABLE "events_images" ADD FOREIGN KEY ("image_id") REFERENCES "images" ("id");
ALTER TABLE "events_images" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id");

ALTER TABLE "events_videos" ADD FOREIGN KEY ("video_id") REFERENCES "videos" ("id");
ALTER TABLE "events_videos" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id");

ALTER TABLE "hosts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "events_hosts" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id");

ALTER TABLE "events_hosts" ADD FOREIGN KEY ("host_id") REFERENCES "hosts" ("id");

CREATE INDEX ON "users" ("id");

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "users" ("usertype");

CREATE INDEX ON "events" ("id");

CREATE INDEX ON "events" ("venue");

CREATE INDEX ON "events" ("start_date");

CREATE INDEX ON "events" ("end_date");

CREATE INDEX ON "events" ("type");



INSERT  INTO users_type ( "desc", "status")
VALUES
    ('Attendee',1),
    ('Host',1),
    ('Sponsor',1);

INSERT  INTO events_type ("desc", "status")
VALUES
    ( 'Free',1),
    ( 'Paid',1),
    ( 'Special',1);

INSERT  INTO categories ( "desc", "status")
VALUES
    ('Education',1),
    ('Food ',1),
    ('Sport',1),
    ('Music',1),
    ('Business',1);


INSERT  INTO subcategories ( "desc","category_id", "status")
VALUES
    ( 'Sleeping',1,1),
    ( 'Eating Habits ',2,1),
    ( 'Marathon',3,1),
    ( 'Running',3,1),
    ( 'P&L',4,1);

INSERT  INTO venues ("name", "address", "postal_code","city","province","country_code",virtual)
VALUES
    ('Eko Hotels','34 TempleBy Way,54532 ','T2A6YG','Calgary','AB','CAN',false ),
    ('Eko Hotels','34 TempleBy Way,54532 ','T2A6YG','Calgary','AB','CAN',false),
    ('Eko Hotels','34 TempleBy Way,54532 ','T2A6YG','Calgary','AB','CAN',false),
    ('Eko Hotels','34 TempleBy Way 54532','T2A6YG','Calgary','AB','CAN',false),
    ('Eko Hotels','34 TempleBy Way,54532','T2A6YG','Calgary','AB','CAN',false);

INSERT  INTO tickets_status ("desc", "status")
VALUES
    ( 'Active',1),
    ( 'Cancelled',2),
    ( 'Renewed',3);


INSERT  INTO events_status ("desc" )
VALUES
    ( 'New'),
    ( 'Published'),
    ( 'Rejected'),
    ( 'Completed'),
    ( 'Deleted');



INSERT  INTO images ("name","url" )
VALUES
    ( 'face', 'https://www.gravatar.com/avatar/205e460b479e2e5b48aec07710c08d50');
   
INSERT  INTO videos ("name","url" )
VALUES
    ( 'face', 'https://www.gravatar.com/avatar/205e460b479e2e5b48aec07710c08d50');