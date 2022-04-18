docker build --force --tag golsim .;
docker rm golsim;
docker run -dp 3000:3000 --name golsim golsim
