ENVIRONMENT:=tmp/environment

$(ENVIRONMENT):
	$(info --- Launching environment)
	docker-compose up --build -d
	mkdir -p $(@D)
	touch $@

.PHONY: clean
clean:
	$(info --- Cleaning everything)
	@docker-compose down
	-@docker rm -f $(NAME)-http
	rm -rf $(ENVIRONMENT)

.PHONY: run
run: $(ENVIRONMENT)
	$(info --- Let the battle begin!)
	@docker-compose run --rm \
		app \
		/go/bin/battleship -input /data/input.txt

.PHONY: test
test: $(ENVIRONMENT)
	@docker-compose run --rm \
		app \
		go test -cover `go list ./... | grep -v vendor`
