
# ----------------------------
.PHONY: stop
stop:
	docker-compose stop


# ----------------------------
.PHONY: removeall
removeall:
	docker-compose stop
	docker system prune
	sudo rm -rf db/engine/mysql/var_lib_mysql/
	sudo rm -rf db/engine/mysql/data/
	sudo rm -rf db/engine/postgres/data
	sudo rm -rf db/tool/phpmyadmin/sessions/
	sudo rm -rf db/tool/pgadmin/root/
	sudo rm -rf backend/yummy/*.json
	sudo rm -rf backend/script/order/register/reserved/*.json


# ----------------------------
.PHONY: build
build:
	docker-compose build
	# docker-compose build --no-cache


# ----------------------------
.PHONY: up
up:
	docker-compose up

# ----------------------------
.PHONY: reup
reup: build up


# ----------------------------
.PHONY: restart
restart: stop up

# ----------------------------
.PHONY: resetall
resetall: removeall build up


# ----------------------------
# github.com/securego/gosec/v2/cmd/gosec
.PHONY: gosec
gosec:
	gosec -exclude=G303 ./backend/...
