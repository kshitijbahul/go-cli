### Adding a command using Cobra
## Eg Adding a command called list use
```go install github.com/spf13/cobra-cli```

``` cobra-cli init github.com/kshitijbahul/go-cli -a "Kshitij Bahul" ```

```cobra-cli add list ```

### To add a new Flag for the commands globally 
-> Add it to the Root file

### To add a new Flag for the commands locally i.e only for a specific command 
-> Add the Flag to the individual file of teh command 

### Add a priority using 
```./go-cli add "Top Priority" -p1```


### Mark a todo as done
```go run main.go done (a todo)```