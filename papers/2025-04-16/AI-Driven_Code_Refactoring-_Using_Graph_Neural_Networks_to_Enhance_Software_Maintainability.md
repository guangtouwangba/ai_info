# AI-Driven Code Refactoring: Using Graph Neural Networks to Enhance Software Maintainability

**URL**: http://arxiv.org/abs/2504.10412v1

## 原始摘要

This study explores Graph Neural Networks (GNNs) as a transformative tool for
code refactoring, using abstract syntax trees (ASTs) to boost software
maintainability. It analyzes a dataset of 2 million snippets from CodeSearchNet
and a custom 75000-file GitHub Python corpus, comparing GNNs against rule-based
SonarQube and decision trees. Metrics include cyclomatic complexity (target
below 10), coupling (target below 5), and refactoring precision. GNNs achieve
92% accuracy, reducing complexity by 35% and coupling by 33%, outperforming
SonarQube (78%, 16%) and decision trees (85%, 25%). Preprocessing fixed 60% of
syntax errors. Bar graphs, tables, and AST visuals clarify results. This offers
a scalable AI-driven path to cleaner codebases, which is crucial for software
engineering.


## AI 摘要

本研究利用图神经网络(GNN)和抽象语法树(AST)改进代码重构，分析了205万代码片段和7.5万Python文件。相比基于规则的SonarQube(78%准确率)和决策树(85%)，GNN达到92%准确率，将代码复杂度降低35%，耦合度降低33%。预处理修复了60%语法错误。通过复杂度(目标<10)、耦合度(目标<5)等指标评估，GNN展现出显著优势，为软件工程提供了可扩展的AI驱动方案，有效提升代码可维护性。研究结果通过柱状图、表格和AST可视化清晰呈现。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-16T03:17:50Z
- **目录日期**: 2025-04-16
