# Benchmarking Generative AI for Scoring Medical Student Interviews in Objective Structured Clinical Examinations (OSCEs)

**URL**: http://arxiv.org/abs/2501.13957v2

## 原始摘要

Objective Structured Clinical Examinations (OSCEs) are widely used to assess
medical students' communication skills, but scoring interview-based assessments
is time-consuming and potentially subject to human bias. This study explored
the potential of large language models (LLMs) to automate OSCE evaluations
using the Master Interview Rating Scale (MIRS). We compared the performance of
four state-of-the-art LLMs (GPT-4o, Claude 3.5, Llama 3.1, and Gemini 1.5 Pro)
in evaluating OSCE transcripts across all 28 items of the MIRS under the
conditions of zero-shot, chain-of-thought (CoT), few-shot, and multi-step
prompting. The models were benchmarked against a dataset of 10 OSCE cases with
174 expert consensus scores available. Model performance was measured using
three accuracy metrics (exact, off-by-one, thresholded). Averaging across all
MIRS items and OSCE cases, LLMs performed with low exact accuracy (0.27 to
0.44), and moderate to high off-by-one accuracy (0.67 to 0.87) and thresholded
accuracy (0.75 to 0.88). A zero temperature parameter ensured high intra-rater
reliability ({\alpha} = 0.98 for GPT-4o). CoT, few-shot, and multi-step
techniques proved valuable when tailored to specific assessment items. The
performance was consistent across MIRS items, independent of encounter phases
and communication domains. We demonstrated the feasibility of AI-assisted OSCE
evaluation and provided benchmarking of multiple LLMs across multiple prompt
techniques. Our work provides a baseline performance assessment for LLMs that
lays a foundation for future research into automated assessment of clinical
communication skills.


## AI 摘要

本研究探索了使用大型语言模型（LLMs）自动评估医学临床技能考试（OSCE）的可行性，测试了GPT-4o、Claude 3.5等四种先进模型在Master Interview评分量表（MIRS）28个项目上的表现。通过零样本、思维链等多种提示技术，模型在174份专家评分数据上的表现显示：完全准确率较低（0.27-0.44），但近似准确率较高（0.67-0.88）。零温度参数确保了高评分者内部一致性（GPT-4o α=0.98）。研究表明，AI辅助OSCE评估具有潜力，为临床沟通技能的自动化评估提供了基准，但需针对特定项目优化提示技术。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-17T03:18:44Z
- **目录日期**: 2025-05-17
