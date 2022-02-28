create table notes (
    "id" serial primary key,
    "title" varchar not null,
    "description" text,
    "checked" BOOLEAN NOT NULL DEFAULT FALSE,
    "created_at" timestamp not null default (now()),
    "updated_at" timestamp not null default (now())
)