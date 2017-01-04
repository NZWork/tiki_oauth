docker build -t tiki_oauth .

# Stop it
docker rm -f tiki_oauth_app

# Start a new app

docker run -itd -p 30003:80 --name tiki_oauth_app tiki_oauth
