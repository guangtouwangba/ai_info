# TxGemma: Efficient and Agentic LLMs for Therapeutics

**URL**: http://arxiv.org/abs/2504.06196v1

## 原始摘要

Therapeutic development is a costly and high-risk endeavor that is often
plagued by high failure rates. To address this, we introduce TxGemma, a suite
of efficient, generalist large language models (LLMs) capable of therapeutic
property prediction as well as interactive reasoning and explainability. Unlike
task-specific models, TxGemma synthesizes information from diverse sources,
enabling broad application across the therapeutic development pipeline. The
suite includes 2B, 9B, and 27B parameter models, fine-tuned from Gemma-2 on a
comprehensive dataset of small molecules, proteins, nucleic acids, diseases,
and cell lines. Across 66 therapeutic development tasks, TxGemma achieved
superior or comparable performance to the state-of-the-art generalist model on
64 (superior on 45), and against state-of-the-art specialist models on 50
(superior on 26). Fine-tuning TxGemma models on therapeutic downstream tasks,
such as clinical trial adverse event prediction, requires less training data
than fine-tuning base LLMs, making TxGemma suitable for data-limited
applications. Beyond these predictive capabilities, TxGemma features
conversational models that bridge the gap between general LLMs and specialized
property predictors. These allow scientists to interact in natural language,
provide mechanistic reasoning for predictions based on molecular structure, and
engage in scientific discussions. Building on this, we further introduce
Agentic-Tx, a generalist therapeutic agentic system powered by Gemini 2.5 that
reasons, acts, manages diverse workflows, and acquires external domain
knowledge. Agentic-Tx surpasses prior leading models on the Humanity's Last
Exam benchmark (Chemistry &amp; Biology) with 52.3% relative improvement over
o3-mini (high) and 26.7% over o3-mini (high) on GPQA (Chemistry) and excels
with improvements of 6.3% (ChemBench-Preference) and 2.4% (ChemBench-Mini) over
o3-mini (high).


## AI 摘要

研究人员开发了TxGemma系列高效通用大语言模型（2B/9B/27B参数），专为药物研发设计。该模型在66项药物开发任务中表现优异：64项超越或持平现有通用模型（45项更优），50项超越专业模型（26项更优）。其特点包括：1）多源数据整合能力；2）小样本微调优势；3）支持自然语言交互解释。基于此还开发了Agentic-Tx智能体系统，在多个化学/生物学基准测试中显著领先（最高提升52.3%）。这套工具通过预测药物属性、提供机理解释和科学讨论功能，显著提升了药物开发的效率和可解释性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-09T23:01:29Z
- **目录日期**: 2025-04-09
