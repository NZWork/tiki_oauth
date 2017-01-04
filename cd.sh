# Build for linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

# rsync
rsync product.toml neo@10.1.1.7:/home/neo/docker/images/tiki_oauth
rsync tiki_oauth neo@10.1.1.7:/home/neo/docker/images/tiki_oauth

# Deploy
ssh neo@10.1.1.7 /home/neo/docker/images/tiki_oauth/deploy.sh
