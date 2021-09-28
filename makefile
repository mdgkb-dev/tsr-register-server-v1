ifeq ($(OS),Windows_NT)
	database := .\database
	migrations := .\database\migrations
	cli := .\cmd\cli
	main := .\cmd\server\main.go
else
	database := database/*.go
	migrations := database/migrations/*.go
	cli := cmd/cli/*.go
	main := cmd/server/main.go
endif

run:
	reflex -r '\.go' -s -- sh -c "go run $(main)"

run_cold:
	go run $(main)

migrate:
	go run $(database) -mode=migration -action=migrate

migrate_create:
	go run $(database) -mode=migration -action=create -name=${name}

seed:
	go run $(database) -mode=seed -action=migrate

seed_create:
	go run $(database) -mode=seed -action=create -name=${name}

migrate_rollback:
	go run $(migrations) rollback

full_migrate: drop_database migrate_init migrate seed

drop_database:
	go run database/*.go -action=dropDatabase

migrate_init:
	go run database/*.go -action=init
