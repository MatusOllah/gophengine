<p align="center">
    <img src="https://github.com/MatusOllah/gophengine/blob/main/docs/bf-gopher_240x320.png" alt="GophEngine logo">
</p>

# 🎤 GophEngine

**[English 🇺🇸](https://github.com/MatusOllah/gophengine/blob/main/README.md)** | **Slovenčina 🇸🇰**

**Stav:** Vo vývoji

[![Go Reference](https://pkg.go.dev/badge/github.com/MatusOllah/gophengine.svg)](https://pkg.go.dev/github.com/MatusOllah/gophengine) [![Go Report Card](https://goreportcard.com/badge/github.com/MatusOllah/gophengine)](https://goreportcard.com/report/github.com/MatusOllah/gophengine) [![GitHub license](https://img.shields.io/github/license/MatusOllah/gophengine)](https://github.com/MatusOllah/gophengine/blob/main/LICENSE) [![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-2.1-4baaaa.svg)](https://github.com/MatusOllah/gophengine/blob/main/CODE_OF_CONDUCT.md) [![Made in Slovakia](https://raw.githubusercontent.com/pedromxavier/flag-badges/refs/heads/main/badges/SK.svg)](https://www.youtube.com/watch?v=UqXJ0ktrmh0)

**GophEngine** (skrátene pre "Gopher Engine") je Go implementácia hry **Friday Night Funkin'** (populárnej FOSS hudobnej hry) s vylepšeniami, ktoré prinášajú nové funkcie, lepší výkon a čistejší zážitok. GophEngine je vytvorený od základov v **Go a Ebitengine** s cieľom poskytnúť ľahko použiteľný framework na modifikáciu a podporovať komunitu fanúšikov FNF.

> [!NOTE]
> Toto nie je originálna hra. Ide o fanúšikovskú prerobenú verziu pôvodnej hry, ktorá nie je nijako spojená alebo schválená tímom Funkin' Crew alebo Newgrounds.
> Predstavte si GophEngine vo vzťahu k FNF tak, ako je Black Mesa k Half-Life - fanúšikovský remake a vylepšenie, vybudované od základu.

## Prečo?

GophEngine som vytvoril z niekoľkých dôvodov:

* Pre vytvorenie jednoducho použitelného Go moddingového MDK pre Friday Night Funkin', podobné MinecraftForge, ale pre FNF.
* Pre opravenie chýb v pôvodnom FNF enginu a pridať nové funkcie a vylepšenia.
* Pre podporu pozitívnej a podporujúcej komunity okolo hry, bez toxicity, ktorá často prevláda v iných moddingových kruhoch.
* Pre odstránenie nevhodného alebo nadmerne grafického obsahu z neskorších verzií FNF a vytvorenie čistejšej a prístupnejšej verzie hry.
* Pre prepísanie hry v Go, mojom obľúbenom programovacom jazyku.

### Prečo Go?

Go je môj obľúbený programovací jazyk a preferujem ho pred učením sa Haxe.
Ponúka lepší výkon, jednoduchšiu prácu s paralelizmom a rýchlejšie časy kompilácie v porovnaní s Haxe, čo ho robí ideálnym pre rýchlu rytmickú hru, ako je FNF.

## Plánované Funkcie

* Zvýšené súkromie bez integrácie Newgrounds
* Úplne napísané v Go so žiadnym použitím Haxe
* Prívetivé pre rodiny
* Drobné úpravy a optimalizácie pre lepší a príjemnejší zážitok
* Nižšia spotreba RAM a celkovo menšia veľkosť
* Riadne menu možností pre lepšie prispôsobenie
* Robustné Go moddingové MDK pre jednoduché modifikácie
* Rýchlejšie časy kompilácie
* ...a mnoho ďalšieho! 😉

## Hardvérové požiadavky

| Komponent        | Minimálne                                                                          | Odporúčané                                                             |
|------------------|------------------------------------------------------------------------------------|------------------------------------------------------------------------|
| Procesor         | Intel Core i3 / AMD Ryzen 3                                                        | Intel Core i5 / AMD Ryzen 5                                            |
| Pamäť            | 4 GB                                                                               | 8 GB                                                                   |
| Grafická karta   | Intel HD Graphics 4000 / NVIDIA GeForce GTX 600 Series / AMD Radeon HD 7000 Series | Intel HD Graphics 5000 / NVIDIA GeForce GTX 750 Ti / AMD Radeon RX 560 |
| DirectX / OpenGL | DirectX 11 / OpenGL 3.0                                                            | DirectX 12 / OpenGL 4.5                                                |
| Úložisko         | 256 MB                                                                             | 512 MB                                                                 |
| Operačný systém  | Windows 7 / macOS 10.12 / Linux Kernel 3.x.x                                       | Windows 11 / macOS 10.15 / Linux Kernel 5.x.x                          |

## Kompilovanie & inštalovanie (zo zdrojového kódu)

Podrobné inštrukcie na buildovanie nájdete v [BUILDING.md](https://github.com/MatusOllah/gophengine/blob/main/docs/BUILDING.md).
Na začiatok budete potrebovať nainštalovaný Go a základné znalosti práce s príkazovým riadkom.

## Menovanie

Názov "GophEngine" kombinuje "Goph" (reprezentujúci Gophera, maskota Go programovacieho jazyka) a "Engine" (reprezentujúci FNF engine).
Tento názov sa mi zdal vhodnejší a prirodzenejší pre tento projekt než alternatívny názov "funkin-go".

## Prispievanie

Radi privítame vaše príspevky! Podrobnosti o tom, ako začať, nájdete v [CONTRIBUTING.md](https://github.com/MatusOllah/gophengine/blob/main/CONTRIBUTING.md).

## Licencia

Licencované podľa **Apache License 2.0** (viď [LICENSE](https://github.com/MatusOllah/gophengine/blob/main/LICENSE))

### Poďakovanie

* **The Funkin' Crew** - Pôvodná hra
* **Hajime Hoshi** - Ebitengine
* [Logo Ebitengine](https://github.com/MatusOllah/gophengine/blob/main/assets/images/ebiten_logo.png) od Hajime Hoshi je licencované podľa [Creative Commons Attribution 4.0](https://creativecommons.org/licenses/by/4.0/).
* Maskot Go Gopher bol vytvorený Renee French a je licencovaný podľa [Creative Commons 4.0 Attribution License](https://creativecommons.org/licenses/by/4.0/).

## 💲 Darujte

Ak vás baví hrať GophEngine a chcete podporiť jeho vývoj, zvážte darovanie. Vaše príspevky a podpora pomáhajú vývoju GophEngine (a mojej láske ku [Kofole](https://kofola.sk)!). Každá podpora je veľmi cenená!

Môžete darovať prostredníctvom nasledujúcich platforiem:

* **Bitcoin (BTC):** `bc1qtykrhm2ar9jreha5rnqve72lutw02jzpu6lcgs`
* **Duino-Coin (DUCO):** `SladkyCitron`
* **Magi (XMG):** `9K8GrfCGEvTK7qjDMVtkGE18UfRyUkv5QT` alebo `SladkyCitron`

Ďakujem za vašu podporu - znamená to pre mňa veľa! 😊❤️
