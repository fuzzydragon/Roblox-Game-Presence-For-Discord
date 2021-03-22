# Roblox-Game-Presence-For-Discord
Uses Discord RPC to give you a rich presence on discord for the current ROBLOX game you are playing.
This is an alternative to "ro-presence". Why? Because ro-presence goes overboard by reading from your registry so that it can get your `.ROBLOSECURITY` cookie to function. 
This just reads from the process' start command line. 

### Build instructions

If you wish to build the project, simply run `go build Main.go`.
If you want to build it without the terminal window showing, try this: `go build -ldflags "-H=windowsgui" Main.go`
