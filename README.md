gosomewhere
----------------

## Install

```
go get -u github.com/sato-s/gosomewhere
```

Make sure you have PATH to `$GOPATH/bin`.  

## Usage

Create your config.yaml  

```
cat << EOF > config.yaml
port: 80
listen: 0.0.0.0
destinations:
  shopping: https://www.amazon.com/
  creditcard: https://www.smbc-card.com/mem/index.jsp
  search: https://www.google.co.jp
  vim: https://vim.rtorr.com/
  ascii: https://en.wikipedia.org/wiki/ASCII#Printable_characters
  cloud: https://developers.digitalocean.com/documentation/v2/
EOF
```

Run server  

```
sudo env "PATH=$PATH" gosomewhere --config config.yaml
# require `sudo` to use port 80
```

Go to `localhost/vim` to check vim cheetsheet!.  

*You need root privilege to use port 80*  


(Optional) Add following entry to `/etc/hosts`  

```
127.0.0.1 go
```

*If you are using windows edit `c:\windows\system32\drivers\etc\hosts`*

Go to `go/vim` to check vim cheetsheet!.  
