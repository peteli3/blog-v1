.PHONY: dev
dev:
	templ generate --watch --proxy="http://localhost:6969" --cmd="go run ."

.PHONY: clean
clean:
	rm components/*_templ.go
