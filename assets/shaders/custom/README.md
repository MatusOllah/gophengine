# Custom shaders

This directory is intented for custom Kage shaders, loaded when `Graphics.EnableCustomShaders` (from options.gecfg) is enabled.

Available uniform variables:

* `UnixTime` - The unix time. (`int`)
* `CursorPosition` - The cursor's position. (`vec2`)
* `Random` - A random floating-point number in the half-open interval $[0.0,1.0)$. (`float`)
