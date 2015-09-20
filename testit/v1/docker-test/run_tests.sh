docker run --rm -v "$(pwd)/my-soln":/app -w /app glot/clang sh local_build.sh my-soln
docker run --rm -v "$(pwd)/primary-soln":/app -w /app glot/clang sh local_build.sh primary-soln
docker run --rm -v "$(pwd)/gen":/app -w /app glot/clang sh local_build.sh gen
docker run --rm -v "$(pwd)/runtest":/app -w /app glot/golang go build runtest.go
mkdir workdir
cp gen/gen workdir
cp my-soln/my-soln workdir
cp primary-soln/primary-soln workdir
cp runtest/runtest workdir

