.PHONY: setup setup-fe setup-be start-be start-fe dev

setup: setup-be setup-fe
	@echo "Full stack dev environment ready"

# ---------- BACKEND ----------
setup-be:
	@echo "🔧 Setting up Go backend"
	mkdir -p backend/internal/handlers backend/internal/utils
	cd backend && go mod init backend
	cd backend && go get github.com/joho/godotenv

	echo 'package main\r\n\r\nimport (\r\n\t"backend/internal/handlers"\r\n\t"backend/internal/utils"\r\n\t"log"\r\n\t"net/http"\r\n)\r\n\r\nfunc main() {\r\n\tenvCfg := utils.SetupEnvCfg()\r\n\r\n\tcfg := &handlers.APIConfig{\r\n\t\tEnv: envCfg,\r\n\t}\r\n\r\n\tmux := http.NewServeMux()\r\n\r\n\t// System\r\n\tmux.HandleFunc("/health", cfg.Health)\r\n\r\n\tsrv := &http.Server{\r\n\t\tAddr:    ":" + cfg.Env.Port,\r\n\t\tHandler: mux,\r\n\t}\r\n\r\n\tlog.Printf("Serving on: http://localhost:%s/\\n", cfg.Env.Port)\r\n\tlog.Fatal(srv.ListenAndServe())\r\n}' > backend/main.go
	
	echo 'package handlers\r\n\r\nimport (\r\n\t"net/http"\r\n)\r\n\r\nfunc (cfg *APIConfig) Health(w http.ResponseWriter, r *http.Request) {\r\n\tw.WriteHeader(http.StatusOK)\r\n\tw.Write([]byte("OK"))\r\n}' > backend/internal/handlers/health.go

	echo 'package handlers\r\n\r\nimport "backend/internal/utils"\r\n\r\ntype APIConfig struct {\r\n\tEnv *utils.EnvCfg\r\n}' > backend/internal/handlers/apiCfg.go

	echo 'package utils\r\n\r\nimport (\r\n\t"log"\r\n\t"os"\r\n\r\n\t"github.com/joho/godotenv"\r\n)\r\n\r\ntype EnvCfg struct {\r\n\tPort string\r\n}\r\n\r\nfunc SetupEnvCfg() *EnvCfg {\r\n\tgodotenv.Load(".env")\r\n\r\n\tport := os.Getenv("PORT")\r\n\tif port == "" {\r\n\t\tlog.Fatal("PORT environment variable is not set")\r\n\t}\r\n\r\n\treturn &EnvCfg{\r\n\t\tPort: port,\r\n\t}\r\n}' > backend/internal/utils/env.go

	echo 'PORT=8080' > backend/.env
	echo '.env' > backend/.gitignore

# ---------- FRONTEND ----------
setup-fe:
	@echo "Setting up React frontend"
	npx create-vite@latest frontend -- --template react-ts
	cd frontend && npm install
	cd frontend && npm install -D tailwindcss postcss autoprefixer
	cd frontend && npx tailwindcss init -p
	cd frontend && sed -i 's/content: \\[\\]/content: ["\\.\\/index.html","\\.\\/src\\/**\\/*.{js,ts,jsx,tsx}"]/' tailwind.config.js
	cd frontend && echo "@tailwind base;\n@tailwind components;\n@tailwind utilities;" > src/index.css

# ---------- DEV ----------
start-be:
	cd backend && go run .

start-fe:
	cd frontend && npm run dev

dev: start-be start-fe