# Minimalistic Hexagonal Architecture

Proyecto de estudio que demuestra una implementación práctica y escalable de la arquitectura hexagonal en Go.

## 🏗️ Estructura del Proyecto

```
src/
├── domain/          # Core business (sin dependencias externas)
│   ├── entities/    # Entidades de negocio
│   ├── events/      # Eventos de dominio
│   ├── ports/       # Interfaces/contratos
│   └── errors/      # Errores específicos de dominio
├── usecases/        # Lógica de aplicación
├── infra/           # Implementaciones externas
└── presentation/    # Capa de entrega
```

## 🚀 Características

- Arquitectura hexagonal limpia y pragmática
- Separación clara de responsabilidades
- Fácil de mantener y escalar
- Estructura autodocumentada
- Imports cortos y claros

## 🛠️ Requisitos

- Go 1.23.1 o superior

## 📦 Instalación

```bash
go mod download
```

## 🏃‍♂️ Ejecución

```bash
go run main.go
```

## 🧪 Testing

```bash
go test ./...
```

## 📝 Licencia

MIT 