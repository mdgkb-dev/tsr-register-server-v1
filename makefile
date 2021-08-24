ifeq ($(OS),Windows_NT)
	database := .\database
	migrations := .\database\migrations
	cli := .\cmd\cli
else
	database := database/*.go
	migrations := database/migrations/*.go
	cli := cmd/cli/*.go
endif

run:
	reflex -r '\.go' -s -- sh -c "go run cmd/server/main.go"

full_migrate: drop_database migrate_init migrate seed

migrate_init:
	go run $(database) -action=init

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

drop_database:
	go run $(database) -action=dropDatabase

create_model:
	go run $(cli) -mode=model -action=create -name=${name}

create_api:
	go run $(cli) -mode=api -action=create -name=${name}
