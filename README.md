<p align="center">
    <img src="https://github.com/MatusOllah/gophengine/blob/main/bf-gopher_240x320.png" alt="GophEngine logo">
</p>

# 🎤 GophEngine

**English 🇬🇧** | **[Slovenčina 🇸🇰](https://github.com/MatusOllah/gophengine/blob/main/README.sk-SK.md)**

Status: work-in-progress

[![Go Reference](https://pkg.go.dev/badge/github.com/MatusOllah/gophengine.svg)](https://pkg.go.dev/github.com/MatusOllah/gophengine) [![Go Report Card](https://goreportcard.com/badge/github.com/MatusOllah/gophengine)](https://goreportcard.com/report/github.com/MatusOllah/gophengine)

**GophEngine** is a Go implementation of **Friday Night Funkin' v0.2.7.1** with improvments.

> [!NOTE]
> This is a mod. This is not the vanilla game and should be treated as a modification.

## Why?

I made this for a couple of reasons:

* I wanted to make a easy-to-use Go (my fav programming language) modding API for Friday Night Funkin' (my fav game). Something like MinecraftForge but for FNF.
* I also wanted to fix everything what's wrong with the vanilla FNF engine and add some more features and improvments.
* I hate the toxic community & developers and I wanted to make my own FNF clone.
* As a protest against FNF v0.3.2 .
  * I made this in response to FNF v0.3.2 (the Weekend 1 update), which introduced highly violent and graphic NSFW/L cutscenes where my favorite character, Boyfriend, is killed, and replaced with Pico. These changes significantly alter the tone and experience of the game in a way that many Boyfriend fans, including myself, find upsetting.
  * The purpose of this "fork" is to rollback FNF v0.3.2, to restore Boyfriend and other narrative elements that were negatively impacted and to create a generally "cleaner" version of the game.
* I just wanted to rewrite Friday Night Funkin' in Go.

### Why Go?

Go is my favorite programming language and I don't wanna learn Haxe.

## Naming

"Goph" means Go Gopher (Go programming language mascot) and "Engine" means FNF engine.
I wanted to call this funkin-go, but "I made a mod with GophEngine" just sounds more natural than "I made a mod with funkin-go".

## Building & installing (from source)

Build instructions for GophEngine are available in [BUILDING.md](https://github.com/MatusOllah/gophengine/blob/main/BUILDING.md).

## Notable tools and libraries

* [Ebitengine](https://github.com/hajimehoshi/ebiten) - graphics and input
* [Beep](https://github.com/gopxl/beep) - audio
* [ganim8](https://github.com/yohamta/ganim8) - animations
* [go-winres](https://github.com/tc-hib/go-winres) - embedding .ico files
