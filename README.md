# Roblox-Game-Presence-For-Discord
Uses Discord RPC to give you a rich presence on discord for the current ROBLOX game you are playing.

![example](https://doy2mn9upadnk.cloudfront.net/uploads/default/original/4X/d/5/7/d57e244c276ffde1f968185e2ac8eff0305ba9c6.png)


This is an alternative to "ro-presence". _Why?_ Because ro-presence goes overboard by reading from your registry so that it can get your `.ROBLOSECURITY` cookie to function. 

This just reads from the process's start command line. 

### Build instructions

If you wish to build the project, simply run `go build Main.go`.
If you want to build it without the terminal window showing, try this: `go build -ldflags "-H=windowsgui" Main.go`
