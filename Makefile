# === Go Backend ===
run-backend:
	cd backend && air

test-backend:
	cd backend && go test ./...

.PHONY: gen-docs
gen-docs:
	@swag init -g ./backend/server/main.go -d cmd,internal && swag fmt

# === Next.js Frontend ===
run-frontend:
	cd frontend && npm run dev

build-frontend:
	cd frontend && npm run build

lint-frontend:
	cd frontend && npm run lint

# === Utils ===
start: run-backend run-frontend

.PHONY: run-backend build-backend test-backend \
        run-frontend build-frontend lint-frontend \
        start
