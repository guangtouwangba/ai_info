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

本文提出了一种名为RIG的端到端通用策略，首次将推理（Reasoning）与想象（Imagination）能力协同整合。通过构建渐进式数据管道，RIG联合学习推理与下一帧图像生成，显式建模环境动态关联，样本效率提升17倍以上。推理时，RIG先预测行动，再想象结果，使代理能在真实行动前自我修正。实验表明，这种协同显著增强了策略的鲁棒性、泛化性和互操作性，并支持测试时扩展以提升整体性能，克服了先前方法中单一能力或模块化整合的局限性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-01T23:01:04Z
- **目录日期**: 2025-04-01
