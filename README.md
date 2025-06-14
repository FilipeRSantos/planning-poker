# Summary

This is a pet project to discover how easy it is to implement a MVP with HTMX and WebSockets.

The selected project is a planning poker, with the goal of using it on my current company as a real test.

To make my life easier, I opted to use this [bootstrap library](https://github.com/Melkeydev/go-blueprint).

# Goals

In order to make my life easier, I have set up a few goals to work towards, to help organize my thoughts in the process.

- [x] Setup a basic webapp
- [x] Configure Basic Templates Support
- [x] Add a HTMX action to test behavior
- [] Create a form that redirects the user to a new area
- [] Add session support to this action
- [x] Add WebSocket support to the webapp
- [x] Add a simple JS flow to test the WS on the new area
- [x] Map the WS calls with HTMX
- [] Throw some numbers on the UI 
- [] Roles 
- [] Analystics

# Extra 

## Roles

### Owners

Should not vote, but rather:
- Set the storie to be voted data
- Configure tags
- Manually trigger the show button
- Get analytics

### DEV / QA

Roles that exist solely to allow for role based voting. AKA, QA only story.

Ideally, when all the required roles present vote, the votes will be made visible by default. 

## Analytics 

TBD, but I am thinking in stuff like:

- User that diverge the most 
- User that average the lowest and highest points
- Quickest and slowest to vote


## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## MakeFile

Run build make command with tests
```bash
make all
```

Build the application
```bash
make build
```

Run the application
```bash
make run
```
Create DB container
```bash
make docker-run
```

Shutdown DB Container
```bash
make docker-down
```

DB Integrations Test:
```bash
make itest
```

Live reload the application:
```bash
make watch
```

Run the test suite:
```bash
make test
```

Clean up binary from the last build:
```bash
make clean
```
