# NOTE about the project status

Recently (February 2024) I've decided to finally revive this project. But I'm not really sure if I'll continue developing it, especially since Ebiten isn't really the best framework out there with it's heavy use of singletons and poor optimization (blank map using almost 500 MB!). Maybe I'll rewrite it in Godot or something else with better community, performance, expressiveness and less imposing of sketchy code standards. But there are presonal projects that matter to me more, like [LibRate](https://codeberg.org/mjh/LibRate) and I want to focus my limited time on that.

# Dopewars 2D
[![Go Report Card](https://goreportcard.com/badge/github.com/154pinkchairs/dopewars2d)](https://goreportcard.com/report/github.com/154pinkchairs/dopewars2d)
[![CircleCI](https://dl.circleci.com/status-badge/img/gh/154pinkchairs/dopewars2d/tree/main.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/gh/154pinkchairs/dopewars2d/tree/main)

This game is being created with the intention of being the best Dopewars game ever. Initially it was meant to be just a text-based adventure,
but I've decided to use [Ebiten](https://github.com/hajimeoshi/ebiten)

You are a small time drug dealer in the city of New York. After failing one job after another, you have decided to start a small business.
This game is going to feature:
- Traveling between districts
- Buying weapons
- Reputation and wanted level
- Combat, where you can deal melee damage, throw weapons and of course shoot guns at rival dealers, customers trying to rob you and law enforcers trying to arrest you.
- A bank, a loan shark and a hospital.
- Various random events.

Roadmap:
- [x] Create main menu
- [x] Declare base game mechanics
- [x] World assets – use [iso tiles by Screaming Brain Studios (public domain)](https://screamingbrainstudios.itch.io/)
- [x] Entity assets: WIP
- [x] Implement a basic map and movement
- [ ] Create maps for other districts, add travel between them
- [ ] Implement trading, reputation and wanted level into frontend
- [ ] Bring combat and weapons trading into fe, add hospitals
- [ ] Online mode!!
- [ ] Web frontend
- [ ] Random events
