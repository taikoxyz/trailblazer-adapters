# Trailblazer Whitelisted Protocols

## Adding a New Protocol to the Whitelist

### 1. Modify `protocols.json`

Open `protocols.json` and add a new protocol entry to the whitelist.

### 2. Example Protocol Entry

A new protocol entry needs atleast a name, slug and its contract addresses. Optionally, add twitter and logo reference. Logo must be stored under `/img`.

```json
{
    "name": "Name of Protocol",
    "slug": "name-of-protocol",
    "contracts": [
        "0x0"
        ...
    ]
}
```

### 3. Create a Pull Request

Create a pull request (PR) on GitHub to merge your changes into the main branch. Provide a clear description of the changes and the protocol added.

### 6. Review and Merge

Wait for the PR to be reviewed by the maintainers. Once approved, your changes will be merged, and the new protocol will be added to the whitelist.
