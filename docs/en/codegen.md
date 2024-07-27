## Codegen
A lot of operations in the project require codegen, such as generating APIs from swagger documentation, generating ent schema, and generating dependency injection implementations. 

The following commands can be used to generate the required code.

### Generate the swagger documentation
```bash
just openapi
```

### Generate the ent schema
```bash
just ent
```

### Generate dependency injection
```bash
just wire
```

### Generate all
All the above codegen commands can be combined into one:
```bash
just
```