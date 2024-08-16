## Welcome to BlueSoft Bank

### Introduction
This development has been created for evaluation purposes. It was designed under **Hexagonal Architecture** (so-called "**Ports and Adapters**").

### Note: Please the document until the end, there are several things WIP.
(Sorry for the above text in bold)


## Technologies used

- Golang 1.22 with Gin Framework
- Docker 27.0.3
- Docker Compose 2.28.1
- PostgreSQL 16.3

## Bootstrapping

At the moment of writting this document, the complete app hasn't been dockerized yet, so we need to setup the database:
```bash
# In the root of the project, I mean ./bluesoft-bank-solution
$ docker compose up -d # Watch out, we're running this docker compose in background
```



## WIP
  1. Write tests.
  2. Develop a frontend project in order to improve the interactivity.
  3. Dockerize the application completely, this includes the main app.