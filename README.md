# Uptime Kuma Bulk Importer

Quick and dirty script to import a list of hosts into Uptime Kuma by generating a backup.json file suitable to restore from the Uptime Kuma web interface.

## Usage

Create a file called `hosts.txt` in the following format:

```txt
host1.example.com
host2.example.com
host3.example.com
...
```

Start the script:

```bash
go run main.go
```

Upload the `backup.json` file to Uptime Kuma, and you're done!
