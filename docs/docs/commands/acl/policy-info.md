# Command: acl policy info

The `acl policy info` command is used to display detailed information about an existing ACL policy.

## Usage

```
drago acl policy info [options] <name>
```

## General Options

- `-address=<addr>`
    The address of the Drago server.
    Overrides the DRAGO_ADDR environment variable if set.
    Defaults to `http://127.0.0.1:8080`


- `-token=<token>`
    The token used to authenticate with the Drago server.
    Overrides the `DRAGO_TOKEN` environment variable if set.
    Defaults to `""`

## List Options

- `-json`: Enable JSON output.