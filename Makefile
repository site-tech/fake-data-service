build:
	@echo "Building subgraph..."
	@go build -o bin/fake-data -v .
	@echo "done."

db:
	@echo "building docker images"
	@docker compose build
	@echo done.

du:
	@echo "starting docker images"
	@docker compose up

da: db du
