.PHONY: setup

# Setup the development environment
setup: setup-air setup-templ setup-tailwind
	
setup-air:
	@echo "Setting up air..."
	@curl -sSfL https://raw.githubusercontent.com/air-verse/air/master/install.sh | sh -s -- -b ./tmp

setup-templ:
	@echo "Setting up templ..."
	@wget https://github.com/a-h/templ/releases/download/v0.2.707/templ_Linux_x86_64.tar.gz -O ./tmp/templ.tar.gz
	@tar xf ./tmp/templ.tar.gz -C ./tmp
	@rm ./tmp/templ.tar.gz

setup-tailwind:
	@echo "Setting up tailwind..."
	@npm i -g tailwindcss

# watch .go files for changes and recompiles the entire application
watch-air:
	@./tmp/air

# watch template files for changes, and recompiles them
watch-templ:
	@./tmp/templ generate -watch -proxy=http://localhost:3000

watch-tailwind:
	@tailwindcss --watch -i ./views/css/app.css -o ./public/styles.css
