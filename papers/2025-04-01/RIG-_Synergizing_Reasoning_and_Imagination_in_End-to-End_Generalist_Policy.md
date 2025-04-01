# RIG: Synergizing Reasoning and Imagination in End-to-End Generalist Policy

**URL**: http://arxiv.org/abs/2503.24388v1

## 原始摘要

Reasoning before action and imagining potential outcomes (i.e., world models)
are essential for embodied agents operating in complex open-world environments.
Yet, prior work either incorporates only one of these abilities in an
end-to-end agent or integrates multiple specialized models into an agent
system, limiting the learning efficiency and generalization of the policy.
Thus, this paper makes the first attempt to synergize Reasoning and Imagination
in an end-to-end Generalist policy, termed RIG. To train RIG in an end-to-end
manner, we construct a data pipeline that progressively integrates and enriches
the content of imagination and reasoning in the trajectories collected from
existing agents. The joint learning of reasoning and next image generation
explicitly models the inherent correlation between reasoning, action, and
dynamics of environments, and thus exhibits more than $17\times$ sample
efficiency improvements and generalization in comparison with previous works.
During inference, RIG first reasons about the next action, produces potential
action, and then predicts the action outcomes, which offers the agent a chance
to review and self-correct based on the imagination before taking real actions.
Experimental results show that the synergy of reasoning and imagination not
only improves the robustness, generalization, and interoperability of
generalist policy but also enables test-time scaling to enhance overall
performance.


## AI 摘要

这篇论文提出了一种名为RIG的端到端通用策略，首次将推理（Reasoning）与想象（Imagination）协同整合到单一模型中，用于复杂开放环境中的智能体决策。通过构建渐进式数据管道联合训练推理与图像生成，RIG显式建模了推理、动作与环境动态的关联性，相比先前方法实现了17倍以上的样本效率提升和泛化能力。推理时，RIG先预测动作，再想象结果，使智能体能在执行前自我修正。实验表明，这种协同机制不仅增强了策略的鲁棒性、泛化性和互操作性，还能通过测试时扩展提升整体性能。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-01T17:01:49Z
- **目录日期**: 2025-04-01
