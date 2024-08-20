#run docker
docker compose up -d

#run mod
go mod tidy

#run migrate
go run migrate/migrate.go

#run seed 
go run seed/seed.go
#run for dev
CompileDaemon -command="./baby-varkin-health"

log terminal 
	// fmt.Printf("Email: %s, Password: %s, Username: %s ,Role: %v ,rolesarr: %v\n", c.PostForm("email"),authInput.Password, authInput.Username ,role ,rolesarr)
    %s string
    %v array