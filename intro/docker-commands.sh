# build dockerfile - up to dev stage of dockerfile
docker build --target dev . -t go-intro-dev
# run dockerfile using the current directory as volume
docker run -it -v ${PWD}:/work go-intro-dev sh
# with ports exposed
docker run -it -v ${PWD}:/work -p 8080:8080 go-intro-dev sh
# check version
go version

# build dockerfile as build step with build file
docker build --target build . -t go-intro-build

# build dockerfile as runnable using alpine
docker build --target runtime . -t go-intro-runtime