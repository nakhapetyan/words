stop:
	docker stop `docker ps -a -q`

restart:
	-docker stop app
	docker run --rm -p 80:80 -v $PWD:/go/src/app --memory="2G" --memory-swap="0G" --name app -t app

rebuild:
	docker build -t app .

bash:
	docker exec -it app /bin/bash

dependence:
	go get -t ./...

update_json:
	~/go/bin/easyjson -all easyjson.go

ps:
	ps aux | grep -v "-" | grep private | grep ___go_build_

pprof_mem:
	go tool pprof --alloc_space --web http://localhost:6060/debug/pprof/heap
    # install brew install graphviz