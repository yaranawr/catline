# catline

## Description

Display the number of lines in files within a directory. That's literally it.

You can choose which file extensions you want the script to read or exclude the extensions of files you don't want to be included.

There's a bash version [here](https://github.com/yannawr/catline-bash).

## Requirements:

Go :)

## Usage

### Instalation

```
go install github.com/yannawr/catline@latest
```

## Commands

To display usage information, run catline with the `-h` or `-help` option:

### Options
```
-e <extensions>   List of extensions to include (comma-separated)
-x <extensions>   List of extensions to exclude (comma-separated)
```
### Examples
```
❯ ls 
catline.go  go.mod  test.txt

❯ catline -e txt
test.txt        2

❯ catline -x txt
catline.go      163
go.mod          3

❯ catline
catline.go      163
go.mod          3
test.txt        2
```
To read all files within a directory, simply run catline without including any additional commands.


That's all :)

