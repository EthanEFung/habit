# habit
simple habit track cli tool.

# how to install

Application relies on your home directory having a `Documents` folder to publish csv
files to. From the command line run the following

```bash
go mod init github.com/ethanefung/habit

# output the binary to your /bin directory
go build -o $GOPATH/bin/habit
```

# usage
You should have everything you need to get started. From the command
line specify a habit to track using the `habit` command

```bash
habit [your habit to track] 
```
The application will save the habit that you've specified, and should show you a
monthly view of every day you've exercised your habit.
