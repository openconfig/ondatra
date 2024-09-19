build_container:
	docker build -t ondatra-dn  .

run_container:
	docker run -v /home/dn/ondatra:/dn/ondatra -v/home/dn/cheetah/prod/dnos_monolith/yangs/:/dn/yangs -it --rm ondatra-dn:latest bash
