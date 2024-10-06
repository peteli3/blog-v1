.PHONY: templ
templ:
	templ generate --watch --proxy="http://localhost:6969" --cmd="go run ."

.PHONY: tailwind
tailwind:
	npx tailwindcss -i ./css/input.css -o ./css/output.css --watch

.PHONY: clean
clean:
	rm -f \
		components/*_templ.go \
		css/output.css \
		personal-v1.tar

.PHONY: build
build:
	templ generate
	npx tailwindcss -i ./css/input.css -o ./css/output.css --minify
	docker build -t personal-v1:latest --platform linux/amd64 .
	docker save personal-v1:latest --output personal-v1.tar
