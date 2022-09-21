# Library Service

* DDD (Domain Driven Design) & Fundamental Clean Architecture 
* Hexagonal Architecture with Database (SQL & No SQL)
* Hexagonal Architecture with Middleware (Log & Authenthication)
<div><img src="../tasks/task-day4.jpg" widht="500" height="500"/></div>

Run Project
for migrate table, use:
```bash
go run main.go -migrate=migrate
 ```
for rollback table, use:
```bash
go run main.go -migrate=rollback
 ```
for get status migration (default), use:
```bash
go run main.go -migrate=status 
 ```


