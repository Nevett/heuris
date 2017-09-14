GIT_COMMIT=$(git rev-parse --short HEAD)

docker build -t nullseed/heuris:latest -t nullseed/heuris:$GIT_COMMIT .
