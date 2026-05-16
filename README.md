# Практические занятия №7 и №8: Docker-контейнеризация и GitHub Actions CI/CD

## Описание

В рамках данных практических работ выполнена контейнеризация Go-сервиса `tasks` с помощью Docker и настроен автоматизированный CI/CD-пайплайн с использованием GitHub Actions.

---

## Структура проекта

.
├── .git
│   ├── COMMIT_EDITMSG
│   ├── HEAD
│   ├── branches
│   ├── config
│   ├── description
│   ├── hooks
│   │   ├── applypatch-msg.sample
│   │   ├── commit-msg.sample
│   │   ├── fsmonitor-watchman.sample
│   │   ├── post-update.sample
│   │   ├── pre-applypatch.sample
│   │   ├── pre-commit.sample
│   │   ├── pre-merge-commit.sample
│   │   ├── pre-push.sample
│   │   ├── pre-rebase.sample
│   │   ├── pre-receive.sample
│   │   ├── prepare-commit-msg.sample
│   │   ├── push-to-checkout.sample
│   │   ├── sendemail-validate.sample
│   │   └── update.sample
│   ├── index
│   ├── info
│   │   └── exclude
│   ├── logs
│   │   ├── HEAD
│   │   └── refs
│   │       ├── heads
│   │       │   └── main
│   │       └── remotes
│   │           └── origin
│   │               └── main
│   ├── objects
│   │   ├── 16
│   │   │   └── 234aefc37f89854dea51d122dbdfee17d11698
│   │   ├── 3e
│   │   │   └── 2e2816e1aab43d7181d15c0c105644236b2f8f
│   │   ├── 5b
│   │   │   └── e8873faf86fdf8ab13a404ba89dab60659ff7c
│   │   ├── 7e
│   │   │   └── 7f6d4bcdf3a05bbbe4bbc1774bd81ab375ad72
│   │   ├── 94
│   │   │   ├── 8724c15e9f1fab7a320b5e86c4a0e6f934b9ce
│   │   │   └── e96951797db816c70c57f1e421dc9963f14d76
│   │   ├── a4
│   │   │   └── 18d1c6ec91298587f5121dea2435367e50cdd2
│   │   ├── b9
│   │   │   └── 8289cd4cf2858bd19de42028a87c2294d531fa
│   │   ├── c9
│   │   │   └── 5f33aa2eb06d6e2fa2a07221bfae4aff4433d4
│   │   ├── e6
│   │   │   └── 9de29bb2d1d6434b8b29ae775ad8c2e48c5391
│   │   ├── e8
│   │   │   └── a9d21efddb8cf89a664231f2f528d9553d8495
│   │   ├── eb
│   │   │   └── 43e8d569a9d6822344f4cb39aeb1852c66a338
│   │   ├── ef
│   │   │   └── c84d9010300625a86cb3ee40cbcf1f9f18c63b
│   │   ├── f2
│   │   │   └── 6cc42f20d00006160a267ece645cc118cc8b00
│   │   ├── f9
│   │   │   └── ef86e9811580afb80fdb9f28e38d05f1e915bf
│   │   ├── fa
│   │   │   └── 4eb842dda745518d36401947f0b9d089918441
│   │   ├── fc
│   │   │   └── 49427dcd99c0e15042567a46c68547feed9746
│   │   ├── info
│   │   └── pack
│   └── refs
│       ├── heads
│       │   └── main
│       ├── remotes
│       │   └── origin
│       │       └── main
│       └── tags
├── .github
│   └── workflows
│       └── ci.yml
├── .gitignore
├── deploy
│   └── docker-compose.yml
└── services
    └── tasks
        ├── .dockerignore
        ├── Dockerfile
        ├── cmd
        │   └── tasks
        │       └── main.go
        ├── go.mod
        ├── go.sum
        └── internal

---

## Практическая работа №7: Docker-контейнеризация

### 1. Сервис tasks

HTTP-сервис на Go с эндпоинтами:
- `GET /health` — проверка работоспособности
- `GET /` — информация о сервисе

### 2. Сборка Docker-образа

![Сборка Docker-образа](screen/docker-built.png)

```bash
docker build -t techip-tasks:0.1 .
```

### 3. Список Docker-образов
![Сборка Docker-образа](screen/docker_images.png)

### 4. Запуск контейнера
![Сборка Docker-образа](screen/start_docker.png)

### 5. Проверка работоспособности
![Сборка Docker-образа](screen/check_local.png)

### 6. Логи контейнера
![Сборка Docker-образа](screen/docker_logs.png)

### 7. Запуск через Docker Compose
![Сборка Docker-образа](screen/compose.png)

## Практическая работа №8: GitHub Actions CI/CD

### 1. Пайплайн CI в GitHub Actions
![Сборка Docker-образа](screen/Actions.png)

### 2. Успешное выполнение пайплайна
![Сборка Docker-образа](screen/Pipeline.png)

### 3. Лог выполнения test-and-build
![Сборка Docker-образа](screen/Test_and_built.png)

### 4. Лог выполнения docker-build
![Сборка Docker-образа](screen/Docker_built.png)

### 5. Файл ci.yml

```bash
name: CI Pipeline

on:
  push:
    branches: [ "main", "master" ]
  pull_request:
    branches: [ "main", "master" ]

jobs:
  test-and-build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version: '1.23'
      - name: Test and build
        run: |
          cd services/tasks
          go mod tidy
          go test ./...
          go build ./cmd/tasks

  docker-build:
    runs-on: ubuntu-latest
    needs: test-and-build
    steps:
      - uses: actions/checkout@v4
      - name: Build Docker image
        run: |
          cd services/tasks
          docker build -t techip-tasks:${{ github.sha }} .
```
