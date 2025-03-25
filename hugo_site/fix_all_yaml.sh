#!/bin/bash

echo "开始扫描并修复所有含转义URL的Markdown文件..."

# 直接使用find命令找到所有包含转义URL的文件
url_files=$(find content/papers -name "*.md" -exec grep -l "https\\\\:" {} \;)
total_files=0
fixed_files=0

# 处理每个文件
for file in $url_files; do
  total_files=$((total_files + 1))
  echo "修复文件: $file"
  
  # 创建临时文件
  temp_file=$(mktemp)
  
  # 替换所有转义URL为普通URL
  sed 's/https\\:/https:/g' "$file" > "$temp_file"
  
  # 将修改后的内容移回原文件
  mv "$temp_file" "$file"
  fixed_files=$((fixed_files + 1))
  echo "已修复: $file"
done

echo "========================================="
echo "处理完成! 共修复了 $fixed_files 个文件。"
echo "=========================================" 