# HTTPS Golang
How to setup HTTPS using [`net/http`](https://pkg.go.dev/net/http) package in Go.
## [Certbot](https://certbot.eff.org/)
For generating SSL certificates, we can use `certbot` which is a tool to manage certificates from Let's Encrypt.
### Install
```bash
sudo apt-get remove certbot
sudo snap install --classic certbot
sudo ln -s /snap/bin/certbot /usr/bin/certbot
```
### Generate
```bash
sudo certbot certonly --standalone
```
Now, you can find the certificates/keys in `/etc/letsencrypt/live/***domain-name***/`.
- `/etc/letsencrypt/live/staging.a110.studio/cert.pem`
- `/etc/letsencrypt/live/staging.a110.studio/privkey.pem`
## Go
[https-golang.go](./https-golang.go)