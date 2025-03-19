#!/bin/bash

echo "安装AWS SDK for Go v2及相关组件..."

# 安装AWS SDK Core
go get github.com/aws/aws-sdk-go-v2
go get github.com/aws/aws-sdk-go-v2/config

# 安装DynamoDB相关依赖
go get github.com/aws/aws-sdk-go-v2/service/dynamodb
go get github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue

echo "AWS SDK安装完成!"
echo 
echo "如何配置AWS DynamoDB:"
echo "1. 在.env文件中添加以下配置"
echo "   APP_STORAGE_PROVIDER=dynamodb          # 用于存储"
echo "   APP_DEDUPLICATION_DATABASE_TYPE=dynamodb  # 用于去重"
echo "   APP_AWS_REGION=us-east-1               # 您的AWS区域"
echo "   APP_AWS_PROFILE=default                # 您的AWS配置文件"
echo "   APP_AWS_DYNAMODB_TABLE=arxiv_papers    # 存储表名"
echo "   APP_DEDUPLICATION_TABLE_PREFIX=dedup_  # 去重表前缀"
echo 
echo "2. 确保您的AWS凭证已配置："
echo "   - 使用AWS CLI: aws configure"
echo "   - 或在~/.aws/credentials文件中设置凭证"
echo 
echo "3. 首次使用时，系统会自动创建必要的DynamoDB表" 