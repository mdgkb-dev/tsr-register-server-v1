include .env

ifeq ($(OS),Windows_NT)
	migrations := .\migrations
	cli := .\cmd\cli
	main := .\cmd\server\main.go
else
	migrations := /migrations
	cli := cmd/cli/*.go
	main := cmd/server/main.go
endif

run: migrate set_git_hooks_dir
	reflex -r '\.go' -s -- sh -c "go run $(main)"

set_git_hooks_dir:
	git config core.hooksPath cmd/githooks/

run_cold:
	go run $(main)

migrate:
	go run $(main) -mode=migrate -action=migrate

migrate_create:
	go run $(main) -mode=migrate -action=create -name=${name}

migrate_rollback:
	go run $(migrates) rollback

drop_database:
	go run $(database) -action=dropDatabase

dump_from_remote:
	@./cmd/dump_pg.sh $(DB_NAME) $(DB_USER) $(DB_PASSWORD) $(DB_REMOTE_USER) $(DB_REMOTE_PASSWORD)

dump: dump_from_remote

deploy:
	./cmd/server/deploy.sh $(SERVER_NAME)

lint:
	./cmd/golangci.sh

read_logs:
	tail -f ./logs/info/_actual.log ./logs/errors/_actual.log

show_sql_error: 
	@./cmd/scripts/write_last_sql_error.sh | sed 's/\\//g'

# update_assister: 
# 	go get github.com/pro-assistance/pro-assister@${tag} 

update_assister: 
	@./cmd/scripts/update_assister.sh
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
	git flow feature start TSR-$n

git_deploy:
	git fetch . develop:master
	git push

########
#DEPLOY#
########

start:
	./cmd/server/deploy.sh $(DEPLOY_PATH)

kill:
	kill -9 `cat ./bundle/"$(shell basename "$(shell pwd)")".pid`
