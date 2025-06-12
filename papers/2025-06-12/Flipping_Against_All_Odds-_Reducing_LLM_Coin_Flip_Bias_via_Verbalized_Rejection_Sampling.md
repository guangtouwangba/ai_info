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

研究发现大型语言模型(LLMs)能准确描述概率分布但难以生成可靠样本，限制了其在蒙特卡洛方法等随机任务中的应用。针对伯努利分布，研究者提出"语言化拒绝采样"(VRS)方法，通过自然语言提示让LLM推理并接受/拒绝样本。实验表明VRS显著减少了采样偏差，理论分析证实其在温和假设下优于直接采样。该方法将经典概率工具融入LLM工作流，无需修改模型内部结构或复杂提示工程即可提升可靠性。研究展示了如何通过算法和提示设计改进LLM的随机采样能力。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-12T23:00:56Z
- **目录日期**: 2025-06-12
