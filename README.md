# nanoflux-plugin-instagram

A [nanoflux](https://github.com/metruzanca/nanoflux)↗ external plugin for
Instagram profile feeds.

Instagram has no public feed. This plugin reads the recent-post grid from the
profile page as a crawler user-agent and turns it into a normal feed. It is
best-effort: a profile without a reachable anonymous grid comes back empty.

## Install

Build the binary into nanoflux's `plugins/` directory (`NF_PLUGINS_DIR`,
`/plugins` in the container):

```bash
cd nanoflux/plugins/instagram
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ../nanoflux-plugin-instagram .
```

The binary lands in `nanoflux/plugins/`. Restart nanoflux; the log should show
`registered external plugin instagram`.

Adjust `GOARCH` if your host is not `amd64` (use `arm64` for ARM).

## Usage

Add a feed with a profile URL:

```
https://www.instagram.com/<handle>
```

## Notes

- Only recent posts are fetched. There is no "load older items" paging.
- Login-restricted, age-gated, and private profiles are not supported.
- Instagram blocks many datacenter IP ranges; a 401/429 may mean the host's
  egress IP is blocked.
