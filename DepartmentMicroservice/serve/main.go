package main

import (
	"flag"
	"github.com/Fring02/HospitalMicroservices/DepartmentMicroservice/api"
	"github.com/Fring02/HospitalMicroservices/DepartmentMicroservice/database"
	"github.com/Fring02/HospitalMicroservices/DepartmentMicroservice/pkg/repositories"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"os"
)

func main() {
	dsn := flag.String("dsn", "postgresql://localhost/hospital?user=postgres&password=Rubin1!!", "PostGreSQL")
	flag.Parse()
	var pool *pgxpool.Pool
	var err error
	pool, err = database.OpenDB(*dsn)
	os.Setenv("connection", *dsn)
	api.DoctorRepository = repositories.NewDoctorRepository(pool)
	api.DepartmentRepository = repositories.NewDepartmentsRepository(pool)
	if err != nil{
		log.Fatalf("Failed to connect to db: ", err)
	}
	router  := gin.Default()
	api.RouteDoctors(router)
	api.RouteDepartments(router)
	router.Run(":4000")
}