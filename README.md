# cff-cli
CLI client to SBB/CFF API (for Swiss trains), whitten in Go

## Usage

### Request for now

```cff <origin> <destination>```

Print 5 next relations from origin station to destination

Example: ```cff court tavannes```

### Request for today, particular time

```cff <origin> <destination> <time>```

Example: ```cff court tavannes 18:00```

Print 5 next relations from origin station to destination from today provided time

### Reques for particular date and time

```cff <origin> <destination> <time> <day>```

Print 5 next relations from origin station to destination from provided time and provided day

Example: ```cff court tavannes 18:00 20240731```
