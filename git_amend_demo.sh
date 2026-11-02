#!/usr/bin/env bash
#
# git_amend_demo.sh
# Demonstrates rewriting git commit author name, email and dates.
# Everything happens inside a throwaway temp repo; nothing outside it is changed.
#
# Every value below can be overridden from the command line, e.g.:
#   NEW_NAME="Ritesh" NEW_EMAIL="ritesh@example.com" ./git_amend_demo.sh
#
set -euo pipefail

# ---------- config ----------
# Identity the demo repo starts with (the "wrong" one we want to fix)
WRONG_NAME="${WRONG_NAME:-Wrong Name}"
WRONG_EMAIL="${WRONG_EMAIL:-wrong@example.com}"

# Identity used to fix the last commit (step 1) and the bulk rewrite (step 7)
NEW_NAME="${NEW_NAME:-Correct Name}"
NEW_EMAIL="${NEW_EMAIL:-correct@example.com}"

# Identity put into git config before --reset-author (step 4)
CONFIG_NAME="${CONFIG_NAME:-Config User}"
CONFIG_EMAIL="${CONFIG_EMAIL:-config@example.com}"

# Identity and date applied to an older commit via rebase (step 5)
REBASE_NAME="${REBASE_NAME:-Rebased Author}"
REBASE_EMAIL="${REBASE_EMAIL:-rebased@example.com}"
REBASE_DATE="${REBASE_DATE:-2026-05-05T09:00:00+0530}"

# Date applied when amending the last commit (steps 2 and 3)
NEW_DATE="${NEW_DATE:-2026-10-30T14:00:00+0530}"

# Initial commits: count and date pattern (day number is appended as 01, 02, ...)
NUM_COMMITS="${NUM_COMMITS:-3}"
INITIAL_DATE_PREFIX="${INITIAL_DATE_PREFIX:-2026-01-}"
INITIAL_TIME="${INITIAL_TIME:-T10:00:00+0530}"

# ---------- helpers ----------
BOLD=$'\033[1m'; GREEN=$'\033[32m'; YELLOW=$'\033[33m'; RESET=$'\033[0m'

section() { echo; echo "${BOLD}${GREEN}=== $* ===${RESET}"; }
note()    { echo "${YELLOW}-> $*${RESET}"; }

show_log() {  # extra args (e.g. -1) are passed to git log
  git log --date=iso-strict \
    --format='%h | %s%n        AUTHOR:    %an <%ae>  %ad%n        COMMITTER: %cn <%ce>  %cd' "$@"
}

# ---------- setup ----------
DEMO_DIR="$(mktemp -d -t git-amend-demo)"
trap 'echo; note "Demo repo left at: $DEMO_DIR  (delete with: rm -rf \"$DEMO_DIR\")"' EXIT
cd "$DEMO_DIR"

note "Using $(git --version)"
git init -q
# Local-only identity (does NOT touch your global ~/.gitconfig)
git config user.name  "$WRONG_NAME"
git config user.email "$WRONG_EMAIL"

for i in $(seq 1 "$NUM_COMMITS"); do
  day=$(printf '%02d' "$i")
  echo "line $i" >> file.txt
  git add file.txt
  GIT_AUTHOR_DATE="${INITIAL_DATE_PREFIX}${day}${INITIAL_TIME}" \
  GIT_COMMITTER_DATE="${INITIAL_DATE_PREFIX}${day}${INITIAL_TIME}" \
  git commit -q -m "commit $i"
done

section "INITIAL HISTORY ($NUM_COMMITS commits by '$WRONG_NAME')"
show_log

# ---------- 1. Amend author of last commit ----------
section "1. Amend AUTHOR of last commit (--author)"
OLD_HASH=$(git rev-parse --short HEAD)
git commit -q --amend --no-edit --author="$NEW_NAME <$NEW_EMAIL>"
NEW_HASH=$(git rev-parse --short HEAD)
show_log -1
note "Hash changed: $OLD_HASH -> $NEW_HASH"
note "Author changed, but COMMITTER is still '$WRONG_NAME' (committer = current git config)"

# ---------- 2. Amend author date only ----------
section "2. Amend AUTHOR DATE only (--date)"
git commit -q --amend --no-edit --date="$NEW_DATE"
show_log -1
note "Author date = $NEW_DATE, committer date = now (amend sets committer date to current time)"

# ---------- 3. Amend both dates ----------
section "3. Amend AUTHOR + COMMITTER dates (GIT_COMMITTER_DATE + --date)"
GIT_COMMITTER_DATE="$NEW_DATE" \
  git commit -q --amend --no-edit --date="$NEW_DATE"
show_log -1
note "Both dates are now $NEW_DATE"

# ---------- 4. --reset-author ----------
section "4. --reset-author (author := current config, author date := now)"
git config user.name  "$CONFIG_NAME"
git config user.email "$CONFIG_EMAIL"
git commit -q --amend --no-edit --reset-author
show_log -1
note "Author AND committer now taken from git config, both dates = now"

# ---------- 5. Older commit via rebase ----------
section "5. Change an OLDER commit (second-to-last) via interactive rebase"
note "Using GIT_SEQUENCE_EDITOR to auto-mark the first todo line as 'edit' (same as doing it by hand)"
GIT_SEQUENCE_EDITOR="sed -i.bak '1s/^pick/edit/'" git rebase -i HEAD~2 >/dev/null 2>&1
git commit -q --amend --no-edit \
  --author="$REBASE_NAME <$REBASE_EMAIL>" \
  --date="$REBASE_DATE"
git rebase --continue >/dev/null 2>&1
show_log
note "Second-to-last commit changed; the last commit also got a NEW hash because its parent changed"

# ---------- 6. --committer-date-is-author-date ----------
section "6. Rebase with --committer-date-is-author-date"
git rebase --force-rebase --committer-date-is-author-date --root >/dev/null 2>&1
show_log
# Check the result instead of assuming it: --root rebases on git < 2.29
# go through the interactive backend, which silently ignores this flag.
if git log --format='%ad %cd' --date=raw | awk '$1!=$3 {bad=1} END {exit bad}'; then
  note "PASS: every commit's COMMITTER date now equals its AUTHOR date"
else
  note "NOT APPLIED: committer dates differ from author dates (needs git >= 2.29 for --root)"
fi

# ---------- 7. Bulk rewrite of email across ALL commits ----------
section "7. Bulk rewrite: replace '$WRONG_EMAIL' everywhere"
# make sure some commit still has the wrong identity
git commit -q --allow-empty -m "commit $((NUM_COMMITS + 1)) (wrong identity)" \
  --author="$WRONG_NAME <$WRONG_EMAIL>"
note "Before bulk rewrite:"
show_log -1

if command -v git-filter-repo >/dev/null 2>&1; then
  note "Using git filter-repo with a mailmap"
  echo "$NEW_NAME <$NEW_EMAIL> <$WRONG_EMAIL>" > ../mailmap.txt
  git filter-repo --force --quiet --mailmap ../mailmap.txt
  rm -f ../mailmap.txt
else
  note "git-filter-repo not installed (brew install git-filter-repo); falling back to built-in filter-branch"
  # Values are exported so the env-filter (run in a subshell per commit) can read them
  export WRONG_EMAIL NEW_NAME NEW_EMAIL
  FILTER_BRANCH_SQUELCH_WARNING=1 git filter-branch -f --env-filter '
    if [ "$GIT_AUTHOR_EMAIL" = "$WRONG_EMAIL" ]; then
      export GIT_AUTHOR_NAME="$NEW_NAME"
      export GIT_AUTHOR_EMAIL="$NEW_EMAIL"
    fi
    if [ "$GIT_COMMITTER_EMAIL" = "$WRONG_EMAIL" ]; then
      export GIT_COMMITTER_NAME="$NEW_NAME"
      export GIT_COMMITTER_EMAIL="$NEW_EMAIL"
    fi
  ' -- --all >/dev/null
fi
note "After bulk rewrite:"
show_log

# ---------- verification ----------
section "VERIFY: any '$WRONG_EMAIL' left?"
if git log --format='%ae %ce' | grep -qF "$WRONG_EMAIL"; then
  echo "FOUND - rewrite incomplete"
else
  echo "None - all commits rewritten"
fi

section "DONE"
note "To push rewritten history to a real remote you would run: git push --force-with-lease"
