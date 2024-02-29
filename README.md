# Grove shell
A simple Linux shell written in Go.

> [!IMPORTANT]
> Grove is still in the early stages of (active) development. Currently, it only works on Linux systems (using it on any others will result in unexpected behaviour or even refuse to launch).

## Installation

- The easiest way to install Grove at the moment is using the script below. It simply builds the program (this requires you to have the **Go** compiler installed) and stores it in `/usr/local/bin/grove`.

```bash
git clone https://github.com/mattishere/grove-shell && cd grove-shell && make install
```
Afterwards, you can also use `make register`, which registers the shell in the /etc/shells file, allowing you to set it as your default shell.

- Another way to install Grove is by installing a binary from [here](https://github.com/mattishere/grove-shell/releases). At the moment, they're only for **amd64** Linux systems.

## Usage
- After installing Grove, you can start a session/REPL with the command `grove`.

## Features
- Commands (use `help` to see them, their usage and their syntax)
- String expansion of environment variables (`$VARIABLE` -> `value`)
- Home directory expansion (`~` -> `/home/user`)
- Strings (`'raw string'`, `"normal string"`)
- Scripts and a launch script (`~/.groverc`)
- Shebang exectuable scripts (`#!/path/to/your/grove/binary`, and then `./your/executable/script`)
- Prompt customization by setting the `$PROFILE` environment variable, with placeholders:
<table>
<tr>
    <th>Placeholder<th>
    <th>Description</th>
</tr>
<tr>
    <td>{username}</td>
    <td>Gets the current username of the user</td>
</tr>
<tr>
    <td>{hostname}</td>
    <td>Gets the hostname of the machine</td>
</tr>
<tr>
    <td>{path}</td>
    <td>Gets the current full path</td>
</tr>
<tr>
    <td>{curr_dir}</td>
    <td>Gets the current directory that you are in (without the full path)</td>
</tr>
<tr>
    <td>{reset|bold|underline}</td>
    <td>Different styling options</td>
</tr>
<tr>
    <td>{black|red|green|yellow|blue|magenta|cyan|white}</td>
    <td>Different color options</td>
</tr>
</table>