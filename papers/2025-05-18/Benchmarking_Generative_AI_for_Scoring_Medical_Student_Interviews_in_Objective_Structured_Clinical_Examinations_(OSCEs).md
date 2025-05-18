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

本研究探讨了使用大型语言模型(LLMs)自动化评估医学生临床技能考试(OSCE)的可行性。研究人员测试了GPT-4o、Claude 3.5等四种先进LLM在零样本、思维链等不同提示技术下对28项MIRS量表的评估表现。结果显示，LLMs在精确匹配准确率(0.27-0.44)上表现一般，但在近似准确率(0.67-0.87)上表现良好，且具有高评分者内部信度(α=0.98)。研究表明，针对特定评估项目定制提示技术可提升性能，为未来临床沟通技能的自动化评估奠定了基础。该研究为LLMs在医学教育评估中的应用提供了基准参考。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-18T18:02:23Z
- **目录日期**: 2025-05-18
