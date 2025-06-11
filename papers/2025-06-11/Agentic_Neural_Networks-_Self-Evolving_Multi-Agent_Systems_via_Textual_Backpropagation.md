# Agentic Neural Networks: Self-Evolving Multi-Agent Systems via Textual Backpropagation

**URL**: http://arxiv.org/abs/2506.09046v1

## 原始摘要

Leveraging multiple Large Language Models(LLMs) has proven effective for
addressing complex, high-dimensional tasks, but current approaches often rely
on static, manually engineered multi-agent configurations. To overcome these
constraints, we present the Agentic Neural Network(ANN), a framework that
conceptualizes multi-agent collaboration as a layered neural network
architecture. In this design, each agent operates as a node, and each layer
forms a cooperative "team" focused on a specific subtask. Agentic Neural
Network follows a two-phase optimization strategy: (1) Forward Phase-Drawing
inspiration from neural network forward passes, tasks are dynamically
decomposed into subtasks, and cooperative agent teams with suitable aggregation
methods are constructed layer by layer. (2) Backward Phase-Mirroring
backpropagation, we refine both global and local collaboration through
iterative feedback, allowing agents to self-evolve their roles, prompts, and
coordination. This neuro-symbolic approach enables ANN to create new or
specialized agent teams post-training, delivering notable gains in accuracy and
adaptability. Across four benchmark datasets, ANN surpasses leading multi-agent
baselines under the same configurations, showing consistent performance
improvements. Our findings indicate that ANN provides a scalable, data-driven
framework for multi-agent systems, combining the collaborative capabilities of
LLMs with the efficiency and flexibility of neural network principles. We plan
to open-source the entire framework.


## AI 摘要

本文提出了Agentic Neural Network（ANN）框架，将多智能体协作建模为分层神经网络架构。每个智能体作为节点，每层组成专注于子任务的协作团队。ANN采用两阶段优化：前向阶段动态分解任务并逐层构建智能体团队；反向阶段通过迭代反馈优化协作，使智能体自我进化角色、提示和协调。这种神经符号方法使ANN能在训练后创建新的专业团队，显著提升准确性和适应性。在四个基准测试中，ANN在相同配置下超越主流多智能体基线，表现出稳定性能提升。该框架结合了LLM的协作能力与神经网络的高效灵活性，具有可扩展性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-11T19:01:23Z
- **目录日期**: 2025-06-11
