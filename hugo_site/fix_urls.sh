#!/bin/bash

# 查找所有包含 https\: 的文件
find content/papers -name "*.md" -exec sed -i '' 's/https\\:/https:/g' {} \;

echo "完成所有文件的URL修复" 