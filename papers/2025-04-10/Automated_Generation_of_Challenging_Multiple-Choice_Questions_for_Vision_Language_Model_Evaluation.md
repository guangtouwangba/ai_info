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

本文提出AutoConverter框架，可将开放式视觉问答(VQA)自动转换为多选题形式，实现更客观的评估。该框架生成的题目具有挑战性，且视觉语言模型(VLMs)在其上的表现与人工编写题目相当或更低。基于此，研究团队构建了VMCBench基准，将20个现有VQA数据集统一转换为9,018道多选题，并对33个前沿VLMs进行全面评估。该方法为可扩展、一致且可复现的VLM评估设立了新标准，解决了开放式问题因自然语言变异性导致的评估困难，同时降低了多选题创建成本。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-10T10:02:23Z
- **目录日期**: 2025-04-10
