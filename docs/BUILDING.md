# 🔨 Building & Installing

The following instructions will guide you through the process of setting up and building GophEngine from source. This setup is designed to be straightforward compared to other FNF engines.
Instead of painstakingly installing dependencies and debugging broken Haxe code, installation is completely pain-free.

## Prerequisites

GophEngine requires these key elements to be present: the **Go toolchain (version go1.22.0 or later)**, a **C Compiler**, **Make**, a **system graphics driver**, and **UPX** (optional, required only for building release builds).

A C Compiler is required as GophEngine uses not only Go, but also C.

### 🪟 Windows

1. Install [Go](https://go.dev) (version go1.22.0 or later).
2. Install a C Compiler. The easiest way to install a C Compiler on Windows is to use something like [MSYS2](https://www.msys2.org/), [TDM-GCC](https://jmeubank.github.io/tdm-gcc/download/) or [w64devkit](https://github.com/skeeto/w64devkit). MSVC isn't supported, see [Go issue #20982](https://github.com/golang/go/issues/20982).
3. On Windows the graphics driver will already be installed, but it is recommended to ensure they are up to date.

#### Steps for installing with MSYS2 (recommended)

1. Install [MSYS2](https://www.msys2.org/).
2. Once installed do not use the MSYS terminal that opens.
3. Open "MSYS2 MinGW 64-bit" from the start menu.
4. Run the following commands:
    * `pacman -Syu`
    * `pacman -S git mingw-w64-x86_64-gcc mingw-w64-x86_64-make`
5. You will need to add `/c/Program\ Files/Go/bin` and `~/go/bin` to your `$PATH`, for MSYS2 you can run the following command:
    * `echo "export PATH=\$PATH:/c/Program\ Files/Go/bin:~/go/bin" >> ~/.bashrc`

### 🍎 macOS

1. Install [Go](https://go.dev) (version go1.22.0 or later).
2. Install Xcode from the [Mac App Store](https://apps.apple.com/us/app/xcode/id497799835?mt=12).
3. Set up the Xcode command line tools by opening a Terminal window and typing the following:
    * `xcode-select --install`
4. On macOS the graphics driver will already be installed.

### 🐧 GNU/Linux

1. Install [Go](https://go.dev) (version go1.22.0 or later).
2. Install a C Compiler and Make using your distribution's package manager. For example, Ubuntu (or other Debian based distros) uses `apt`.
    * `sudo apt install gcc make`
3. Install the graphics library header files using your distribution's package manager. Use the appropriate command for your distro:
    * **Debian / Ubuntu / Linux Mint:** `sudo apt install libc6-dev libgl1-mesa-dev libxcursor-dev libxi-dev libxinerama-dev libxrandr-dev libxxf86vm-dev libasound2-dev pkg-config`
    * **Fedora / RHEL:** `sudo dnf install mesa-libGL-devel mesa-libGLES-devel libXrandr-devel libXcursor-devel libXinerama-devel libXi-devel libXxf86vm-devel alsa-lib-devel pkg-config`
    * **Solus:** `sudo eopkg install libglvnd-devel libx11-devel libxrandr-devel libxinerama-devel libxcursor-devel libxi-devel libxxf86vm-devel alsa-lib-devel pkg-config`
    * **Arch / Manjaro:** `sudo pacman -S mesa libxrandr libxcursor libxinerama libxi pkg-config`
    * **Alpine:** `sudo apk add alsa-lib-dev libx11-dev libxrandr-dev libxcursor-dev libxinerama-dev libxi-dev mesa-dev pkgconf`

## Building

### Using `go install`

Run the following command:

```sh
go install -v github.com/MatusOllah/gophengine/cmd/gophengine@latest
```

### Using Makefile

1. Clone this repo:
    * `git clone https://github.com/MatusOllah/gophengine.git`
2. Run `make` in the cloned repo's directory:
    * `cd ./gophengine/ && make`

### Using the build scripts

1. Clone this repo:
    * `git clone https://github.com/MatusOllah/gophengine.git`
2. Move to the cloned repo's directory:
    * `cd ./gophengine/`
3. Run the script:
    * for a debug build: `scripts/build-debug.sh`
    * for a release build (with `-ldflags="-s -w"` and UPX compression): `scripts/build-release.sh`

For more details, visit the [Ebitengine installation guide](https://ebitengine.org/en/documents/install.html).
