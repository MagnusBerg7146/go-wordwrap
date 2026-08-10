# Greedy word wrap

```
wordwrap.go
```

Run the Go Wordwrap test next to the implementation for concrete examples.

This is a small utility that wraps text to a fixed column width. It never splits a word mid-way, so your output stays readable whether you're formatting logs, terminal output, or plain-text emails. The algorithm is greedy: it packs as many words as fit on each line, then moves on.

There's no magic here. The code relies solely on the Go standard library — no external packages, no services to spin up, no dependencies to vendor. You can drop it into any project and it just works.

The test file sits right next to the implementation, so you can see exactly how the wrapping behaves with different inputs: long words, short lines, empty strings, and edge cases where a single word exceeds the column width. If you've ever debugged an OTP email that wrapped mid-code or a log line that got mangled by a naive formatter, you know why these cases matter.

Use it wherever you need predictable, word-safe wrapping without pulling in a heavyweight text library.