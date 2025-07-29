# 🎯 MinHex - Demo Arquitectura Hexagonal
## Microservicios Orientados a Negocio

> **¡Demo súper simple!** Cambia de tecnología en **2 líneas de código**. Perfecta demostración de arquitectura hexagonal para microservicios orientados a negocio.

## 🚀 Quick Start

```bash
git clone <repo>
cd minhex
go run cmd/main.go
```

**¡Ya está!** Servidor corriendo en http://localhost:8080

## 🔄 Cambiar Tecnologías (SÚPER FÁCIL)

Edita `cmd/main.go` líneas 23-24:

```go
// 📨 MESSAGING: Descomenta la que quieras usar
var eventPublisher ports.EventPublisher
eventPublisher = sqs.NewPublisher()        // ✅ AWS SQS
// eventPublisher = rabbitmq.NewPublisher() // 🔄 RabbitMQ
```

**¡2 líneas = Cambio completo de tecnología!**

## 🏗️ Estructura SÚPER Organizada

```
cmd/                        # 🚀 Entry points
└── main.go                 # Main application
src/
├── domain/                 # 🧠 DOMINIOS DE NEGOCIO
│   ├── users/              # Users completo
│   ├── commerces/          # Commerces completo  
│   └── shared/             # Shared kernel
├── usecases/               # 🎯 CASOS DE USO PLANOS
│   ├── create_user/        # Directo aquí
│   ├── get_user/           # Sin anidación
│   ├── create_commerce/    # Súper plano
│   └── activate_commerce/  # Simple
├── infra/                  # 🔌 INFRAESTRUCTURA POR TIPO
│   ├── persistence/        # 💾 Base de datos
│   │   ├── memory/         # ├─ In-memory  
│   │   └── postgres/       # └─ PostgreSQL
│   ├── messaging/          # 📨 Eventos
│   │   ├── sqs/            # ├─ AWS SQS
│   │   └── rabbitmq/       # └─ RabbitMQ
│   ├── cache/              # 🗄️ Cache
│   └── external/           # 🌐 APIs externas
└── presentation/           # 🎭 HTTP server + handlers
    ├── handlers/           # Request/response handling
    └── server/             # HTTP server logic
```

## 🎭 Demo en Vivo

### **1. API Funcionando**
```bash
# Crear usuario
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"email":"demo@test.com","name":"Demo User"}'

# Health check
curl http://localhost:8080/health
```

### **2. Cambio de Tecnología**
```bash
# 1. Para el servidor (Ctrl+C)
# 2. Edita cmd/main.go líneas 23-24:
#    Comenta SQS, descomenta RabbitMQ
# 3. Reinicia: go run cmd/main.go
# 4. ¡Misma API, diferente tecnología!
```

### **3. Cross-Domain**
```bash
# Crear comercio (valida que owner exista)
curl -X POST http://localhost:8080/commerces \
  -H "Content-Type: application/json" \
  -d '{"name":"Demo Store","description":"Test","owner_id":"user_123"}'
```

## 💪 Ventajas Demostradas

### ✅ **Cambio de Tecnología = 2 Líneas**
- **SQS** ↔ **RabbitMQ**: Solo comentar/descomentar
- **Memory** ↔ **PostgreSQL**: Solo comentar/descomentar
- **Sin configuración compleja**: Todo hardcodeado para la demo

### ✅ **Arquitectura Hexagonal Pura**
- **Dominios separados**: users/, commerces/
- **Casos de uso planos**: directos en usecases/
- **Infraestructura organizada**: por tipo + tecnología
- **Interfaces múltiples**: HTTP API (+ CLI preparado)

### ✅ **Microservicios Orientados a Negocio**
- **Un dominio = un microservicio potencial**
- **Cross-domain controlado**: través de ports
- **Escalabilidad real**: agregar domains sin impacto
- **Deploy independiente**: mismo código, diferentes servicios

## 📊 VS Otras Arquitecturas

| **Aspecto** | **Vertical Slicing** | **Micro-Micros** | **MinHex** |
|-------------|---------------------|------------------|------------|
| **Cambio Tech** | Muchos archivos | Muchos servicios | **2 líneas** |
| **Estructura** | Por capas técnicas | Muy fragmentada | **Por negocio** |
| **Setup** | Complejo | Muy complejo | **go run cmd/main.go** |
| **Demo** | 30+ min | 60+ min | **< 5 min** |

## 🎯 Para la Presentación

### **Guion (3 minutos):**

1. **"Miren qué simple"** (30s)
   ```bash
   tree src/  # Mostrar estructura
   go run cmd/main.go  # Levantar API
   ```

2. **"API funcionando"** (60s)
   ```bash
   curl -X POST http://localhost:8080/users -H "Content-Type: application/json" -d '{"email":"demo@test.com","name":"Demo"}'
   curl http://localhost:8080/health
   ```

3. **"Cambio de tecnología"** (90s)
   ```bash
   # Ctrl+C para parar
   # Mostrar líneas 23-24 en cmd/main.go
   # Comentar SQS, descomentar RabbitMQ  
   # go run cmd/main.go
   # Misma API, diferente tech!
   ```

### **Argumentos de Cierre:**
- ✅ **"Arquitectura hexagonal en su forma más pura"**
- ✅ **"Cambio de tecnología = 2 líneas de código"**
- ✅ **"Microservicios orientados a negocio"**
- ✅ **"Estructura súper organizada y navegable"**

## 🛠️ Requisitos

- Go 1.23.1+
- **Eso es todo.** Sin Docker, sin k8s, sin configuración.

**La demo más simple y potente de arquitectura hexagonal.**
**¡2 líneas = Cambio completo de tecnología!** 