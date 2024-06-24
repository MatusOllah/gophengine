<p align="center">
    <img src="https://github.com/MatusOllah/gophengine/blob/main/bf-gopher_240x320.png" alt="GophEngine logo">
</p>

# 🎤 GophEngine

**English 🇬🇧** | **[Slovenčina 🇸🇰](https://github.com/MatusOllah/gophengine/blob/main/README.sk-SK.md)**

**Status:** work-in-progress

[![Go Reference](https://pkg.go.dev/badge/github.com/MatusOllah/gophengine.svg)](https://pkg.go.dev/github.com/MatusOllah/gophengine) [![Go Report Card](https://goreportcard.com/badge/github.com/MatusOllah/gophengine)](https://goreportcard.com/report/github.com/MatusOllah/gophengine)

**GophEngine** is a Go implementation of **Friday Night Funkin' v0.2.7.1** with improvments.

> [!NOTE]
> This is a mod. This is not the vanilla game and should be treated as a modification.

## Why?

I created GophEngine for several reasons:

* To develop an easy-to-use Go modding API for Friday Night Funkin', similar to MinecraftForge but for FNF.
* To fix issues in the vanilla FNF engine and add new features and improvements.
* To foster a positive and non-toxic community around the game without the toxic community and developers, supporting Boyfriend fans and avoiding anti-Boyfriend sentiments.
* To remove violent NSFW/L content introduced in FNF v0.3.2, and create a cleaner version of the game.
* To rewrite Friday Night Funkin' in Go, my favorite programming language.

### Why Go?

Go is my favorite programming language, and I prefer it over learning Haxe.

## Naming

The name "GophEngine" combines "Goph" (representing the Go programming language mascot) and "Engine" (representing the FNF engine). It felt more natural and fitting for this project than "funkin-go."

## Features

* Enhanced privacy with no Newgrounds integration
* Entirely written in Go with little to no Haxe
* Family-friendly and welcoming to Boyfriend fans (no NSFW/L content)
* Small tweaks and optimizations for a smoother and more enjoyable experience
* A proper options menu for better customization
* Robust Go modding API for easy modification

## Building & installing (from source)

Build instructions for GophEngine are available in [BUILDING.md](https://github.com/MatusOllah/gophengine/blob/main/BUILDING.md).

## Notable tools and libraries

* [Ebitengine](https://github.com/hajimehoshi/ebiten) - graphics and input
* [Beep](https://github.com/gopxl/beep) - audio
* [ganim8](https://github.com/yohamta/ganim8) - animations
* [go-winres](https://github.com/tc-hib/go-winres) - embedding .ico files
