package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"minhex/src/presentation/handlers"
	"minhex/src/presentation/middleware"
	"minhex/src/usecases/activate_commerce"
	"minhex/src/usecases/create_commerce"
	"minhex/src/usecases/create_user"
	"minhex/src/usecases/get_user"
)

type HTTPServer struct {
	userHandler     *handlers.UserHandler
	commerceHandler *handlers.CommerceHandler
	port            string
}

func NewHTTPServer(
	createUserUC create_user.CreateUserUseCase,
	getUserUC get_user.GetUserUseCase,
	createCommerceUC create_commerce.CreateCommerceUseCase,
	activateCommerceUC activate_commerce.ActivateCommerceUseCase,
) *HTTPServer {
	userHandler := handlers.NewUserHandler(createUserUC, getUserUC)
	commerceHandler := handlers.NewCommerceHandler(createCommerceUC, activateCommerceUC)

	return &HTTPServer{
		userHandler:     userHandler,
		commerceHandler: commerceHandler,
		port:            ":8080",
	}
}

func (s *HTTPServer) setupRoutes() {
	// ===============================================
	// 📋 ENDPOINTS - MICROSERVICIOS ORIENTADOS A NEGOCIO
	// ===============================================

	// 👤 USERS DOMAIN
	http.HandleFunc("/users", middleware.Logging(middleware.CORS(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			s.userHandler.CreateUser(w, r)
		case "GET":
			s.userHandler.GetUser(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})))

	// 🏪 COMMERCES DOMAIN
	http.HandleFunc("/commerces", middleware.Logging(middleware.CORS(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			s.commerceHandler.CreateCommerce(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})))

	http.HandleFunc("/commerces/activate", middleware.Logging(middleware.CORS(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			s.commerceHandler.ActivateCommerce(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})))

	// 💚 HEALTH CHECK
	http.HandleFunc("/health", middleware.Logging(middleware.CORS(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		response := map[string]string{
			"status": "healthy",
			"demo":   "microservicios-orientados-negocio",
		}
		json.NewEncoder(w).Encode(response)
	})))
}

func (s *HTTPServer) Start() {
	s.setupRoutes()

	fmt.Println("🎯 MinHex - Arquitectura Hexagonal Demo")
	fmt.Println("📊 Microservicios Orientados a Negocio")
	fmt.Println("=====================================")
	fmt.Printf("🌐 Server: http://localhost%s\n", s.port)
	fmt.Println("📋 Endpoints:")
	fmt.Println("   POST /users          - Create user")
	fmt.Println("   GET  /users?id=ID    - Get user")
	fmt.Println("   POST /commerces      - Create commerce")
	fmt.Println("   POST /commerces/activate - Activate commerce")
	fmt.Println("   GET  /health         - Health check")
	fmt.Println("")

	log.Fatal(http.ListenAndServe(s.port, nil))
}
