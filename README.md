# ebitengine-turn-concurrency

A test using [Ebitengine](https://ebitengine.org/) exploring concurrency patterns in Go and game development for a strategy game

## Current Features
- [x] Grid based because i'd like to keep it simple without the complexity of hex grids
- [x] Player movement locked to the gird cells

## Goals
- [ ] Add backend server to handle game state and player movement. gRPC???
- [ ] Add enemy movement system that is run by a backend server
- [ ] Implement attack/action system
- [ ] Add health/stats system

## Stretch Goals
- [ ] Add turn count???
- [ ] Movement range indicators (nice to have)
- [ ] Explore save states
- [ ] Shut the door on this and make something else entirely using what ive learned