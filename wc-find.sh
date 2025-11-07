# 
# Usage:
#
# 1 - Set some regex filter for `find` with `REGEX="..."`. Examples:
# ```
# REGEX=".*\.\(go\|rs\|js*\|ts*\|md\)"   # Go, Rust, JS, TS, Markdown
# REGEX=".*\.\(go\|rs\|md\)"             # Go, Rust, Markdown
# ```
#
# 2 - Run the script with REGEX as an env var.
# ```
# REGEX=$REGEX bash wc-find.sh
# ```

echo "REGEX: $REGEX"

find . -regex "$REGEX" -exec wc -l {} + | tail -1

# Directory analysis
for dir in */; do
  if [ -d "$dir" ]; then
    lines=$(find "$dir" -regex "$REGEX" -exec wc -l {} + 2>/dev/null | tail -1 | awk '{print $1}' || echo "0")
    files=$(find "$dir" -regex "$REGEX" | wc -l)
    echo "$dir: $files files, $lines lines"
  fi
done | sort -k4 -nr
