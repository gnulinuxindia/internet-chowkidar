# Migrations
Every time you change the database schema, you need to create a new migration. This is done by running the following command:

```bash
just migration-create <migration_name>
```

This will create a new migration file in the `migrations` directory. The migration file will contain two functions: `Up` and `Down`. The `Up` function is used to apply the migration, while the `Down` function is used to rollback the migration.

Migrations will be applied automatically when you run the application.