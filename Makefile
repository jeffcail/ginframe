help:
	@echo "usage: make <option>"
	@echo "options and effects:"
	@echo "    mod    : go set env and go install package"
	@echo "    clean  : clean this mod file"
	@echo "    env    : install project env"

env:
	chmod +x server-env/mysql-redis-nacos/deploy_docker.sh
	server-env/mysql-redis-nacos/deploy_docker.sh

	chmod +x server-env/mysql-redis-nacos/init_env.sh
	server-env/msyql-redis-nacos/init_env.sh

mod:
	@if [ !-f go.mod ]; then go mod init github.com/jeffcail/ginframe;fi
	@go env -w GOPROXY=https://goproxy.cn,direct
	@go mod tidy
clean:
	rm -f go.mod
