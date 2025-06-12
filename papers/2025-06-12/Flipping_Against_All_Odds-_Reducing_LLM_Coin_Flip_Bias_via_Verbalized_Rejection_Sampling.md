# Flipping Against All Odds: Reducing LLM Coin Flip Bias via Verbalized Rejection Sampling

**URL**: http://arxiv.org/abs/2506.09998v1

## 原始摘要

Large language models (LLMs) can often accurately describe probability
distributions using natural language, yet they still struggle to generate
faithful samples from them. This mismatch limits their use in tasks requiring
reliable stochasticity, such as Monte Carlo methods, agent-based simulations,
and randomized decision-making. We investigate this gap between knowledge and
sampling in the context of Bernoulli distributions. We introduce Verbalized
Rejection Sampling (VRS), a natural-language adaptation of classical rejection
sampling that prompts the LLM to reason about and accept or reject proposed
samples. Despite relying on the same Bernoulli mechanism internally, VRS
substantially reduces sampling bias across models. We provide theoretical
analysis showing that, under mild assumptions, VRS improves over direct
sampling, with gains attributable to both the algorithm and prompt design. More
broadly, our results show how classical probabilistic tools can be verbalized
and embedded into LLM workflows to improve reliability, without requiring
access to model internals or heavy prompt engineering.


## AI 摘要

大型语言模型(LLMs)能准确描述概率分布，但在生成可靠样本时仍存在困难。研究者针对伯努利分布提出"语言化拒绝采样"(VRS)方法，通过自然语言提示让LLM推理并接受/拒绝候选样本。实验表明VRS显著降低了采样偏差，理论分析证明该方法在温和假设下优于直接采样，其优势来自算法和提示设计的双重改进。研究表明，将经典概率工具语言化并整合到LLM工作流中，无需访问模型内部或复杂提示工程，即可提升可靠性。这一发现对需要可靠随机性的任务(如蒙特卡洛模拟)具有重要意义。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-12T22:00:54Z
- **目录日期**: 2025-06-12
