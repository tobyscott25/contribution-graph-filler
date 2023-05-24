# Contribution Graph Filler

A fun little project that fills out your GitHub contributions graph.

> **Note:** I don't use this personally, as I take pride in genuine commit history. This was a fun experiment, but using it undermines the integrity of authentic contributions. As such, I strongly discourage the use of this tool for any reason other than learning and experimenting.

![Example](screenshot.png)

## Usage

1. Create a new private repository on GitHub
2. Ensure your contributions graph is configured to show private contributions
3. Clone this repository and run the following commands:

```bash
go run main.go <path/to/private/repository>
```
4. Push the generated commits to your private repository
