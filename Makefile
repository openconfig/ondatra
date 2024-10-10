build_container:
	docker build -t ondatra-dn  .

run_container:
	docker run -v/home/dn/ondatra:/dn/ondatra \
		-v/home/dn/cheetah/prod/dnos_monolith/yangs/:/dn/yangs \
		-v/home/dn/ondatra/dn_cert.crt:/dn/dn_cert.crt \
		-v/var/tmp/dtest/host_fs/run/netns/test_ns:/var/run/netns/test_ns \
		-it --rm ondatra-dn:latest bash
