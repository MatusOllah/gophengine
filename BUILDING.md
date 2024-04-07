# Building & Installing

Instead of painstakingly installing dependencies and debugging broken Haxe code, installation is completely pain-free.

## Prerequisites

GophEngine requires 3 basic elements to be present: the **Go toolchain (version 1.22 or later)**, a **C compiler** and a **system graphics driver**.
A C Compiler is required as GophEngine uses not only Go but also C.

Note that **this is just required for development**; your GophEngine FNF mods **will not require any setup or dependency installation for end users!**

### Windows

1. Install [Go](https://go.dev).
2. Install a C Compiler. The easiest way to install a C Compiler on Windows is to use something like [MSYS2](https://www.msys2.org/), [TDM-GCC](https://jmeubank.github.io/tdm-gcc/download/) or [w64devkit](https://github.com/skeeto/w64devkit).
3. On Windows the graphics driver will already be installed, but it is recommended to ensure they are up to date.

The steps for installing with MSYS2 (recommended) are as follows:

1. Install MSYS2 from [msys2.org](https://www.msys2.org/).
2. Once installed do not use the MSYS terminal that opens
3. Open "MSYS2 MinGW 64-bit" from the start menu
4. Run the following commands (if asked for install options be sure to choose "all"):
    * `pacman -Syu`
    * `pacman -S git mingw-w64-x86_64-toolchain`

5. You will need to add `/c/Program\ Files/Go/bin` and `~/Go/bin` to your `$PATH`, for MSYS2 you can run the following command:
    * `echo "export PATH=\$PATH:/c/Program\ Files/Go/bin:~/Go/bin" >> ~/.bashrc`

### macOS X

1. Install [Go](https://go.dev).
2. Install Xcode from the [Mac App Store](https://apps.apple.com/us/app/xcode/id497799835?mt=12).
3. Set up the Xcode command line tools by opening a Terminal window and typing the following:
    * `xcode-select --install`
4. On macOS the graphics driver will already be installed.

### Linux

1. Install [Go](https://go.dev).
2. Install a C Compiler using your distribution's package manager. For example, Ubuntu (or other Debian based distros) uses `apt`.
    * `sudo apt install gcc automake`
3. Install the graphics library header files using your distribution's package manager. One of the following commands should work.
    * **Debian / Ubuntu / Linux Mint:** `sudo apt install libc6-dev libgl1-mesa-dev libxcursor-dev libxi-dev libxinerama-dev libxrandr-dev libxxf86vm-dev libasound2-dev pkg-config`
    * **Fedora / RHEL:** `sudo dnf install mesa-libGL-devel mesa-libGLES-devel libXrandr-devel libXcursor-devel libXinerama-devel libXi-devel libXxf86vm-devel alsa-lib-devel pkg-config`
    * **Solus:** `sudo eopkg install libglvnd-devel libx11-devel libxrandr-devel libxinerama-devel libxcursor-devel libxi-devel libxxf86vm-devel alsa-lib-devel pkg-config`
    * **Arch / Manjaro:** `sudo pacman -S mesa libxrandr libxcursor libxinerama libxi pkg-config`
    * **Alpine:** `sudo apk add alsa-lib-dev libx11-dev libxrandr-dev libxcursor-dev libxinerama-dev libxi-dev mesa-dev pkgconf`

## Building

### using `go install`

Run `go install -v github.com/MatusOllah/gophengine`

### using Makefile

Clone this repo and run `make` in the cloned repo's directory.

Go [here](https://ebitengine.org/en/documents/install.html) for more details.
