build_container:
	docker build -t ondatra-dn  .

run_container:
	docker run -v /home/dn/ondatra:/dn/ondatra -it --rm ondatra-dn:latest bash

build:
	docker run -v /home/dn/ondatra:/dn/ondatra -it --rm ondatra-dn:latest /dn/ondatra/build.sh
