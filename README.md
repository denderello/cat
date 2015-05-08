# cat
A very simple version of the UNIX tool cat re-written in Go only using syscalls

## Usage
No fancy flags are supported. Just plain `read(2)` of a file's content and
`write(2)` onto `stdout` with some error processing:

```
cat <file>
```
