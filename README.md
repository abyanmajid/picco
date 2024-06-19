# codemore.io (WIP)

<p>
  <a href="https://github.com/abyanmajid/codemore.io/blob/main/LICENSE"><img alt="GPL-3.0 License" src="https://img.shields.io/badge/License-GPL%203.0-blue.svg"></a>
  <img alt="cicd-badge" src="https://github.com/abyanmajid/codemore.io/actions/workflows/cicd.yml/badge.svg">
</p>

**codemore.io** is a distributed application designed to offer *free programming courses* that attempts to help you get out of *"tutorial hell"*, or simply *speed up your learning*, all by putting great emphasis on *writing more code*.

<!-- ### Live App  🚀 | Documentation 🔍 | Source 📦 -->

## Architecture

All client requests are sent to the `broker` service (which serves as an API gateway) with a JSON payload. The `broker` service will then redirect this request via `gRPC` to the correct microservice. The following `mermaid` visualizer depicts the architecture:

```mermaid
graph TD
    Client["<b>Client</b><br>(codemore.io)"] <-->|REST| Broker["<b>Broker</b><br>(Docker)"]
    Broker <-->|gRPC| User["<b>User</b><br>(Docker)"]
    Broker <-->|gRPC| Notification["<b>Notification</b><br>(Docker)"]
    Broker <-->|gRPC| Mail["<b>Mail</b><br>(Docker)"]
    Broker <-->|gRPC| Judge["<b>Judge</b><br>(Docker)"]
    Broker <-->|gRPC| Compiler["<b>Compiler</b><br>(Docker)"]
    User -->|gRPC| PostgreSQL["<b>PostgreSQL</b><br>(Docker)"]
    Notification -->|gRPC| MongoDB["<b>MongoDB</b><br>(Docker)"]
    Mail -->|gRPC| Mailhog["<b>Mailhog</b><br>(Docker)"]
```

There are currently 6 API microservices:

- `broker`: An API gateway to process all user requests via the `REST` communication protocol
- `user`: A microservice responsible for authentication, authorization, and changing user information and metrics
- `notification`: A microservice responsible for logging and retrieving notifications
- `mail`: A microservice responsible for sending mails
- `judge`: A microservice responsible for running test cases on code outputs
- `compiler`: A microservice responsible for compiling user-submitted code

## Contributing

1. Get started by forking this repository, clone it to your local device, and create a new branch by running `git checkout -b <branch_name>`
2. Open `docker-compose-yml` and make sure all `ENVIRONMENT` environment variables are set to `"development"`
3. Run `make up-build` to pull all required docker images, build binaries, and run all backend microservices in docker containers
4. To set up the frontend client, run `cd web/ui && npm i` from the root to install dependencies
5. Run `cp .env.default .env && rm -rf .env.default` to copy over default environment variables. Make sure `ENVIRONMENT` is set to `"development"`
6. Run `npm run dev` to start the frontend client at `localhost:3000`
7. Refer to the documentation as you make changes
8. Submit a pull request to `staging` when you are done.

**Note on languages:** Since all backend microservices communicate via `gRPC`, all languages that have support for compiling `.proto` files are welcome. To compile `.proto` in Go, `cd` to the directory where your `.proto` file lives in and run the following:
```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative <filename>.proto
```
