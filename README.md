# LuluFrame

LuluFrame is a lightweight local photo player built with Wails. It plays images from a folder configured by `config.json`, including static images and GIF animations.

## Usage

Run LuluFrame once to generate the default files next to the executable:

- `config.json`
- `config说明.txt`
- `photos/`

Put images in `photos/`, or edit `photo_path` in `config.json` to point to another folder.

## Development

Run in live development mode:

```bash
wails dev
```

Build a redistributable package:

```bash
wails build
```
