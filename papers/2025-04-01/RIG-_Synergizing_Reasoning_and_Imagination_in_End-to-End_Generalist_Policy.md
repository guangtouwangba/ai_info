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

本文提出了RIG框架，首次将推理与想象能力整合到端到端的通用策略中。通过构建渐进式数据管道，联合训练推理和下一帧生成，显式建模了推理、动作与环境动态的关联性，相比之前工作提升了17倍以上的样本效率和泛化能力。推理时，RIG先推理动作，生成潜在动作，再预测结果，使代理能在执行前通过想象进行自我修正。实验表明，推理与想象的协同不仅提升了策略的鲁棒性、泛化性和互操作性，还能通过测试时扩展增强整体性能。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-01T21:00:58Z
- **目录日期**: 2025-04-01
