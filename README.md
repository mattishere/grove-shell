# Grove shell
A simple Linux shell written in Go.

> [!IMPORTANT]
> Grove is still in the early stages of (active) development. Therefore, it is **not production-ready**.

## Installation
> [!WARNING]
> We don't take accountability for any damage caused by the program in its current state. It is simply **not production-ready**.s

- The easiest way to install Grove at the moment is using the script below. It simply builds the program (this requires you to have the **Go** compiler installed) and stores it in `/usr/local/bin/grove`.

```bash
git clone https://github.com/groveshell/grove-shell && cd grove-shell && sudo make install
```

- Another way to install Grove is by installing a binary from [here](https://github.com/groveshell/grove-shell/releases). At the moment, they're only for **amd64** Linux systems.
> [!NOTE]
> At the moment, releases are only made periodically. If you want the bleeding-edge, then the method above is superior.

## Usage
- After installing Grove, you can start a session/REPL with the command `grove`. After that, it's fairly straight forward as it is a very similar experience to most other Linux shells.

## Features
> [!IMPORTANT]
> At the time of writing this, Grove is still missing many important features. Due to the current **frequency** of updates to the project, this list will not be updated until it reaches a *"bare minimum"* in functionality and featureset. After that threshold is reached, expect this list to be accurate again!

- File navigation with `cd`
- A basic implemenation of `echo`
- String expansion of environment variables (`$VARIABLE` -> `value`)
- Home directory expansion (`~` -> `/home/user`)
- Strings (`'raw string'`, `"normal string"`)