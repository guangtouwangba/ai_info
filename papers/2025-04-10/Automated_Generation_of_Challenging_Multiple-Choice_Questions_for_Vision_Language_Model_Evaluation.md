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

针对视觉语言模型(VLM)评估的挑战，研究者提出AutoConverter框架，可将开放式视觉问答(VQA)自动转换为多选题形式，实现更客观的评估。实验表明，该方法生成的多选题准确且具有挑战性，模型在其上的表现与人工创建题目相当或更低。基于此构建的VMCBench基准整合了20个VQA数据集共9,018道多选题，并对33个前沿VLM进行全面评估，为可扩展、一致且可复现的VLM评估设立了新标准。该研究解决了开放式问题评估的局限性，降低了多选题创建成本。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-10T16:02:36Z
- **目录日期**: 2025-04-10
