create table notes (
    "id" serial primary key,
    "title" varchar not null,
    "description" text,
    "check" bool not null,
    "created_at" timestamp not null default (now()),
    "updated_at" timestamp not null default (now())
)