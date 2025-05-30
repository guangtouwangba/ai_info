# SAFEPATH: Preventing Harmful Reasoning in Chain-of-Thought via Early Alignment

**URL**: http://arxiv.org/abs/2505.14667v1

## 原始摘要

Large Reasoning Models (LRMs) have become powerful tools for complex problem
solving, but their structured reasoning pathways can lead to unsafe outputs
when exposed to harmful prompts. Existing safety alignment methods reduce
harmful outputs but can degrade reasoning depth, leading to significant
trade-offs in complex, multi-step tasks, and remain vulnerable to sophisticated
jailbreak attacks. To address this, we introduce SAFEPATH, a lightweight
alignment method that fine-tunes LRMs to emit a short, 8-token Safety Primer at
the start of their reasoning, in response to harmful prompts, while leaving the
rest of the reasoning process unsupervised. Empirical results across multiple
benchmarks indicate that SAFEPATH effectively reduces harmful outputs while
maintaining reasoning performance. Specifically, SAFEPATH reduces harmful
responses by up to 90.0% and blocks 83.3% of jailbreak attempts in the
DeepSeek-R1-Distill-Llama-8B model, while requiring 295.9x less compute than
Direct Refusal and 314.1x less than SafeChain. We further introduce a zero-shot
variant that requires no fine-tuning. In addition, we provide a comprehensive
analysis of how existing methods in LLMs generalize, or fail, when applied to
reasoning-centric models, revealing critical gaps and new directions for safer
AI.


## AI 摘要

SAFEPATH是一种轻量级对齐方法，通过微调大型推理模型(LRMs)，使其在遇到有害提示时首先生成一个简短的8-token安全提示，而后续推理过程不受监督。实验表明，SAFEPATH在保持推理性能的同时，显著减少有害输出，最高可降低90%的有害回答，并拦截83.3%的越狱攻击，计算成本比现有方法低数百倍。此外，还提出了无需微调的零样本变体。研究还分析了现有LLM方法在推理模型上的适用性，揭示了关键差距，为更安全的AI提供了新方向。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-22T01:29:11Z
- **目录日期**: 2025-05-22
