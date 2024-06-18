# Possibilities

1. forbidden flags
2. flag name priority
   - example: `--file` should be `-f`, but can be `-F` (maybe a warning if preferred isn't used?)
3. one to many relationships
   - example: `--output` can be `-o` or `-f` or `-w` (no preference)
4. limit number of flags that can be associated with a single variable
5. limit similarity between flags ([Levenshtein](https://en.wikipedia.org/wiki/Levenshtein_distance) or [Damerau-Levenshtein](https://en.wikipedia.org/wiki/Damerau%E2%80%93Levenshtein_distance))
6. Spelling (is this already covered by other spelling linters?)
7. character encoding (other linters for this, though)
8. flags that require a short version
   - if `--help` exists, it must have `-h`
9. forbidden flag imports (already covered by forbidigo)