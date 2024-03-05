run_parallel:
		curl -X POST -H "Content-Type: application/json" --silent \
		-d '{ "Email": "test69@69.com", "Username":"smoketrdees69", "Password":"smoketredes69"}' \
		http://localhost:3000/api/auth/login

clear_port:
	sudo kill -9 `sudo lsof -t -i:3000`
