# go-notes-api

Golang REST api for taking notes

You can start/build app locally with `go` command (see below)

Or you can use `run` script which provide facilities for start/build app locally or inside containers

## Dependencies

App use [Gin](https://github.com/gin-gonic/gin) as the web server

If you want to execute/build app locally, please install needed dependencies

For docker execution, you can skip this step

```bash
go get all
```

## Port

App listen on port 8080 (default), you can change this setting with the environment variable `PORT`

example `PORT=3000`

If you use `run` script, `.env` file present in project root directory is automatically sourced

# Start

```bash
go run main.go		# basic execution
./run start			# with the help of the run script
./run docker start	# execution inside container
```

# Build

```Bash
go build				# basic execution
# or
go build -o build/

./run build			# with the help of the run script
./run docker build 	# execution inside container
```

## Bonus, run in watch mode

Run the app in watch mode (dev mode), each update in source files restart the app

```Bash
./run dev
./run docker dev
```

# Test App

Api has one enpoint `/notes` for create/get notes

Notes are in the form:

```json
{
  "message": "the message",
  "tag": "<TAG>"
}
```

Tag is optionnal

You can :

- create a note &rightarrow; `POST /notes` with a Note, JSON formated, in the body
- get all notes &rightarrow; `GET /notes`
- get notes filtered &rightarrow; `GET /notes?tag=<TAG>`
