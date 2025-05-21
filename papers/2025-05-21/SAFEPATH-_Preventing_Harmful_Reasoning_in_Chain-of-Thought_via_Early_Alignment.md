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

SAFEPATH是一种轻量级对齐方法，用于提升大型推理模型(LRMs)的安全性。它通过在有害提示下生成8个token的"安全引导词"来减少有害输出，同时保持无监督的后续推理能力。实验表明，SAFEPATH在DeepSeek-R1-Distill-Llama-8B模型上能减少90%有害响应，并阻挡83.3%的越狱攻击，计算量比直接拒绝方法减少295.9倍。该方法还提供了无需微调的零样本变体。研究同时揭示了现有LLM安全方法在推理模型上的局限性，为更安全的AI发展提供了新方向。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-21T07:01:44Z
- **目录日期**: 2025-05-21
