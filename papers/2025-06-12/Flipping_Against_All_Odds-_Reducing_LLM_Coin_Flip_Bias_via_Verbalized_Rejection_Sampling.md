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

该研究探讨了大语言模型(LLMs)在描述概率分布与生成可靠样本之间的差距，针对伯努利分布提出了"语言化拒绝采样"(VRS)方法。VRS通过自然语言提示让LLM对提议样本进行推理判断，显著降低了采样偏差。理论分析表明，在温和假设下，VRS优于直接采样，改进来自算法和提示设计的双重作用。研究表明，无需修改模型内部或复杂提示工程，通过将经典概率工具语言化并嵌入LLM工作流，即可提高可靠性。这对需要可靠随机性的任务(如蒙特卡洛方法)具有重要意义。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-12T11:01:04Z
- **目录日期**: 2025-06-12
