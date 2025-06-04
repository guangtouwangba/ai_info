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

研究人员提出了CURE框架，这是一种新颖的强化学习方法，通过协同进化代码生成和单元测试能力，无需真实代码作为监督。该方法基于交互结果设计专用奖励机制，使单元测试器能从编码错误中学习。优化后的ReasonFlux-Coder-7B/14B模型在Qwen2.5-Instruct上实现了5.3%的代码生成准确率提升和9.0%的Best-of-N提升，优于同类模型。该框架还支持测试时扩展和智能编码等下游任务，性能提升8.1%。ReasonFlux-Coder-4B在长链思维任务中超越Qwen3-4B，同时保持64.8%的测试生成效率。该模型还可作为基础模型的强化学习奖励模型。项目已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-04T22:01:09Z
- **目录日期**: 2025-06-04
