migrateup:
	migrate -path migration -database "postgresql://bucketeer:bucketeer_pass@localhost:5432/noteapp?sslmode=disable" -verbose up

migratedown:
	migrate -path migration -database "postgresql://bucketeer:bucketeer_pass@localhost:5432/noteapp?sslmode=disable" -verbose down

.PHONY: migrateup migratedown