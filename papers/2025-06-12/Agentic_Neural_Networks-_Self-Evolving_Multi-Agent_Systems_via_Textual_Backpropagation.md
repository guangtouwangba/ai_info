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

本文提出了Agentic Neural Network (ANN)框架，将多智能体协作建模为分层神经网络结构。每个智能体作为节点，每层形成专注于特定子任务的协作团队。ANN采用两阶段优化：前向阶段动态分解任务并逐层构建协作团队；反向阶段通过迭代反馈优化全局和局部协作，使智能体自我进化。这种神经符号方法使ANN能在训练后创建新的专业团队，显著提升准确性和适应性。在四个基准数据集上，ANN在相同配置下优于主流多智能体基线，性能持续提升。ANN为多智能体系统提供了可扩展、数据驱动的框架，结合了LLM的协作能力和神经网络的高效灵活性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-12T01:28:58Z
- **目录日期**: 2025-06-12
