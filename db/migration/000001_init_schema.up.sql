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
                         "date_of_birth" timestamp NOT NULL,
                         "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "usertype" (
                            "id" SERIAL PRIMARY KEY,
                            "desc" varchar,
                            "status" int
);

CREATE TABLE "event_type" (
                              "id" SERIAL PRIMARY KEY,
                              "desc" varchar,
                              "status" int
);

CREATE TABLE "category" (
                            "id" SERIAL PRIMARY KEY,
                            "desc" varchar,
                            "status" int
);

CREATE TABLE "subcategory" (
                               "id" SERIAL PRIMARY KEY,
                               "desc" varchar,
                               "status" int
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
                          "status" varchar,
                          "image1" varchar,
                          "image2" varchar,
                          "image3" varchar,
                          "video1" varchar,
                          "video2" varchar,
                          "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "venue" (
                         "id" SERIAL PRIMARY KEY,
                         "name" varchar NOT NULL,
                         "address" varchar NOT NULL,
                         "city" varchar NOT NULL,
                         "province" varchar NOT NULL,
                         "country_code" varchar NOT NULL
);

ALTER TABLE "users" ADD FOREIGN KEY ("usertype") REFERENCES "usertype" ("id");

ALTER TABLE "events" ADD FOREIGN KEY ("venue") REFERENCES "venue" ("id");

ALTER TABLE "events" ADD FOREIGN KEY ("type") REFERENCES "event_type" ("id");

ALTER TABLE "events" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "events" ADD FOREIGN KEY ("category") REFERENCES "category" ("id");

ALTER TABLE "events" ADD FOREIGN KEY ("subcategory") REFERENCES "subcategory" ("id");

CREATE INDEX ON "users" ("id");

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "users" ("usertype");

CREATE INDEX ON "events" ("id");

CREATE INDEX ON "events" ("venue");

CREATE INDEX ON "events" ("start_date");

CREATE INDEX ON "events" ("end_date");

CREATE INDEX ON "events" ("type");


INSERT  INTO usertype ( "desc", "status")
VALUES
    ('Attendee',1),
    ('Host',1),
    ('Sponsor',1);

INSERT  INTO event_type ("desc", "status")
VALUES
    ( 'Free',1),
    ( 'Paid',1),
    ( 'Special',1);

INSERT  INTO category ( "desc", "status")
VALUES
    ('Education',1),
    ('Food ',1),
    ('Sport',1),
    ('Music',1),
    ('Arts',1),
    ('Business',1);


INSERT  INTO subcategory ( "desc", "status")
VALUES
    ( 'Sleeping',1),
    ( 'Eating ',1),
    ( 'Running',1);

INSERT  INTO venue ("name", "address", "city","province","country_code")
VALUES
    ('Eko Hotels','34 TempleBy Way,54532 ','Calgary','AB','CAN'),
    ('Eko Hotels','34 TempleBy Way,54532 ','Calgary','AB','CAN'),
    ('Eko Hotels','34 TempleBy Way,54532 ','Calgary','AB','CAN'),
    ('Eko Hotels','34 TempleBy Way 54532','Calgary','AB','CAN'),
    ('Eko Hotels','34 TempleBy Way,54532','Calgary','AB','CAN');
