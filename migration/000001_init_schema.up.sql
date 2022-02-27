create table notes (
    "id" int primary key,
    "title" text not null,
    "description" text not null,
    "created_at" timestamp not null default (now()),
    "updated_at" timestamp not null default (now())
)