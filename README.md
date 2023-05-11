# contribution-graph-filler
Want to improve the look of your contributions graph? It's easy!

```zsh
yesterday=$(date -v -1d +%s) GIT_AUTHOR_DATE="$yesterday" GIT_COMMITTER_DATE="$yesterday" gca
```
