package main

import (
	"fmt"
	commercePorts "minhex/src/domain/commerces/ports"
	userPorts "minhex/src/domain/users/ports"
	"minhex/src/infra/messaging/sqs"
	"minhex/src/infra/persistence/memory"
	"minhex/src/presentation/server"
	"minhex/src/usecases/activate_commerce"
	"minhex/src/usecases/create_commerce"
	"minhex/src/usecases/create_user"
	"minhex/src/usecases/get_user"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	locationNow := now.Location()
	fmt.Println(locationNow)

	utcNow := now.UTC()
	fmt.Println(utcNow)

	colombiaNow := now.In(time.FixedZone("America/Bogota", -5*3600))
	fmt.Println(colombiaNow)

	salvadorNow := now.In(time.FixedZone("America/El_Salvador", -6*3600))
	fmt.Println(salvadorNow)

	chileNow := now.In(time.FixedZone("America/Santiago", -4*3600))
	fmt.Println(chileNow)

	// 🎯 DEMO: Arquitectura Hexagonal - Microservicios Orientados a Negocio

	// ===============================================
	// 🔧 CONFIGURACIÓN - CAMBIAR TECNOLOGÍAS AQUÍ
	// ===============================================

	// 📨 MESSAGING: Descomenta la que quieras usar
	eventPublisher := sqs.NewPublisher() // ✅ AWS SQS
	// eventPublisher = rabbitmq.NewPublisher() // 🔄 RabbitMQ

	// 💾 PERSISTENCE: Descomenta la que quieras usar
	var userRepo userPorts.UserRepository
	var commerceRepo commercePorts.CommerceRepository
	userRepo = memory.NewUserRepository()         // ✅ In-Memory
	commerceRepo = memory.NewCommerceRepository() // ✅ In-Memory
	// userRepo = postgres.NewUserRepository()      // 🔄 PostgreSQL
	// commerceRepo = postgres.NewCommerceRepository() // 🔄 PostgreSQL

	// ===============================================
	// 🎯 CASOS DE USO - CORE DE NEGOCIO (INTERFACES)
	// ===============================================
	var createUserUC create_user.CreateUserUseCase
	var getUserUC get_user.GetUserUseCase
	var createCommerceUC create_commerce.CreateCommerceUseCase
	var activateCommerceUC activate_commerce.ActivateCommerceUseCase

	createUserUC = create_user.New(userRepo, eventPublisher)
	getUserUC = get_user.New(userRepo)
	createCommerceUC = create_commerce.New(commerceRepo, userRepo, eventPublisher)
	activateCommerceUC = activate_commerce.New(commerceRepo, eventPublisher)

	// ===============================================
	// 🌐 HTTP SERVER - CAPA DE PRESENTACIÓN
	// ===============================================
	httpServer := server.NewHTTPServer(
		createUserUC,
		getUserUC,
		createCommerceUC,
		activateCommerceUC,
	)

	// 🚀 INICIAR SERVIDOR
	httpServer.Start()
}
