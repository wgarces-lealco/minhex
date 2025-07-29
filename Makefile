.PHONY: help demo test clean

help:
	@echo "🎯 MinHex - Demo Arquitectura Hexagonal"
	@echo "======================================"
	@echo ""
	@echo "✨ DEMO: Cambio de tecnología en 2 líneas de código"
	@echo ""
	@echo "Comandos:"
	@echo "  make demo     - Iniciar servidor API"
	@echo "  make test     - Probar API con curl"
	@echo "  make clean    - Limpiar binarios"
	@echo ""
	@echo "🔄 Cambiar tecnología:"
	@echo "  1. Ctrl+C para parar servidor"
	@echo "  2. Editar cmd/main.go líneas 23-24"
	@echo "  3. make demo"

demo:
	@echo "🚀 Iniciando demo..."
	@echo "🌐 API disponible en: http://localhost:8080"
	@echo "💡 Cambiar tecnología: Editar cmd/main.go líneas 23-24"
	go run cmd/main.go

test:
	@echo "🧪 Probando API..."
	@echo "1. Creando usuario..."
	@curl -X POST http://localhost:8080/users \
		-H "Content-Type: application/json" \
		-d '{"email":"demo@test.com","name":"Demo User"}' \
		2>/dev/null && echo "" || echo "❌ Servidor no está corriendo"
	@echo ""
	@echo "2. Health check..."
	@curl http://localhost:8080/health 2>/dev/null && echo "" || echo "❌ Servidor no está corriendo"
	@echo ""
	@echo "3. Creando comercio..."
	@curl -X POST http://localhost:8080/commerces \
		-H "Content-Type: application/json" \
		-d '{"name":"Demo Store","description":"Test store","owner_id":"user_20240101120000"}' \
		2>/dev/null && echo "" || echo "❌ Servidor no está corriendo"

clean:
	@echo "🧹 Limpiando..."
	@rm -rf bin/ 2>/dev/null || true 