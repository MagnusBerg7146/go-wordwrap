# Greedy word wrap

```
wordwrap.go
```

Check the test next to the source for usage examples.

Wraps text to a given column width, keeping words intact — no external packages, no setup, just the standard library.

The algorithm walks through the input word by word, tracking how much space remains on the current line. If the next word fits, it gets appended; otherwise, the line is flushed and the word starts a new one. This is the classic greedy approach: fast, simple, and good enough for most formatting needs.

Edge cases are handled explicitly. A word longer than the width will still be emitted on its own line rather than silently dropped or mangled. Empty lines in the input are preserved, so paragraph breaks don't collapse. Tabs and multiple spaces are treated as single separators, which keeps the output predictable without over-engineering the tokenizer.

The function returns a list of strings, one per line. If you want a single string with newlines, join it yourself — that keeps the API minimal and lets the caller decide on line endings. The implementation is deliberately short; there's no configuration object, no callbacks, no cleverness. Just give it text and a width, get lines back.

For anything more elaborate — hyphenation, justification, or CJK-aware breaking — you'd want a dedicated library. But for plain prose, config files, or log output, this covers the common case without pulling in a dependency.