# Trailblazer Whitelisted Protocols and Adapters

## Adding a new Protocol to the Whitelist

### 1. Modify `protocols.json`

Open [`protocols.json`](./whitelist/protocols.json) and add a new protocol entry to the whitelist.

### 2. Example Protocol Entry

A new protocol entry needs atleast a name and a slug to be recognized as part of the ecosystem. When competing in trailblazer DApp competition add your contracts' addresses. Optionally, add a short description, ecosystem category, twitter, website and logo reference. For existing ecosystem categories, see [`whitelist.go`](./whitelist/whitelist.go) Logo must be stored under `./img`. Ensure your slug is consistent with your project's slug used in Defillama if you are integrated.

```json
{
    "name": "Name of Protocol",
    "slug": "name-of-protocol",
    "contracts": [
        "0x0"
        ...
    ],
    "twitter": "@optional-protocol-twitter",
    "logo": "optional-protocol-logo.jpg",
    "description": "An optional short description",
    "wesbite": "optional-website.xyz",
    "category": "optional-ecosystem-category"
}
```

### 3. Create a Pull Request

Create a pull request (PR) on GitHub to merge your changes into the main branch. Provide a clear description of the changes and the protocol added.

### 4. Review and Merge

Wait for the PR to be reviewed by the maintainers. Once approved, your changes will be merged, and the new protocol will be added to the whitelist.

## Adding a new Adapter

### 1. Add your Project under `adapters/projects`

After the details of the indexer has been confirmed with the team, add your project's indexer in the `adapters/projects` folder. See `adapters/adapters.go` for interfaces and types.

### 2. Write your Indexer

An example adapter for tracking OrderFulfilled Event on the OKX marketplace can be seen [here](./adapters/projects/okx/order_fulfilled.go) along with its accompanying test file.

### 3. Manually Test your Adapter by adding it to the CLI

First, add a new adapter in `cmd/adapter.go`. Second, add adapter to switch case of `cmd/cmd.go`'s function `executeCommand`. Test out the CLI with `go run .` from project root.

### 4. Add a Test to your Indexer

Add at least one test case to your project's adapter. For an example, see [here](./adapters/projects/okx/order_fulfilled_test.go).

### 5. Create a Pull Request

Create a pull request (PR) on GitHub to merge your changes into the main branch. Provide a clear description of the changes and the protocol added.

### 6. Review and Merge

Wait for the PR to be reviewed by the maintainers. Once approved, your changes will be merged, and the protocol info will be added to trailblazers.
