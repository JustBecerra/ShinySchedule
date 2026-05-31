# ShinySchedule

CLI tool for viewing and managing a weekly schedule from the terminal.

## Run

```bash
go run . <command>
```

Or build once and run the binary:

```bash
go build -o shinyschedule .
./shinyschedule <command>
```

## Commands

| Command | Description |
|---------|-------------|
| `today` | Show today's tasks |
| `day <day>` | Show tasks for a day (e.g. `day wednesday`) |
| `week` | Show the full week |
| `add <day> <start> <end> <activity>` | Add a task for this week (e.g. `add friday 14:00 15:00 dentist`) |
| `remove <day> <index>` | Remove a task by its number from the day view |
| `reset` | Delete all tasks (asks for confirmation if defaults exist) |
| `help` | List available commands |

## Notes

- **Default tasks** live in `defaults.json` and repeat every week.
- **Extra tasks** are saved in `extratasks.json` and reset automatically after 7 days.
- Tasks are sorted by start time. Overlapping times are not allowed when adding.
- Removing a default task asks for confirmation. Use `reset` to clear everything.
