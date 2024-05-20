<p align="center">
    <img src="https://github.com/MatusOllah/gophengine/blob/main/bf-gopher_240x320.png" alt="GophEngine logo">
</p>

# 🎤 GophEngine

**[English 🇬🇧](https://github.com/MatusOllah/gophengine/blob/main/README.md)** | **Slovenčina 🇸🇰**

[![Go Reference](https://pkg.go.dev/badge/github.com/MatusOllah/gophengine.svg)](https://pkg.go.dev/github.com/MatusOllah/gophengine) [![Go Report Card](https://goreportcard.com/badge/github.com/MatusOllah/gophengine)](https://goreportcard.com/report/github.com/MatusOllah/gophengine)

**GophEngine** je Go implementácia hry **Friday Night Funkin'** s vylepšeniami.

> [!NOTE]
> Toto je mód. Toto není originálna verzia a mal by byť považovaný za modifikáciu.

## Prečo?

Urobil som to z niekoľkých dôvodov:

* Chcel som vytvoriť jednoducho použiteľné Go (moj obľúbený programovací jazyk) moddingové API pre Friday Night Funkin' (moju obľúbenú hru). Niečo ako MinecraftForge, ale pre FNF.
* Tiež som chcel opraviť všetko, čo je zlé vo vanilla FNF engine a pridať niekoľko ďalších funkcií a vylepšení.
* Ako protest proti FNF v0.3.2 .
  * Vytvoril som toto ako odpoveď na FNF v0.3.2 (Weekend 1 update), ktorá zaviedla veľmi násilné a grafické cutscény, v ktorých je môj obľúbený postava, Boyfriend, zabitá. Tieto zmeny významne menia tón a zážitok z hry spôsobom, ktorý mnohí fanúšikovia postavy Boyfriend, vrátane mňa, považujú za nepríjemný.
  * Účelom tohto "forku" je odstrániť tieto násilné a grafické cutscény zavedené vo FNF v0.3.2 a obnoviť postavu Boyfriend a ďalšie naratívne prvky, ktoré boli negatívne ovplyvnené.

### Prečo Go?

Go je môj obľúbený programovací jazyk a nechce sa mi učiť Haxe.

## Menovanie

"Goph" znamená Go Gopher (maskot Go programovacieho jazyku) a "Engine" znamená FNF engine.
Chcel som toto nazvať funkin-go, ale "Spravil som mód s GophEngine" znie prirodzenejšie než "Spravil som mód s funkin-go".

## Kompilovanie & inštalovanie

Pokyny na kompilovanie pre GophEngine sú zatiaľ dostupné v [BUILDING.md](https://github.com/MatusOllah/gophengine/blob/main/BUILDING.md) (slovenský preklad som ešte nespravil).

## Pozoruhodné nástroje a knižnice

* [Ebitengine](https://github.com/hajimehoshi/ebiten) - grafika a vstup
* [Beep](https://github.com/gopxl/beep) - audio
* [ganim8](https://github.com/yohamta/ganim8) - animácie
* [go-winres](https://github.com/tc-hib/go-winres) - vkladanie .ico súborov
