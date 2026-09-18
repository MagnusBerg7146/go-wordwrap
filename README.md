# Greedy word wrap

```
wordwrap.go
```
Run the Go test suite right next to the source files to see how it handles tricky edge cases. You will spot coverage for long URLs and odd unicode characters that usually break naive splitters.

Wrapping plain text to strict column widths without mangling words is a specific pain point. It matters when you format plain text email fallbacks or calculate SMS segment boundaries.

The tool operates entirely offline and relies strictly on the Go standard library. No third-party modules or external services required. This keeps your build pipeline clean and avoids the usual supply chain headaches.