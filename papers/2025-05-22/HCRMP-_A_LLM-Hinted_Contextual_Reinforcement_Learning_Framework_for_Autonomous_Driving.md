# HCRMP: A LLM-Hinted Contextual Reinforcement Learning Framework for Autonomous Driving

**URL**: http://arxiv.org/abs/2505.15793v1

## 原始摘要

Integrating Large Language Models (LLMs) with Reinforcement Learning (RL) can
enhance autonomous driving (AD) performance in complex scenarios. However,
current LLM-Dominated RL methods over-rely on LLM outputs, which are prone to
hallucinations.Evaluations show that state-of-the-art LLM indicates a
non-hallucination rate of only approximately 57.95% when assessed on essential
driving-related tasks. Thus, in these methods, hallucinations from the LLM can
directly jeopardize the performance of driving policies. This paper argues that
maintaining relative independence between the LLM and the RL is vital for
solving the hallucinations problem. Consequently, this paper is devoted to
propose a novel LLM-Hinted RL paradigm. The LLM is used to generate semantic
hints for state augmentation and policy optimization to assist RL agent in
motion planning, while the RL agent counteracts potential erroneous semantic
indications through policy learning to achieve excellent driving performance.
Based on this paradigm, we propose the HCRMP (LLM-Hinted Contextual
Reinforcement Learning Motion Planner) architecture, which is designed that
includes Augmented Semantic Representation Module to extend state space.
Contextual Stability Anchor Module enhances the reliability of multi-critic
weight hints by utilizing information from the knowledge base. Semantic Cache
Module is employed to seamlessly integrate LLM low-frequency guidance with RL
high-frequency control. Extensive experiments in CARLA validate HCRMP's strong
overall driving performance. HCRMP achieves a task success rate of up to 80.3%
under diverse driving conditions with different traffic densities. Under
safety-critical driving conditions, HCRMP significantly reduces the collision
rate by 11.4%, which effectively improves the driving performance in complex
scenarios.


## AI 摘要

该研究提出了一种新型LLM-Hinted RL范式，通过将大型语言模型（LLM）与强化学习（RL）结合来提升自动驾驶性能。针对现有LLM主导方法存在57.95%幻觉率的问题，新方法保持LLM与RL的相对独立性：LLM生成语义提示辅助RL进行运动规划，而RL通过策略学习纠正潜在错误。基于此提出的HCRMP架构包含状态增强、上下文稳定锚和语义缓存三大模块，在CARLA测试中实现了80.3%的任务成功率，并在安全关键场景下将碰撞率降低11.4%，显著提升了复杂场景下的驾驶性能。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-22T16:01:47Z
- **目录日期**: 2025-05-22
