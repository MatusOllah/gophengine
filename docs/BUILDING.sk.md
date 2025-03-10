# 🔨 Budovanie a Inštalácia

Nasledujúce inštrukcie Vás prevedú procesom nastavenia a budovania GophEngine zo zdrojového kódu. Toto nastavenie je navrhnuté tak, aby bolo jednoduchšie v porovnaní s inými FNF enginami.
Namiesto námahy s inštaláciou závislostí a ladením nefunkčného Haxe kódu je inštalácia úplne bezproblémová.

## Závislosti

GophEngine vyžaduje, aby boli prítomné tieto kľúčové prvky: **nástroje pre Go (verzia go1.23.0 alebo novšia)**, **C kompilátor**, **Make**, **grafický ovládač systému** a **UPX** (voliteľné, povinné iba pre budovanie finálnej verzie).

C kompilátor je potrebný, pretože GophEngine používa nielen Go, ale aj jazyk C.

### 🪟 Windows

1. Nainštalujte [Go](https://go.dev) (verzia go1.23.0 alebo novšia).
2. Nainštalujte C kompilátor. Najjednoduchší spôsob, ako nainštalovať C kompilátor na Windowse, je použiť niečo ako [MSYS2](https://www.msys2.org/), [TDM-GCC](https://jmeubank.github.io/tdm-gcc/download/) alebo [w64devkit](https://github.com/skeeto/w64devkit). MSVC nie je podporovaný, viď [Go issue #20982](https://github.com/golang/go/issues/20982).
3. Na Windowse bude grafický ovládač už nainštalovaný, ale odporúča sa overiť, či je aktuálny.

#### Postup pre inštaláciu cez MSYS2 (odporúčané)

1. Nainštalujte [MSYS2](https://www.msys2.org/).
2. Po inštalácii neotvárajte terminál MSYS.
3. Otvorte "MSYS2 MinGW 64-bit" z ponuky štartovacieho menu.
4. Spustite nasledujúce príkazy:
    * `pacman -Syu`
    * `pacman -S git mingw-w64-x86_64-gcc mingw-w64-x86_64-make`
5. Budete musieť pridať `/c/Program\ Files/Go/bin` a `~/go/bin` do vášho `$PATH`, pre MSYS2 môžete použiť nasledujúci príkaz:
    * `echo "export PATH=\$PATH:/c/Program\ Files/Go/bin:~/go/bin" >> ~/.bashrc`

### 🍎 macOS

1. Nainštalujte [Go](https://go.dev) (verzia go1.23.0 alebo novšia).
2. Nainštalujte Xcode z [Mac App Store](https://apps.apple.com/us/app/xcode/id497799835?mt=12).
3. Nastavte Xcode otvorením terminálu a napísaním nasledujúceho:
    * `xcode-select --install`
4. Na macOS bude grafický ovládač už nainštalovaný.

### 🐧 GNU/Linux

1. Nainštalujte [Go](https://go.dev) (verzia go1.23.0 alebo novšia).
2. Nainštalujte C kompilátor a Make pomocou správcu balíkov vašej distribúcie. Napríklad Ubuntu (alebo iné distribúcie založené na Debiane) používa `apt`.
    * `sudo apt install gcc make`
3. Nainštalujte hlavičkové súbory grafickej knižnice pomocou správcu balíkov vašej distribúcie. Použite príslušný príkaz pre vašu distribúciu:
    * **Debian / Ubuntu / Linux Mint:** `sudo apt install libc6-dev libgl1-mesa-dev libxcursor-dev libxi-dev libxinerama-dev libxrandr-dev libxxf86vm-dev libasound2-dev pkg-config`
    * **Fedora / RHEL:** `sudo dnf install mesa-libGL-devel mesa-libGLES-devel libXrandr-devel libXcursor-devel libXinerama-devel libXi-devel libXxf86vm-devel alsa-lib-devel pkg-config`
    * **Solus:** `sudo eopkg install libglvnd-devel libx11-devel libxrandr-devel libxinerama-devel libxcursor-devel libxi-devel libxxf86vm-devel alsa-lib-devel pkg-config`
    * **Arch / Manjaro:** `sudo pacman -S mesa libxrandr libxcursor libxinerama libxi pkg-config`
    * **Alpine:** `sudo apk add alsa-lib-dev libx11-dev libxrandr-dev libxcursor-dev libxinerama-dev libxi-dev mesa-dev pkgconf`

## Budovanie

### Pomocou `go install`

Spustite nasledujúci príkaz:

```sh
go install -v github.com/MatusOllah/gophengine/cmd/gophengine@latest
```

### Pomocou Makefile

1. Naklonujte tento repozitár:
    * `git clone https://github.com/MatusOllah/gophengine.git`
2. Spustite `make` v priečinku naklonovaného repozitára:
    * `cd ./gophengine/ && make`

### Pomocou budovacích (build) skriptov

1. Naklonujte tento repozitár:
    * `git clone https://github.com/MatusOllah/gophengine.git`
2. Prejdite do priečinka naklonovaného repozitára:
    * `cd ./gophengine/`
3. Spustite skript:
    * pre vývojovú verziu (debug): `scripts/build-debug.sh`
    * pre finálnu verziu (release) (s `-ldflags="-s -w"` a UPX kompresiou): `scripts/build-release.sh`

Pre viac informácií navštívte [sprievodcu inštalácie Ebitengine](https://ebitengine.org/en/documents/install.html).
