Go standard-library web app

Quick start

1. Build and run:

```powershell
cd d:/code/gohtmx
go run .
```

2. Open http://localhost:8080/

What it includes

- Routing with `http.ServeMux`
- Templating with `html/template` (templates in `templates/`)
 - Static file serving from `static/` (exposed at `/static/`)
 - Local copy of `htmx` is included at `static/js/htmx.min.js` and loaded on every page.

Notes

- Only uses Go standard library. No third-party packages.
- Change the port in `main.go` if needed.
