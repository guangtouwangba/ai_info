#!/bin/bash

# 查找所有包含 https\: 的文件
files=$(grep -l --include="*.md" "https\\:" content/papers/)

# 遍历每个文件并修复URL
for file in $files; do
    echo "修复文件: $file"
    # 将 https\: 替换为 https:
    sed -i '' 's/https\\:/https:/g' "$file"
done

echo "完成所有URL修复！" 