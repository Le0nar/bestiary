1. ### For run docker postgresql:

    Pull docker image of postgres docker pull postgres
   ` docker run --name=bestiary -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d postgres`

    For next run use `docker ps -a` for view list of all container
    After use `docker start {container_id}`

2. ### Create migrate file (alredy created in this repo)

    -   Create 'up' and 'down' files
    `migrate create -ext sql -dir ./schema -seq init`

    - Write sql command for 'up' and 'down' files

3. ### Apply migrations

    `migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up`

4. ### Open postgres in docker 
    - `docker exec -it {container_id} /bin/bash `
    - `root@3d52773f7ff2:/# psql -U postgres `
    - ` postgres=# \d`


5. ### Run app
    - Run psql container  (if it dont run already)
    - Run golang app `go run main.go`