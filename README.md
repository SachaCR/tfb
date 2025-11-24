# Neck

Neck is CLI that display Guitar chords and scales diagrams. It is built to allow supporting any kind of instrument that has string with frets.

It's a tool for the true nerds. Those who never quit their terminal. Now they will be able to stay in the terminal even when practicing guitar.

# Chords

```bash
$ neck chord 0-2-2-0-0-0 
```
Will display:

```bash
E 0╓────┬────┬────┬─
B 0╟────┼────┼────┼─
G 0╟────┼────┼────┼─
D  ╟────┼─⬤─┼────┼─
A  ╟────┼─⬤─┼────┼─
E 0╙────┴────┴────┴─
```

Print chords that you pass in argument. Example: 0-2-2-0-0-0

Usage:
  neck chord [flags]

Flags:

  `-h`, `--help`          help for chord

  `-n`, `--name` string   Give a chord name like Major7 or m7b5

  `-r`, `--root` string   Set the root of your chord


# Scales 

```bash
$ ./neck scale C-D-E-F-G-A-B
```
Will display:
```bash
E  ╟─⬤─┼────┼─⬤─┼────┼─⬤─┼────┼─⬤─┼─🔴─┼────┼─⬤─┼────┼─⬤─┼─
B  ╟─🔴─┼────┼─⬤─┼────┼─⬤─┼─⬤─┼────┼─⬤─┼────┼─⬤─┼────┼─⬤─┼─
G  ╟────┼─⬤─┼────┼─⬤─┼─🔴─┼────┼─⬤─┼────┼─⬤─┼─⬤─┼────┼─⬤─┼─
D  ╟────┼─⬤─┼─⬤─┼────┼─⬤─┼────┼─⬤─┼────┼─⬤─┼─🔴─┼────┼─⬤─┼─
A  ╟────┼─⬤─┼─🔴─┼────┼─⬤─┼────┼─⬤─┼─⬤─┼────┼─⬤─┼────┼─⬤─┼─
E  ╟─⬤─┼────┼─⬤─┼────┼─⬤─┼────┼─⬤─┼─🔴─┼────┼─⬤─┼────┼─⬤─┼─
               .         .         .         .              ..
```

Print scale that you pass in argument. Example: C-D-E-F-G-A-B

Usage:
  neck scale [flags]

Flags:

  `-h`, `--help`          help for scale

  `-n`, `--name` string   Set the scale name. Example: C Major

  `-r`, `--root` string   Set the root note of your scale

