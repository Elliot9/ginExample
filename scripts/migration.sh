#!/bin/bash

# 獲取腳本所在的目錄
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# 切換到項目根目錄
# 假設 scripts 目錄在項目根目錄下
cd "$SCRIPT_DIR/.." || exit

echo "Start to Migrate..."
if [ -f "./cmd/migrate/main.go" ]; then
    time go run ./cmd/migrate/main.go
else
    echo "Error: ./cmd/migrate/main.go not found"
    echo "Current directory: $(pwd)"
    echo "Contents of ./cmd directory:"
    ls -R ./cmd
    exit 1
fi