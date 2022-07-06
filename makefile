ifeq ($(OS),Windows_NT)
	migrates := .\database\migrates
	cli := .\cmd\cli
	main := .\cmd\server\main.go
else
	migrates := database/migrates/*.go
	cli := cmd/cli/*.go
	main := cmd/server/main.go
endif

run:
	reflex -r '\.go' -s -- sh -c "go run $(main)"

run_cold:
	go run $(main)

migrate:
	go run $(main) -mode=migrate -action=migrate

migrate_create:
	go run $(main) -mode=migrate -action=create -name=${name}

seed:
	go run $(database) -mode=seed -action=migrate

seed_create:
	go run $(database) -mode=seed -action=create -name=${name}

migrate_rollback:
	go run $(migrates) rollback

full_migrate: drop_database migrate_init migrate seed

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