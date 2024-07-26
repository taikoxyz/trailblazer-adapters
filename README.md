# Trailblazer Adapters

## Adding a New Protocol Adapter

### 1. Add your project under adapters/projects folder

### 2. Write an indexer in the new folder

An example adapter for tracking OrderFulfilled Event on the OKX marketplace can be seen [here](./adapters/projects/okx/order_fulfilled.go)

### 3. Test out the adapter by adding it to the cli

An example on how to add a new adapter to the cli is [here](./cmd/README.md)

### 4. Create a Pull Request

Create a pull request (PR) on GitHub to merge your changes into the main branch. Provide a clear description of the changes and the protocol added.

### 5. Review and Merge

Wait for the PR to be reviewed by the maintainers. Once approved, your changes will be merged, and the protocol info will be added to trailblazers. 

For further details, refer to the official documentation or contact the maintainers for support.
