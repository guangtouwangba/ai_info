# Co-Evolving LLM Coder and Unit Tester via Reinforcement Learning

**URL**: http://arxiv.org/abs/2506.03136v1

## 原始摘要

We propose CURE, a novel reinforcement learning framework with a dedicated
reward design that co-evolves coding and unit test generation capabilities
based on their interaction outcomes, without any ground-truth code as
supervision. This approach enables flexible and scalable training and allows
the unit tester to learn directly from the coder's mistakes. Our derived
ReasonFlux-Coder-7B and 14B models improve code generation accuracy by 5.3% and
Best-of-N accuracy by 9.0% after optimization on Qwen2.5-Instruct models,
outperforming similarly sized Qwen-Coder, DeepSeek-Coder, and Seed-Coder. They
naturally extend to downstream tasks such as test-time scaling and agentic
coding-achieving a 8.1% improvement over the base model. For the long-CoT
model, our ReasonFlux-Coder-4B consistently outperforms Qwen3-4B while
achieving 64.8% inference efficiency in unit test generation. Notably, we also
find that our model can serve as an effective reward model for reinforcement
learning on base models. Project: https://github.com/Gen-Verse/CURE


## AI 摘要

研究人员提出了CURE框架，一种通过强化学习协同进化代码生成与单元测试能力的新方法。该框架采用专用奖励机制，无需真实代码作为监督，使测试器能直接从编码错误中学习。基于Qwen2.5-Instruct优化的ReasonFlux-Coder-7B/14B模型，代码生成准确率提升5.3%，Best-of-N准确率提高9%，超越同类模型。在长链思维任务中，4B版本保持64.8%的测试生成效率且优于Qwen3-4B。该模型还可作为强化学习的奖励模型，并在代理编码等下游任务中实现8.1%的性能提升。项目已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-04T08:01:28Z
- **目录日期**: 2025-06-04
