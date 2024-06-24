<p align="center">
    <img src="https://github.com/MatusOllah/gophengine/blob/main/bf-gopher_240x320.png" alt="GophEngine logo">
</p>

# 🎤 GophEngine

**[English 🇺🇸](https://github.com/MatusOllah/gophengine/blob/main/README.md)** | **Slovenčina 🇸🇰**

[![Go Reference](https://pkg.go.dev/badge/github.com/MatusOllah/gophengine.svg)](https://pkg.go.dev/github.com/MatusOllah/gophengine) [![Go Report Card](https://goreportcard.com/badge/github.com/MatusOllah/gophengine)](https://goreportcard.com/report/github.com/MatusOllah/gophengine)

**GophEngine** je Go implementácia hry **Friday Night Funkin' v0.2.7.1** s vylepšeniami.

> [!NOTE]
> Toto je mód. Toto není originálna verzia a mal by byť považovaný za modifikáciu.

## Prečo?

GophEngine som vytvoril z niekoľkých dôvodov:

* Pre vytvorenie jednoducho použitelného Go moddingového API pre Friday Night Funkin', podobné MinecraftForge, ale pre FNF.
* Pre opravenie chýb v pôvodnom FNF enginu a pridať nové funkcie a vylepšenia.
* Pre podporu pozitívnej a netoxickej komunity okolo hry bez toxickej komunity a vývojárov, pre podporu fanúšikov postavy Boyfriend a vyhýbanie sa anti-Boyfriend postojom.
* Pre odstránenie násilného NSFL obsahu zavedené v FNF v0.3.2 a pre vytvorenie čistejšej verzie hry.
* Pre prepísanie hry Friday Night Funkin' v Go, mojom obľúbenom programovacom jazyku.

### Prečo Go?

Go je môj obľúbený programovací jazyk a preferujem ho pred učením sa Haxe.

## Funkcie

* Zvýšené súkromie bez integrácie Newgrounds
* Úplne napísané v Go s minimálnym alebo žiadnym použitím Haxe
* Prívetivé pre rodiny a fanúšikov postavy Boyfriend (bez obsahu NSFW/L)
* Drobné úpravy a optimalizácie pre lepší a príjemnejší zážitok
* Riadne menu možností pre lepšie prispôsobenie
* Robustné Go moddingové API pre jednoduché modifikácie

## Hardvérové požiadavky

| Komponent        | Minimálne                                                                          | Odporúčané                                                             |
|------------------|------------------------------------------------------------------------------------|------------------------------------------------------------------------|
| Procesor         | Intel Core i3 / AMD Ryzen 3                                                        | Intel Core i5 / AMD Ryzen 5                                            |
| Pamäť            | 4 GB                                                                               | 8 GB                                                                   |
| Grafická karta   | Intel HD Graphics 4000 / NVIDIA GeForce GTX 600 Series / AMD Radeon HD 7000 Series | Intel HD Graphics 5000 / NVIDIA GeForce GTX 750 Ti / AMD Radeon RX 560 |
| DirectX / OpenGL | DirectX 11 / OpenGL 3.0                                                            | DirectX 11 / OpenGL 4.5                                                |
| Úložisko         | 256 MB                                                                             | 512 MB                                                                 |
| Operačný systém  | Windows 7 / macOS 10.12 / Linux kernel 2.6.32                                      | Windows 10 / macOS 10.15 / Linux kernel 5.x.x                          |

## Kompilovanie & inštalovanie (zo zdrojového kódu)

Pokyny na kompilovanie pre GophEngine sú zatiaľ dostupné v [BUILDING.md](https://github.com/MatusOllah/gophengine/blob/main/BUILDING.md) (slovenský preklad som ešte nespravil).

## Menovanie

Názov "GophEngine" kombinuje "Goph" (reprezentujúci maskot Go programovacieho jazyka) a "Engine" (reprezentujúci FNF engine).
Pôsobí to prirodzenejšie a lepšie zapadá do tohto projektu než "funkin-go."

## Pozoruhodné nástroje a knižnice

* [Ebitengine](https://github.com/hajimehoshi/ebiten) - grafika a vstup
* [Beep](https://github.com/gopxl/beep) - audio
* [ganim8](https://github.com/yohamta/ganim8) - animácie
* [go-winres](https://github.com/tc-hib/go-winres) - vkladanie .ico súborov
