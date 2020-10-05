# connect4

Simple Connect 4 game written in Go

## Features

* Connect N for any number of N, not just 4
* Arbitrary board sizes, not just 7 x 6
* Two or more players

## Run

### Build docker image

```
docker build -t connect4 .
```

### Run game

```
> echo 0 1 0 1 0 1 0 | docker container run -i connect4
WINNER: Player 1
```

## Not included

* Early draw detection
* Fancy game config option handling: game supports it, but it's not exposed externally
* Game is not a state machine
