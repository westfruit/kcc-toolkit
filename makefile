PROJECT_NAME = kcc-toolkit
PROJECT_FOLDER = kcc
PK = 1bf94e7993753b61441d6f5b6f510829c5cfd267
PROJECT_GIT = http://$(PK)@git.bemular.net/$(PROJECT_FOLDER)/$(PROJECT_NAME).git

MESSAGE = $(m)
ifeq ("$(MESSAGE)","")
MESSAGE = submit
endif

# 构建
build:
	@echo -----------build-----------
	go build -o bin/$(PROJECT_NAME).exe main.go

# 清理
clean:
	@echo -----------clean-----------
	rm -rf ./log/*
	rm -rf ./bin/*
	rm -rf ./bin_release/*
	rm -rf ./$(PROJECT_NAME)
	rm -rf ./$(PROJECT_NAME).exe	
	rm -rf ./__debug_bin

# 运行
run:
	@echo -----------run-----------
	go run main.go

# 编译
compile:
	@echo -----------compile-----------
	@# go build -o bin/$(PROJECT_NAME).exe main.go
	gox -osarch=linux/amd64 -output=bin/$(PROJECT_NAME)

git:clean
	cd /d/GoCode/$(PROJECT_FOLDER)/$(PROJECT_NAME)
	git add .
	git commit -am '$(MESSAGE)' || echo '$(PROJECT_NAME) Commit failed. There is probably nothing to commit.'
	git pull $(PROJECT_GIT)
	git push $(PROJECT_GIT)

pull:
	git pull $(PROJECT_GIT)

# 使用方法：make git m='注释'