# On Path to Multimodal Historical Reasoning: HistBench and HistAgent

**URL**: http://arxiv.org/abs/2505.20246v1

## 原始摘要

Recent advances in large language models (LLMs) have led to remarkable
progress across domains, yet their capabilities in the humanities, particularly
history, remain underexplored. Historical reasoning poses unique challenges for
AI, involving multimodal source interpretation, temporal inference, and
cross-linguistic analysis. While general-purpose agents perform well on many
existing benchmarks, they lack the domain-specific expertise required to engage
with historical materials and questions. To address this gap, we introduce
HistBench, a new benchmark of 414 high-quality questions designed to evaluate
AI's capacity for historical reasoning and authored by more than 40 expert
contributors. The tasks span a wide range of historical problems-from factual
retrieval based on primary sources to interpretive analysis of manuscripts and
images, to interdisciplinary challenges involving archaeology, linguistics, or
cultural history. Furthermore, the benchmark dataset spans 29 ancient and
modern languages and covers a wide range of historical periods and world
regions. Finding the poor performance of LLMs and other agents on HistBench, we
further present HistAgent, a history-specific agent equipped with carefully
designed tools for OCR, translation, archival search, and image understanding
in History. On HistBench, HistAgent based on GPT-4o achieves an accuracy of
27.54% pass@1 and 36.47% pass@2, significantly outperforming LLMs with online
search and generalist agents, including GPT-4o (18.60%), DeepSeek-R1(14.49%)
and Open Deep Research-smolagents(20.29% pass@1 and 25.12% pass@2). These
results highlight the limitations of existing LLMs and generalist agents and
demonstrate the advantages of HistAgent for historical reasoning.


## AI 摘要

大型语言模型（LLMs）在人文领域，尤其是历史推理方面仍存在不足。为评估AI的历史推理能力，研究者开发了包含414个问题的HistBench基准测试，涵盖多模态源解释、跨语言分析等任务。测试显示，现有LLMs表现不佳（如GPT-4o准确率18.6%）。为此，团队推出专用历史智能体HistAgent，整合OCR、翻译、档案搜索等工具，在HistBench上达到27.54%（pass@1）和36.47%（pass@2）的准确率，显著优于通用模型。该研究揭示了通用LLMs在历史领域的局限，展示了领域专用方案的优势。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-27T13:10:27Z
- **目录日期**: 2025-05-27
