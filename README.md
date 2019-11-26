# Go macOS

Golang exploration exposing macOS UI via CGO

This currently an exploration, using CGO against the NSInvocation API.

## What Works

- A basic app with a window and menu bar, that is all.

## Known Issues

- Non-primitives in the argument list breaks the void** parameter.
