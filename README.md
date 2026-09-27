# go-tebanare-sample

A small Go module for trying out [go-tebanare](https://github.com/miyamo2/go-tebanare): a Chrome extension that hides Go code from the "Files changed" tab of a GitHub pull request when a team has agreed that the code needs no line-by-line review.

The repository is a toy task-tracker API, layered the way a real service would be: domain entities, use cases, a repository port with an in-memory adapter, and an HTTP handler. `.gotebanare.yml` at the root enables all three built-in presets:

```yaml
version: 1
presets:
  - getter
  - noop
  - iferr
```

## Try it

1. Install the go-tebanare Chrome extension.
2. Open a pull request against this repository that touches one of the files below, for example by adding a field and its getter to `domain.Task`, or by adding an `if err != nil { return err }` guard to a use case.
3. Look at the "Files changed" tab: lines the presets recognize are folded, and everything else — including the contrasting examples below — stays visible.

## Layers

| Layer | Package | Role |
|---|---|---|
| Domain | `domain` | `Task` and `Project` entities, the `TaskRepository` port, and domain errors. No dependencies. |
| Use case | `usecase` | `TaskUseCase` (create, get, complete) and `RecentlyViewedUseCase`, both depending only on `domain` and the generic `pkg/collection.Stack`. |
| Repository | `repository` | `InMemoryTaskRepository`, an adapter implementing `domain.TaskRepository`. |
| Handler | `handler` | `TaskHandler`, an adapter exposing the use cases over `net/http`. |
| Infrastructure | `infra/logger` | `Noop`, a logging adapter that discards every message. |
| Composition root | `cmd/server` | Wires the layers together and starts the HTTP server. |

## What each file demonstrates

| File | Preset | Notes |
|---|---|---|
| `domain/task.go` | `getter` | `ID`, `ProjectID`, `Title`, `Description`, `Status`, `AssigneeID`, `CreatedAt`, `DueAt` match. `Assign` (takes a parameter) and `IsOverdue` (compares two values) do not. |
| `domain/project.go` | `getter` | `ID`, `Name`, `OwnerID`, `CreatedAt` all match. |
| `usecase/task_usecase.go` | `iferr` | `CreateTask` and `CompleteTask` each have one plain `if err != nil { return err }` guard that matches, and `CompleteTask`'s second guard wraps the error with `fmt.Errorf`, so it stays visible. |
| `usecase/recently_viewed_usecase.go` | `getter` | `Count` does not match: it calls `Len()` on a field instead of returning a field directly, the same shape as the `u.cfg.Limit()` example in the docs. |
| `repository/memory_task_repository.go` | `noop` | `Migrate` matches; `Close` does not, since it returns a value. |
| `infra/logger/noop.go` | `noop` | `Flush` matches; `Debug`, `Info`, `Error` do not, since each takes a parameter. |
| `pkg/collection/stack.go` | `getter` | `Len` matches on a generic receiver; `Peek` does not (two results, indexes a field). |
| `handler/task_handler.go` | `iferr` | Both handlers' guards call `http.Error` before `return`, so neither matches: `iferr` only hides a body of exactly one `return` statement. |
