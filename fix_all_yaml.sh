#!/bin/bash

echo "开始扫描并修复所有Markdown文件中的转义字符..."

# 获取所有Markdown文件列表
files=$(find content/papers -name "*.md")
total_files=0
fixed_files=0

# 处理每个文件
for file in $files; do
  total_files=$((total_files + 1))
  
  # 检测文件是否包含转义字符
  if grep -q '\\:' "$file" || grep -q 'https\\:' "$file" || grep -q 'http\\:' "$file"; then
    echo "发现需要修复的文件: $file"
    
    # 创建临时文件
    temp_file=$(mktemp)
    
    # 替换URL中的转义冒号 (https\: 和 http\:)
    sed 's/https\\:/https:/g; s/http\\:/http:/g; s/\\:/:/g' "$file" > "$temp_file"
    
    # 将修改后的内容移回原文件
    mv "$temp_file" "$file"
    fixed_files=$((fixed_files + 1))
    echo "已修复: $file"
  fi
  
  # 每处理50个文件输出一次进度
  if [[ $((total_files % 50)) -eq 0 ]]; then
    echo "已处理 $total_files 个文件..."
  fi
done

echo "========================================="
echo "处理完成! 共扫描 $total_files 个文件，修复了 $fixed_files 个文件。"
echo "=========================================" 