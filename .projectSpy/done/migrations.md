Migrations
===

I guess it's finally time to build that simple SQLite migration module.

Use sqlite version column to store the migration version in.

Rollback functionality shouldn't be necessary.

---

Ended up going with a schema based migration - based on a [module I was working on](https://github.com/gregdaynes/sqlite-go-migration). Likely the changes here will be merged upstream and then this app can utilize the module instead of the local code.

---

2025-04-19 21:09	Created task
2025-04-24 22:34	Updated task
2025-04-24 22:34	Moved task from backlog to done