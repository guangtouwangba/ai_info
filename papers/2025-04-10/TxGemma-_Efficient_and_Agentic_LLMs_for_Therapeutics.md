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

TxGemma是一套高效、通用的治疗开发大语言模型(LLM)套件，包含2B、9B和27B参数模型，在66项治疗开发任务中表现优异。相比专业模型，TxGemma在50项任务中表现相当或更好(26项更优)，且在小样本下游任务(如临床试验不良事件预测)中表现突出。该模型支持自然语言交互，能基于分子结构提供预测机制解释。研究还推出了基于Gemini 2.5的Agentic-Tx智能代理系统，在多个化学和生物学基准测试中显著超越先前最佳模型，最高提升达52.3%。这套工具显著提升了药物研发的效率和可解释性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-10T01:28:54Z
- **目录日期**: 2025-04-10
