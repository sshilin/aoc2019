Run each day:

```sh
for dir in "$(go env GOPATH)"/src/github.com/sshilin/aoc2019/day*; do \
  echo "$dir"; (cd "$dir" && exec go run .); \
done
```