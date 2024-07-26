# Trailblazer Adapters

## Adding a New Protocol to the Whitelist

### 1. Modify `whitelist.go`

Open `whitelist.go` add the new protocol entry to the whitelist.

### 2. Example Protocol Entry

An example entry format for adding a new protocol is:

- Name: Name of Protocol
- Slug: name-of-protocol
- Contracts: 0x0

```json
{
    "name": "Name of Protocol",
    "slug": "name-of-protocol",
    "contracts": [
        "0x0"
    ]
}
```

### 3. Create a Pull Request

Create a pull request (PR) on GitHub to merge your changes into the main branch. Provide a clear description of the changes and the protocol added.

### 6. Review and Merge

Wait for the PR to be reviewed by the maintainers. Once approved, your changes will be merged, and the new protocol will be added to the whitelist.

For further details, refer to the official documentation or contact the maintainers for support.
