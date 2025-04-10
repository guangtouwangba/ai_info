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

研究人员开发了AutoConverter框架，将开放式视觉问答(VQA)自动转换为多选题形式，以解决现有VQA评测因回答多样性导致的评估困难问题。实验表明，该框架能生成正确且具挑战性的多选题，且视觉语言模型(VLM)在其上的表现与人工编写题目相当或更低。基于此，研究团队构建了包含20个VQA数据集转换而来的9,018道多选题的VMCBench基准，并对33个先进VLM进行了全面评估，为可扩展、一致且可复现的VLM评估设立了新标准。该方法显著降低了多选题创建成本，提升了评测客观性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-10T15:02:38Z
- **目录日期**: 2025-04-10
