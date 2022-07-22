include .env

ifeq ($(OS),Windows_NT)
	migrates := .\database\migrates
	cli := .\cmd\cli
	main := .\cmd\server\main.go
else
	migrates := database/migrates/*.go
	cli := cmd/cli/*.go
	main := cmd/server/main.go
endif

run: migrate
	reflex -r '\.go' -s -- sh -c "go run $(main)"

run_cold:
	go run $(main)

migrate:
	go run $(main) -mode=migrate -action=migrate

migrate_create:
	go run $(main) -mode=migrate -action=create -name=${name}

migrate_rollback:
	go run $(migrates) rollback

dump_from_remote: migrate
	@./cmd/dump_pg.sh $(DB_NAME) $(DB_USER) $(DB_PASSWORD) $(DB_REMOTE_USER) $(DB_REMOTE_PASSWORD)

dump: dump_from_remote migrate

drop_database:
	go run database/*.go -action=dropDatabase

migrate_init:
	go run database/*.go -action=init


#####
#GIT#
#####

git_push: git_commit
	git push -u origin HEAD

git_commit:
	git pull origin develop
	git add .
	git commit -m "$m"

git_merge: git_push
	git checkout develop
	git pull
	git merge @{-1}
	git push

# example: make git_feature n=1
git_feature:
	git flow feature start PORTAL-$n