# Contribution Graph Filler

A fun little project that fills out your GitHub contributions graph.

> **Warning:** This was a fun experiment, I discourage actually using this unless you are happy for people to question the integrity of your authentic contributions.

![Example](screenshot.png)

## Guide

1. Create a new private repository on GitHub
2. Ensure your contributions graph is configured to show private contributions
3. Clone this repository and run the following commands:

```bash
go run main.go <path/to/private/repository>
```
4. Push the generated commits to your private repository
