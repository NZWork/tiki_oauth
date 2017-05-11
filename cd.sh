# Build for linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

# rsync
rsync -e 'ssh -p 1027' -v product.toml tiki_oauth neo@amoy.layer.nevoz.com:/home/neo/docker/images/tiki_oauth/

# Deploy
ssh neo@amoy.layer.nevoz.com -p 1027 /home/neo/docker/images/tiki_oauth/deploy.sh
