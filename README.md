# Getting set up with a back end for pact testing (go lang)

## Install Go

```
curl https://storage.googleapis.com/golang/go1.9.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.9.linux-amd64.tar.gz
echo 'PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
```

## Install echo server
```
go get -u github.com/labstack/echo/...
```

## Install pact for golang

Follow the instructions [here]: (https://github.com/pact-foundation/pact-go)

