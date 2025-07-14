package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Abhyuday04/wyp/handlers"
	"github.com/Abhyuday04/wyp/infra/redis"
	"github.com/Abhyuday04/wyp/infra/sms"
	"github.com/Abhyuday04/wyp/internal/app"
	repository "github.com/Abhyuday04/wyp/layers/repository"
	"github.com/Abhyuday04/wyp/layers/services"
	"github.com/Abhyuday04/wyp/layers/transport"
	_ "github.com/lib/pq"
	"github.com/twilio/twilio-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// TODO move to config file and refactor
const (
	user     = "postgres.xqjolmvnjxhlktqvjtrs"
	password = "12345678"
	host     = "aws-0-ap-south-1.pooler.supabase.com"
	port_db  = 6543
	dbname   = "postgres"
)

const (
	accountSid = "ACf5c3f52a851beff684f452dd78bd0332"
    authToken = "507ff438005c12cab2fd92eb5f05b417"
)

var db *gorm.DB
var tc *twilio.RestClient

func main() {
	// Initialize Twilio client
	tc = twilio.NewRestClientWithParams(twilio.ClientParams{
        Username: accountSid,
        Password: authToken,
    })


	// Get port from environment or use default
	var port string
	if port == "" {
		port = "8000"
	}
	log.Println("port", port)
	// refactor
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port_db, user, password, dbname)
	var err error
	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  psqlInfo,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		PrepareStmt: false,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(" DBSuccessfully connected!")

	config := redis.RedisConfig{
		Host:         "redis",
		Port:         "6379",
		Password:     "",
		DB:           0,
		PoolSize:     10,
		MinIdleConns: 5,
		MaxRetries:   3,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}

	redis.RedisClient, err = redis.NewRedisClientWithPool(config)
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
		panic(err)
	}

	defer func() {
		sqlDB, err := db.DB()
		if err == nil {
			sqlDB.Close()
		}
		redis.RedisClient.Close()
	}()

	// make server provider
	makeServerProvider()

	// Initialize router
	router := handlers.NewRouter()

	// Create server
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		fmt.Printf("Server starting on port %s...\n", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Server shutting down...")
}

func makeServerProvider() {
	sms := sms.New(tc)
	repository := repository.New(db)
	services := services.New(repository, sms)
	transport := transport.New(services)
	app.Srv = app.Server{
		Service:    services,
		Transport:  transport,
		Repository: repository,
		SmsService: sms,
	}
}
