# Automated Generation of Challenging Multiple-Choice Questions for Vision Language Model Evaluation

**URL**: http://arxiv.org/abs/2501.03225v2

## 原始摘要

The rapid development of vision language models (VLMs) demands rigorous and
reliable evaluation. However, current visual question answering (VQA)
benchmarks often depend on open-ended questions, making accurate evaluation
difficult due to the variability in natural language responses. To address
this, we introduce AutoConverter, an agentic framework that automatically
converts these open-ended questions into multiple-choice format, enabling
objective evaluation while reducing the costly multiple-choice question
creation process. Our experiments demonstrate that AutoConverter can generate
correct and challenging multiple-choice questions, with VLMs demonstrating
consistently similar or lower accuracy on these questions compared to
human-created ones. Using AutoConverter, we construct VMCBench, a benchmark
created by transforming 20 existing VQA datasets into a unified multiple-choice
format, totaling 9,018 questions. We comprehensively evaluate 33
state-of-the-art VLMs on VMCBench, setting a new standard for scalable,
consistent, and reproducible VLM evaluation.


## AI 摘要

当前视觉语言模型(VLM)评估面临开放性问题导致评价困难的挑战。研究者提出AutoConverter框架，能自动将开放性问题转换为多选题格式，实现客观评估并降低人工成本。实验表明，该框架能生成正确且具挑战性的多选题，且VLM在新题上的准确率与人工题相当或更低。基于此构建的VMCBench基准整合了20个VQA数据集共9,018道多选题，并对33个前沿VLM进行全面评估，为可扩展、一致且可复现的VLM评估设立了新标准。该研究解决了VLM评估中的关键难题。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-10T05:02:25Z
- **目录日期**: 2025-04-10
