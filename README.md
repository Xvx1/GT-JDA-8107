# GT-JDA-8107
Junior Design for Team 8107

## Setup
You must install GoLang in order to run this command-line application:
`
https://golang.org/doc/install
`

## Install this Application
In order to install this application, use the `go get` command as follows:

`go get github.com/Xvx1/GT-JDA-8107`

Once you have performed this operation, navigate to your `GOPATH` directory for this application.

## Usage
Once you are in your `GOPATH` directory, navigate to the source of this project, likely located at `src/github.com/Xvx1/GT-JDA-8107`

To run the application, use the following commands:

`go run generateBrackets.go`

The `--single` flag will print out the single elimination brackets generated from the given example list of players.
The `--roundrobin` flag takes a positive integer number of groups and will print out the generated round robin groups.