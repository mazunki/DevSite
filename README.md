# DevSite
My developer and public face website.

Using GoLang as a websocket, configured on top of a Google Cloud Instance, running Debian 10.

I'm going to be adding different aspects of what I do over time. I've built the backbone now, and I'm happy that I got SSL/TLS running. What a hassle.

The idea is to use the site as a portfolio, but also as a place to place thoughts. I have another site, on another TLD, which I will be using to play with different things. For this page, I will /try/ to act professionally. Ha. Ha. Ha.

To make a copy of the page on your local machine, because I know you will:
```
git clone https://github.com/mazunki/DevSite.git
cd DevSite
go run websocket.go
```

## Dependencies: 
  `iptables-persistent iptables golang go`

On Debian:
```
sudo apt update; sudo apt upgrade;
sudo apt install -y iptables-persistent iptables golang
```

On Arch-based
```
sudo pacman -Syu iptables go
```

## GitHub integration -- shortcuts
- My GitHub page: `git.mazunki.dev`
- A GitHub repository: `git.mazunki.dev/DevSite`
- A GitHub file, on default branch: `git.mazunki.dev/DevSite/content/index.html`
- A raw file, since some browsers will try to download files: `git.mazunki.dev/f/uio/EXPHIL03E/oblig2/argumentative_singer.pdf`

In general: `git.mazunki.dev/[[f/]repo/[file.extension]]`

Also, `git.mazunki.dev` is synonymous to `mazunki.dev/git`.

Have a nice day! 

PS: ILoveCandy
