# Trailblazer Adapters CLI

## Adding a New CLI Command

### 1. Add a new option under ```adapterOptions``` in ```root.go```

### 2. Import adapter 

An example on how to add import a newly added adapter [here](./order_fulfilled.go)

### 3. Test out the CLI

Test out the CLI with the following command and input your parameters accordingly.
```
go run main.go process
```

